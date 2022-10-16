# HAProxy
## 基础配置

```haproxy
# 第一部分，全局配置，一般和操作系统配置有关
global
  # 全局日志配置，local0是日志设备，info日志级别
  # 级别还有err, warning, info, debug
  # 该配置表示使用127.0.0.1上的rsylog服务中的local0日志设备，记录日志等级为info
  log 127.0.0.1 local0 info

  # 设定每个haproxy进程可接受的最大并发连接数
  maxconn 4096

  # 运行haproxy进程的用户和组
  user nobody
  group nobody

  # haproxy进程后台运行
  daemon

  # 启动时可创建的进程数。需将daemon设置为true
  nbproc 1

  # haproxy进程pid文件。启动进程的用户必须有访问此文件的权限
  pidfile /usr/local/haproxy/logs/haproxy.pid

# 默认参数的配置部分，此处的配置会默认引用到之后的frontend、backend和listen中
# 如果某些参数属于公共配置，则可配置在此。
# frontend、backend和listen配置了和defaults一样的配置，则会覆盖掉此处的配置
defaults
  # 有tcp、http、health
  # tcp 经常用于ssl, ssh, smtp
  # http 不用多说了
  # health 已经被废弃
  mode http

  # 连接后端服务器的失败重试次数。超过此数后服务器会被标记为不可用。
  retries 3

  # 成功连接到一台服务器的最长等待时间，默认单位是毫秒
  timeout connect 10s
  # 连接客户端发送数据时最长等待时间，默认单位是毫秒
  timeout client 20s
  # 服务器端回应客户端数据发送的最长等待时间，默认单位是毫秒
  timeout server 30s
  # 对后端服务器的检测超时时间，默认单位是毫秒
  timeout check 5s

# 接收用户请求的前端虚拟节点
# 此处定义一个名为www的虚拟节点
frontend www
  bind *:80
  mode http
  # 启用日志记录http请求。默认是关闭的
  option httplog
  # 配置后可以让后端服务器获得客户端真实的ip
  # 通过X-Forwarded-For获得
  option forwardfor
  # 在客户端和服务器端完成一次请求后，HAProxy将主动关闭tcp连接。有助于提升性能
  option httpclose
  # 使用全局的日志配置格式。即使用global中的log的配置
  log global
  # 指定后端服务器池
  default_backend htmpool

# 设置集群后端服务器集群的配置
# 此处定义一个名为htmpool的后端真实服务器组
backend htmpool
  mode http
  # 将客户的请求在后端服务器出现故障时，将客户的请求强制定向到另外一台健康的后端服务器
  option redispatch
  # 在服务器负载很高的情况下，自动结束当前队列中处理时间比较长的连接
  option abortonclose
  # 负载算法
  balance roundrobin
  # 允许向cookie中插入SERVERID，可在后续的server字段中使用cookie
  cookie SERVERID
  # 启用http服务状态检测功能
  # option httpchk method uri version
  #   method http请求的方式
  #   uri 检测的地址
  #   version 指定心跳检测时的http的版本号
  option httpchk GET /index.php

  # 定义多台后端真实服务器
  # server name address[:port] [param*]
  #   name 服务器内部名称
  #   address 后端真实服务器地址或者主机名
  #   param 参数
  #     check 启用对此后端服务器执行健康状态检查
  #     inter 健康检查的时间间隔，单位为毫秒
  #     rise 设置从故障状态转换到正常状态需要成功检查的次数。此处表示检查正确2次就认为服务器可用
  #     fail 设置从正常状态到故障状态转换需要成功检查的次数。此处表示检查失败3次就认为服务器不可用
  #     cookie 为服务器设置cookie值。cookie server1表示web1的serverid为server1
  #     weight 设置权重
  #     backup 设置后端服务器的备份服务器
  server web1 10.200.34.181:80 cookie server1 weight 6 check inter 2000 rise 2 fail 3
  server web2 10.200.34.182:8080 cookie server2 weight 6 check inter 2000 rise 2 fail 3
```

## 监控页面
haproxy提供了一个监控平台，我们只需在haproxy的配置文件做以下简单配置即可

```
frontend admin
  bind *:9999
  # 开启监控
  stats enable
  # 监控页面地址
  stats uri /stats
  # 开启页面认证， 用户名:密码
  stats auth test:test
  # 隐藏haproxy版本号
  stats hide-version
  # 管理界面，如果认证成功了，可通过webui管理节点
  stats admin if TRUE
  # 统计页面自动刷新时间
  stats refresh 30s
```

之后重启haproxy访问`http://xxxx:9999/admin`，出入用户名密码即可

## 启动与重启

### 启动

```
/usr/local/haproxy/sbin/haproxy [-f 配置文件] -vdVD [-n 最大并发连接数] [-N 默认的连接数] 
```
-v：显示当前版本信息  
-d：让进程运行在debug模式，“-db”表示禁用后台模式，让程序在前台运行  
-D：让程序以daemon模式启动，此选项也可以在haproxy配置文件中设置  
-sf：程序启动后向pid文件里的进程发送FINISH信号，这个参数要放在命令行最后  
-st：程序启动后向pid文件里的进程发送TERMINATE信号，这个参数要放在命令行最后，经常用于重启haproxy进程  

### 关闭

```
killall -9 haproxy
```

### 平滑重启

```
/usr/local/haproxy/sbin/haproxy \ 
  -f ./logs/xxx.conf \
  -st ../logs/haproxy/logs/haproxy.pid
  
```

## ACL规则实现智能负载均衡
通过acl规则，可以检查客户端的请求是否合法，如果符合acl规则，则放行；否则，则中断请求

另外通过acl规则，符合规则的请求将被提交到后端的backend服务器集群，进而实现基于acl规则的负载均衡

使用方式
```
acl 自定义的acl的名称 acl方法 -i [匹配的文件或者路径]
use_backend backend名称
default_backend backend名称
```
acl方法：用来定义实现acl的方法，详见下面事例
```
ACL derivatives :
  hdr([<name>[,<occ>]])     : exact string match
  hdr_reg([<name>[,<occ>]]) : regex match
  hdr_beg([<name>[,<occ>]]) : prefix match
  hdr_dir([<name>[,<occ>]]) : subdir match
  hdr_dom([<name>[,<occ>]]) : domain match
  hdr_end([<name>[,<occ>]]) : suffix match
  hdr_sub([<name>[,<occ>]]) : substring match
  url_beg : prefix match
  url_dir : subdir match
  path_beg : prefix match
  path_reg : regex match
  path_end : suffix match
```  
-i表示不区分大小写，后面需要跟上匹配的路径或文件或正则表达式

```
# hdr_reg使用正则匹配，这里匹配以www.z.cn或者z.cn开头的请求
acl www_policy hdr_reg(host) -i ^(www.z.cn|z.cn)
# hdr_beg，匹配以某路径开头的请求
acl host_policy hdr_beg(host) -i www
acl host_static hdr_beg(host) -i img. video. download. ftp.
# hdr_dom，通过某域名发送的请求
acl bbs_policy hdr_dom(host) -i bbs.z.cn

# url_sub，通过请求中包含buy_sid=的字符串
acl url_policy url_sub -i buy_sid=

# path_end，通过请求的url中以指定字符串结尾的
acl url_static path_end .gif .png .jpg .css .js


use_backend server_www if www_policy
use_backend server_app if url_policy
use_backend server_bbs if bbs_policy
# 满足host_static与url _static 或者 满足 host_www与url _static
use_backend static if host_static || host_www url _static
use_backend www if host_www
default_backend server_cache
```