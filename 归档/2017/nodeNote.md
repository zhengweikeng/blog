> 一些关于node进阶的笔记

# 模块
## require
require是一个函数，同时下面也挂着很多属性

```javascript
{
  cache: [],
  extensions: {
    '.js',
    '.json',
    '.node',
  },
  main: {},
  name: 'require',
  prototype: Function,
  resolve: Function
}
```

1. cache是缓存每个require进来模块  
1. extensions很明显是require支持的文件后缀名
1. main是当前模块本身。
1. resolve 即为require函数

其中require.main可以用来检测当前文件是直接使用node运行的还是require进来的

```javascript
// test.js
console.log(require.main === module)

// app.js
var test = require('./test')
```

运行
```shell
$ node test.js
$ true

$ node app.js
$ false
```

## 热加载
我们知道node在require一个文件的时候会优先从v8的缓存里面获取，即使你修改了文件，也无法获取到修改后的文件

这个缓存信息便存放在require.cache中，只要它存在，每次require都会拿v8中的旧代码

修改cache的确能拿到新的代码，但是并不能保证旧的引用还跑着旧的代码的问题。

## 循环引用
```javascript
// a.js
var b = require('./b.js')
b.f2()
module.exports = {
  f1: () => console.log(`seed: f1`)
}

// b.js
var a = require('./a.js')
a.f1()
module.exports = {
  f2: () => console.log(`seed: f2`)
}
```

无论是执行哪个文件，都会f1或者f2函数没找到的异常。

猜想是根本没有拿到暴露出来的对象

我们修改下
```javascript
var b = require('./b.js')
b.f2()
module.exports = {
  f1: () => console.log(`seed: f1`)
}

// b.js
var a = require('./a.js')
console.log(a)
module.exports = {
  f2: () => console.log(`seed: f2`)
}
```

此时我们运行a.js
```shell
$ node a.js
$ {}
$ seed: f1
```
打印出了在b.js文件中，a是个空对象{}，而b正常暴露出来了。

这下我们明白了，a.js在require('./b.js')后，由于a尚未暴露出对象，所以此时在b中用require引进来的的a.js便是个空对象。  
然后b.js成功暴露出了对象，在a.js文件中便可以成功调用b中暴露出来的方法

这里又可以牵扯到一个知识点，为什么b.js没有执行module.exports时暴露出来的是个空对象，而不是null。那是因为Node原码中是这样写的：

```javascript
function require(...) {
  var module = { exports: {} };
  ((module, exports) => {
    // Your module code here. In this example, define a function.
    function some_func() {};
    exports = some_func;
    // At this point, exports is no longer a shortcut to module.exports, and
    // this module will still export an empty default object.
    module.exports = some_func;
    // At this point, the module will now export some_func, instead of the
    // default object.
  })(module, module.exports);
  return module.exports;
}
```

由此可见模块在被require时，如果你自己没有暴露任何对象出去，就是个空对象，也即暴露的是module下的exports的值

## vm
node可以使用vm来创建一个js的上下文，避免上下文被污染，用法和eval很像。  
只是vm是通过创建新的上下文沙盒（sandbox）来避免这种污染的。

看过node官网的案例来和eval做个对比
```javascript
const vm = require('vm')
var localVar = 'initial value'

const vmResult = vm.runInThisContext('localVar = "vm"')
console.log('vmResult:', vmResult)
console.log('localVar:', localVar)
// vmResult: 'vm', localVar: 'initial value'

const evalResult = eval('localVar = "eval"')
console.log('vmResult:', vmResult)
console.log('localVar:', localVar)
// evalResult: 'eval', localVar: 'eval'
```

由此可见通过vm执行的代码，不会污染当前的上下文，安全的起到了隔离的作用。

