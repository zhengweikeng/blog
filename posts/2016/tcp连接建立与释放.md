### tcp建立连接
tcp连接的建立需要经历”三次握手“的过程。过程如下

1. client发送SYN包（值为j）以及SEQ包到server端，此时client进入SYN_SEND状态。此为第一次握手。
2. server端收到SYN包后，发送一个ACK（值为seq+1）确认包和SYN（值为k）给client，此时server进入SYN_RECV状态。此为第二次握手。
3. client收到SYN+ACK包后，向server发送一个ACK（值为k+1），该包发送完成后，client和server均进入ESTABLISH状态。此为第三次握手。

client和server两端状态变化如下：

client:  
CLOSED->SYN_SEND->ESTABLISH  
server:  
CLOSED->LISTEN->SYN_RECV->ESTABLISH

为什么是3次握手，不是2次或者4次呢？

建立一个可靠的tcp连接，至少是要3次的，既然3次就足够，为什么还要4次。  
那2次可以么。假设2次就建立连接，那么会出现如下情况。  
client发送一个A请求，由于网络原因，没有那么快送达server端。此时client便废弃该请求，重新发起请求B，请求B成功送达server端，并得到响应。  
如果此时便建立连接，若之前的请求A终于到达server端了，由于请求的格式都是正常的，于是也发送响应回去client，而client则认为该请求已经被废弃。于是server端就一直挂着这个请求，造成资源浪费。

因此，才需要让client再发送一次确认给server端，server接收确认后才建立连接，便可以避免这种情况了。

### tcp连接释放
Tcp释放连接的过程需要经历“四次挥手”的过程，为什么建立连接只需要3次握手，而释放连接需要进行4次挥手呢？

很简单，因为TCP连接是全双工（Full Duplex）的，因此造成了两个方向都需要进行关闭。

怎么理解呢？

client和server，需要关闭连接，此时client通知server我要关闭连接了，此时关闭的只会是client这一端的连接，而server端并未关闭，它依旧能够向client发送数据。

当然，关闭连接也可以是server作为主动方的。

接下来以client主动断开与server端的连接为场景来描述整个过程，我们把它分为两个阶段，分别为client端关闭连接和server端关闭连接。

