## docker 查询容器、镜像、日志
```
	docker top <container> #显示容器内运行的进程   
	docker images #查询所有的镜像，默认是最近创建的排在最上。  
	docker ps #查看正在运行的容器  
	docker ps -l #查看最后退出的容器的ID  
	docker ps -a #查看所有的容器，包括退出的。  
	docker logs {容器ID|容器名称} #查询某个容器的所有操作记录。  
	docker logs -f {容器ID|容器名称} #实时查看容易的操作记录。
```

## docker 删除容器与镜像
```
	docker rm$(docker ps -a -q)  #删除所有容器
 	docker rm <containerName or id>  #删除单个容器
  	docker rmi id #删除单个容器
	docker rmi$(docker images | grep none | awk '${print $3}' | sort -r) #删除所有镜像	
```

## 启动停止容器

```
	docker stop <容器名or ID> #停止某个容器  
	docker start <容器名or ID> #启动某个容器  
	docker kill <容器名or ID> #杀掉某个容器
```

## docker 参数详解
```
	docker  
	useage of docker  
	-D 默认false 允许调试模式(debugmode)  
	-H 默认是unix:///var/run/docker.sock tcp://[host[:port]]来绑定 或者unix://[/path/to/socket]来使用(二进制文件的时候)，当主机ip host=[0.0.0.0],(端口)port=[4243] 或者 path=[/var/run/docker.sock]是缺省值，做为默认值来使用  
	-api-enable-cors 默认flase 允许CORS header远程api  
	-b 默认是空，附加在已存在的网桥上，如果是用'none'参数，就禁用了容器的网络  
	-bip 默认是空，使用提供的CIDR（ClasslessInter-Domain Routing-无类型域间选路）标记地址动态创建网桥(dcoker0),和-b参数冲突  
	-d 默认false 允许进程模式(daemonmode)  
	-dns 默认是空，使docker使用指定的DNS服务器  
	-g 默认是"/var/lib/docker":作为docker使用的根路径  
	-icc 默认true，允许inter-container来通信  
	-ip 默认"0.0.0.0"：绑定容器端口的默认Ip地址  
	-iptables 默认true 禁用docker添加iptables规则  
	-mtu 默认1500 : 设置容器网络传输的最大单元(mtu)  
	-p 默认是/var/run/docker.pid进程pid使用的文件路径  
	-r 默认是true 重启之前运行的容器  
	-s 默认是空 ，这个是docker运行是使用一个指定的存储驱动器  
	-v 默认false 打印版本信息和退出 
```

## docker run命令详解
```
	Usage: docker run [OPTIONS] IMAGE[:TAG] [COMMAND] 	[ARG...]  
	Run a command in a new container  
	-a=map[]: 附加标准输入、输出或者错误输出  
	-c=0: 共享CPU格式（相对重要）  
	-cidfile="": 将容器的ID标识写入文件  
	-d=false: 分离模式，在后台运行容器，并且打印出容器ID  
	-e=[]:设置环境变量  
	-h="": 容器的主机名称  
	-i=false: 保持输入流开放即使没有附加输入流  
	-privileged=false: 给容器扩展的权限  
	-m="": 内存限制 (格式:<number><optional unit>, unit单位 = b, k, m or g)  
	-n=true: 允许镜像使用网络  
	-p=[]: 匹配镜像内的网络端口号  
	-rm=false:当容器退出时自动删除容器 (不能跟 -d一起使用)  
	-t=false: 分配一个伪造的终端输入  
	-u="": 用户名或者ID  
	-dns=[]: 自定义容器的DNS服务器  
	-v=[]: 创建一个挂载绑定：[host-dir]:[container-dir]:[rw|ro].如果容器目录丢失，docker会创建一个新的卷  
	-volumes-from="": 挂载容器所有的卷  
	-entrypoint="": 覆盖镜像设置默认的入口点  
	-w="": 工作目录内的容器  
	-lxc-conf=[]: 添加自定义-lxc-	conf="lxc.cgroup.cpuset.cpus = 0,1" 
	-sig-proxy=true: 代理接收所有进程信号(even in non-tty mode)  
	-expose=[]: 让你主机没有开放的端口  
	-link="": 连接到另一个容器(name:alias)  
	-name="": 分配容器的名称，如果没有指定就会随机生成一个  
	-P=false: Publish all exposed ports to thehost interfaces 公布所有显示的端口主机接口
```

### Dockerizing a application
	docker run image cmd param
	exp:
		此处会打印出 hello world
		docker run ubuntu:14.04 /bin/echo 'hello world'
	
	docker run -t -i ubuntu:14.04 /bin/bash
		-t 在容器指定一个伪终端或者终端
		-i 进行命令交互
	ext:
		docker run -t -i ubuntu:14.04 /bin/echo
	
	让docker在后台运行：
	docker run -d ...
		-d: daemon
	docker ps 查询当前运行的container
	docker logs id docker打印的日志
	docker stop id 停止container
	
	打包image
	docker save ubuntu:14.04 > /root/unbuntu.rar
	
	载入image
	docker load < ubuntu.rar

### Image
	docker会先查找本地是否有image，若没有则从网上拉去image。	image格式为 name:tag
	可以事先拉image到本地，日后备用
	docker pull image
	
	docker提供image搜索
	docker search ...
	
	修改image有2种方式
	1. 使用container，commit the result 	
		docker run -t -i training/sinatra /bin/bash
		gem install json
		docker commit -m "added json gem" -a "kate smith" 0b2616 outuser/sinatra:v2
			-m: 提交信息
			-a: 提交作者
	2. 使用Dockerfile
		touch Dockerfile
		docker build -t ouruser/sinatra:v2
