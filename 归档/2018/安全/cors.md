如果网站在访问一个协议、域名和端口中任何一个都不相同的地址时，就会产生跨域。

同源策略限制以下几种行为：
1. Cookie、LocalStorage 和 IndexDB 无法读取
2. DOM 和 Js对象无法获得
3. AJAX 请求不能发送

解决跨域的方式有很多种，从传统的jsonp、iframe、html5 postMessage，到目前广泛使用的cors。重点看下cors

## jsonp跨域
利用了script标签允许跨域的特性，获取json数据
```html
<script>
function doSomeThing() {
  // some code
}
</script>
<script src="http://example.com/data.php?callback=doSomeThing">
```
假设上述页面a.html想要通过ajax获取一个不同域上的json数据，地址为example.com/data.php。我们约定了一个名为callback的请求参数，当然你也可以用其他参数名。

由于是作为script标签引入的，所以data.php必须返回一个能够执行的js脚本。
```php
$callback = $_GET['callback']
$data = array('a', 'b', 'c')
echo $callback.'('.json_encode($data).')'
```
该php相当于执行了callback参数中指定的方法，即doSomeThing。

需要注意的是：
1. jsonp只能发起get请求
1. 会有引发xss漏洞的可能

## 通过修改document.domain来跨子域
当页面有父子页面的时候，如果父子页面是不同域时，虽然能够获取到子页面的window对象，却无法获取到其几乎任何属性和方法

为什么说几乎呢，因为有些方法是能拿到的，如后面要介绍的postMessage。 

假设有有个页面`www.example.com/a.html`，其引入了一个`example.com/b.html`的子页面。
```html
<script>
function loadPage() {
  var iframe = document.getElementById('iframe');
  var win = iframe.contentWindow;
  var doc = win.document; // 无法获取到
  var name = win.name; // 无法获取到
}
</script>
<iframe id="iframe" src="http://example.com/b.html" onload="loadPage()">
```
由于跨域的原因，导致doc和name都为空。

解决这个问题，只需要将两个页面的document.domain设置成一致即可。但要注意的是，document.domain的设置是有限制的，我们只能把document.domain设置成自身或更高一级的父域，且主域必须相同。例如：a.b.example.com 中某个文档的document.domain 可以设成a.b.example.com、b.example.com 、example.com中的任意一个，但是不可以设成 c.a.b.example.com,因为这是当前域的子域，也不可以设成baidu.com,因为主域已经不相同了。

我们先设置www.example.com/a.html
```html
<iframe id="iframe" src="http://example.com/b.html" onload="loadPage()">
<script>
document.domain = 'example.com'
</script>
```
这时还是得显式设置example.com/b.html，虽然其域名已经是example.com
```javascript
document.domain = 'example.com'
```
通过这种方式，a.html中就能通过iframe访问b.html中的属性和方法了。

不过如果你想在www.example.com/a.html页面中通过ajax直接请求example.com/b.html页面，即使你设置了相同的document.domain也还是不行的，所以修改document.domain的方法只适用于不同子域的框架间的交互。

可以通过jsonp的方式来实现ajax的跨域

但是我们可以换成方式间接实现该功能：
1. 将iframe隐藏，且src的与ajax请求同域，此时ifame的页面用于同域的原因是能够发起请求的；
1. 设置两个父子页面的document.domain，完成父页面对子页面的控制，获得js操作权，达到跨域。

使用document.domain的缺陷：
1. 如果修改了document.domain，则在某些机器上的IE678里，获取location.href有权限异常
1. 设置document.doamin，也会影响到其它跟iframe有关的功能，如富文本编辑器、ajax的前进后退
1. 设置document.doamin，导致ie6下无法向一个iframe提交表单

## 通过window.name
window对象有个name属性，该属性有个特征：即在一个窗口(window)的生命周期内,窗口载入的所有的页面都是共享一个window.name的，每个页面对window.name都有读写的权限，window.name是持久存在一个窗口载入过的所有页面中的，并不会因新页面的载入而进行重置。

