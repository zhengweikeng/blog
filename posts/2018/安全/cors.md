如果网站在访问一个协议、域名和端口中任何一个都不相同的地址时，就会产生跨域。

解决跨域的方式有很多种，从传统的jsonp、iframe、html5 postMessage，到目前广泛使用的cors。我们可以逐个看看

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
        * application/x-www-form-urlencoded
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
表示接受请求的origin的备选值，如果是*说明接受任何origin的请求。

#### cookie操作
```
Access-Control-Allow-Credentials: true
```
默认情况下跨域发送请求时，服务器端是不接受cookie的，如果希望cookie包含在请求中，需要设置`Access-Control-Allow-Credentials: true`

另外浏览器端，需要设置ajax请求的withCredentials为true，不然即使服务器同意发送cookie，浏览器也可能不会发送。
```javascript
const request = new AJAXRequest();
request.withCredentials = true;
```

#### header
```
Access-Control-Expose-Headers: Foo
```
默认情况下，在跨域时，浏览器端在获取响应头只能是如下6个：
1. Cache-Control
1. Content-Language
1. Content-Type
1. Expires
1. Last-Modified
1. Pragma

如果想要获取其他字段，就需要在`Access-Control-Expose-Headers`里面指定

总的来说，简单请求时，服务端会返回如下响应头
```
Access-Control-Allow-Origin: http://my.web.com
Access-Control-Allow-Credentials: true
Access-Control-Expose-Headers: Foo
Content-Type: text/html; charset=utf-8
```

### 非简单请求:
除了简单请求之外的请求就是非简单请求了。  
非简单请求会在正式通信前多发送一个请求，该请求为“预检”请求（preflight）

浏览器会先询问，当前发起的请求是否符合服务器的要求。服务器会通过Origin、http method、headers等字段来检验请求，如果确认合法，浏览器正式发起请求。如果不合法，服务器不会返回任何cors相关的头信息，这时浏览器会报错，控制台会打印无法跨域请求的信息。

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
可以看到`Access-Control-Allow-Origin`这个是每个响应都一定要返回的

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