那vm执行的代码，和require的方式有什么区别呢，sf有个人问了这个 [问题](http://stackoverflow.com/questions/9867069/node-js-vm-runinnewcontext-vs-require-and-eval)

也即其实require也是在执行vm的过程，但是它也额外做了一些其他的事情，如cache

# 事件
## emit是同步还是异步的
同步的
详见 [官网解释](https://nodejs.org/dist/latest-v6.x/docs/api/events.html#events_asynchronous_vs_synchronous)

> This is important to ensure the proper sequencing of events and to avoid race conditions or logic errors

> 这对于确保事件的正确排序以及避免竞争条件或逻辑错误很重要

调用emit的时候，只会调用那些注册了相应事件名的监听函数。而且由于emit是同步的，所以之后注册的监听是不会被调用的

例子1：
```javascript
const EventEmitter = require('events');

let emitter = new EventEmitter();

emitter.on('myEvent', () => {
  console.log('hi 1');
});

emitter.on('myEvent', () => {
  console.log('hi 2');
});

emitter.emit('myEvent');
// h1 1
// hi 2
```
此时调用emit，之前注册的两个监听都会被调用

例子2：
```javascript
const EventEmitter = require('events');

let emitter = new EventEmitter();

emitter.on('myEvent', () => {
  console.log('hi');
  emitter.emit('myEvent');
});

emitter.emit('myEvent');
```
程序进入死循环，最后堆栈溢出，程序退出

例子3：
```javascript
const EventEmitter = require('events');

let emitter = new EventEmitter();

emitter.on('myEvent', function sth () {
  emitter.on('myEvent', sth);
  console.log('hi');
});

emitter.emit('myEvent');
// hi
```
此时只打印出了一个hi，sth函数中的监听并不会被触发，因为emit只是触发了在它之前注册的监听。  
像sth函数中的监听是在触发时才注册的，监听函数是不会被调用的。

这也充分说了emit是同步的作用，如果是异步，sth中的监听很有可能会被调用，这种调用顺序和是否调用不明确的结果，会引发很多异常。

# Timer
## setTimeout(fn, 0) vs setImmediat and process.nextTick
这几个方法都是用来创建异步

setTimeout和setImmediat都是事件循环的一部分  
process.nextTick不是事件循环的一部分，nextTick的队列会在当前操作结束后得到处理，也即会运行nextTick中的事件。它也是在事件循环（event loop）被处理之前执行的。

根据官方解释，如果使用nextTick会出现一个不好的现象：  
如果你一直在循环的执行nextTick，会使得I/O一直得不到处理，因为nextTick某种程度上阻止了事件循环的处理
```javascript
// 没完没了的执行foo，主线程不会去处理事件队列
process.nextTick(function foo() {
  process.nextTick(foo)
})
```

setTimeout和setImmediat应该优先使用哪个？
应该优先使用setImmediat，因为如果在I/O的循环内部（应该说是在递归调用的时候），setImmediat会先执行
```javascript
var fs = require('fs')

fs.readFile(__filename, () => {
  setTimeout(() => {
    console.log('timeout')
  }, 0)
  setImmediate(() => {
    console.log('immediate')
  })
})
```
```shell
$ node timeout_vs_immediate.js
immediate
timeout

$ node timeout_vs_immediate.js
immediate
timeout
```

而process.nextTick会在时间循环的时候执行，性能会优于setImmediat，因为不用检查"任务队列"

官方推荐使用setImmediat，因为会使得代码兼容性更好

详见这几篇文章
1. [setImmediate vs nextTick vs setTimeout(fn, 0)](https://cnodejs.org/topic/5556efce7cabb7b45ee6bcac)
1. [The Node.js Event Loop, Timers, and ](https://nodejs.org/en/docs/guides/event-loop-timers-and-nexttick/)
1. [JavaScript 运行机制详解：再谈Event Loop](http://www.ruanyifeng.com/blog/2014/10/event-loop.html)
1. [理解 Node.js 里的 process.nextTick()](https://www.oschina.net/translate/understanding-process-next-tick)

# 进程
## console.log是异步还是同步的
首先console.log并不是标准的javascript正式的一部分，而是由各个宿主环境添加到javascript中的。  

从浏览器的方面来说，每个浏览器都有自己的实现方式。有些浏览器可能并不会把它同步输出，因为对于浏览器来说I/O是个耗时的操作，所以被浏览器放在了后台执行。
```javascript
var a = {
  index: 1
}

console.log(a)

a.index++
```
这段代码在浏览器大部分都能按照我们的意愿被显示执行，打印{a:1}。但有时可能会是{a:2}。

在Node中，由于console.log是采用了process.stdout标准输出，而process.stdout是否异步取决于操作系统，参照node文档中关于process.stdout的解释
1. 在文件中，在windows和posix系统是同步的
2. 在terminal，如果是windows是异步的，如果是posix系统是同步的
3. 在管道，如socket中，如果是window系统下是同步的，在posix系统下是异步的

## 创建子进程
child_process提供了几个方法用于创建子进程
1. spawn(command[, args][, options])，启动一个子进程来执行命令
1. exec(command[, options][, callback]) 同上，不过它提供了一个回调来获知子进程的状况
1. execFile(file[, args][, options][, callback]) 启动一个子进程来执行可执行文件（不一定要是Node文件）
1. fork(modulePath[, args][, options]) 加强版的 spawn(), 返回值是 ChildProcess 对象可以与子进程交互

其中exec和execFile都可以设置超时

领完spawn、exec、execFile都有自己的同步方法spawnAsync、execAsync和execFileAsync

```javascript
var cp = require('child_process')
cp.spawn('node', ['worker.js'])
cp.exec('node worker.js', function (err, stdout, stderr) {
  // ...
})
cp.execFile('worker.js', function (err, stdout, stderr) {
  // ...
})
cp.fork('./worker.js')
```

那么通过fork创建的进程和传统unix机器fork进程是一样的么？  
unix机器fork一个进程的时候，它是基于父进程创建的，会继承一些父进程的属性。  
而node去fork一个进程的时候，并不会去继承父进程的属性

## 进程间通信（IPC）
当父子进程建立IPC通道后，便可以进行通信

```javascript
// parent.js
const cp = require('child_process')
const n = cp.fork('./sub.js')

n.on('message', (msg) => {
  console.log('PARENT got message:', m)
})

n.send({hello: 'world'})
```

```javascript
// sub.js
process.on('message', (msg) => {
  console.log('CHILD got message:', m)
})

process.send({foo: 'bar'})
```

node是怎么实现进程间的通信的？  
```
node通过管道的方式实现进程间通信。

当父进程准备创建子进程时，父进程会先创建IPC通道并且监听它，然后才真正去创建子进程。
同时会通过环境变量（NODE_CHANNEL_FD）的方式告诉子进程这个IPC通道的文件描述符，子进程通过这个文件描述符去连接IPC通道。


【父进程】  --------生成------> 【子进程】
    |                             |
    |                             |
  监听/接收                       连接
    |                             |
    |---------> 【IPC】 <----------|

而这个通道在windows系统下面，使用的是命名管道，而在*nix系统采用的是Unix Domain Socket技术实现。
这种socket与网络socket很详细，都是双向通信，只是不用经过网络层，非常高效。
```

> 注意的是，只有子进程是node进程，才可以通过上述方式进行进程间通信。除非其他类型的子进程也去连接这个通道。

## child.kill 和 child.send的差别
kill是通过发送信号的方式给子进程，如果没有指定信号，则会默认是SIGTERM  
而send是通过IPC通道发送的的。

```javascript
const spawn = require('child_process').spawn
const grep = spawn('grep', ['ssh'])

grep.on('close', (code, signal) => {
  console.log(
    `child process terminated due to receipt of signal ${signal}`)
})

// Send SIGHUP to process
grep.kill('SIGHUP')
```

注意，当使用kill方法时并不是去杀死子进程，只是去发送信号给子进程时。子进程收到信号时需要自己去做相应的处理，如果退出进程

子进程死亡不会影响父进程, 不过 node 中父进程会收到子进程死亡的信号. 反之父进程死亡, 一般情况下子进程也会跟着死亡, 如果子进程需要死亡却没有随之终止而继续存在的状态, 被称作孤儿进程. 另外, 子进程死亡之后资源没有回收的情况被称作僵死进程.

## 如何实现一个守护进程
使用linux的都知道守护进程（daemon），简单点说就是在后台默默运行的进程。

那node既然可以创建进程，那么是不是也可以创建守护进程呢，答案是可以的，而且也很简单。具体可以参看这篇文档：[Nodejs编写守护进程](https://cnodejs.org/topic/57adfadf476898b472247eac)

重点就是使用spawn创建子进程的时候传入的一个备选参数detached:true。
```javascript
child_process.spawn('node', 'app.js', {
  detached: true
})
```
设置detached为true后，可以使得即使父进程退出了，子进程依旧可以运行。对于非windows系统而言，子进程会成为成为会话首进程和组长进程

值得注意的是，虽然子进程已经脱离父进程了，但是父进程依旧会不断监听等待子进程退出。如果想完全将子进程“抛弃”，完全不接受detached的子进程，可以调用unref方法。

```javascript
const child = child_process.spawn('node', 'app.js', {
  detached: true
})
child.unref()
```

# 内存控制
v8对内存做了限制，在64位机器下，可用内存约为1.4GB，在32位机器下约为0.7GB。

因此即使你的机器有再大的内存，也无法全部得到使用，所以编写程序时都应该注意到这一点，避免内存泄漏造成程序奔溃。

## v8对象分配
可以通过`process.memoryUsage()`来查看内存信息

```javascript
process.memoryUsage()
{
  rss: 25440256, // 进程常驻内存大小
  heapTotal: 10522624, // 申请到的堆内存
  heapUsed: 5586008 // 当前堆的使用情况
}
```

也可以通过os.totalmen和freemem()来查看内存使用情况
```javascript
os.totalmem()
// 8589934592
os.freemem()
// 1967972352
```

当我们在代码中声明变量并赋值时，所使用的对象就会分配在堆中。  
如果申请的堆空闲内存不够分配新的对象，将继续申请堆内存

v8将内存分为新生代和老生代
1. 新生代中的对象为存活时间较短对象
1. 老生代中的对象为存活时间较长或者常驻内存的对象

### 新生代内存回收算法
新生代中的对象主要通过Scavenge算法进行垃圾回收。该算法将堆内存一分为二，每一部分空间称为semispace  
两个semispace空间，一个处于使用中（from空间），一个处于闲置状态（to空间）  
每次分配对象，先是在from空间中分配  

垃圾回收时，会检查from空间的存活对象，并被复制到to空间，而非存活对象占用的空间将会被释放  
这时from空间和to空间的角色会进行调换
当一个对象经过多次复制之后依然存活，将会被认为是声明周期较长的对象，会被移到老生代中

新生代垃圾回收主要流程如下：  
1. from空间复制到to空间时会检查该对象是否经历过Scavenge回收，如果经历过则从from空间移到老生代，否则移入to空间
1. 如果当要移入to空间时，to空间已经使用了超过25%，也会将该对象移入老生代。主要是因为之后to空间会变为from空间，如果from空间占比过高，会影响后续的内存分配

### 老生代内存回收算法
采用了Mark-Sweep和Mark-Compact相结合的方式进行垃圾回收

Mark-Sweep算法会做对象的标记和清除，先将对象进行标记活着的对象，在之后的清除阶段，只清除没有标记的对象

Mark-Sweep算法容易造成内存空间不连续，如果要分配一个大对象，会导致提前出发垃圾回收，而这次回收是不必要的

为了解决这种情况，引入了Mark-Compact算法，它会在对象被标记为死亡后，将活着的对象往一端移动，完成后，直接清理掉边界外的内存

## 查看垃圾回收日志
启动时增加--trace_gc参数，可以查看垃圾回收的日志
```javascript
// test.js
var a = []
for (var i = 0; i < 1000000; i++) {
  a.push(new Array(100))
}

node --trace_gc test.js > gc.log
```

## 堆外内存
通过process.memoryUsage()可以看出，每次堆中的内存总用量总是小于进程的常驻内存用量，即heapTotal < rss。  
这意味着node中的内存使用，还有部分是不经过v8分配的，这些不经过v8分配的内存称之为堆外内存

```javascript
let arr = []
while(true)
  arr.push(new Buffer(1000))
```

我们会发现内存并不会被撑爆，可以尝试去打印内存使用情况，会发现heapTotal和heapUsed的变化极小，rss的变化很大，而且值已经远远超过v8的限制值了。

没有爆掉的原因是Buffer对象的分配不同于其他对象，它是不经过v8的内存分配机制的，所以也不会有内存大小的限制
