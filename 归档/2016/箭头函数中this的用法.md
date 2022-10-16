> es6中有一种新的定义函数的形式，称之为arrow function，即箭头函数。其带来了很多便捷性，本文不打算阐述arrow function带来的好处，而是想来说说它的this。

学习javascript的人，一般来说会遇到两道坎，其中一道是原型链，另一道可能就是this问题了。js的函数不断变化的this经常让人无法摸不着头脑，很多时候我们会使用bind、call、apply来强制指定函数的this。

本文假设你已经掌握了js中的this问题了，如果你不懂，可以访问如下链接，了解下js中的this问题

[图解javascript this指向什么？](http://www.cnblogs.com/isaboy/archive/2015/10/29/javascript_this.html)

### 词法作用域
Arrow Function中的this机制和一般的函数是不一样的。  
本质来说Arrow Function并没有自己的this，它的this是派生而来的，根据“词法作用域”派生而来。

因此Arrow Function中的this是遵守“词法作用域”的

什么是词法作用域？

一般来说，作用域有两种常见的模型，一种叫做词法作用域（Lexical Scope），一种叫做动态作用域（Dynamic Scope）。而javascript采取的便是词法作用域

简单来说，所谓的词法作用域就是一个变量的作用在定义的时候就已经被定义好，当在本作用域中找不到变量，就会一直向父作用域中查找，直到找到为止。

说到这一点，我相信大家都明白了，不明白？上个例子
```javascript
function fn() {
  var a = 'hello'
  var b = 'javascript'
  
  function innerFn() {
    var b = 'world'
    console.log(`${a} ${b}`)
  }
  innerFn()
}
fn() // hello world
```
因为在innerFn中已经定义了b所以，因此在查找b时便不会去使用父作用域中的b了。

Arrow Function中的this便遵循了这个含义

### Arrow Function中的this
先来看一个案例
```javascript
function taskA() {
  this.name = 'hello'
  
  var fn = function() {
    console.log(this)
    console.log(this.name)
  }
  
  var arrow_fn = () => {
    console.log(this)
    console.log(this.name)
  }
  fn()
  arrow_fn()
}
taskA()
```
最终我们会发现，两个内部函数的this都是window，而且this.name都是hello。

好像没什么区别。其实两个函数的this的产生流程是不一样的。

fn的this是在运行时产生的，由于我们是直接调用fn()，所以其this就是指向window。如果将其调用改成
```javascript
function taskA() {
  this.name = 'hello'
  
  var fn = function() {
    console.log(this)
    console.log(this.name)
  }
  var obj = {
    name: 'haha',
    fn: fn
  }
  obj.fn()
}
taskA()
```
这时this就是obj对象，name是haha。这个符合我们对一般函数this的理解。

接下来看看Arrow Function中的this。它是怎么产生的呢，首先根据“词法作用域”，由于它本身没有this，于是便向上查找this，于是发现taskA是有this的，于是便直接继承了taskA的作用域。

那taskA的this又是什么？很简单，taskA是一个普通的函数，普通函数的this是在运行时决定的，由于我们是直接调用taskA的，即taskA()，所以其this便是window。

这下我们便明白了，arrow_fn中的this是window的原因了。我们稍微修改下案例
```javascript
function taskA() {
  var arrow_fn = () => {
    console.log(this)
    console.log(this.name)
  }
  arrow_fn()
}
var obj = {name: 'Jack'}
taskA.bind(obj)()
```
这时候，Arrow Function中的this便变成了obj对象了，name便是Jack。

可能有人会说，不是说Arrow Function中的this是定义的时候就决定了么，怎么现在又变成了运行的时候决定了呢。

Arrow Function中的this是定义的时候就决定的，这句话是对的。

该案例中，Arrow Function中，即arrow_fn的this便是taskA的this，在定义这个arrow_fn时候便决定了，于是又回到了上面说的，taskA是一个普通的函数，普通函数的this是在运行时决定的，而此时由于bind的原因，taskA的this已经变为obj，因此arrow_fnd的this便是obj了。

说到这里，相信大家应该已经明白了Arrow Function的this的含义和具体指向了。

所以我们才说Arrow Function的this是遵守“词法作用域”的。

### 其他案例
我们再来看看其他案例
```javascript
var obj = {
  field: 'hello',
  getField: () => {
    console.log(this.field)
  }
}
obj.getField() // undefined
```
这里最终会打出undefined，因为getField中的this就是window，而window是没有field这个属性的，所以就是undefined了。

所以我们一般不建议对象中定义函数的时候使用Arrow Function，毕竟this就会造成错误了。所以应该这么写
```javascript
var obj = {
  field: 'hello',
  getField(){
    console.log(this.field)
  }
}
```
这样this就是obj了。
