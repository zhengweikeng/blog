# 添加api
EndPoint: 'POST /apis'

Request Body

ATTRIBUTE  DESCRIPTION
name  api的名字
hosts  指向api的host，多个host通过逗号分隔。注意的是，hosts、uris和methods至少有一个必须指定
uris   指向api的uri前缀，通过该路径访问api，如`/my-path`
methods   指向api的http method，例如GET、POST，多个method通过逗号分隔
upstream_url   api服务的地址，如http://api.com
strip_uri     当通过一个uris前缀匹配一个API时，从上游的URI中去掉匹配的前缀。 默认值：true。
preserve_host    是否将上游的服务地址作为头部Host传入上游服务，默认情况下为false；