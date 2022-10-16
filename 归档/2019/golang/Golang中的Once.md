# Golang中的Once
有些时候我们创建一些单例的对象，即在程序的生命周期里只会产生一个，之后都是重复使用这个对象。

在golang里，采用once就可以很方便的实现这个需求。

```go
type Config struct {
	AppName string
}

var config Config

var once sync.Once

func CreateSingleInstance() {
	once.Do(func() {
		fmt.Println("CreateSingleInstance")
		config = Config{
			AppName: "test",
		}
	})
}

func main() {
	for i := 0; i < 2; i++ {
		go CreateSingleInstance()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(config.AppName)
}
```

代码运行结果为：
```
CreateSingleInstance
test
```

可见，虽然并发的调用了2次CreateSingleInstance，但是Config实例的创建只执行了一次。

## Once的实现
golang中Once的源码非常简单，主要的就是Do函数的实现过程。

而Do方法的执行流程也很简单：
1. 先检查once.done的值是否为1，如果是则直接返回，不做任何事情。
2. 如果once.done的值为0，则：
    1. 先将once.done改为1
    2. 再执行函数参数

但是为了处理并发调用Do时导致的并发问题，例如两个goroutine同时调用Do，发现都为0，然后就都调用了函数参数，这是不行的。因此Once中采用了原子操作和锁来解决这个问题。

```go
type Once struct {
    m    Mutex
    done uint32
}

func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    o.m.Lock()
    defer o.m.Unlock()
    if (o.done == 0) {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

其中atomic.LoadUint32导致了Once结构体中的done属性类型必须是uint32的，虽然它的值只会是0或者1。

通过原子操作，让所有goroutine中读取done值时都是原子的，即不会导致done修改了一半被读取到了。
但是这样还是不够的，因为有可能两个goroutine都是读取到0，都开始执行后续的修改操作。因此接下来需要加锁。

加了锁后，在临界区中继续判断done的值（因为有可能其他goroutine已经修改了done的值并且返回），再通过原子操作修改done的值为1，最终执行函数参数。

这种先在临界区外，判断一次关键条件，若条件不满足则立即返回。这叫做*快路径*，或者叫*快速失败路径*。  
加了锁后，多少会有性能损失，因此在进入临界区后，第二次判断以及后续的操作都被称为*慢路径*（*Slow-path*），或者叫*常规路径*。

## Once的中Do的特点
1. 如果Do中的参数函数执行后，一直不返回，就会导致相关的goroutine阻塞。
2. 这里Do中采用defer的方式来修改done的值，也就是说无论参数函数是以什么方式结束（例如抛出panic），done都会被修改为1，也就是说以后也不能重复执行参数函数了。


