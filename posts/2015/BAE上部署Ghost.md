搞了一天，终于还是把ghost blog给部署起来了。ghost博客属于bootstrap风格，还是比较不错的，而且支持的皮肤和插件都挺多，不过这些日后再研究了解吧。先把这次的部署过程记录下。  
流程如下  

* BAE环境搭建
* Ghost源代码下载
* 依据BAE要求配置ghost
* 七牛云储存配置
* 引入第三方评论系统Disqus
* 提交代码到bae

## BAE环境搭建
首先当然是要成为百度的开发者，然后登陆[BAE](http://login.bce.baidu.com/)创建应用引擎。ghost默认使用sqlite3作为存储引擎，这种方式会将博文存储在本地，而Bae会定时清除资源，所以我们必须更换数据库，可以使用mysql。bae也提供了mysql，而且免费。  
创建扩展服务，选择创建mysql服务。
git clone代码到本地，至此到此环境搭建完毕。

## Ghost源代码下载
可以到[官网](http://blog.ghost.org/)下载，也可以到[中文官网](http://www.ghostchina.com/)下载中文版的ghost，我选择中文版的，因为他支持七牛等云储存引擎，可以存储图片什么的。具体如何运行，可以参照官网。
之后便将该源码全部复制拷贝到从git clone下来的目录里面，package照样覆盖替换。

## 依据BAE要求配置ghost
首先配置ghost项目根目录下的config.js，若没有该文件，直接把config.example.js复制改名即可。将配置文件中的production下的数据库配置改为mysql的方式，如下图：
![mysql配置](http://7xjw3r.com1.z0.glb.clouddn.com/QQ20150623-2@2x.png)

接着将配置文件中production下的url修改成域名，如：http://lnwayne.duapp.com ，然后所有2368端口全部改成18080。

ghost默认启动为开发环境，我们可以修改`/core/index.js`中的`process.env.NODE_ENV = process.env.NODE_ENV || 'development';`为`process.env.NODE_ENV＝'production'`，这样启动则为生产环境的方式

## 七牛云储存配置
博客会经常上传图片，而bae也会定期清理静态资源，因此我们使用[七牛云储存](https://portal.qiniu.com/)。注册登陆后创建一个空间，接着配置conf.js文件，如下图所示
![七牛云存储配置](http://7xjw3r.com1.z0.glb.clouddn.com/QQ20150623-3@2x.png)

## 引入第三方评论系统Disqus
博客需要有评论系统，这里使用了国外的disqus评论系统。需要先到其[官网](https://disqus.com/)注册账户。点击官网右上角的setting图标，下拉菜单中点击add disqus to site。
将创建的代码复制到项目`/content/themes/casper-zh/post.hbs`处，只需要粘贴在`{{post}}`和`{{/post}}`即可。
此时评论系统已经可用，其他配置可参照http://support.ghost.org/add-disqus-to-my-ghost-blog/

## 提交代码到bae
最后就是提交代码到bae  
      
    git add -A
    git commit -m "commit"
    git push
