众所周知，node.js有很多底层代码是调用c++库的，因此node.js必然提供了调用c++类库的方式。

[node-gyp](https://github.com/nodejs/node-gyp)便是提供了这种方式，它可以让你为node编写c++类库，经过它编译之后，便可以提供给node调用。

很明显，你的c++代码必须按照node-gyp的规范来写，但是很多时候我们是要直接调用其他第三方c++类库，这些类库不可能是一开始就设计给node用的。那应该怎么去调用这些类库呢？

这就是[node-ffi](https://github.com/node-ffi/node-ffi)这个工具所做的事情，它可以让你使用Node调用编译好的c或者c++类库。

来看看它怎么使用。

### Installation
首先需要安装[node-gyp](https://github.com/nodejs/node-gyp#installation)，推荐大家还是直接用linux，在mac上折腾了很久，总会出现各种问题，毕竟它需要安装xcode。

我是在mac里用vagrant开了个虚拟机，装的系统是Debian。

接下来就是`npm install ffi`

### 案例
先来一个简单的案例

有一个非常简单的c文件，hello.c
```c
#include <stdio.h>

void print(int num) {
  printf("%d", num);
}
```

接下来我们来编译这个文件
```shell
gcc -shared -fpic hello.c -o libhello.so
```

这样我们便要使用node-ffi来调用这个文件了
```javascript
var ffi = require('ffi')

var libhello = ffi.Library('./libhello', {
  'print': ['void', ['int']]
})

libhello.print(100)
```

### 剖析一下
fi.Library函数签名
```
ffi.Library(libraryFile, { functionSymbol: [ returnType, [ arg1Type, arg2Type, ... ], ... ]);
```
第一个参数是调用类库的路径。第二个参数是个对象，这个对象描述的就是该类库包含的方法。

这个函数会返回一个对象，该对象就包含你所要调用的类库的方法了，当然这些方法你必须先在上面的Library方法声明过。
