很多时候，我们需要集中的管理多个docker容器，例如我们会希望自己不用一个个的去启动容器，毕竟这样会很耗费时间和精力。我们希望可以通过配置文件，将启动容器时需要的配置信息写好，然后通过一条命令就可以启动所有容器。

目前github上已经有这样的开源工具，如crane。

而docker官方也实现和开源了这样的一套工具，即编排系统，docker-compose。它的前身就是Fig。

接下来我们就来学习使用这样的一个工具，我们将会以Node.js为例子来学习使用。

### Getting Started
首先你需要安装docker-compose。如何安装直接就看[官网](https://docs.docker.com/compose/install/)了，已经很清楚。

#### 创建node.js项目
1.创建一个项目目录

```bash
$ mkdir compose-demo
$ cd compose-demo
$ npm init
$ npm i redis --save 
```
2.创建一个简单的node.js服务器
```javascript
// app.js
const http = require('http')
const redis = require('redis')

const redisCli = redis.createClient(6379, 'redis')

const server = http.createServer((req, res) => {
  const num = Math.floor(Math.random() * 100)
  redisCli.set('num', num)

  res.statusCode = 200
  res.setHeader('Content-Type', 'text/plain')

  res.end(`Hello, you set a random num ${num} to redis \n`)
})

server.listen(8080)
```

3.创建Dockerfile
```dockerfile
FROM node:6.2

COPY ./package.json /tmp/package.json
RUN cd /tmp && npm install
RUN mkdir -p /compose-demo && cp -a /tmp/node_modules /compose-demo/

WORKDIR /compose-demo
ADD . /compose-demo/

EXPOSE 8080

CMD ["node", "app.js"]
```

4.创建docker镜像
```bash
$ docker build -t compose-demo:latest .
```

5.创建compose配置文件docker-compose.yml
```yaml
version: '2'

services:
  compose-demo:
    build: .
    container_name: compose-demo
    ports:
      - "18080:8080"
    volumes:
      - .:/compose-demo
    depends_on:
      - redis
  redis:
    image: redis 
```
整个过程其实最重要的就是这份配置文件了，分析一下这份文件。

首行就是说明了这份文件使用compose版本2来解析，版本1也会在不远的将来被淘汰掉。

接下来就是服务的配置了，compose将要编排的的容器称之为服务，多个服务组成了一个服务组，compose便是对这些服务进行管理。

这里我们将会启动两个容器，第一个是我们的Node项目容器，它会将容器中的8080端口暴露在宿主机器的18080端口上面。

我们还会将容器中的/compose-demo目录挂在到宿主机器的当前目录下面，这样我们修改源代码的时候就不用重新构建镜像了。

最后我们使用ddpends_on关键字，说明这个容器依赖于另外一个容器redis，配置之后容器之间便可以进行交互，又不会暴露在外部宿主机器，有点像link命令。

6.使用docker-compose构建镜像和容器
```bash
$ docker-compose up
```
会发现首先构建的是redis镜像，并首先启动redis容器。之后才是我们自己的容器，因为我们的容器依赖于redis容器。

启动成功后便可以直接访问`http://localhost:18080`，会看到类似这样的页面
```
Hello, you set a random num 35 to redis
```

上面使用的启动命令，会使compose在前台运行，如果想让它在后台运行可以使用`docker-compose up -d`

### 环境变量
#### docker-compose.yml变量替换
compose配置文件允许我们动态的传入环境变量。

还是上面的配置文件，不过我们让宿主端口动态的绑定
```yaml
version: '2'

services:
  compose-demo:
    build: .
    container_name: compose-demo
    ports:
      - "${WEB_PORTS}:8080"
    volumes:
      - .:/compose-demo
    depends_on:
      - redis
  redis:
    image: redis
```

compose使用$符号来引入环境变量，$var和${var}都是支持的。

运行`docker-compose up`的时候，compose会先去我们当前的shell的环境变量中查找WEB_PORTS，找到了则为它赋值。如果没找到则compose会抛出错误。

compose也支持使用环境变量配置文件的方式引入变量，docker-compose.yml的目录下面创建一个.env的文件，compose运行的时候会去读取它。
```
# .env file
WEB_PORTS=28080
```
不过，如果我们当前的shell环境中也含有同名的变量，则compose会以shell环境变量优先，即shell环境变量会覆盖.env配置文件中的同名变量。

#### 容器环境变量
我们在运行容器的时候经常会传入一些环境变量，例如`docker run -e NODE_ENV=development ...`。compose支持多种方式给容器传递环境变量。

1.在compose配置文件中定义环境变量
```yaml
services:
  compose-demo:
    environment:
      - NODE_ENV=development
```
2.在compose配置文件中配置env_file选项
```yaml
services:
  compose-demo:
    env_file:
      - web-variables.env
```
这时，应该在docker-compose.yml下应该有一个web-variables.env文件
```
# web-variables.env
NODE_ENV=development
```
3.运行docker-compose run 为单独的容器配置环境变量
```bash
$ docker-compose up -e NODE_ENV=development compose-demo node app.js
```

### compose配置文件
compose默认的配置文件是docker-compose.yml，当然我们也可以使用`docker-compose -f`指定配置文件。接下来我们来了解下配置文件中每一项的配置方法。

注意：我们的配置都是基于版本2的，即version:2

#### build
该命令会在镜像构建的时候被应用到，它可以是个字符串，代表了镜像构建时的上下文的路径。也可以是个对象，包含了context，dockerfile和args三个选项。
```yaml
build: ./dir

build:
  context: ./dir
  dockerfile: Dockerfile-alternate
  args:
    buildno: 1
```
其中context代表了dockerfile的路径或者git仓库的路径。dockerfile则是使用的dockerfile的文件名，该选项也可以独立出来，不做为build的子选项。

#### cap_add, cap_drop
添加或者删除容器能力。linux下有capabilities命令，自己尚未研究过，不做讨论。
```bash
cap_add:
  - ALL

cap_drop:
  - NET_ADMIN
  - SYS_ADMIN
```

#### command
覆盖容器运行时的CMD指令

#### cgroup_parent
配置容器的父组

#### container_name
配置容器的名称。默认compose会自己生成一个名字给容器。

#### devices
将容器的设备映射到宿主机器

#### depends_on
compose服务间的依赖关系，被依赖的服务会优先启动
```yaml
version: '2'
services:
  web:
    build: .
    depends_on:
      - db
      - redis
  redis:
    image: redis
  db:
    image: postgres
```
无论使用`docker-compose up`还是`docker-compose up web`，redis和db都会优先启动。  

depends_on还有个好处就是，我们可以直接使用service的名字当作访问地址来访问容器，例如连接redis可以如下所示
```javascript
redis.createClient({host: 'redis', port: 6379})
```

#### dns
设置dns

#### dns_search
设置dns搜索域名

#### tmpfs
挂在一个临时目录到容器中

#### entrypoint
覆盖默认的entrypoint

#### env_file
指定环境变量文件，前面环境变量例子已经提到

#### environment
配置容器环境变量，前面环境变量例子已经提到

#### expose
暴露端口。只有互联的容器才能使用，宿主机器无法访问暴露的端口。

#### extends
用于扩展服务。可以对指定的service进行扩展，因此该选项中必须包含一个新的compose配置文件
```yaml
extends:
  file: common.yml
  service: webapp
```
一般来说，但各个服务之间有很多共有的配置项时，会把它们抽取成独立的文件，这样各个服务间便可以相互引用
```yaml
# common.yml
app:
  build: .
  environment:
    CONFIG_FILE_PATH: /code/config
    API_KEY: xxxyyy
  cpu_shares: 5


# docker-compose.yml
webapp:
  extends:
    file: common.yml
    service: app
  command: /code/run_web_app
  ports:
    - 8080:8080
  links:
    - queue
    - db

queue_worker:
  extends:
    file: common.yml
    service: app
  command: /code/run_worker
  links:
    - queue
```

#### external_links
用于连接其他compose配置文件的容器

#### extra_hosts
为容器添加host
```yaml
extra_hosts:
  - "somehost:162.242.195.82"
  - "otherhost:50.31.209.229"
```
相当于在容器的/etc/hosts中添加了
```
162.242.195.82  somehost
50.31.209.229   otherhost
```

#### image
指定了容器启动时的镜像。如果该镜像在本地不存在，则会执行pull拉取。

如果和build一起使用，则会使用该配置项的值作为镜像的名字进行镜像的构建。

#### labels
设置label标签

#### links
和docker的links一致，用于连接到其他的容器，容器之间可以使用service的名称相互访问，和depends_on用法一样，也会决定service的启动顺序。

#### logging
为service配置日志
```yaml
logging:
  driver: syslog
  options:
    syslog-address: "tcp://192.168.0.42:123"
```
其中drive的值可以为json-file、syslog和none。当想使用`docker-compose logs`必须使用json-file的格式，日志才可以看到。

#### net/network_mode 
配置网络的模式
```yaml
net: "bridge"
net: "host"
net: "none"
net: "container:[service name or container name/id]"
```
#### networks
要加入service的网络

#### pid
应用也是用于和其他容器交互的

#### ports
同docker配置ports一样，用于暴露容器端口，并和宿主端口进行映射。  
需要注意的是，yaml的语法会将`xx:yy`这种格式的数字转化为60进制的格式，所以应该将数字用双引号括起来
```yaml
ports:
  - "3000"
  - "3000-3005"
  - "8000:8000"
  - "9090-9091:8080-8081"
  - "49100:22"
  - "127.0.0.1:8001:8001"
  - "127.0.0.1:5000-5010:5000-5010"
```

#### security_opt
覆盖默认的标签模式

#### volumes/volumes_from
同docker的volumes是一致的
