# redis cluster实践
在本机通过手动的方式搭建cluster，理解cluster搭建流程
## 准备节点
准备搭建6个节点，为后续的3主3从做准备。由于是在本机搭建，因此采用端口进行区分即redis-7000.conf、redis-7001.conf、redis-7002.conf、redis-7003.conf、redis-7004.conf、redis-7005.conf

配置redis配置文件
```
port 7000
daemonize yes
pidfile /Users/seed/work_space/tools/redis-5.0.3/cluster-demo/redis_7000.pid
logfile "7000.log"
dbfilename dump-7000.rdb
dir "/Users/seed/work_space/tools/redis-5.0.3/cluster-demo"
cluster-enabled yes
cluster-config-file nodes-7000.conf
cluster-require-full-coverage no
```
其他配置文件只需要把7000修改即可
```bash
$ sed 's/7000/7001/g' redis-7000.conf > redis.7001.conf
$ sed 's/7000/7002/g' redis-7000.conf > redis.7002.conf
$ sed 's/7000/7003/g' redis-7000.conf > redis.7003.conf
$ sed 's/7000/7004/g' redis-7000.conf > redis.7004.conf
$ sed 's/7000/7005/g' redis-7000.conf > redis.7005.conf
```
然后启动各个节点
```bash
$ redis-server redis-7000.conf
$ redis-server redis-7001.conf
$ redis-server redis-7002.conf
$ redis-server redis-7003.conf
$ redis-server redis-7004.conf
$ redis-server redis-7005.conf
```

## 建立各个节点的通信
通过meet指令，让各个节点互相知道彼此的存在
```bash
$ redis-cli -p 7000 cluster meet 127.0.0.1 7001
$ redis-cli -p 7000 cluster meet 127.0.0.1 7002
$ redis-cli -p 7000 cluster meet 127.0.0.1 7003
$ redis-cli -p 7000 cluster meet 127.0.0.1 7004
$ redis-cli -p 7000 cluster meet 127.0.0.1 7005
```
通过cluster nodes和cluster info指令查看节点通信情况
```
$ redis-cli -p 7000 cluster nodes
$ redis-cli -p 7000 cluster info
$ redis-cli -p 7001 cluster nodes
$ redis-cli -p 7001 cluster info
$ redis-cli -p 7002 cluster nodes
$ redis-cli -p 7002 cluster info
$ redis-cli -p 7003 cluster nodes
$ redis-cli -p 7003 cluster info
$ redis-cli -p 7004 cluster nodes
$ redis-cli -p 7004 cluster info
$ redis-cli -p 7005 cluster nodes
$ redis-cli -p 7005 cluster info
```

## 分配槽
首先写一个分配槽的脚本addslot.sh
```shell
start=$1
end=$2
port=$3

for slot in `seq ${start} ${port}`
do
  echo "slot: ${slot}"
  redis-cli -p ${port} cluster addslots ${slot}
done
```
然后开始分配槽
```
$ sh addslot.sh 0 5461 7000
$ sh addslot.sh 5462 10922 7001
$ sh addslot.sh 10923 16383 7002
```

## 主从节点分配
```
$ redis-cli -p 7003 cluster replicate ${node0 id}
$ redis-cli -p 7004 cluster replicate ${node1 id}
$ redis-cli -p 7005 cluster replicate ${node2 id}
```