# == 和 ===
== 允许在相等比较中进行强制类型转换，而 === 不允许

换句话说，使用==的时候，会进行一些强制类型转换，然后进行比较；而===是不会做这层转换的

首先有几点要注意的：
1. NaN不等于NaN
1. +0等于-0

## 比较的两者都是对象
这里的对象还包括函数和数组，无论是使用==还是===，当两个对象指向同一个值时则视为相等

注意不是对象里面的值相等，是对象的引用相等
```javascript
var obj1 = {
  a: 1
}
var obj2 = obj1
var obj3 = {
  a: 1
}

console.log(obj1 == obj2) // true
console.log(obj1 === obj2) // true

console.log(obj1 == obj3) // false
console.log(obj1 === obj3) // false
```

## 使用==比较字符串和数字
当比较`x == y`根据ES5的定义：  
1. 如果 Type(x) 是数字，Type(y) 是字符串，则返回 x == ToNumber(y) 的结果
1. 如果 Type(x) 是字符串，Type(y) 是数字，则返回 ToNumber(x) == y 的结果

```javascript
var a = 42
var b = '42'

a === b // 没有强制转换，所以为false 

a == b // ---> 42 == ToNumber(42) ---> 42 == 42 ---> true
```

## 使用==比较其他类型和布尔类型
使用==来比较true或者false是最容易出错的地方

ES5中是这样定义的:
1. 如果 Type(x) 是布尔类型，则返回 ToNumber(x) == y 的结果
1. 如果 Type(y) 是布尔类型，则返回 x == ToNumber(y) 的结果

```javascript
var a = '42'
var b = true

a == b // false
```

很多人以为`a=42`是真值，就会==true，其实是错误的。套用上面的转化规则：
1. Type(b)是布尔类型，则 '42' == ToNumber(true)
1. 转化后为 '42' == 1，测试便转化为比较字符串和数字，在转换一遍 ToNumber('42') == 1
1. 42 == 1很明显示不相等，所以返回false

```javascript
var a = true
var b = '42'

a == b // false
```
该例子同理， `ToNumber(true) == '42'`，之后`1 == '42'`，再`1 == 42`，不相等

再看一个
```javascript
var a = '42'
var b = false

a == b // false
```
依旧不相等， `'42' == ToNumber(false)`，之后`'42' == 0`，再`42 == 0`，明显不相等

由此可看出，无论是和true比较还是false比较都不相等，这是个很容易忽视的问题

应该怎么用呢？
```javascript
var a = '42'

// 不要这样用，因为条件不成立
if (a == true) {

}

// 也不要这样用，因为条件不成立
if (a === true) {

}

// 这样是没问题的
if (a) {

}

// 这样更好
if (!!a) {

}

// 这样也很好
if (Boolean(a))
```

## null和undefined的比较
这两者使用==比较时，相应的转化规则为：
1. 如果 x 为 null，y 为 undefined，则结果为 true
1. 如果 x 为 undefined，y 为 null，则结果为 true

```javascript
var a = null
var b
a == b     // true
a == null  // true
b == null  // true

a == false // false
b == false // false
a == "" // false
b == "" // false
a == 0 // false
b == 0 // false
```

由此可见，如果你是想等于null或者undefined的，使用==是安全可靠的，其他值如""、false、0都不会使你的判断成立

## 对象和非对象之间的比较
当使用==比较对象（对象/函数/数组）和标量基本类型（字符串/数字/布尔值）时，会按照下面的规则进行转化：
1. 如果 Type(x) 是字符串或数字，Type(y) 是对象，则返回 x == ToPromitive(y) 的结果
1. 如果 Type(x) 是对象，Type(y) 是字符串或数字，则返回 ToPromitive(x) == y 的结果

至于布尔值，会转化为数字，所以依旧套用上述规则

*ToPrimitive是什么？*  
抽象操作ToPrimitive会首先(通过内部操作 DefaultValue)检查该值是否有`valueOf()`方法。  
如果有并且返回基本类型值，就使用该值进行强制类型转换。如果没有就使用`toString()`的返回值(如果存在)来进行强制类型转换  
如果 valueOf() 和 toString() 均不返回基本类型值，会产生`TypeError`错误。例如`Object.create(null)`

