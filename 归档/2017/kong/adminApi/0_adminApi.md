Kong提供了restful的接口来配置和管理我们的api，其默认情况下暴露了8001和8444两个端口来管理api，其中8444是https的。  
因此我们的kong admin api必须配置在内网。

# 路由信息
## 检索kong节点信息
Endpoint: 'get /'

Response:
```json
{
  "hostname": "ad8a242cf365",
  "lua_version": "LuaJIT 2.1.0-alpha",
  "plugins": {
    "available_on_server": [
        ...
    ],
    "enabled_in_cluster": [
        ...
    ]
  },
  "configuration" : {
    ...
  },
  "tagline": "Welcome to Kong",
  "version": "0.11.0"
}
```
* available_on_server: 目前安装的所有插件
* enabled_in_cluster: 目前启用的插件，也即被所有kong节点启用的插件

## 检索节点状态
获取kong节点的使用状态信息，例如查看一些连接信息。kong是基于nginx封装的，所以一些基于nginx的监控工具都可以用来监控kong

EndPoint: 'GET /status'

Response:
```json
{
  "database": {
    "reachable": true
  },
  "server": {
  "connections_writing": 1,
  "total_requests": 182,
  "connections_handled": 181,
  "connections_accepted": 181,
  "connections_reading": 0,
  "connections_active": 2,
  "connections_waiting": 1
  }
}
```
* server: HTTP/S的相关信息
  * total_requests: 客户端请求总数
  * connections_active: 当前活跃的客户端连接数（包含等待着的连接）
  * connections_accepted: 被接收的连接数量
  * connections_handled: 被处理的连接数量
  * connections_reading: 正在读取请求头部的连接数量
  * connections_writing: 正在响应客户端的连接数量
  * connections_waiting: 正在等待请求的连接数量
* database: 数据库的信息
  * reachable: 数据库连接状态，它是个布尔值，需要注意的是，这个值并不反映数据库的健康状态