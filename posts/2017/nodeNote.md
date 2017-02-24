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

