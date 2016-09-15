## nginx中http模块相关属性的配置

```nginx
http {
  # 引用其他配置文件 
  include conf/*.conf;
  # 默认的文件类型。当文件类型为定义时使用这种方式
  # 例如如果没有配置php环境，nginx不会去解析它，此时访问php文件会出现下载窗口 
  default_type application/octet-stream;
  # 指定日志的输出格式。 main为日志输出格式名称，可在access_log中引用
  log_format main '$remote_addr - $remote_user [$time_local] '
    '"$request" $status $bytes_sent '
    '"$http_referer" "$http_user_agent" '
    '"$gzip_ratio"';
  log_format download '$remote_addr - $remote_user [$time_local] '
    '"$request" $status $bytes_sent '
    '"$http_referer" "$http_user_agent"'
    '"$http_range" "$sent_http_content_range"';
  # 访问日志 存储的位置和所使用的日志格式
  access_log logs/www.test_nginx.access.log main;

  # 允许客户端请求的最大单个文件字节数
  client_max_body_size 20m;
  # 指定客户端请求头的headerbuffer大小
  client_header_buffer_size 32k;
  # 指定客户端请求中较大的消息头的缓存最大数量和大小。 4个128k
  large_client_header_buffers 4 128k;
  # 开启高效文件传输模式
  sendfile on;
  # 防止网络阻塞
  tcp_nopush on;
  tcp_nodelay on;
  # 设置客户端连接保持活动的超时时间
  keeplive_timeout 60;
  # 设置客户端请求头的超时时间。超时将返回Request time out(408)
  client_header_timeout 10;
  # 设置客户端请求主体超时时间。超时将返回Request time out(408)
  client_body_timeout 10;
  # 指定响应客户端的超时时间。
  send_timeout 10;

  # 开启gzip模块
  gzip on;
  # 允许压缩的页面最小字节数，该数从header头的Content_Length中获取。
  # 默认值是0。建议设置大于1KB，小于1KB可能会越压缩越大。
  gzip_min_length 1k;
  # 申请4个单位为16KB的内存作为压缩结果流缓存。
  gzip_buffers 4 16k;
  # 设置识别http协议版本
  gzip_http_version 1.1;
  # 指定压缩比。 
  # 1表示压缩比最小，处理速度最快；
  # 9表示压缩比最大，传输速度最快，但处理速度最慢，也比较消耗CPU资源
  gzip_comp_level 2;
  # 指定压缩类型。无论是否指定，text/html都会被压缩
  gzip_types text/plain application/x-javascript text/css application/xml;
  # 让前端的缓存服务器缓存经过gzip压缩的页面
  gzip_vary on;

  # 虚拟主机配置。后面会详细解释
  server {
    ...
  }
}
```

## 关于日志配置
### log_format 配置日志格式
语法： log_format name string ...  
配置段：http

nginx有个默认的日志格式，名称为combined，格式如下：

```nginx
log_format combined '$remote_addr - $remote_user [$time_local] '
    '"$request" $status $body_bytes_sent '
    '"$http_referer" "$http_user_agent" ';
```

如果nginx作为反向代理，此时web服务器无法获得客户端真实的ip地址。$remote_addr获取反向代理的ip地址。

反向代理的服务器在转发请求的Http头信息中，可以增加X-Forwarded-For信息，用来记录客户端IP地址和客户端请求的服务器地址。

关于X-Forwarded-For可以参考这篇[文章](https://imququ.com/post/x-forwarded-for-header-in-http.html)

此时我们可以定义如下日志格式：

```nginx
log_format proxy '$http_x_forward_for - $remote_user [$time_local] '
    '"$request" $status $body_bytes_sent '
    '"$http_referer" "$http_user_agent" ';
```

日志格式可使用的变量如下：
```  
$remote_addr，$http_x_forward_for用于记录用户的ip地址
$remote_user 客户端用户名称
$request 请求的URL和HTTP协议名称
$status 请求状态
$body_bytes_sent 发送给客户端的字节数,不包括响应头的大小
$bytes_sent 发送给客户端的总字节数。
$connection 连接的序列号。
$connection_requests 当前通过一个连接获得的请求数量。
$msec 日志写入时间。单位为秒,精度是毫秒
$pipe 如果请求是通过 HTTP 流水线(pipelined)发送,pipe 值为“p”,否则为“.”
$http_referer 记录从哪个页面链接访问过来的
$http_user_agent 记录客户端浏览器相关信息
$request_length 请求的长度(包括请求行,请求头和请求正文)。
$request_time 请求处理时间,单位为秒,精度毫秒; 从读入客户端的第一个字节开始,直到把最后一个字符发送 给客户端后进行日志写入为止。
$time_iso8601 ISO8601 标准格式下的本地时间。
$time_local 通用日志格式下的本地时间
```

### access_log 访问日志
语法: access_log path [format [buffer=size [flush=time]]]  
默认值: access_log logs/access.log combined;  
配置段: http, server, location, if in location, limit_except

gzip 压缩等级。  
buffer 设置内存缓存区大小。  
flush 保存在缓存区中的最长时间。  
不记录日志:access_log off

### error_log 配置错误日志
语法: error_log file | stderr | syslog:server=address[,parameter=value] [debug | info | notice | warn | error | crit | alert | emerg];  
默认值: error_log logs/error.log error;  
配置段: main, http, server, location

### rewrite_log 记录重写日志

语法: rewrite_log on | off;  
默认值: rewrite_log off;  
配置段: http, server, location, if  
启用时将在 error log 中记录 notice 级别的重写日志

未完待续...
