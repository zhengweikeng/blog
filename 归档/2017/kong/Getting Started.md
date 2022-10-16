# KONG安装
Kong的安装很简单，Kong[官网](https://getkong.org/install/)提供了不通平台的安装方式，最简单的我们可以使用docker的方式来安装kong。

使用Kong前，需要为其配置数据库，KONG支持9.4以上的版本的Postgresql，也可以选择使用3.0以上版本的Cassandra，我们选择使用postgresql，也是Kong的默认选择

对关系型数据熟悉的可以选择使用Postgresql，很简单的一款数据库，教程参考阮一峰[PostgreSQL新手入门](http://www.ruanyifeng.com/blog/2013/12/getting_started_with_postgresql.html)

安装postgresql也很简单，我们依旧可以采用docker的方式来安装，参见[postgresql镜像](https://hub.docker.com/_/postgres/)
```bash
$ docker pull postgres

# 为Kong创建数据库kong，用户名为kong，容器名postgres
$ docker run --name postgres -e POSTGRES_USER=kong -e POSTGRES_DB=kong -p 5432:5432 -d postgres
```

接下来安装Kong，[Kong镜像](https://hub.docker.com/_/kong/)
```bash
docker pull kong

# 数据库初始化
docker run --rm \
  --link postgres:kong-database \
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-database" \
  kong:latest kong migrations up

# 启动kong
docker run -d --name kong \
  --link postgres:kong-database \
  -e "KONG_DATABASE=postgres" \
  -e "KONG_PG_HOST=kong-database" \
  -p 8000:8000 \
  -p 8443:8443 \
  -p 8001:8001 \
  -p 8444:8444 \
  kong:latest
```

默认情况下，kong会占用如下几个端口
* :8000，该端口用于接受客户端的http请求，然后转发到你的相应的服务
* :8443，该端口和8000端口一样，只是它用于接受https请求
* :8001，该端口是用于调用kong的一些admin api，用于配置kong
* :8444，同上，只不是使用https

# 快速入门
经过上面的流程，Kong已经启动，接下来将我们尝试将应用快速接入Kong，我们使用mockbin作为我们要接入的应用。

Kong提供了restful admin api的方式来配置kong，接下来添加一个api
```bash
$ curl -i -X post \
  --url http://localhost:8001/apis \
  --data 'name=mockbin' \
  --data 'hosts=mockbin' \
  --data 'upstream_url=http://mockbin.org'
```
成功会返回类似下面的数据
```json
HTTP/1.1 201 Created
Content-Type: application/json
Connection: keep-alive

{
  "created_at": 1488830759000,
  "hosts": [
      "mockbin"
  ],
  "http_if_terminated": false,
  "https_only": false,
  "id": "6378122c-a0a1-438d-a5c6-efabae9fb969",
  "name": "example-api",
  "preserve_host": false,
  "retries": 5,
  "strip_uri": true,
  "upstream_connect_timeout": 60000,
  "upstream_read_timeout": 60000,
  "upstream_send_timeout": 60000,
  "upstream_url": "http://mockbin.org"
}
```

接下来就可以通过Kong来请求应用了
```bash
$ curl -i -X GET \
  --url http://localhost:8000 \
  --header 'Host: mockbin'
```