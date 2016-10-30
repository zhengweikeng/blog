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

目前该插件2.0版本的api有一些变化，这里还是用了1.0的api

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

举个栗子

```javascript
// entry.js

require('./main.css')
let hello = null

// 代码分离
require.ensure([], function(require) {
  hello = require('./hello').default
})

$("#hello").click(() => {
  hello()
})

```

在hello.js中我们引用一个css文件

```javascript
// hello.js

require('./normal.css')
export default () => {
  console.log('hello~~')
}
```

假设我们两个css文件内容如下：

```css
// main.css
body {
  color: red;
}

// normal.css
h1 {
  color: blue
}
```

此时我们在webpack配置loader时这么写

```javascript
...
module: {
    loaders: [
      { 
        test: /\.js$/, 
        loader: 'babel', 
        query: {
          presets: ['es2015']
        }
      }, { 
        test: /\.css$/, 
        loader: ExtractTextPlugin.extract("style-loader", ["css-loader"]) 
      }
    ]
  }，
  plugins: [
    // allChunks 默认是false
    new ExtractTextPlugin("[name].styles.[id].[contenthash].css", {allChunks: false}),
  ]
...
```

注意我们在 ExtractTextPlugin.extract第一个参数传入了style-loader，这个参数便是给像被代码分离后的chunk用的，表明这些无法被抽取的chunk的处理方式，既使用style-loader，将其嵌入到页面的style标签中。

这时到页面上看，就会发现h1标签的文字是**蓝色**的（normal.css定义的），而不是红色的。

这时我们尝试把ExtractTextPlugin.extract的第一个参数删除掉，会发现，这时h1标签的文字变为**红色**（main.css定义的）了，normal.css的样式也没有被嵌入到页面中。

这便是第一个参数的作用

如果我们也想要将normal.css的样式代码抽取到文件中，我们可以将allChunks设置为true，即

```javascript
plugins: [
  // allChunks 默认是false
  new ExtractTextPlugin("[name].styles.[id].[contenthash].css", {allChunks: true})
]
```
---

### ExtractTextPlugin配置项
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

### postcss
以前我们写样式时，有些样式不同浏览器需要加不同的前缀，如-webkit-。现在有了构建工具，我们便不需要再去关注这些前缀了，构建工具会自动帮我们加上这些前缀。

对于webpack我们自然想到需要使用loader或者plugin来帮助我们做这些事情，查了下发现autoprefixer-loader已经废弃不再维护了，推荐使用[posscss](https://github.com/postcss/postcss)

postcss是用于在js中转换css样式的js插件，需要搭配其他插件一起使用，这点和babel6一样，本身只是个转换器，并不提供代码解析功能。

这里我们需要[autoprefixer](https://github.com/postcss/autoprefixer)插件来为我们的样式添加前缀。首先下载该模块。
```javascript
npm install autoprefixer --save-dev
```
接着便可以配置webpack了
```javascript
var autoprefixer = require('autoprefixer')
module.exports = {
  ...
  module: {
    loaders: [
      ...
      {
        {
          test: /\.css$/,
          loader: ExtractTextPlugin.extract(["css-loader", "postcss-loader"])
        },
      }
    ]
  },
  postcss: [autoprefixer()],
  ...
}
```
查看一下抽取出来的样式文件便可以发现已经加上了前缀
```css
a {
    display: flex;
}
/*compiles to:*/
a {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex
}
```
另外autoprefixer还可以根据目标浏览器版本生成不同的前缀个数，例如你的应用的使用用户如果大多数是使用比较新版本的浏览器，那么便可以做如下配置。
```
postcss: [autoprefixer({ browsers: ['last 2 versions'] })]
```
这是生成的样式便会有些不一样，还是上面的例子
```css
a {
    display: flex;
}
/*compiles to:*/
a {
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
}
```

### postcss后记
这里再说一个问题，有些童鞋可能会在css文件中使用@import引入其他样式文件，但是使用autoprefixer发现，import进来的样式没有处理，如下面所示：
```css
/*myStyle.css:*/
body {
  background-color: gray;
}
.flex {
  display: flex;
}

/*myStyle2.css:*/
@import "./myStyle.css";
.div {
  color: red;
}

/*autoprefixer之后*/
body {
  background-color: gray;
}
.flex {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
}
body {
  background-color: gray;
}
.flex {
  display: flex;
}
.div {
  color: red;
}
```
要解决这个问题，postcss有个[解释](https://github.com/postcss/postcss-loader#integration-with-postcss-import)，它让我们使用[postcss-import](https://github.com/postcss/postcss-import)插件，再配合autoprefixer
```javascript
postcss: function(webpack) {
  return [
    postcssImport({
      addDependencyTo: webpack
    }),
    autoprefixer
  ]
},
```
其实我们是不推荐使用@import的，心细的童鞋可以看到最后生成的样式文件有样式是重复的。  
所以一般我们应该是在js中使用require来引入样式文件。可以参考的说法[这里](https://github.com/postcss/postcss-loader/issues/35)

### 样式压缩
压缩代码我们可以使用webpack的内置插件UglifyJsPlugin来做，它既可以压缩js代码也可以压缩css代码。
```javascript
plugins: [
  ...
  new webpack.optimize.UglifyJsPlugin({
    compress: {
      warnings: false
    }
  }),
  ...
]
```
其实并不能说是在压缩css代码，本质来说还是压缩js代码，再将这块代码输出到css文件中。

### 使用[CommonsChunkPlugin](http://webpack.github.io/docs/list-of-plugins.html#commonschunkplugin)抽取公共代码
首先要明确一点CommonsChunkPlugin是在有多个entry时使用的，即在有多个入口文件时，这些入口文件可能会有一些共同的代码，我们便可以将这些共同的代码抽取出来成独立的文件。明白这一点非常重要。（搞了很久才明白的一点，唉～～～～）

如果在多个entry中require了相同的css文件，我们便可以使用CommonsChunkPlugin来将这些共同的样式文件抽取出来为独立的样式文件。
```javascript
module.exports = {
  entry: {
    "A": "./src/entry.js",
    "B": "./src/entry2.js"
  },
  ...
  plugins: [
    new webpack.optimize.CommonsChunkPlugin({name: "commons", filename: "commons.js"}),
    ...
  ]
}
```
当然，这里不止会抽取共同的css，如果有共同的js代码，也会抽取成为commons.js。  
这里有个有趣的现象，抽取出来的css文件的命名将会是参数中name的值，而js文件名则会是filename的值。

CommonsChunkPlugin好像只会将所有chunk中都共有的模块抽取出来，如果存在如下的依赖
```javascript
// entry1.js
var style1 = require('./style/myStyle.css')
var style2 = require('./style/style.css')

// entry2.js
require("./style/myStyle.css")
require("./style/myStyle2.css")

// entry3.js
require("./style/myStyle2.css")
```
使用插件后会发现，根本没有生成commons.css文件。

如果我们只需要取前两个chunk的共同代码，我们可以这么做
```javascript
module.exports = {
  entry: {
    "A": "./src/entry.js",
    "B": "./src/entry2.js",
    "C": "./src/entry3.js"
  },
  ...
  plugins: [
    new webpack.optimize.CommonsChunkPlugin({name: "commons", filename: "commons.js", chunks: ['A', 'B']}),
    ...
  ]
}
```


### 补充一下
根据webpack官网中关于[stylesheet](http://webpack.github.io/docs/stylesheets.html)的说法，建议是不要将allChunks设为true，即只是将样式嵌入到分离文件中。  
这个可能还是需要具体问题具体分析了。