先上图  
![tcp建立连接和释放连接](https://github.com/zhengweikeng/blog/blob/master/posts/2016/images/tcp%E5%BB%BA%E7%AB%8B%E8%BF%9E%E6%8E%A5%E5%92%8C%E9%87%8A%E6%94%BE%E8%BF%9E%E6%8E%A5.png?raw=true)

第一阶段  

1. 首先client会发送一个FIN包给server（同时还有ack和seq包），这是要告诉server，我已经没有数据要发给你了，此时client处于FIN_WAIT_1状态。接收到FIN包的server处于CLOSE_WAIT的状态。
2. server发回一个ACK（值为client传过来的seq+1）和seq（值为client传过来的ack的值）给client。client收到server发过来的包后确认关闭连接，此时client处于FIN_WAIT_2。

第二阶段  

1. server在接收到client的FIN后，得知client要断开tcp连接了，于是在发送完ack和seq给client后，自己发送一个FIN包给client（也带有ack和seq包），告诉client我也要断开连接了，此时server处于LAST_ACK状态。  
2. client接收到server的FIN信息后，会回复server一个ack包，并且会进入TIME_WAIT状态，持续2个MSL（Max Segment Lifetime），这个时间windows下为240s。而server接收到client后便关闭连接。client在指定时间过后仍然没有接收到server的数据，确认server已经没有数据过来，也关闭了连接。

此时双方都进入CLOSED状态。

### 分析一下

server在接收到client的FIN后，会返回ACK给client，此时关闭的就是读通道，也就是说不能再从这个连接中读信息了。

client收到server发送过来的对自己的FIN的确认包ack后，便关闭了写通道，不能再向连接中写信息了。

server也开始关闭连接，发送FIN给client，而client收到后便回复ack给server，同时关闭读通道，自己则进入TIME_WAIT状态。

server收到client发送过来的对自己的FIN的确认包ack后，便关闭了写通道，状态转化为CLOSED。

而client在TIME_WAIT结束后也进入CLOSED状态。

因此client和server会经历如下状态的转移  
client:
FIN_WAIT_1->FIN_WAIT_2->TIME_WAIT->CLOSED  
server:
CLOSE_WAIT->LAST_ACK->CLOSED

### 其中有几个关键点需要注意：

1. 这个TIME_WAIT的作用是什么？  
   这个[博客](http://www.cnblogs.com/Jessy/p/3535612.html)是这么解释的。
   ```
   原因有二：
  一、保证TCP协议的全双工连接能够可靠关闭
  二、保证这次连接的重复数据段从网络中消失

  先说第一点，如果Client直接CLOSED了，那么由于IP协议的不可靠性或者是其它网络原因，导致Server没有收到Client最后回复的ACK。那么Server就会在超时之后继续发送FIN。  
  此时由于Client已经CLOSED了，就找不到与重发的FIN对应的连接，最后Server就会收到RST而不是ACK，Server就会以为是连接错误把问题报告给高层。  
  这样的情况虽然不会造成数据丢失，但是却导致TCP协议不符合可靠连接的要求。  
  所以，Client不是直接进入CLOSED，而是要保持TIME_WAIT，当再次收到FIN的时候，能够保证对方收到ACK，最后正确的关闭连接。

  再说第二点，如果Client直接CLOSED，然后又再向Server发起一个新连接，我们不能保证这个新连接与刚关闭的连接的端口号是不同的。也就是说有可能新连接和老连接的端口号是相同的。  
  一般来说不会发生什么问题，但是还是有特殊情况出现：假设新连接和已经关闭的老连接端口号是一样的，如果前一次连接的某些数据仍然滞留在网络中，这些延迟数据在建立新连接之后才到达Server，由于新连接和老连接的端口号是一样的，又因为TCP协议判断不同连接的依据是socket pair。  
  于是，TCP协议就认为那个延迟的数据是属于新连接的，这样就和真正的新连接的数据包发生混淆了。所以TCP连接还要在TIME_WAIT状态等待2倍MSL，这样可以保证本次连接的所有数据都从网络中消失。  
  ```   

2. CLOSE_WAIT的解释。  
   在以上事例，我们知道server在接收到FIN后，发送ACK之前会进入CLOSE_WAIT，如果长期处于这个状态，或者说服务器出现大量CLOSE_WAIT，说明ACK包一直没有发出，这时候就应该检查代码了。 

3. TIME_WAIT注意事项  
   从事例我们知道，主动关闭连接的一方会经历TIME_WAIT状态，在该状态下的socket是不会被回收的。而如果是服务器端主动关闭连接，则可能会面临处于大量TIME_WAIT的情况（因为连接很多嘛），会严重影响服务器的处理能力。  
   怎么解决呢，那就减少服务器端TIME_WAIT的时间咯。

综上所述，client和server从建立连接到断开连接，整个状态的变化如下：

client:  
CLOSED->SYN_SEND->ESTABLISH->FIN_WAIT_1->FIN_WAIT_2->TIME_WAIT->CLOSED  
server:  
CLOSED->LISTEN->SYN_RECV->ESTABLISH->CLOSE_WAIT->LAST_ACK->CLOSED

### websocket
tcp是传输层的协议，tcp三次握手后，应用层协议http也便建立了连接。而对于当今web的发展情况，http仍有许多瓶颈。

1. 一条连接只能发送一个请求。
2. 请求只能从客户端开始。客户端不可以接收除响应以外的指令。
3. 请求/响应首部未经压缩发送，首部信息越多延迟越大。
4. 发送冗长的首部。每次互相发送相同的首部造成较多的浪费。
5. 可任意选择数据压缩格式。非强制压缩发送。

虽然已经出现了很多解决方案，如ajax、comet，但是他们最终使用的都是http协议，因此也无法从根本上解决这些瓶颈。  
因此也就诞生了一个新的通信协议,WebSocket协议，一种全双工通信协议。

该通信协议建立在http协议的基础之上，因此连接的发起方仍然是客户端，在http连接建立之后，再将协议升级为webSocket连接，在webSocket连接建立之后，客户度和服务器端都可以主动向对方发送报文信息了。

建立webSocket连接，需要先建立http连接，并在此基础上再进行一次”握手“。

client会发起一个”握手“的请求，请求首部含有upgrade:websocket（还有其他首部，具体看如下示例）。服务器端返回一个101状态码，确认转换协议。完成握手后便可以使用websocket协议进行通信。

Client（request）
```
GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: AQIDBAUGBwgJCgsMDQ4PEC==
Origin: http://example.com
Sec-WebSocket-protocol: chat, superchat
Sec-WebSocket-Version: 13
```
Sec-WebSocket-Key其值采用base64编码的随机16字节长的字符序列，服务器端根据该域来判断client确实是websocket请求而不是冒充的，如http
Sec-WebSocket-protocol使用的子协议
Sec-WebSocket-Version该值必须是13

Server（response）
```
HTTP/1.1 101 switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: dGhlIHNhbXBsZSBub25jZQ==
Sec-WebSocket-Protocol: chat
```
Sec-WebSocket-Accept值是由client请求首部中的Sec-WebSocket-Key做SHA-1 hash计算，然后再把得到的结果通过base64加密而来

连接建立后，通信的url格式如下：  
ws://example.com/  
wss://example.com/

#### 参考资料
1. [TCP的三次握手(建立连接）和四次挥手(关闭连接）](http://www.cnblogs.com/Jessy/p/3535612.html)
2. [服务器TIME_WAIT和CLOSE_WAIT详解和解决办法](http://www.cnblogs.com/sunxucool/p/3449068.html)
