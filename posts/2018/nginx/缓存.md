> 利用nginx可以做一些缓存的功能，主要是利用expire、cache-control、etag等响应头来做实现缓存功能

# ngx_http_headers_module
该模块能够添加`Expires`和`Cache-Control`响应头
```nginx
expires    24h;
expires    modified +24h;
expires    @24h;
expires    0;
expires    -1;
expires    epoch;
expires    $expires;
add_header Cache-Control private;
```

## 案例
`/Users/seed/work_space/demo/nginxDemo`路径下有个app.html页面
```html
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <title>Page Title</title>
  <script src="app.js"></script>
</head>
<body>
  <h1>app</h1>
</body>
</html>
```

### 案例一：不配置exipre
```nginx
server {
  location /app.html {
    root /Users/seed/work_space/demo/nginxDemo;
  }
}
```
采用chrome浏览器，采用隐私模式（确保没有缓存），打开开发者工具，访问app.html，响应中http状态码为200，响应头部信息如下：
```
Response Headers
  Accept-Ranges: bytes
  Connection: keep-alive
  Content-Length: 130
  Content-Type: text/html
  Date: Mon, 21 May 2018 07:13:39 GMT
  ETag: "5b027187-82"
  Last-Modified: Mon, 21 May 2018 07:13:11 GMT
  Server: nginx/1.13.12
```
可见这里我们nginx给我们返回了ETag和Last-Modified

此时再次访问页面，此时http响应的状态码为304，响应头信息如下：
```
Response Headers
  Connection: keep-alive
  Date: Mon, 21 May 2018 07:18:15 GMT
  ETag: "5b027187-82"
  Last-Modified: Mon, 21 May 2018 07:13:11 GMT
  Server: nginx/1.13.12

Request Headers
  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
  Accept-Encoding: gzip, deflate, br
  Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
  Cache-Control: max-age=0
  Connection: keep-alive
  Host: localhost:8888
  If-Modified-Since: Mon, 21 May 2018 07:13:11 GMT
  If-None-Match: "5b027187-82"
  Upgrade-Insecure-Requests: 1
  User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36
```

另外浏览器还塞了两个请求头：  
1. 依据第一次访问时，从响应头`Last-Modified`中获取的值放置到本次请求的请求头`If-Modified-Since`中
2. 依据第一次访问时，从响应头`etag`中获取的值放置到本次请求的请求头`If-None-Match`中

服务器通过判断`If-Modified-Since`和`If-None-Match`，发现文件没有被修改，因此返回了304

注意的是，我们请求的app.html页面中，查看app.js文件的响应，会发现第一次请求是200，第二次是200(from memory cache)  
这个可能跟浏览器有关

### 案例一：配置exipre
```nginx
server {
  location /app.html {
    root /Users/seed/work_space/demo/nginxDemo;
    expires 3m;
    # expires modified 3m;
    # expires epoch;
    # expires max;
    # expires -1;
  }
}
```
如上面配置expires中  
1. 第一种配置是当期时间+有效时间
2. 第二种配置是文件的修改时间+有效时间
3. 第三种配置expire的值为Thu, 01 Jan 1970 00:00:01 GMT
4. 第四种配置expire的值为Thu, 31 Dec 2037 23:55:55 GMT，Cache-Control的值为10年
4. 第五种配置expire的值为服务器当前时间-1s,即永远过期，此时Cache-Control的值为no-cache

这里我们采用第一种，有效期3分钟

依旧采用隐私模式访问app.html，http状态码为200，响应头如下
```
Response Headers
  Accept-Ranges: bytes
  Cache-Control: max-age=180
  Connection: keep-alive
  Content-Length: 234
  Content-Type: text/html
  Date: Mon, 21 May 2018 08:13:22 GMT
  ETag: "5b027c46-ea"
  Expires: Mon, 21 May 2018 08:16:22 GMT
  Last-Modified: Mon, 21 May 2018 07:59:02 GMT
  Server: nginx/1.13.12
```
这里我们可以看到响应头返回了`Cache-Control: max-age=180`和`Last-Modified: Mon, 21 May 2018 07:59:02 GMT`，和我们设想的一直

3分钟内，再次访问页面，app.htmlhttp状态依旧为304
```
Response Headers
  Cache-Control: max-age=180
  Connection: keep-alive
  Date: Mon, 21 May 2018 08:31:56 GMT
  ETag: "5b027c46-ea"
  Expires: Mon, 21 May 2018 08:34:56 GMT
  Last-Modified: Mon, 21 May 2018 07:59:02 GMT
  Server: nginx/1.13.12

Request Headers
  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
  Accept-Encoding: gzip, deflate, br
  Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
  Cache-Control: max-age=0
  Connection: keep-alive
  Host: localhost:8888
  If-Modified-Since: Mon, 21 May 2018 07:59:02 GMT
  If-None-Match: "5b027c46-ea"
  Upgrade-Insecure-Requests: 1
  User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36
```
其实nginx已经按照http标准返回了对应的头部信息了，只是浏览器自身并没有按照http标准执行，理论上来说，此时并不应该发起请求，而是应该采用缓存中的数据，对应的状态码是200(from memory cache)

查看app.js的响应为200(from memory cache)，这个便符合我们的设想了。  
可能浏览器对html文件和其他静态文件的策略不一样。

3分钟后再次浏览器地址栏回车访问，app.html的http状态码304，这时请求头和响应头如下
```
Response Headers
  Cache-Control: max-age=180
  Connection: keep-alive
  Date: Mon, 21 May 2018 08:21:00 GMT
  ETag: "5b027c46-ea"
  Expires: Mon, 21 May 2018 08:24:00 GMT
  Last-Modified: Mon, 21 May 2018 07:59:02 GMT
  Server: nginx/1.13.12

Request Headers
  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
  Accept-Encoding: gzip, deflate, br
  Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
  Cache-Control: max-age=0
  Connection: keep-alive
  Host: localhost:8888
  If-Modified-Since: Mon, 21 May 2018 07:59:02 GMT
  If-None-Match: "5b027c46-ea"
  Upgrade-Insecure-Requests: 1
  User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36
```

此时app.js的http状态码也为304，符合设想。