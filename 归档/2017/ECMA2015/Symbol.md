# 基本概念
ES6引入了一种新的原始数据类型Symbol，表示独一无二的值。它是JavaScript语言的第七种数据类型，前六种是：Undefined、Null、布尔值（Boolean）、字符串（String）、数值（Number）、对象（Object）

```javascript
var s1 = Symbol()
var s2 = Symbol()
s1 === s2 // false

var s3 = Symbol('test')
var s4 = Symbol('test')
s3 === s4 // false
```
由此可见，用它来定义一些属性或者一些独一无二的值非常好
```javascript
var host = Symbol('host')

var config = {
  [host]: 'localhost'
}

console.log(config[host])
```

注意几点：
1. Symbol前不能用new，即不能使用`new Symbol()`
1. Symbol值不能与其他类型的值进行运算，即不能`var s = Symbol(); var res = s + 1`
1. Symbol对象可以转化为字符串。`var s = Symbol('test'); String(s)`
1. Symbol对象可以转化为布尔值。`var s = Symbol('test'); Boolean(s) // true`
1. Symbol对象不能转化为数字，即不能`var s = Symbol(); Number(s)`

# 使用案例
## 作为配置
```javascript
function doSomeThing(type) {
  swtich (type) {
    case 'test1':
      // doSomeThing
      break
    case 'test2':
      // doSomeThing
      break
  }
}

doSomeThing('test1')
```
这种方式就一旦要修改test1的值，就得改两处，甚至多处

较好的实现方式
```javascript
var types = {
  test1: Symbol('test1'),
  test2: Symbol('test2'),
}

function doSomeThing(type) {
  swtich (type) {
    case types.test1:
      // doSomeThing
      break
    case types.test2:
      // doSomeThing
      break
  }
}

doSomeThing(types.test1)
```

## 实现非私有，只用于内部的方法
Symbol 作为属性名，该属性不会出现在for...in、for...of循环中，也不会被Object.keys()、Object.getOwnPropertyNames()、JSON.stringify()返回。  
但是，它也不是私有属性，有一个Object.getOwnPropertySymbols方法，可以获取指定对象的所有 Symbol 属性名。

```javascript
var size = Symbol('size');

class Collection {
  constructor() {
    this[size] = 0;
  }

  add(item) {
    this[this[size]] = item;
    this[size]++;
  }

  static sizeOf(instance) {
    return instance[size];
  }
}

var x = new Collection();
Collection.sizeOf(x) // 0

x.add('foo');
Collection.sizeOf(x) // 1

Object.keys(x) // ['0']
Object.getOwnPropertyNames(x) // ['0']
Object.getOwnPropertySymbols(x) // [Symbol(size)]
```

目前class没有提供仅供内部类使用的私有属性，上述例子对于size属性，依旧被外部使用`x[size]`访问到。  
但是使用Symbol的key不会被一般的遍历方法访问到，某种程度上还是起到了保护作用吧。  
如果真的想实现私有变量，估计只能使用闭包了。