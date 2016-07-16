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

const redisCli = redis.createClient(6379)

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

接下来就是服务的配置了，compose将要编排的的容器组称之为服务，每个服务都可以启动多个容器。

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
