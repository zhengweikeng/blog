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

## 虚拟主机配置
一般来说，我们会将虚拟主机的配置写在另外一个独立的配置文件中，并在nginx.conf中使用include引用

```nginx
# 定义虚拟主机开始的关键字
server {
  # 虚拟主机端口
  listen 80;
  # 指定ip或者域名，多个域名用空格隔开
  server_name 192.168.12.188 www.test.net;
  # 设定访问的默认首页
  index index.html index.htm;
  # 指定虚拟主机的网页根目录
  root /www/wwwroot/wwww.test.net;
  # 网页默认的编码格式
  charset utf8;
  # 此虚拟主机的访问日志的存放路径
  access_log logs/www.test.net.access.log main

  # url匹配设置，后续会讲
  location
}
```

## location配置
语法规则：location [= | ~ | ~* | ^~] /uri/ {...}

1. =为精确匹配，优先级最高的匹配
2. ~区分大小写的正则匹配
3. ~*不区分大小写的正则匹配
4. ^~普通字符匹配，如果此选项匹配成功，忽略其他匹配选项，一般用来匹配目录 
5. /通用匹配，匹配任何请求，因为所有请求都是以"/"开始

匹配顺序
1. 首先进行精确匹配，匹配成功则停止其他匹配
2. ^~匹配
3. 最后是通用匹配/

例子A

```nginx
# A
location = / {...}

# B
location / {...}

# C
location ^~ /image/ {...}

# D
location ~* \.(gif|jpg|jpeg)$ {...}
```

请求url
```
# 匹配A
/

# 匹配B
/hello/world

# 匹配C
/image/hello.gif

# 匹配D
/hello/world.jpg
```

例子B

```nginx
# A
location / {
  echo "/"; //需要安装 echo 模块才行,这边大家可以改成各自的规则
}

#B
location = / {...}

# C
location = /hello {...}

# D
location ~ \.(gif|jpg|png|js|css)$ {...}

# E
location ~* \.png$ {...}

# F
location ^~ /static/ {...}
```

请求url
```
# 完全匹配 B
curl http://www.test.com/

# 完全匹配C
curl http://www.test.com/hello

# 匹配E
curl http://www.test.com/world/test.PNG

# 匹配F
curl http://www.test.com/static/test.jpg
```

### location中root和alias的差异

```nginx
location ~ ^/weblogs/ {
  root /data/weblogs/www.ttlsa.com; 
}
```
此时若请求为：  
/weblogs/httplogs/www.ttlsa.com-access.log  
web服务器将会返回服务器上的/data/weblogs/www.ttlsa.com//weblogs/httplogs/www.ttlsa.com-access.log文件

```nginx
location ^~ /binapp/ {
  alias /data/statics/bin/apps/ 
}
```
此时若请求为：  
/binapp/a.ttlsa.com/favicon  
web服务器将会返回服务器上的/data/statics/bin/apps/a.ttlsa.com/favicon文件  
即会舍弃location后的路径

### 利用location也可以实现访问控制
使用ngx_http_access_module可以实现访问控制

限制某些ip的访问

```nginx
location / {
  deny 192.168.66.90;
  allow 192.168.66.91;
  deny all;
}
```

限制访问某个目录

```nginx
location ~ ^/WEB-INF/ {
  deny all;
}
```

禁止访问doc和txt文件

```nginx
location ~* \.(txt|doc)$ {
  root /data/www/wwwroot;
  deny all;
}
```

### 请求代理
经常会有访问不同域名的将请求代理到不同服务器上

例如

一个网站有两个域名，分别是www.hello.com和w.hello.com。  
要实现当访问www.hello.com是通过nginx代理到192.168.66.90的8080端口的web上。  
当访问www.hello.com/admin是通过nginx代理到192.168.66.90的8080端口的admin上。  
当访问w.hello.com是通过nginx代理到192.168.66.90的8080端口的wap上。

