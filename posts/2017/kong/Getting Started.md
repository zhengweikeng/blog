# KONG安装
kong的安装很简单，直接下载[官网](https://getkong.org/install/)其安装包，安装即可

使用KONG，需要为其配置数据库，KONG支持9.4以上的版本的Postgresql，也可以选择使用3.0以上版本的Cassandra

对关系型数据熟悉的可以选择使用Postgresql，很简单的一款数据库，教程参考阮一峰[PostgreSQL新手入门](http://www.ruanyifeng.com/blog/2013/12/getting_started_with_postgresql.html)

为KONG创建一个数据库用户kong，并且创建数据库kong
```sql
CREATE USER kong;
CREATE DATABASE kong OWNER kong;
```

# KONG启动
启动kong很简单，直接`kong start`

一般来说，会提示你需要使用密码，这里的密码就是数据库的密码。  
需要通过指定KONG配置文件的方式，指定数据库和密码

在`/etc/kong/`目录下会有一份kong.conf.default，这是一份配置模板文件，我们将其拷贝为kong.conf
```
cp kong.conf.default kong.conf
```
修改kong.conf的配置，将下列注释打开
```
database = postgres
pg_host = 127.0.0.1
pg_port = 5432     
pg_user = kong     
pg_password = kong 
pg_database = kong
```

在重新启动kong
```
kong start -c /etc/kong/kong.conf
```

默认情况下，kong会占用如下几个端口
* :8000，该端口用于接受客户端的http请求，然后转发到你的相应的服务
* :8443，该端口和8000端口一样，只是它用于接受https请求
* :8001，该端口是用于调用kong的一些admin api，用于配置kong
* :8444，同上，只不是使用https

停止kong也很简单
```
kong stop
```