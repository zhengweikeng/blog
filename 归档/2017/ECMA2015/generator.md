> 纯粹想弄清楚generator

# 基本概念
```javascript
// 声明generator的方式
function * foo() {
  const res1 = yield 'hello'
  const res2 = yield 'world'
  const res3 = yield 'ending'
}

// 调用generator函数会返回一个遍历器对象，代表Generator函数的内部指针。函数并未开始执行
const g = foo()
// 每次调用next都会返回一个对象，value是yield的的返回值，done代表generator是否结束
// 第一次调用，函数内部开始运行，直到遇到第一个yield
g.next() // {value: 'hello', done: false}
g.next() // {value: 'world', done: false}
g.next() // {value: 'ending', done: true}
g.next() // {value: 'undefined', done: true}
g.next() // {value: 'undefined', done: true}
```

需要注意的是：
1. generator函数内部可以没有yield语句，此时函数便成为了一个暂缓执行的函数，直到调用next才会执行函数
1. yield只能出现在generator函数中
1. yield如果用在一个表达式中，必须放在圆括号里面
    *. console.log('hello' + yield) // 报错
    *. console.log('hello' + yield 123) // 报错
    *. console.log('hello' + (yield)) // 正常
    *. console.log('hello' + (yield 123)) // 正常
1. yield语句用作函数参数或赋值表达式的右边，可以不加括号。如：foo(yield 'a', yield 'b')

# 与Iterator接口的关系
> 任意一个对象的Symbol.iterator方法，等于该对象的遍历器生成函数，调用该函数会返回该对象的一个遍历器对象

```javascript
var myIter = {}
myIter[Symbol.iterator] = function* () {
  yield 1
  yield 2
  yield 3
}

// 此时myIter便可以被...运算符遍历了
[...myIter] // [1,2,3]
```

> Generator函数执行后，返回一个遍历器对象。该对象本身也具有Symbol.iterator属性，执行后返回自身
```javascript
function* gen() {}

var g = gen()

g[Symbol.iterator]() === g // true
```

# next方法
yield默认总是返回undefined，next方法可以传递参数，参数即为上一个yield的返回值。

注意：next方法的返回值是{value: '', done: true/false}；此处说的是yield的返回值

```javascript
function* foo(x) {
  var y = 2 * (yield (x + 1))
  var z = yield (y / 3)
  return x + y + z
}

var a = foo(5)
a.next() // {value: 6, done: false}
a.next() // {value: NaN, done: false}
a.next() // {value: NaN, done: true}

var b = foo(5)
b.next() // {value: 6, done: false}
b.next(12) // {value: 8, done: false}
b.next(13) // {value: 42, done: done}
```

# for...of循环
Generator函数执行后，返回一个遍历器对象。该对象本身也具有Symbol.iterator属性，所以generator可以使用for...of进行遍历，达到自动执行next()

```javascript
function* foo() {
  yield 1
  yield 2
  yield 3
  yield 4
  yield 5

  return 6
}

for (let v of foo()) {
  console.log(v)
}
// 1 2 3 4 5
```

需要注意的是执行return的时候，done已经是true，循环终止，所以返回值不在for...of之中

# yield* 语句
如果一个generator函数里，调用了其他generator函数，默认情况下是没有效果的，需要使用yield*来使其执行

```javascript
function* foo() {
  yield 'a'
  yield 'b'
}

function* bar() {
  yield 'x'
  foo()
  yield 'y'
}

for (let v of bar()) {
  console.log(v)
}
// x
// y

function* bar2() {
  yield 'x'
  yield* foo()
  yield 'y'
}

for (let v of bar2()) {
  console.log()
}
```

# generator的this
Generator函数总是返回一个遍历器，ES6规定这个遍历器是Generator函数的实例，也继承了Generator函数的prototype对象上的方法

```javascript
function* g() {}

g.prototype.hello = function() {
  return 'hi'
}

let obj = g()
obj instanceof g // true
obj.hello() // 'hi'

// 由于g总返回遍历器对象，而不是this对象，所有在g构造函数里使用this是不会生效的
function* g {
  this.a = 'test'
}
let obj = g()
obj.a // undefined
```

generator函数不能和new命令一起使用
```javascript
function* F() {
  yield this.x = 2
}
new F()
// TypeError: F is not a constructor
```

