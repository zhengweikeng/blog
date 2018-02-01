# 添加api
EndPoint: 'POST /apis'

Request Body

ATTRIBUTE | DESCRIPTION
----------|------------
name | api的名字
hosts | 指向api的host，多个host通过逗号分隔。注意的是，hosts、uris和methods至少有一个必须指定
uris | 指向api的uri前缀，通过该路径访问api，如`/my-path`
methods | 指向api的http method，例如GET、POST，多个method通过逗号分隔
upstream_url | api服务的地址，如http://api.com
strip_uri | 当通过一个uris前缀匹配一个API时，从上游的URI中去掉匹配的前缀。 默认值：true。
preserve_host | 默认值为false，此时会将api的hostname作为头部Host的值传入上游upstream_url的服务。默认情况下是将调用kong时所指定的Host传入上游upstream_url（查看案例三）
retry | 重试次数，默认5次
upstream_connect_timeout | 连接上游服务的超时时间，默认60000ms
upstream_send_timeout | 定义两个连续向上游服务发送的写入操作的超时时间，单位是毫秒，默认值为60000
upstream_read_timeout | 定义两个连续向上游服务发送读取操作的超时时间，单位是毫秒，默认值为60000
https_only | 是否只允许使用https，默认为false
http_if_terminated | 当强制使用https时，是否使用X-Forwarded-Proto头

## 案例一
hosts、uris和methods这三个值，只需要指定其中一个即可：
```bash
$ curl -i -X POST http://localhost:8001/apis \
  -d 'name=my-api' \
  -d 'upstream_url=http://my-api.com' \
  -d 'hosts=example.com'
```

返回成功后，即可用kong请求api
```bash
$ curl http://localhost:8000 \
  -H 'Host:example.com'
```

## 案例二
```bash
$ curl -i -X POST http://localhost:8001/apis \
  -d 'name=my-api' \
  -d 'upstream_url=http://my-api.com' \
  -d 'hosts=example.com,service.com' \
  -d 'uris=/foo,/bar' \
  -d 'methods=GET'
```

有如下请求
```bash
# 请求一，请求成功
$ curl http://localhost:8000/foo \
  -H 'Host:example.com' 

# 请求二，请求成功
$ curl http://localhost:8000/bar \
  -H 'Host:service.com' 

# 请求三，请求成功
$ curl http://localhost:8000/foo/hello/world \
  -H 'Host:example.com' 

# 请求四，请求失败
$ curl http://localhost:8000 \
  -H 'Host:example.com'

# 请求五，请求失败
$ curl http://localhost:8000/foo \
  -X POST
  -H 'Host:example.com'

# 请求六，请求失败
$ curl http://localhost:8000/foo \
  -H 'Host:foo.com'
```

通过上述请求，说明请求必须全部符合创建api时hosts、uris和methods的配置时，请求才能成功。

host还可以用通配符：
```bash
$ curl -i -X POST http://localhost:8001/apis \
  -d 'name=my-api' \
  -d 'upstream_url=http://my-api.com' \
  -d 'hosts=*.example.com'
```

## 案例三
如果是通过如下配置创建的api
```json
{
  "name": "my-api",
  "upstream_url": "http://my-api.com",
  "hosts": ["service.com"]
}
```
此时通过如下方式访问时
```
curl http://localhost:8000 -H "Host:service.com"
```
此时请求到达api服务，可以获取到如下头部
```
GET / HTTP/1.1
Host: my-api.com
```
即取到的头部是从api的url抽取出来的host

当我们将preserver_host设置为true
```json
{
  "name": "my-api",
  "upstream_url": "http://my-api.com",
  "hosts": ["service.com"],
  "preserve_host": true
}
```
此时再访问，会发现头部如下：
```bash
GET / HTTP/1.1
Host: service.com
```
即变成了调用kong时所使用的host