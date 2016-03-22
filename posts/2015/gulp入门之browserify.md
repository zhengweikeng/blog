> 看了gulp的文档后，觉得比起Grunt，gulp使用起来实在简单方便。光看不练是无法理解这货的，所以果断尝试将项目中使用的grunt改成用gulp实现。这里抽了其中关于browserify的部分。

关于gulp的基础知识，实在没必要讲，毕竟[官网](http://gulpjs.com/)已经有，还有中文版的[网站](http://www.gulpjs.com.cn/)。不过重点是gulp的API实在是太简单了，需要记住的也就5个，哪5个自行脑补。

[browserify](https://github.com/substack/node-browserify)是个神奇的东西，它让我们可以在写前端代码的时候，也可以使用require这样的代码，据说也几乎已经把node.js的大部分类库都包含进来。我们可以不在使用`<script src="jquery.js"></script>`这样的方式来引入jquery代码，而是使用`var $ = require('jquery')`这样的方式，当然前提是你需要`npm install jquery`。

一般我们可以在命令行如此使用browserify，`browserify main.js > bundle.js`  
如果是coffeescript文件，可以这么使用
```
npm install coffeeify
browerify -t coffeeify main.coffee > bundle.js
```

上面是用命令行的方式，如果要使用gulp来进行自动构建了，该怎么做呢，以下将细细详说。  
既然是使用gulp，那就应该找browerify的gulp的插件了，到gulp的[插件](http://gulpjs.com/plugins/)库搜索，卧槽，竟然没有。到github找吧，[gulp-browserify](https://github.com/deepak1556/gulp-browserify)。说明里已经说了，该模块已经不再维护，看来已经被gulp拉黑了。我们从gulp的[秘籍](https://github.com/gulpjs/gulp/tree/master/docs/recipes)中可以看到，官网还是推荐使用原生browserify模块的。OK，那我们就根据文档所说的来做。

项目中使用的是coffeescript来进行开发的，因此我们在使用browserify时也需要将coffee转化成js。
```
npm install coffeeify
npm install browerify
npm install gulp
npm install vinyl-source-stream
npm install vinyl-buffer
npm install gulp-sourcemaps
```
以上是需要使用到的模块，最后3个是用来干嘛，等下会说。接下来看看gulpfile代码，语法是coffeescript的语法
```
gulp = require 'gulp'
browserify = require 'browserify'
source = require 'vinyl-source-stream'
buffer = require 'vinyl-buffer'
sourcemaps = require 'gulp-sourcemaps'

gulp.task 'browserify', () ->
  b = browserify({
     entries: 'app/script/app.coffee'
     extensions: '.coffee'
     debug: true
  })
  b.bundle()
  .pipe(source('bundle.js'))
  .pipe(buffer())
  .pipe(sourcemaps.init({loadMaps: true}))
  .pipe(sourcemaps.write('.'))
  .pipe(gulp.dest('.tmp/scripts/'))
```
以上便是全部代码，现在我们来说说具体的含义。  
首先就是browserify的初始化了，具体配置方式到github看即可。这里有个小问题要说明一下，指定文件名时，根据官网所说的，可以是字符串，对象，或者数组。一开始我使用的字符串是[glob](https://github.com/isaacs/node-glob)格式的字符串，即'app/script/**/*.coffee'，发现无法使用模糊匹配，说明了，还是需要写具体的文件名。
  
[vinyl-source-stream](https://github.com/hughsk/vinyl-source-stream)是用来将常规流转换为包含 Stream 的 vinyl 对象，因为browserify使用的流并非和gulp的流一样，因此为了后续中gulp的使用，需要将一般的stream转化为gulp指定的流，即vinyl流对象（一种虚拟文件格式）。这样`.pipe(source('bundle.js'))`就解释通了。这里的bunle.js就是生成的文件名。  

[vinyl-buffer](https://github.com/hughsk/vinyl-buffer)用于将vinyl流转化为buffered vinyl文件（gulp-sourcemaps及大部分Gulp插件都需要这种格式）。因为后续的sourcemaps需要使用buffer化的vinyl流，所以第二步就需要`.pipe(buffer())`  

[sourcemaps](https://github.com/floridoo/gulp-sourcemaps)看名字就知道，这是用来生成sourcemap文件的。在browserify的配置选项中，我们配置了`debug:true`，这样便会在生成一个sourcemap代码内联在app.coffee文件中，方便调试。而`.pipe(sourcemaps.init({loadMaps: true}))`就是将其转写为一个单独的sourcemap文件，`.pipe(sourcemaps.write('.'))`写在指定的目录，此处为和budle.js同一个目录。这两步如果是在上产环境是不用的，在开发阶段还是应该加上，方便调试  

最后一个就很简单了，就是将bundle.js文件输出到指定的目录。注意这里接的是目录，不用包含文件名。  

最后还需要在package.json中加多一个配置，用于browserify来转化coffeescript
```
"browserify": {
    "transform": [
      "coffeeify"
    ]
}
```  

至此，所有配置都已经结束。其实还可以对此进行uglify，这些就是后话了，现在我们已经知道怎么使用gulp来进行browserify，之后的还有什么难度呢。
