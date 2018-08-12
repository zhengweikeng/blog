# golang
### 从GitHub或者bitbucket导入代码的命令是什么？
go get 和 go install

### 如何理解slice
slice，即切片，它自己不拥有元素，本质上它只是一个对数组的引用，即它底层就是一个数组，它包含三个属性：指针、长度和容量。   
1. 指针指向数组的第一个可以从slice中访问的元素，但是这种元素不一定是数组的第一个元素
2. 长度是slice中元素的个数，但是不能超过slice的容量。
3. 容量的大小通常是从slice的其实元素到底层数组的最后一个元素

```go
var sli = []int{1, 2, 3, 4, 5}
// [2] 1 4
fmt.Println(sli[1:2], len(sli[1:2]), cap(sli[1:2]))
```

多个切片可以拥有相同的底层数组，因此每个切片所做的修改也都会反应到底层数组中
```go
arr := [4]int{1, 2, 3, 4}
sli1 := arr[:]
sli2 := arr[:]
sli3 := arr[:]

sli1[0] = 10
sli2[1] = 20
sli3[2] = 30
fmt.Println(arr) // [10 20 30 4]
```

slice操作符s[i:j] (其中0<=i>=cap(s))，会创建一个新的slice，引用序列s中i到j-1索引位置的所有元素，所以这个新的slice对元素的操作依旧会影响底层数组。

### 关于append
如果slice的容量足够，那么会定义一个新的slice，但是仍然引用原始底层数组，然后将新元素复制到新的位置，返回这个新的slice

如果slice的容量不够，会创建一个拥有足够容量的新的底层数组来存储新的元素，然后将参数的slice复制到这个数组，再将新元素追加到数组的后面，返回这个新的slice，此时这个新的slice和参数slice是引用两个不同的底层数组。

### 一个通过make()命令创建的缓冲区被分配了一块内存后。如何销毁缓冲区并收回内存？
buffer = nil

### 参数传递是怎么传递？
go参数传递采用的是值传递。

在go中，分为值类型和引用类型
1. 值类型除了基本数据类型（各种int和uint、各种float、string、complex、rune、byte、bool），还有数组、结构体、函数、方法、接口
2. 引用类型只有三种，分别是切片slice、字典map、管道channel。

那么函数传参的时候，会有如下几种情况：
1. 如果是值类型，会将实参的值拷贝作为函数的形参，因此在函数中对参数进行改变时，不会影响到原来的值。
2. 如果是引用类型在传参的时候，会将参数的内存地址拷贝一份传给函数，因此在函数中改变参数也会影响到实参。
3. 如果函数的参数是一个指针类型的话，那么指针变量传入的时候，会将参数的内存地址拷贝一份，传入函数，因此在函数中对参数的变化，也就影响到了实参了。

注意：以上规则也应用于赋值操作
```go
var arr = [3]int{1,2,3}
var sli = []int{1,2,3}

newArr := arr
newSli := sli

newArr[0] = 10
newSli[0] = 10

fmt.Println(arr[0]) // 1
fmt.Println(sli[0]) // 10
```

### golang中make和new的区别
make用于内建类型（只能用于创建map、slice 和channel）的内存分配。并且返回一个有初始值(非零)的T类型，而不是*T。

new用于各种类型的内存分配。new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：*new返回指针

### main和init函数的区别
1. init函数只能应用于所有的package，会按照顺序执行，main只能用于package main
2. 一个函数可以有任意多个init
3. 这两个函数都不能有参数和返回值

![main&initn](../images/main&init.png)

### 垃圾回收
在go1.5以前，go采用的是标记清除法，即从程序的根节点开始，给每个对象打上标记，将所有没有被标记的对象作为垃圾回收。但是这种方式有个问题，就是STW问题，即stop the world，算法在标记对象时需要暂停整个程序，否则其他线程的代码可能会改变对象，而且需要遍历对象树，当这个树很大时，相当于把程序卡死了。

因此到了1.5后，go采用了新的GC机制，三色标记法，它是一个并发的GC算法，能够和主程序并行的，这就可以避免程序的长时间暂停

具体可以参考几篇博客
1. [Go 垃圾回收](https://studygolang.com/articles/11904)
2. [golang 垃圾回收机制](https://blog.csdn.net/u010230794/article/details/78909780)

# linux和os
### 网络操作
netstat，用于显示网络状态，`netstat -an | grep 8080`

tcpdump，主要是截获通过本机网络接口的数据，用以分析，`tcpdump tcp port 9090 host 210.27.48.1`

ipcs，检查系统上共享内存的分配，`ipcs -a`

ipcrm，手动解除系统上共享内存的分配，
```
ipcrm [ -m SharedMemoryID ] [ -M SharedMemoryKey ] [ -q MessageID ] [ -Q MessageKey ] [ -s SemaphoreID ] [ -S SemaphoreKey ]

-m SharedMemory id 删除共享内存标识 SharedMemoryID。与 SharedMemoryID 有关联的共享内存段以及数据结构都会在最后一次拆离操作后删除。
-M SharedMemoryKey 删除用关键字 SharedMemoryKey 创建的共享内存标识。与其相关的共享内存段和数据结构段都将在最后一次拆离操作后删除。
-q MessageID 删除消息队列标识 MessageID 和与其相关的消息队列和数据结构。
-Q MessageKey 删除由关键字 MessageKey 创建的消息队列标识和与其相关的消息队列和数据结构。
-s SemaphoreID 删除信号量标识 SemaphoreID 和与其相关的信号量集及数据结构。
-S SemaphoreKey 删除由关键字 SemaphoreKey 创建的信号标识和与其相关的信号量集和数据结构。
```

telnet，用于登录远程主机，对远程主机进行管理，`telnet 192.168.33.10 9999`  

nc，和telnet一样，`nc 192.168.33.10 9999`

### 系统信息
df，查看硬盘信息，`df -lh`

查看cpu信息，`cat /proc/cpuinfo`

ps，查看进程信息，`ps aux | grep tomcat`

lsof，列出当前系统打开文件的工具，`lsof -i tcp:9999`