```nginx
server {
  listen 80;
  server_name www.hello.com;
  location / {
    proxy_pass 192.168.66.90:8080/web/;
  }
  location /admin {
    proxy_pass 192.168.66.90:8080/admin;
  }
}

server {
  listen 80;
  server_name w.hello.com;
  location / {
    proxy_pass 192.168.66.90:8080/wap/;
  }
}

```

## 重写规则rewrite
可根据一定的规则，由一个location跳转到另外一个location

例如当访问www.helloworld.com时，自动转向www.hw.com

```nginx
server {
  server_name www.helloworld.com;
  rewrite ^/(.*)$ http://www.hw.com/$1 permanent;
}

# or
server {
  server_name www.hw.com www.helloworld.com;
  if($host != 'www.hw.com') {
    rewrite ^/(.*)$ http://www.hw.com/$1 permanent;
  }
}
```

### if命令
判断指令，可以使用如下判断

1. 一个变量的名称:空字符传”“或者一些“0”开始的字符串为 false。
2. 字符串比较:使用=或!=运算符
3. 正则表达式匹配:使用~(区分大小写)和~*(不区分大小写),取反运算!~和!~*。 4. 文件是否存在:使用-f 和!-f 操作符
5. 目录是否存在:使用-d 和!-d 操作符
7. 文件、目录、符号链接是否存在:使用-e 和!-e 操作符
8. 文件是否可执行:使用-x 和!-x 操作符

### rewrite
语法:rewrite regex replacement flag  
默认值:none  
使用字段:server, location, if  
last – 停止处理重写模块指令,之后搜索 location 与更改后的 URI 匹配。  
break – 完成重写指令。  
redirect – 返回 302 临时重定向,如果替换字段用 http://开头则被使用。  
permanent – 返回 301 永久重定向。

### set
可用于设置一些变量

```nginx
server {
  server_name www.hw.com www.helloworld.com;
  set $query $query_string;
  rewrite /dede /wordpress?$query
}
```

### break
完成当前设置的规则后，不在匹配后面的重写规则

```nginx
# or
server {
  server_name www.hw.com www.helloworld.com;
  if($host != 'www.hw.com') {
    rewrite ^/(.*)$ http://www.hw.com/error.txt break;
  }

  rewrite ^/(.*) http://www.hw.cn/$1 permanent
}

```

## 负载均衡
通过upstream命令可以实现负载均衡

```nginx
upstream myserver {
  server 192.168.12.181:80 weight=3 max_fails=3 fail_timeout=20s;
  server 192.168.12.182:80 weight=1 max_fails=3 fail_timeout=20s;
  server 192.168.12.183:80 weight=4 max_fails=3 fail_timeout=20s;
}

server {
  listen 80;
  server_name www.hello.com;
  index index.html;
  root /hello/wwwroot/;

  location / {
    proxy_passs http://myserver;
    proxy_next_upstream http_500 http_502 http_503 error timeout invalid_header;
    include /opt/nginx/conf/proxy.conf;
  }
}
```

其中weight为负载均衡的调度算法，还有如下几种算法

轮询（默认算法）  
weight，权值，值越大分配到的概率越大  
ip_hash，按ip的哈希结果分配，来自同一台ip的客户端可以固定访问一台机器  
fair，根据页面和加载的时间长短智能的进行负载均衡。需要下载upstream_fair模块  
url_hash，按访问url的哈希结果来分配请求。需要安装hash软件包。

max_fails和fail_timeout为服务器在负载均衡调度中的状态，还有如下集中状态

down，当前server暂时不参与负载均衡  
backup，备份机器。当非backup机器出现故障或者忙的时候，才会请求backup机器  
max_fails，允许请求失败的次数，默认为1。当超过最大次数，返回proxu_next_upstream模块定义的错误  
fail_timeoout 在经历了max_fails失败后，暂停服务的时间。

这里配置了proxy_next_upstream会将500、502、503等错误发生时会自动将请求转移到负载均衡中的另外一台机器。