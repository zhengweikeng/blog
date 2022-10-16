第一个指令：ping  
ping也属于一个通信协议，是TCP/IP协议的一部分。利用“ping”命令可以检查网络是否连通，可以很好地帮助我们分析和判定网络故障。  
不打算将ping怎么使用，因为这些网络上都有资料，而且可以自行`man`，而是来讲讲ping所返回的数据

###### icmp_seq ＝ 1
ICMP，即Internet Control Message Protocol, Internet控制报文协议，这里icmp_sqp=1表示接收到的来自xxx地址的第1个icmp报文

###### ttl
即Time To Live的缩写，该字段指定IP包被路由器丢弃之前允许通过的最大网段数量。TTL是IPv4包头的一个8 bit字段。  
作用是限制IP数据包在计算机网络中的存在的时间。TTL的最大值是255，TTL的一个推荐值是64。TTL 是由发送主机设置的，以防止数据包不断在 IP 互联网络上永不终止地循环。转发 IP 数据包时，要求路由器至少将 TTL 减小 1

###### rtt min/avg/max/mdev
rtt，即Round-Trip Time往返时延，表示从发送端发送数据开始，到发送端收到来自接收端的确认（接收端收到数据后便立即发送确认），总共经历的时延。  
各部分分别为：最小/平均/最大/算术平均差  
算数平均差越大，说明偏离平均值越大。

###### Request Timed Out
表示对方主机可以到达到TIME OUT，这种情况通常是为对方拒绝接收你发给它的数据包造成数据包丢失。大多数的原因可能是对方装有防火墙或已下线

###### Destination Net Unreachable
对方主机不存在或者没有跟对方建立连接  
"destination host unreachable"和"time out"的区别，如果所经过的路由器的路由表中具有到达目标的路由，而目标因为其它原因不可到达，这时候会出现"time out"，如果路由表中连到达目标的路由都没有，那就会出现"destination host unreachable"。

###### Bad IP address
可能没有连接到DNS服务器所以无法解析这个IP地址，也可能是IP地址不存在

###### Source quench received
表示对方或中途的服务器繁忙无法回应

