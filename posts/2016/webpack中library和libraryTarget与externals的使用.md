> 在看webpack文档的[LIBRARY AND EXTERNALS](http://webpack.github.io/docs/library-and-externals.html)一节时，对output.libraryTarget和output.library还有externals总是无法理解，于是在segmentFault提问后结合实践总算多少理解它是怎么用的。

### externals
官网文档解释的很清楚，就是webpack可以不处理应用的某些依赖库，使用externals配置后，依旧可以在代码中通过CMD、AMD或者window/global全局的方式访问。

怎么理解呢？我们先通过官网说的那个jquery的案例来理解。

有时我们希望我们通过script引入的库，如用CDN的方式引入的jquery，我们在使用时，依旧用require的方式来使用，但是却不希望webpack将它又编译进文件中。

```html
<script src="http://code.jquery.com/jquery-1.12.0.min.js"></script>
```
jquery的使用如下
```javascript
// 我们不想这么用
// const $ = window.jQuery

// 而是这么用
const $ = require("jquery")
$("#content").html("<h1>hello world</h1>")
```
这时，我们便需要配置externals
```javascript
module.exports = {
  ...
  output: {
    ...
    libraryTarget: "umd"
  },
  externals: {
    jquery: "jQuery"
  },
  ...
}
```
我们可以看看编译后的文件
```javascript
({
  0: function(...) {
    var jQuery = require(1);
    /* ... */
  },
  1: function(...) {
    // 很明显这里是把window.jQuery赋值给了module.exports
    // 因此我们便可以使用require来引入了。
    module.exports = jQuery;
  },
  /* ... */
});
```

我们自己可以写个例子，当然这个例子没什么实际意义，只是为了演示如何使用而已。

假设我们自己有个工具库，tools.js，它并没有提供给我们UMD的那些功能，只是使用window或者global的方式把工具的对象tools暴露出来
```javascript
window.Tools = {
  add: function(num1, num2) {
    return num1 + num2
  }
}
```

接下来把它放在任何页面能够引用到的地方，例如CDN，然后用script的方式引入页面
```html
<script src="http://xxx/tools.min.js"></script>
```
一般来说我们可能会直接就这么用了
```javascript
const res = Tools.add(1,2)
```
但是既然我们是模块化开发，当然要杜绝一切全局变量了，我们要用require的方式。
```javascript
const tools = require('mathTools')
const res = tools.add(1,2)
```
这是我们再来配置一些externals即可
```javascript
module.exports = {
  ...
  output: {
    ...
    libraryTarget: "umd"
  },
  externals: {
    mathTools: "tools"
  },
  ...
}
```

### 关于externals的配置
首先是libraryTarget的配置，我们上面的例子都是umd  
当然它还有其他配置方式，具体就看官网[文档](http://webpack.github.io/docs/configuration.html#externals)吧，我们可以看到配置成umd便可以使用任何一种引入的方式了

### library和libraryTarget使用场景
接下来我们来说说library和libraryTarget的使用场景。

有些时候我们想要开发一个库，如lodash、underscore这些，这些库既可以用commonjs和amd的方式使用，也可以通过script标签的方式引入使用，目前很多库都是支持这几种使用方式的。

这时候我们就可以使用library和libraryTarget了，我们只需要用用es6的方式写代码，如何编译成umd就交给webpack了。

还是上面那个tools的例子
```javascript
exports default {
  add: function(num1, num2) {
    return num1 + num2
  }
}
```
接下来配置webpack
```javascript
module.exports = {
  entry: { myTools: "./src/tools.js" },
  output: {
    path: path.resolve(__dirname, "build"),
    filename: '[name].js',
    chunkFilename: "[name].min.js",
    libraryTarget: "umd",
    library: "tools"
  }
}
```
library指定的就是你使用require时的模块名，这里便是require("tools")
通过配置不同的libraryTarget会生成不同umd的代码，例如可以只是commonjs标准的，也可以是指amd标准的，也可以只是通过script标签引入的。

开发测试之后，便可以发布到npm给别人使用了。
