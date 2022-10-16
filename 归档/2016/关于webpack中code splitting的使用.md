### 何为code splitting
code splitting是webpack里一个非常重要的功能，利用它可以实现按需加载，减少首次加载的时间。

我们知道，如果将所有代码全部都放在一个文件中会很不恰当，尤其是打开页面的时候不需要使用到的代码块。

而很多人进行前端打包时都会把所有代码都打成一个bundle包，这样虽然开发人员是省事了，但是用户使用时很可能会因为js文件过大而造成页面卡顿的现象，尤其是大型网站。

webpack的code splitting则可以让我们做到按需加载。它的做法就是做出一个分离点，将本次加载不会用到的代码快放到分离点，一个分离点就是一个文件。需要分离点的时候再异步加载进来。
```javascript
// 第一个参数是依赖列表，webpack会加载模块，但不会执行
// 第二个参数是一个回调，在其中可以使用require载入模块
// 下面的代码会把module-a，module-b，module-c打包一个文件中，虽然module-c没有在依赖列表里，但是在回调里调用了，一样会被打包进来
require.ensure(["module-a", "module-b"], function(require) {
  var a = require("module-a");
  var b = require("module-b");
  var c = require('module-c');
});
```

### 使用范例
我们通过一个简单的例子来了解如何使用code splitting。

假设有如下的页面（index.html）
```html
<h1>code splitting</h1>
<div id="content"></div>
<button id="hello">hello</button>
<button id="world">world</button>
```

webpack入口文件代码如下（entry.js）
```javascript
$("#hello").click(() => {
  $("#content").html("<h1>hello</h1>")
})

$("#world").click(() => {
  $("#content").html("<h1>world</h1>")
})
```

这里为两个按钮注册了两个点击事件，点击后分别在页面显示hello和world。虽然很简单，但是足够拿来说明问题了。

#### 第一次优化
当应用变大庞大复杂的时候，事件处理函数会很复杂，所以正常来说，我们会把它独立成单独的文件。
```javascirpt
// hello.js
export default () => {
  $("#content").html("<h1>hello</h1>")
}

// world.js
export default () => {
  $("#content").html("<h1>world</h1>")
}
```
这时候entry.js则为如下
```javascript
import hello from './hello'
import world from './world'

$("#hello").click(() => {
  hello()
})

$("#world").click(() => {
  world()
})
```

页面在加载的时候会去加载js文件，但是我们发现，这两个事件的处理函数，在页面加载的时候压根不会被调用。  

但是webpack默认会将所有js文件打包成一个文件，这样页面就得去加载这个大而冗余的文件。

这里代码简单，所以没问题。但是当应用复杂的时候，就会带来很多性能上的问题。

我们接下来使用webpack的code splitting来优化这段代码。

#### 第二次优化
修改entry.js如下
```javascript
let hello = null
let world = null
require.ensure([], function(require) {
  hello = require('./hello').default
})
require.ensure([], function(require) {
  world = require('./world').default
})

$("#hello").click(() => {
  hello()
})

$("#world").click(() => {
  world()
})
```
我们使用require.ensure来进行代码分离，一次分离会产生一个文件，因此这里会产生两个文件，再加上webpack打包后的文件（bundle.js），总共会有3个文件。

很明显，这样之后我们页面加载的时候会加载3个文件，虽然http的请求次数多了，但是会首先加载bundle.js，再加载另外两个分离文件。

当然从这段代码看，是没法看出是否变快的。但是应用变大了，这个性能还是有得考量的。

#### 第三次优化
我们发现，其实并没有做到按需加载，毕竟还是产生了另外两次http请求去加载文件。那能不能连这两次请求都省去呢。当然是可以的。

继续来看entry.js的修改
```javascript
const clickHandler = (id) => {
  require.ensure([], function(require) {
    const src = `./${id}`
    const handler = require(src).default
    handler()
  })
}

$("#hello").click(() => {
  clickHandler("hello")
})

$("#world").click(() => {
  clickHandler("world")
})
```
再次运行webpack打包命令后，我们会发现除了bundle文件外，只会产生一个分离文件，即总共只有2个文件，不再是3个文件。

我们会发现分离文件已经有hello.js和world.js的全部代码了。

再打开页面看会发现只加载了一个bundle.js文件，另外一个js文件的http请求并没有发起。当有对按钮进行点击后才会去发起请求加载另外一个文件。

这样我们便做到了，打开页面，对于页面不需要使用的代码不再加载，并且在需要时（即点击的时候）才进行加载，即按需加载。