假设有个www.example.com/a.html的页面如下：
```html
<script>
window.name = 'this is from a.html'

setTimeout(function () {
  window.location = 'b.html'
}, 3000)
</script>
```
a.html会在3秒后跳转到b.html，其js代码如下:
```javascript
alert(window.name)
```
会发现弹框内容为`this is from a.html`，说明即使已经跳转到b.html页面，window.name已经为之前设置的值。  
当然你可以在b.html中改变window.name的值，这时便可以被修改掉了。这个值最大能到大约2M的容量，取决于不同浏览器。

上述例子中，a.html和b.html是同域的，即使换成不同域也是成立的。因此我们可以借助这个特点来完成跨域。

假设有一个www.example.com/a.html页面,需要通过a.html页面里的js来获取另一个位于不同域上的页面example.com/data.html里的数据。实现的方式就是在a.html页面中使用一个隐藏的iframe来充当一个中间人角色，由iframe去获取data.html的数据，然后a.html再去得到iframe获取到的数据。

example.com/data.html中将获取到的数据，设置到window.name中即可。

而中间人iframe想要获取到data.html中的通过window.name设置的数据，只需要把这个iframe的src设为examples.com/data.html就行了。然后a.html想要得到iframe所获取到的数据，也就是想要得到iframe的window.name的值，只要把这个iframe的src设成跟a.html页面同一个域即可。

```html
<script>
function loadPage() {
  var iframe = document.getElementById('iframe');
  iframe.onload = function() {
    var data = iframe.contentWindow.name; // 获取到data.html中的数据
    console.log(data)
  }

  iframe.src = 'b.html'; // 设置成任何一个和a.html同域的即可，即使该页面不存在也可以。
}
</script>

<iframe id="iframe" src="http://example.com/data.html" onload="loadPage()" style="display:none">
```
在http://example.com/data.html中代码如下：
```html
<script>
var data = 'Some datas';
window.name = data;
</script>
```

window.name的弊端：
1. 可能有的旧版本的浏览器没有window.name
1. 实现方式还是有些复杂

## 通过window.postMessage
window.postMessage(message,targetOrigin)是html5中引入的方法，通过该方法可以向任何域发送消息，无论是否同域。

其中window指的是要接收该message的window对象，我们知道即使不同域，我们也能通过js获取到window对象，只是其失去了大部分属性和方法。

message是一个字符串，而targetOrigin是我们要限制的域，如果不想限定可以采用*

看个例子，如果我们www.example.com/a.html想要给example.com/b.html发送数据
```html
<script>
function loadPage() {
  var iframe = document.getElementById('iframe');
  var win = iframe.contentWindow; // 获取window对象
  win.postMessage('hello, i am from a.html'); // 向http://example.com/b.html发送消息
}
</script>
<iframe id="iframe" src="http://example.com/b.html" onload="loadPage()" style="display:none">
```

http://example.com/b.html的代码如下：
```html
<script>
window.onmessage = function(e) {
  var message = e.data
  // do some things
}
</script>
```

由此可见，在使用ifame的方式中，这种方式是显得很简单的一种了，但是它的缺点就是很多旧版本浏览器不支持，如IE6和IE7等。

## cors（Cross-Origin Resource Sharing）
跨域将请求分为`简单请求`和`非简单请求`  
### 简单请求:  
1. 请求方式只能为：GET、POST、PUT
1. 请求头只能为：
    * Accept
    * Accept-Language
    * Content-Language
    * Last-Event-ID
    * Content-Type只能为：
        * application/x-www-form-urlencoded
        * multipart/form-data
        * text/plain

当浏览器发现请求是简单请求时，会自动加上该头
```
Origin: 'http://my.web.com'
```
origin说明来该请求的协议、域名和端口，服务器根据这个值判断是否同意该请求。

这时服务器需要返回一个是否接受的请求头
```
Access-Control-Allow-Origin: http://my.web.com
```
浏览器收到响应后，如果响应头没有包含`Access-Control-Allow-Origin`或者请求的origin不在该值里面，会被请求的onerror回调函数捕获。

需要注意的是，这时请求可能是状态码200，但是却是个异常的响应。

