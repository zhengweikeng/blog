### webpack对样式的处理
我们可以在js中引入样式文件
```javascript
require('myStyle.css')
```
这时我们便需要引入相应的webpack loader来帮助我们解析这段代码。

一般来说需要引入css-loader和style-loader，其中css-loader用于解析，而style-loader则将解析后的样式嵌入js代码。
```javascript
// webpack配置如下
{
  module: {
    loaders: [
      { test: /\.$/, loader: "style-loader!css-loader" }
    ]
  }
}
```
可以发现，webpack的loader的配置是从右往左的，从上面代码看的话，就是先使用css-loader之后使用style-loader。

同理，如果你使用less来写样式的话，则需要先用less-loader来编译样式文件为css文件，再继续使用css-loader与style-loader。
```
{
  module: {
    loaders: [
      { test: /\.$/, loader: "style-loader!css-loader!less-loader" }
    ]
  }
}
```

我们知道，webpack配置loader时是可以不写loader的后缀明-loader，因此css-loader可以写为css。

### 将样式抽取出来为独立的文件
将require引入的样式嵌入js文件中，有好处也有坏处。好处是减少了请求数，坏处也很明显，就是当你的样式文件很大时，造成编译的js文件也很大。

我们可以使用插件的方式，将样式抽取成独立的文件。使用的插件就是[extract-text-webpack-plugin](https://github.com/webpack/extract-text-webpack-plugin)

基本用法如下
```javascript
var ExtractTextPlugin = require("extract-text-webpack-plugin");
module.exports = {
  module: {
    loaders: [
      { test: /\.css$/, loader: ExtractTextPlugin.extract("style-loader", "css-loader") }
    ]
  },
  plugins: [
    new ExtractTextPlugin("styles.css")
  ]
}
```

根据插件在github上的解释，ExtractTextPlugin.extract可以有三个参数。  
第一个参数是可选参数，传入一个loader，当css样式没有被抽取的时候可以使用该loader。  
第二个参数则是用于编译解析的css文件loader，很明显这个是必须传入的，就像上述例子的css-loader。  
第三个参数是一些额外的备选项，貌似目前只有传入publicPath，用于当前loader的路径。

那什么时候需要传入第一个参数呢，那就得明白什么时候样式不会被抽取出来。  
了解过code splittiog的同学便会知道，我们有些代码在加载页面的时候不会被使用时，使用code splitting，可以实现将这部分不会使用的代码分离出去，独立成一个单独的文件，实现按需加载。

那么如果在这些分离出去的代码中如果有使用require引入样式文件，那么使用ExtractTextPlugin这部分样式代码是不会被抽取出来的。  
这部分不会抽取出来的代码，可以使用loader做一些处理，这就是ExtractTextPlugin.extract第一个参数的作用。  

根据上面的案例，ExtractTextPlugin需要配合plugin使用。

```javascript
new ExtractTextPlugin([id: string], filename: string, [options])
```
1. 该插件实例的唯一标志，一般是不会传的，其自己会生成。
2. 文件名。可以是[name]、[id]、[contenthash]  
  [name]：将会和entry中的chunk的名字一致  
  [id]：将会和entry中的chunk的id一致  
  [contenthash]：根据内容生成hash值
3. options  
  allchunk： 是否将所有额外的chunk都压缩成一个文件  
  disable：禁止使用插件
  
这里的参数filename里如何理解呢？上述案例指定了一个固定的名字，因此便会生成一个styles.css文件。

那么像[name]、[id]这些如何理解。这个在你有多个entry的时候，便需要使用这种方式来命名。
```javascript
var ExtractTextPlugin = require("extract-text-webpack-plugin");
module.exports = {
  entry: {
    "script": "./src/entry.js",
    "bundle": "./src/entry2.js",
  },
  ...
  module: {
    loaders: [
      { test: /\.css$/, loader: ExtractTextPlugin.extract("style-loader", "css-loader") }
    ]
  },
  plugins: [
    new ExtractTextPlugin("[name].css")
  ]
}
```
这时候便会生成两个css文件，一个是script.css，另一个便是bundle.css。那些[id]、[contenthash]也是一个道理。  
只要明白，在你有多个entry是，一定要使用这种方式来命名css文件。

最后还有那个allchunks又是什么呢？很简单，还记得前面提到的code splitting么？将该参数配置为true，那么所有分离文件的样式也会全部压缩到一个文件上。
```
plugins: [
  new ExtractTextPlugin("[name].css", {allChunks: true})
]
```

### 补充一下
根据webpack官网中关于[stylesheet](http://webpack.github.io/docs/stylesheets.html)的说法，建议是不要将allChunks设为true，即只是将样式嵌入到分离文件中。  
这个可能还是需要具体问题具体分析了。

另外，我们可以使用CommonsChunkPlugin插件，将共有的样式抽取成独立样式文件，common.css，可以减少样式文件的大小。