```javascript
var a = 42
var b = [42]

a == b // true
```
由于 ToPromitive(b)会转化为'42'，因此就转化为比较42 == '42'的比较，因此相等

另外使用==也会对对象进行拆封，将其转化为基本数据类型
```javascript
var a = 'abc'
var b = Object(a) // 和 new String(a) 一样

a === b // false
a == b // true 此处将b进行拆封，转化为'abc'，故与a相等

var a = null
var b = Object(a) // 和Object()一样
a == b // false

var c = undefined 
var d = Object(c) // 和Object()一样
c == d // false

var e = NaN
var f = Object(e) // 和new Number(e)一样
e == f // false
```
因为没有对应的封装对象，所以 null 和 undefined 不能够被封装(boxed)，Object(null) 和 Object() 均返回一个常规对象  
NaN能够被封装为数字封装对象，但拆封之后NaN == NaN返回false，因为NaN不等于NaN


# 闭包
代码解释
```javascript
function foo() { 
  var a = 2;
  function bar() { 
    console.log(a)
  }
  return bar
}
var baz = foo()
baz() // 2 
```

foo执行结束后，正常来说变量就要被垃圾回收给回收掉，但是此处却没有。这便是闭包的功力

由于bar函数声明的位置的原因，导致foo内存作用域并没有被回收，而bar对这个作用域的引用就是闭包

## 一道经典的闭包面试题
```javascript
for (var i = 1; i <= 5; i++) {
  setTimeout(function () {
    console.log(i)
  }, i * 1000)
}
```
运行后，结果是每隔1s会打印出6，总共打印出5次

每隔1s，总共打印6次很好理解，因为我们循环从i=1到i=5，setTimeout总共运行了5次  
每次传入的i的参数（setTimeout的第二个参数）分别是1000、2000、...、5000

那为什么每次打印出来都是6？  
其实即使是setTimeout(..., 0)它也是打印出5次6。主要原因是setTimeout中指定的回调函数，都是在循环结束时运行的。  
循环结束时i的值是6，而由于闭包的原因，回调函数依然保持了对外部变量的引用，所以获取到的i便是6，也即打印出了6。  
换句话说，此处相当于创建了五个回调函数，但是完成的作用域只有一个，也即只有一个i，所以每个回调函数取到的i都是同一个对象。

那有什么方式可以让其每次都打印循环中i的值呢？  
其实我们只需要每次迭代中都产生一个闭包作用域即可，让每个闭包作用域中的i都是不同的。
```javascript
for (var i = 1; i <= 5; i++) {
  (function(j) {
    setTimeout(function () {
      console.log(j)
    }, j * 1000)
  })(i)
}
```
此时回调函数调用的j变量，来自外层作用域。该外层作用域中的j，是每次循环传入进来的。这也使得每次循环都会产生一个新的作用域。  
而每次产生的作用域不会被释放掉，也是因为内存回调函数造成的闭包的原因。

既然我们是想让每次循环产生的作用域都是独特的，那便还可以这样使用
```javascript
for (let i = 1; i <= 5; i++) {
  setTimeout(function () {
    console.log(i)
  }, i * 1000)
}
```
与最终的代码的变化只是`var`变成了`let`，具体let的用法可以参考 [let和const命令](http://es6.ruanyifeng.com/#docs/let)

## 闭包的作用
通过使用闭包，我们可以访问局部变量，同时让这些变量不被释放。

我们可以通过闭包实现数据私有化，即外部无法直接使用局部变量
```javascript
function f1(){
  var n=999
　nAdd= function(){
    n+=1
  }
　function f2(){
    alert(n)
　}
  return f2
}
var result=f1()
result() // 999
nAdd()
result() // 1000
```

外部无法直接获取变量，进而修改变量，起到了保护变量的作用。只能通过调用内存函数修改和获取变量

## 闭包和内存泄漏
很多人认为闭包中的变量没有被释放，便夸大其词去证明使用闭包不好，会造成内存泄漏。

我个人理解是，通过闭包，让一些局部变量不用被反复创建，这种方式和你定义为全局变量有相似之处。

但是因为闭包中引用的外层变量过多，然后又没有释放变量，最终造成内存泄漏的问题，这只能说明一个问题，你代码写得不好罢了。

参考几篇文章：
1. [关于js闭包是否真的会造成内存泄漏？](https://www.zhihu.com/question/31078912)