#### origin
```
Access-Control-Allow-Origin: *|url[, moreUrls]
```
表示接受请求的origin的备选值，如果是*说明接受任何origin的请求。

#### cookie操作
```
Access-Control-Allow-Credentials: true
```
默认情况下跨域发送请求时，服务器端是不接受cookie的，如果希望cookie包含在请求中，需要设置`Access-Control-Allow-Credentials: true`

另外浏览器端，需要设置ajax请求的withCredentials为true，不然即使服务器同意发送cookie，浏览器也可能不会发送。
```javascript
const request = new AJAXRequest();
request.withCredentials = true;
```

#### header
```
Access-Control-Expose-Headers: Foo
```
默认情况下，在跨域时，浏览器端在获取响应头只能是如下6个：
1. Cache-Control
1. Content-Language
1. Content-Type
1. Expires
1. Last-Modified
1. Pragma

如果想要获取其他字段，就需要在`Access-Control-Expose-Headers`里面指定

总的来说，简单请求时，服务端会返回如下响应头
```
Access-Control-Allow-Origin: http://my.web.com
Access-Control-Allow-Credentials: true
Access-Control-Expose-Headers: Foo
Content-Type: text/html; charset=utf-8
```

### 非简单请求:
除了简单请求之外的请求就是非简单请求了。  
非简单请求会在正式通信前多发送一个请求，该请求为“预检”请求（preflight）

浏览器会先询问，当前发起的请求是否符合服务器的要求。服务器会通过Origin、http method、headers等字段来检验请求，如果确认合法，浏览器正式发起请求。如果不合法，服务器不会返回任何cors相关的头信息，这时浏览器会报错，控制台会打印无法跨域请求的信息。

假设我们发起如下请求
```javascript
var url = 'http://some.server.com/cors';
var xhr = new XMLHttpRequest();
xhr.open('PUT', url, true);
xhr.setRequestHeader('Custom-Header', 'hello');
xhr.send();
```

这时浏览器发现这是个非简单请求，于是发起预检请求
```http
OPTIONS /cors HTTP/1.1
Origin: http://my.web.com
Access-Control-Request-Method: PUT
Access-Control-Request-Headers: Custom-Header
Host: some.server.com
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0
```

这时服务器返回:
```http
HTTP/1.1 200 OK
Date: Mon, 16 Apr 2018 12:21:33 GMT
Server: Apache/2.0.61 (unix)
Access-Control-Allow-Origin: http://my.web.com
Access-Control-Allow-Methods: GET,POST,PUT
Access-Control-Allow-Headers: Custom-Header
Content-Type: text/html; charset=utf-8
Content-Encoding: gzip
Content-Length: 0
Keep-Alive: timeout=2, max=100
Connection: Keep-Alive
```
这时服务器接受了该请求，允许跨域，于是浏览器可以正式发起请求。

```http
PUT /cors HTTP/1.1
Origin: http://my.web.com
Host: some.server.com
Custom-Header: hello
Accept-Language: en-US
Connection: keep-alive
User-Agent: Mozilla/5.0...
```

服务器正常响应
```http
Access-Control-Allow-Origin: http://my.web.com
Content-Type: text/html; charset=utf-8
```
可以看到`Access-Control-Allow-Origin`这个是每个响应都一定要返回的

#### Request-Method
```
Access-Control-Request-Method: PUT
```
非简单请求发送请求时需要加上该请求头，用于说明浏览器的cors请求会用到的http方法。

#### Request-Headers
```
Access-Control-Request-Headers: Custom-Header
```
通过该头说明cors请求会额外发送的头信息。

#### allow header
```
Access-Control-Allow-Headers: Custom-Header
```
如果非简单请求中包含了`Access-Control-Request-Headers`，则需要返回该头说明服务器接受的头信息字段，多个用逗号分隔。

#### max age
```
Access-Control-Max-Age: 2592000
```
用于指定该预检请求的有效期，单位为秒。在这个时间内，不需要再发出预检请求。

## 其他跨域方式
除了上述几种，还有document.hash的方式，nginx代理的方式等等，具体可以参考下：  
[前端常见跨域解决方案](https://www.cnblogs.com/roam/p/7520433.html)
