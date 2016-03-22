##### 1. 安装docker
Jenkins所在服务器为mac系统，使用docker需要先安装boot2docker，具体安装流程参考[官网](https://docs.docker.com/installation/mac/)。此次安装是使用mobileDev用户进行全局安装boot2docker。

##### 2. Jenkins项目构建脚本——docker build
安装完docker后，便可以在jenkins中使用，在构建脚本处使用`docker build -t xxx .`进行打包生成images，直接构建运行会遇到一个问题。    

```
Sending build context to Docker daemon 
Post http:///var/run/docker.sock/v1.19/build?cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=Dockerfile&memory=0&memswap=0&rm=1&t=ytxdockerhub%2Fpractice: dial unix /var/run/docker.sock: no such file or directory. Are you trying to connect to a TLS-enabled daemon without TLS?
Build step 'Execute shell' marked build as failure
Finished: FAILURE
```
这是因为jenkins在运行构建脚本时，没有读取到环境变量`DOCKER_TLS_VERIFY`和`DOCKER_HOST`和`DOCKER_CERT_PATH`。因此需要往系统中加入这三个环境变量。  
完整事例脚本如下：  

```
export DOCKER_TLS_VERIFY=1
export DOCKER_HOST=tcp://192.168.59.103:2376
export DOCKER_CERT_PATH=/Users/Shared/Jenkins/.boot2docker/certs/boot2docker-vm

docker build -t ytxdockerhub/chat .
```
这里有个疑问：  
在1安装docker的流程中有一步骤是将上述环境变量写入系统，但是Jenkins使用时还是提示没有获取到变量。  
自行写入`.zshrc`文件，`source`后，使用 `env`也显示环境变量写入成功。但是jenkins还是无法查找到变量。怪异～  
所以只能在写构建脚本的时候把上述环境变量加上去。有解决方式的请在此处补充。  
```
解决方式：  
目前只找到一种解决方式，即在【系统管理】---> 【系统设置】，手动添加变量
```

##### 3. Jenkins项目构建脚本——docker push
push操作需要`docker login`，现在我们使用的用户是mobileDev，即使该用户已经login了，jenkins依然是未登陆状态。所以我们需要将系统用户切换到jenkins，在该用户下进行登陆。不过该用户并没有安装docker，所以我们需要再次使用如下命令安装docker  

```
su jenkins  ①  # 密码已经修改为superwolf 
boot2docker init   ②
boot2docker start  ③
boot2docker shellinit ④
eval "$(boot2docker shellinit)" ⑤
```
然而在③执行过程中竟然需要输入docker的密码，不知为何。如下所示：

```
.docker@localhost's password:
odocker@localhost's password:
docker@localhost's password:
```
密码全部是`tcuser`  
最后使用`docker login`。构建脚本如下:

```
export DOCKER_TLS_VERIFY=1
export DOCKER_HOST=tcp://192.168.59.103:2376
export DOCKER_CERT_PATH=/Users/Shared/Jenkins/.boot2docker/certs/boot2docker-vm

docker build -t ytxdockerhub/chat .
docker push ytxdockerhub/chat
```
这里有个疑问：  
我尝试将安装在mobileDev的docker删除掉，使用`boot2docker delete`即可。然后只在jenkins用户下安装docker，安装成功后使用`docker images`等命令时会产生如下错误：  

```
An error occurred trying to connect: Get https://192.168.59.103:2376/v1.19/images/json: dial tcp 192.168.59.103:2376: i/o timeout
```
docker无法使用，虽然已经安装成功。最终只能在mobileDev中安装docker后，再在jenkins用户下再安装一遍。有解决方案，请贴出来
