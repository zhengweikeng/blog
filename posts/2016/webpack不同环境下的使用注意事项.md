> webpack在不同环境下，需要注意的事项会有些不同，本文根据webpack文档中[list of hints](http://webpack.github.io/docs/list-of-hints.html)整理而成。

### App performance
需要编译出一个性能好的app应用需要注意的事项

1. 使用UglifyJsPlugin插件压缩你的所有文件
2. 使用code splitting，它可以帮助你提高首屏加载的速度，但是可能会消费额外的http请求。（如果是react app，可以使用react-proxy-loader）
3. 文件名最好使用hash命名，如[hash].bundle.js或者[chunkhash].bundle.js，而且要设置一个很长的缓存时间。
4. 当有些静态文件不再使用到时，不要马上把它删除掉。可以等待一段时间再删除。马上删掉的话可能会导致一些404页面。
5. 你应该使用DefinePlugin插件来定义一些环境变量，app可以根据不同的环境变量做不同的处理。（EnvironmentPlugin插件可以将process.env传入到你的app中）
6. 你应该使用一些分析工具来获取一些app的使用数据，这样可以帮助你定位问题。如stats-webpack-plugin插件或者使用profile选项获取更加完善的数据。
7. 将共有的或者公共的模块抽取独立出来。
8. css的处理  
  * css中会引用很多的静态文件，如图片和字体，这些都应该使用url-loader做处理
  * 可以使用extract-text-webpack-plugin插件来将样式抽取出来，而不是直接插入到页面中
9. 合理的使用code splitting，可以配合插件(LimitChunkCountPlugin, MinChunkSizePlugin, AggressiveMergingPlugin)一起使用。
10. 可以通过插入script标签或者延迟加载chunk的方式预加载一些额外的chunk。

### Developer performance
开发的时候，可以做一些适当的配置，提高我们的开发效率

1. 你应该用配置文件（webpack.config.js）的方式，而不是命令行的方式来使用webpack。这样维护和配置起来也方便
2. 不要尝试去重写那些不兼容的js模块，而是使用imports-loader/exports-loader去让这些模块变得兼容webpack。
3. 你应该开启devtools，这会让你很方便的在浏览器端进行调试。
4. 使用ES6模块化的方式进行开发，然后使用babel-loader对其进行解析。
5. 如果你是在开发一个js库，如lodash这些，你应该去使用output.library和output.libraryTarget，它可以帮助你依据umd里不同的标准构建代码。
6. 使用externals来声明一些想用script引入的依赖包。
7. 启用模块热替换（Hot Module Replacement）实现页面实时自动刷新
8. 可以写一份webpack的基础配置文件，然后开发和生产的配置文件引用这份基础配置文件。
9. 使用resolve.root配置你应用中引用的外部模块的路径，也可以减少写很长的引用路径。
10. 使用karam或者karam-webpack去测试你的模块
11. 使用target在浏览器端以外的环境进行编译
12. 使用BannerPlugin为你的静态文件添加注释
13. 你应该启用debug，它会提供更多的调试信息
14. 对于包的引用，使用include而不是exclude，可以减少错误，路径配置也方便。
