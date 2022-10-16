# csrf(Cross Site Request Forgery)跨站点请求伪造
假设你的博客有个删除博文的功能，url为  
`http://myBlog.com?m=delete&id=123456`

访问该地址便能将id=123456的博文删除（当然一般都不会这么做）

此时攻击者提供一个自己的页面
```
http://www.attacker/csrf.html
```
其内容为
```
<img src="http://myBlog.com?m=delete&id=123456" />
```
当博主用chrome登录了自己的博客，之后访问了该攻击页面，会发现自己的博文被删除了。

刚才的图片向这个删除博文的地址发起了请求，导致被删除。这便是csrf攻击。

攻击者通过诱导用户访问一个攻击页面，便以用户的身份在第三方站点里执行了一次操作。

## 关于cookie
我们看到上述例子，用户是在登录了自己博客后，访问的攻击页面，被诱导发起了一个删除博文的请求，这个请求把用户登录产生的cookie发送过去，导致了服务器端运行该删除操作。

这是csrf攻击的其中一种方式，即利用cookie的方式。我们需要先认识cookie。

浏览器的cookie分为两种：
1. session cookie，又称为“临时cookie”，它会在浏览器关闭后就失效。
1. third-party cookie，又称为“本地cookie”，它是一种被指定了有效期的cookie，直到过期之前都会保存在浏览器本地。

由于安全的原因，有一些浏览器会阻止third-party cookie的发送，如ie,safari。

但是firefox、chrome并不会阻止其发送。

因此当我们即使用chrome打开新的tab来浏览攻击网站，也会被它利用cookie的特性，对我们的删除接口发起了请求，将两种cookie都发送过去。

虽然有些浏览器能做到对本地cookie的拦截，但是现在有种叫`P3P`的请求头。通过添加P3P请求头便可以做到对所有域名开放，这也等于将浏览器的拦截机制给废弃了。

## 不是只有GET请求
html标签，像a、img等都能发起get请求，很多csrf也都是利用这一点来做的。但是不是说只能用get请求来做，其他类型的请求也是可以的，如post、put。

只需要利用表单，提交一个method为post的请求即可。

## csrf防御
### 验证码
网站在提交操作的时候，添加一个验证码，服务端做验证码校验，可以做到防止csrf攻击。但是这种方式就增加了用户的操作步骤，影响了用户体验。

### 添加请求来源
发起请求的时候在请求头加上请求的来源referer，服务端检验来源是否合法。  
但是有些时候用户会禁止发送来源，而有些浏览器也会不发送来源头给服务端。

### csrf token
通过随机的算法（例如uuid）生成一个token，服务端检验该token来判断请求是否合法。  

具体的做法是：  
1. 用户登录后，服务端生成一个随机的token，写入session中，并返回给客户端。
2. 客户端在之后的请求中带上该token。如果是get请求，可以作为url参数传递过来。如果是非get请求，如表单的话，可以作为表单的隐藏参数放置在body中传递过来。
3. 服务端检验session中的token和参数中的token是否一致。

上面是服务端将token存储在session中，也可以存放在cookie中。请求到达服务端后，从cookie中读取token和表单隐藏的token（或者url参数中的token）参数进行对比。

当然，我们可以联合使用referer和token，校验token的时候也校验referer，判断这个token是否来自一个合法的域。同时也可以校验user-agent、ip等等。

需要注意的是：
1. 如果token是存在cookie中的，一旦同一个浏览器发生用户切换，可能会导致新用户依旧使用旧token，因此在每次用户登录都应该重新分配token
1. token在一个有效生命周期内都可以重复使用，但是一旦表单提交，服务端应该重新分配token
1. 如果网站有xss漏洞的话，csrf token将会失效，因为xss可以模拟客户端对网站做任何操作，因此也能拿到隐藏在网站中的token。
 
# 参考资料  
[登录表单是否需要 CSRF 保护？](https://www.chrisyue.com/login-form-csrf-protection.html)  