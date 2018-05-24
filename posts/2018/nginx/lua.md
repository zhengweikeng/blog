# nginx调用lua
### 模块指令
1. set_by_lua和set_by_lua_file 设置nginx变量，可以实现复杂的复制逻辑
2. access_by_lua和access_by_lua_file 请求访问阶段处理，用于访问控制
3. content_by_lua和content_by_lua_file 内容处理器，接收请求处理并输出响应
4. content_by_lua_block 内容处理器，接收请求处理并输出响应

### nginx lua api
1. ngx.var nginx变量
2. ngx.req.get_headers 获取请求头
3. ngx.req.get_uri_args获取url请求参数
4. ngx.redirect 重定向
5. ngx.print 输出响应内容体
6. ngx.say 同ngx.print，但是最后会输出一个换行符
7. ngx.header 输出响应头

```nginx
server {
  root 8081;
  server_name 127.0.0.1;

  location /hello_lua {
    default_type 'text/plain';
    content_by_lua 'ngx.say("hello lua")';
  }

  location /nginx_var {
    default_type 'text/plain';
    # /nginx_var?a=hello,world
    content_by_lua_block {
      ngx.say(ngx.var.arg_a)
    }
  }
}
```

# 事例
根据客户端ip，判断ip是否是白名单，是的话访问server1，否则访问server2。白名单在redis中

用node起了两个服务
```javascipt
// server1.js
const http = require('http');

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end(`I am ${port} port\n`);
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});

// server2
const http = require('http');

const hostname = '127.0.0.1';
const port = 3001;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end(`I am ${port} port\n`);
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
```

配置nginx
```nginx
server {
  listen 8081;
  server_name 127.0.0.1;

  location /lua {
    default_type "text/html";
    content_by_lua_file "/home/vagrant/nginx/1.13.12/lua/ip.lua";
  }

  location @server1 {
    proxy_pass http://127.0.0.1:3000;
  }

  location @server2 {
    proxy_pass http://127.0.0.1:3001;
  }
}
```

ip.lua文件
```lua
-- ip.lua
clientIp = ngx.req.get_headers()["X-Real-IP"]

if clientIp == nil then
	clientIP = ngx.req.get_headers()["x_forwarded_for"]
end

if clientIp == nil then
	clientIp = ngx.var.remote_addr
end

local Red = require "resty.redis"
local redis = Red:new()

 local ok, err = redis:connect("127.0.0.1", 6379, {
	database = 0
})

local res, err = redis:auth("test")

if not res then
	ngx.say("failt to auth: ", err)
	return
end

if not ok then
 	ngx.say("fail to connect redis: ", err)
	return
end

local res, err = redis:get(clientIp)
if not res then
	ngx.say("fail to get data: ", err)
	return
end

if res == "1" then
	ngx.exec("@server2")
	return
end

ngx.exec("@server1")
```
