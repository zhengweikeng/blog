- [Database](#database)
  - [Mysql](#mysql)
    - [MySQL 的存储引擎有哪些?（InnoDB）为什么选 InnoDB?](#mysql-%E7%9A%84%E5%AD%98%E5%82%A8%E5%BC%95%E6%93%8E%E6%9C%89%E5%93%AA%E4%BA%9Binnodb%E4%B8%BA%E4%BB%80%E4%B9%88%E9%80%89-innodb)
    - [知道mysql的索引算法么？](#%E7%9F%A5%E9%81%93mysql%E7%9A%84%E7%B4%A2%E5%BC%95%E7%AE%97%E6%B3%95%E4%B9%88)
    - [为什么mysql要用b+树而不是b树或者其他树？](#%E4%B8%BA%E4%BB%80%E4%B9%88mysql%E8%A6%81%E7%94%A8b%E6%A0%91%E8%80%8C%E4%B8%8D%E6%98%AFb%E6%A0%91%E6%88%96%E8%80%85%E5%85%B6%E4%BB%96%E6%A0%91)
    - [MySQL 的聚簇索引和非聚簇索引有什么区别?](#mysql-%E7%9A%84%E8%81%9A%E7%B0%87%E7%B4%A2%E5%BC%95%E5%92%8C%E9%9D%9E%E8%81%9A%E7%B0%87%E7%B4%A2%E5%BC%95%E6%9C%89%E4%BB%80%E4%B9%88%E5%8C%BA%E5%88%AB)
    - [聊聊事务的隔离级别。你们生产用的什么事务隔离级别？](#%E8%81%8A%E8%81%8A%E4%BA%8B%E5%8A%A1%E7%9A%84%E9%9A%94%E7%A6%BB%E7%BA%A7%E5%88%AB%E4%BD%A0%E4%BB%AC%E7%94%9F%E4%BA%A7%E7%94%A8%E7%9A%84%E4%BB%80%E4%B9%88%E4%BA%8B%E5%8A%A1%E9%9A%94%E7%A6%BB%E7%BA%A7%E5%88%AB)
    - [如何保证数据库的主从一致性](#%E5%A6%82%E4%BD%95%E4%BF%9D%E8%AF%81%E6%95%B0%E6%8D%AE%E5%BA%93%E7%9A%84%E4%B8%BB%E4%BB%8E%E4%B8%80%E8%87%B4%E6%80%A7)
    - [分库分表后怎么保证主键仍然是递增的?](#%E5%88%86%E5%BA%93%E5%88%86%E8%A1%A8%E5%90%8E%E6%80%8E%E4%B9%88%E4%BF%9D%E8%AF%81%E4%B8%BB%E9%94%AE%E4%BB%8D%E7%84%B6%E6%98%AF%E9%80%92%E5%A2%9E%E7%9A%84)
    - [瞬时写入量很大可能会打挂存储, 怎么保护?](#%E7%9E%AC%E6%97%B6%E5%86%99%E5%85%A5%E9%87%8F%E5%BE%88%E5%A4%A7%E5%8F%AF%E8%83%BD%E4%BC%9A%E6%89%93%E6%8C%82%E5%AD%98%E5%82%A8-%E6%80%8E%E4%B9%88%E4%BF%9D%E6%8A%A4)
    - [分库分表的数据源中假如存在主键冲突要怎么解决？](#%E5%88%86%E5%BA%93%E5%88%86%E8%A1%A8%E7%9A%84%E6%95%B0%E6%8D%AE%E6%BA%90%E4%B8%AD%E5%81%87%E5%A6%82%E5%AD%98%E5%9C%A8%E4%B8%BB%E9%94%AE%E5%86%B2%E7%AA%81%E8%A6%81%E6%80%8E%E4%B9%88%E8%A7%A3%E5%86%B3)
    - [怎么保证下游对 Binlog 的消费顺序？](#%E6%80%8E%E4%B9%88%E4%BF%9D%E8%AF%81%E4%B8%8B%E6%B8%B8%E5%AF%B9-binlog-%E7%9A%84%E6%B6%88%E8%B4%B9%E9%A1%BA%E5%BA%8F)
    - [如何在下游保证消费时的事务原子性？](#%E5%A6%82%E4%BD%95%E5%9C%A8%E4%B8%8B%E6%B8%B8%E4%BF%9D%E8%AF%81%E6%B6%88%E8%B4%B9%E6%97%B6%E7%9A%84%E4%BA%8B%E5%8A%A1%E5%8E%9F%E5%AD%90%E6%80%A7)
    - [分库分表后怎么查询分页?](#%E5%88%86%E5%BA%93%E5%88%86%E8%A1%A8%E5%90%8E%E6%80%8E%E4%B9%88%E6%9F%A5%E8%AF%A2%E5%88%86%E9%A1%B5)
    - [需要支持深分页, 页码直接跳转, 怎么实现?](#%E9%9C%80%E8%A6%81%E6%94%AF%E6%8C%81%E6%B7%B1%E5%88%86%E9%A1%B5-%E9%A1%B5%E7%A0%81%E7%9B%B4%E6%8E%A5%E8%B7%B3%E8%BD%AC-%E6%80%8E%E4%B9%88%E5%AE%9E%E7%8E%B0)
  - [Redis](#redis)
    - [redis的底层数据结构了解多少](#redis%E7%9A%84%E5%BA%95%E5%B1%82%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%BA%86%E8%A7%A3%E5%A4%9A%E5%B0%91)
    - [知道动态字符串sds的优缺点么？](#%E7%9F%A5%E9%81%93%E5%8A%A8%E6%80%81%E5%AD%97%E7%AC%A6%E4%B8%B2sds%E7%9A%84%E4%BC%98%E7%BC%BA%E7%82%B9%E4%B9%88)
    - [redis有哪些数据结构，分别使用在什么场景？](#redis%E6%9C%89%E5%93%AA%E4%BA%9B%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%88%86%E5%88%AB%E4%BD%BF%E7%94%A8%E5%9C%A8%E4%BB%80%E4%B9%88%E5%9C%BA%E6%99%AF)
    - [什么是缓存穿透？如何避免？什么是缓存雪崩？何如避免？](#%E4%BB%80%E4%B9%88%E6%98%AF%E7%BC%93%E5%AD%98%E7%A9%BF%E9%80%8F%E5%A6%82%E4%BD%95%E9%81%BF%E5%85%8D%E4%BB%80%E4%B9%88%E6%98%AF%E7%BC%93%E5%AD%98%E9%9B%AA%E5%B4%A9%E4%BD%95%E5%A6%82%E9%81%BF%E5%85%8D)
    - [redis分布式锁](#redis%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81)
    - [简述Redis分布式锁的缺陷？](#%E7%AE%80%E8%BF%B0redis%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E7%9A%84%E7%BC%BA%E9%99%B7)
    - [加锁机制，锁互斥机制，watch dog自动延期机制，可重入加锁机制，锁释放机制是什么？](#%E5%8A%A0%E9%94%81%E6%9C%BA%E5%88%B6%E9%94%81%E4%BA%92%E6%96%A5%E6%9C%BA%E5%88%B6watch-dog%E8%87%AA%E5%8A%A8%E5%BB%B6%E6%9C%9F%E6%9C%BA%E5%88%B6%E5%8F%AF%E9%87%8D%E5%85%A5%E5%8A%A0%E9%94%81%E6%9C%BA%E5%88%B6%E9%94%81%E9%87%8A%E6%94%BE%E6%9C%BA%E5%88%B6%E6%98%AF%E4%BB%80%E4%B9%88)
    - [Redis 的Setnx命令是如何实现分布式锁的？](#redis-%E7%9A%84setnx%E5%91%BD%E4%BB%A4%E6%98%AF%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E7%9A%84)
    - [说说对Setnx 的实现锁的原理的理解？](#%E8%AF%B4%E8%AF%B4%E5%AF%B9setnx-%E7%9A%84%E5%AE%9E%E7%8E%B0%E9%94%81%E7%9A%84%E5%8E%9F%E7%90%86%E7%9A%84%E7%90%86%E8%A7%A3)
    - [如何避免死锁的出现？](#%E5%A6%82%E4%BD%95%E9%81%BF%E5%85%8D%E6%AD%BB%E9%94%81%E7%9A%84%E5%87%BA%E7%8E%B0)
    - [Redis里面有1亿个key，其中有10w个key是以某个固定的已知的前缀开头的，如何将它们全部找出来？](#redis%E9%87%8C%E9%9D%A2%E6%9C%891%E4%BA%BF%E4%B8%AAkey%E5%85%B6%E4%B8%AD%E6%9C%8910w%E4%B8%AAkey%E6%98%AF%E4%BB%A5%E6%9F%90%E4%B8%AA%E5%9B%BA%E5%AE%9A%E7%9A%84%E5%B7%B2%E7%9F%A5%E7%9A%84%E5%89%8D%E7%BC%80%E5%BC%80%E5%A4%B4%E7%9A%84%E5%A6%82%E4%BD%95%E5%B0%86%E5%AE%83%E4%BB%AC%E5%85%A8%E9%83%A8%E6%89%BE%E5%87%BA%E6%9D%A5)
    - [如何使用redis实现队列。又如何实现延时队列。](#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8redis%E5%AE%9E%E7%8E%B0%E9%98%9F%E5%88%97%E5%8F%88%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E5%BB%B6%E6%97%B6%E9%98%9F%E5%88%97)
    - [如何实现持久化](#%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E6%8C%81%E4%B9%85%E5%8C%96)
    - [主从间的同步机制](#%E4%B8%BB%E4%BB%8E%E9%97%B4%E7%9A%84%E5%90%8C%E6%AD%A5%E6%9C%BA%E5%88%B6)
    - [说说Redis哈希槽的概念？](#%E8%AF%B4%E8%AF%B4redis%E5%93%88%E5%B8%8C%E6%A7%BD%E7%9A%84%E6%A6%82%E5%BF%B5)
    - [Redis集群的主从复制模型是怎样的？](#redis%E9%9B%86%E7%BE%A4%E7%9A%84%E4%B8%BB%E4%BB%8E%E5%A4%8D%E5%88%B6%E6%A8%A1%E5%9E%8B%E6%98%AF%E6%80%8E%E6%A0%B7%E7%9A%84)
    - [Redis集群会有写操作丢失吗？为什么？](#redis%E9%9B%86%E7%BE%A4%E4%BC%9A%E6%9C%89%E5%86%99%E6%93%8D%E4%BD%9C%E4%B8%A2%E5%A4%B1%E5%90%97%E4%B8%BA%E4%BB%80%E4%B9%88)
    - [Redis集群之间是如何复制的？](#redis%E9%9B%86%E7%BE%A4%E4%B9%8B%E9%97%B4%E6%98%AF%E5%A6%82%E4%BD%95%E5%A4%8D%E5%88%B6%E7%9A%84)
    - [Redis集群方案什么情况下会导致整个集群不可用？](#redis%E9%9B%86%E7%BE%A4%E6%96%B9%E6%A1%88%E4%BB%80%E4%B9%88%E6%83%85%E5%86%B5%E4%B8%8B%E4%BC%9A%E5%AF%BC%E8%87%B4%E6%95%B4%E4%B8%AA%E9%9B%86%E7%BE%A4%E4%B8%8D%E5%8F%AF%E7%94%A8)
    - [怎么理解Redis事务？](#%E6%80%8E%E4%B9%88%E7%90%86%E8%A7%A3redis%E4%BA%8B%E5%8A%A1)
    - [Redis如何做内存优化，如何回收进程？](#redis%E5%A6%82%E4%BD%95%E5%81%9A%E5%86%85%E5%AD%98%E4%BC%98%E5%8C%96%E5%A6%82%E4%BD%95%E5%9B%9E%E6%94%B6%E8%BF%9B%E7%A8%8B)
- [数据结构和算法](#%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%92%8C%E7%AE%97%E6%B3%95)
  - [数据结构](#%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)
    - [链表和数组的优缺点？](#%E9%93%BE%E8%A1%A8%E5%92%8C%E6%95%B0%E7%BB%84%E7%9A%84%E4%BC%98%E7%BC%BA%E7%82%B9)
    - [解决hash冲突的方法有哪些？](#%E8%A7%A3%E5%86%B3hash%E5%86%B2%E7%AA%81%E7%9A%84%E6%96%B9%E6%B3%95%E6%9C%89%E5%93%AA%E4%BA%9B)
  - [算法](#%E7%AE%97%E6%B3%95)
    - [从无限的字符流中, 随机选出 10 个字符](#%E4%BB%8E%E6%97%A0%E9%99%90%E7%9A%84%E5%AD%97%E7%AC%A6%E6%B5%81%E4%B8%AD-%E9%9A%8F%E6%9C%BA%E9%80%89%E5%87%BA-10-%E4%B8%AA%E5%AD%97%E7%AC%A6)
    - [M*N 横向纵向均递增的矩阵找指定数](#mn-%E6%A8%AA%E5%90%91%E7%BA%B5%E5%90%91%E5%9D%87%E9%80%92%E5%A2%9E%E7%9A%84%E7%9F%A9%E9%98%B5%E6%89%BE%E6%8C%87%E5%AE%9A%E6%95%B0)
    - [如何判断两个无环单链表有没有交叉点](#%E5%A6%82%E4%BD%95%E5%88%A4%E6%96%AD%E4%B8%A4%E4%B8%AA%E6%97%A0%E7%8E%AF%E5%8D%95%E9%93%BE%E8%A1%A8%E6%9C%89%E6%B2%A1%E6%9C%89%E4%BA%A4%E5%8F%89%E7%82%B9)
    - [如何判断两个有环单链表有没有交叉点](#%E5%A6%82%E4%BD%95%E5%88%A4%E6%96%AD%E4%B8%A4%E4%B8%AA%E6%9C%89%E7%8E%AF%E5%8D%95%E9%93%BE%E8%A1%A8%E6%9C%89%E6%B2%A1%E6%9C%89%E4%BA%A4%E5%8F%89%E7%82%B9)
- [网络](#%E7%BD%91%E7%BB%9C)
  - [网络基础](#%E7%BD%91%E7%BB%9C%E5%9F%BA%E7%A1%80)
    - [http1.0和1.1的区别？](#http10%E5%92%8C11%E7%9A%84%E5%8C%BA%E5%88%AB)
    - [描述下http2.0有哪些特性](#%E6%8F%8F%E8%BF%B0%E4%B8%8Bhttp20%E6%9C%89%E5%93%AA%E4%BA%9B%E7%89%B9%E6%80%A7)
    - [一个请求的过程（或者说从浏览器地址栏回车后发生的所有过程）](#%E4%B8%80%E4%B8%AA%E8%AF%B7%E6%B1%82%E7%9A%84%E8%BF%87%E7%A8%8B%E6%88%96%E8%80%85%E8%AF%B4%E4%BB%8E%E6%B5%8F%E8%A7%88%E5%99%A8%E5%9C%B0%E5%9D%80%E6%A0%8F%E5%9B%9E%E8%BD%A6%E5%90%8E%E5%8F%91%E7%94%9F%E7%9A%84%E6%89%80%E6%9C%89%E8%BF%87%E7%A8%8B)
    - [http请求包含哪些数据结构？](#http%E8%AF%B7%E6%B1%82%E5%8C%85%E5%90%AB%E5%93%AA%E4%BA%9B%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)
    - [什么是http的长连接和短连接？如何配置使用长连接](#%E4%BB%80%E4%B9%88%E6%98%AFhttp%E7%9A%84%E9%95%BF%E8%BF%9E%E6%8E%A5%E5%92%8C%E7%9F%AD%E8%BF%9E%E6%8E%A5%E5%A6%82%E4%BD%95%E9%85%8D%E7%BD%AE%E4%BD%BF%E7%94%A8%E9%95%BF%E8%BF%9E%E6%8E%A5)
    - [TCP的keep-alive的作用？它和http的keep-alive有什么差别](#tcp%E7%9A%84keep-alive%E7%9A%84%E4%BD%9C%E7%94%A8%E5%AE%83%E5%92%8Chttp%E7%9A%84keep-alive%E6%9C%89%E4%BB%80%E4%B9%88%E5%B7%AE%E5%88%AB)
    - [谈谈tcp的三次握手和四次挥手。为什么建立连接需要三次，而不是两次？](#%E8%B0%88%E8%B0%88tcp%E7%9A%84%E4%B8%89%E6%AC%A1%E6%8F%A1%E6%89%8B%E5%92%8C%E5%9B%9B%E6%AC%A1%E6%8C%A5%E6%89%8B%E4%B8%BA%E4%BB%80%E4%B9%88%E5%BB%BA%E7%AB%8B%E8%BF%9E%E6%8E%A5%E9%9C%80%E8%A6%81%E4%B8%89%E6%AC%A1%E8%80%8C%E4%B8%8D%E6%98%AF%E4%B8%A4%E6%AC%A1)
    - [tcp有哪些状态，相应状态的含义。](#tcp%E6%9C%89%E5%93%AA%E4%BA%9B%E7%8A%B6%E6%80%81%E7%9B%B8%E5%BA%94%E7%8A%B6%E6%80%81%E7%9A%84%E5%90%AB%E4%B9%89)
    - [如果服务端没有收到最后的ack包，客户端可以开始发数据么？](#%E5%A6%82%E6%9E%9C%E6%9C%8D%E5%8A%A1%E7%AB%AF%E6%B2%A1%E6%9C%89%E6%94%B6%E5%88%B0%E6%9C%80%E5%90%8E%E7%9A%84ack%E5%8C%85%E5%AE%A2%E6%88%B7%E7%AB%AF%E5%8F%AF%E4%BB%A5%E5%BC%80%E5%A7%8B%E5%8F%91%E6%95%B0%E6%8D%AE%E4%B9%88)
    - [为什么接收方在FIN包后不能一次性发送ACK和FIN包给发送方，就像建立连接时一次性发送SYN和ACK包一样。](#%E4%B8%BA%E4%BB%80%E4%B9%88%E6%8E%A5%E6%94%B6%E6%96%B9%E5%9C%A8fin%E5%8C%85%E5%90%8E%E4%B8%8D%E8%83%BD%E4%B8%80%E6%AC%A1%E6%80%A7%E5%8F%91%E9%80%81ack%E5%92%8Cfin%E5%8C%85%E7%BB%99%E5%8F%91%E9%80%81%E6%96%B9%E5%B0%B1%E5%83%8F%E5%BB%BA%E7%AB%8B%E8%BF%9E%E6%8E%A5%E6%97%B6%E4%B8%80%E6%AC%A1%E6%80%A7%E5%8F%91%E9%80%81syn%E5%92%8Cack%E5%8C%85%E4%B8%80%E6%A0%B7)
    - [如果大量出现CLOSE_WAIT状态，说明什么？](#%E5%A6%82%E6%9E%9C%E5%A4%A7%E9%87%8F%E5%87%BA%E7%8E%B0closewait%E7%8A%B6%E6%80%81%E8%AF%B4%E6%98%8E%E4%BB%80%E4%B9%88)
    - [TIME_WAIT的作用？以及出现大量TIME_WAIT的原因。](#timewait%E7%9A%84%E4%BD%9C%E7%94%A8%E4%BB%A5%E5%8F%8A%E5%87%BA%E7%8E%B0%E5%A4%A7%E9%87%8Ftimewait%E7%9A%84%E5%8E%9F%E5%9B%A0)
    - [如何优化time_wait状态的发生？](#%E5%A6%82%E4%BD%95%E4%BC%98%E5%8C%96timewait%E7%8A%B6%E6%80%81%E7%9A%84%E5%8F%91%E7%94%9F)
    - [如果被断开的一方在收到FIN包后就跑路或者回复完ACK就跑路了，会怎么样？](#%E5%A6%82%E6%9E%9C%E8%A2%AB%E6%96%AD%E5%BC%80%E7%9A%84%E4%B8%80%E6%96%B9%E5%9C%A8%E6%94%B6%E5%88%B0fin%E5%8C%85%E5%90%8E%E5%B0%B1%E8%B7%91%E8%B7%AF%E6%88%96%E8%80%85%E5%9B%9E%E5%A4%8D%E5%AE%8Cack%E5%B0%B1%E8%B7%91%E8%B7%AF%E4%BA%86%E4%BC%9A%E6%80%8E%E4%B9%88%E6%A0%B7)
    - [如果出现大量的LAST_ACK状态，说明什么原因？](#%E5%A6%82%E6%9E%9C%E5%87%BA%E7%8E%B0%E5%A4%A7%E9%87%8F%E7%9A%84lastack%E7%8A%B6%E6%80%81%E8%AF%B4%E6%98%8E%E4%BB%80%E4%B9%88%E5%8E%9F%E5%9B%A0)
    - [301和302有什么区别](#301%E5%92%8C302%E6%9C%89%E4%BB%80%E4%B9%88%E5%8C%BA%E5%88%AB)
    - [504和500有什么区别](#504%E5%92%8C500%E6%9C%89%E4%BB%80%E4%B9%88%E5%8C%BA%E5%88%AB)
- [架构](#%E6%9E%B6%E6%9E%84)
  - [微服务](#%E5%BE%AE%E6%9C%8D%E5%8A%A1)
    - [微服务数据一致性问题，如何解决？](#%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%95%B0%E6%8D%AE%E4%B8%80%E8%87%B4%E6%80%A7%E9%97%AE%E9%A2%98%E5%A6%82%E4%BD%95%E8%A7%A3%E5%86%B3)
  - [分布式](#%E5%88%86%E5%B8%83%E5%BC%8F)
    - [分布式系统的唯一id生成算法](#%E5%88%86%E5%B8%83%E5%BC%8F%E7%B3%BB%E7%BB%9F%E7%9A%84%E5%94%AF%E4%B8%80id%E7%94%9F%E6%88%90%E7%AE%97%E6%B3%95)
    - [某一个业务中现在需要生成全局唯一的递增 ID, 并发量非常大, 怎么做](#%E6%9F%90%E4%B8%80%E4%B8%AA%E4%B8%9A%E5%8A%A1%E4%B8%AD%E7%8E%B0%E5%9C%A8%E9%9C%80%E8%A6%81%E7%94%9F%E6%88%90%E5%85%A8%E5%B1%80%E5%94%AF%E4%B8%80%E7%9A%84%E9%80%92%E5%A2%9E-id-%E5%B9%B6%E5%8F%91%E9%87%8F%E9%9D%9E%E5%B8%B8%E5%A4%A7-%E6%80%8E%E4%B9%88%E5%81%9A)
    - [我现在要做一个限流功能, 怎么做?](#%E6%88%91%E7%8E%B0%E5%9C%A8%E8%A6%81%E5%81%9A%E4%B8%80%E4%B8%AA%E9%99%90%E6%B5%81%E5%8A%9F%E8%83%BD-%E6%80%8E%E4%B9%88%E5%81%9A)
    - [这个限流要做成分布式的, 怎么做?](#%E8%BF%99%E4%B8%AA%E9%99%90%E6%B5%81%E8%A6%81%E5%81%9A%E6%88%90%E5%88%86%E5%B8%83%E5%BC%8F%E7%9A%84-%E6%80%8E%E4%B9%88%E5%81%9A)
    - [分布式锁设置超时后，有没可能在没有释放的情况下, 被人抢走锁。有的话，怎么解决？](#%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E8%AE%BE%E7%BD%AE%E8%B6%85%E6%97%B6%E5%90%8E%E6%9C%89%E6%B2%A1%E5%8F%AF%E8%83%BD%E5%9C%A8%E6%B2%A1%E6%9C%89%E9%87%8A%E6%94%BE%E7%9A%84%E6%83%85%E5%86%B5%E4%B8%8B-%E8%A2%AB%E4%BA%BA%E6%8A%A2%E8%B5%B0%E9%94%81%E6%9C%89%E7%9A%84%E8%AF%9D%E6%80%8E%E4%B9%88%E8%A7%A3%E5%86%B3)
    - [不用zk的心跳, 可以怎么解决这个问题呢?](#%E4%B8%8D%E7%94%A8zk%E7%9A%84%E5%BF%83%E8%B7%B3-%E5%8F%AF%E4%BB%A5%E6%80%8E%E4%B9%88%E8%A7%A3%E5%86%B3%E8%BF%99%E4%B8%AA%E9%97%AE%E9%A2%98%E5%91%A2)
  - [并发](#%E5%B9%B6%E5%8F%91)
    - [CAS](#cas)
- [语言](#%E8%AF%AD%E8%A8%80)
  - [golang](#golang)
    - [如何实现CAS。](#%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0cas)
    - [关于golang for-range 里 goroutine 闭包捕获的代码](#%E5%85%B3%E4%BA%8Egolang-for-range-%E9%87%8C-goroutine-%E9%97%AD%E5%8C%85%E6%8D%95%E8%8E%B7%E7%9A%84%E4%BB%A3%E7%A0%81)
    - [goroutine 是怎么调度的？](#goroutine-%E6%98%AF%E6%80%8E%E4%B9%88%E8%B0%83%E5%BA%A6%E7%9A%84)
    - [goroutine 和 kernel thread 之间是什么关系？](#goroutine-%E5%92%8C-kernel-thread-%E4%B9%8B%E9%97%B4%E6%98%AF%E4%BB%80%E4%B9%88%E5%85%B3%E7%B3%BB)
    - [golang 的 gc 算法](#golang-%E7%9A%84-gc-%E7%AE%97%E6%B3%95)
    - [Golang 里的逃逸分析是什么？怎么避免内存逃逸？](#golang-%E9%87%8C%E7%9A%84%E9%80%83%E9%80%B8%E5%88%86%E6%9E%90%E6%98%AF%E4%BB%80%E4%B9%88%E6%80%8E%E4%B9%88%E9%81%BF%E5%85%8D%E5%86%85%E5%AD%98%E9%80%83%E9%80%B8)
  - [node.js](#nodejs)

# Database
## Mysql
### MySQL 的存储引擎有哪些?（InnoDB）为什么选 InnoDB?
Memory、MyISAM、InnoDB。  
选择InnoDB:
1. 支持事务
2. 具有聚集索引
3. MVCC
4. 更完善的奔溃恢复，借助redo log和bin log，能实现更细粒度的数据恢复，基本能够恢复任意时刻的数据。

### 知道mysql的索引算法么？
B+树

### 为什么mysql要用b+树而不是b树或者其他树？
B+树是多叉树，深度更小，B+树可以对叶子节点进行顺序遍历，B+树能够更好地利用磁盘扇区；二叉树：实现简单

### MySQL 的聚簇索引和非聚簇索引有什么区别?
1. 聚簇索引的叶子节点是数据节点，例如mysql里定义了主键后就会有主键索引。
2. 非聚簇索引叶子节点是指向数据块的指针。mysql里是存储的是主键索引id。因此如果要查询的数据不在非聚集索引中的话，就得通过这个id做回表查询，即回到主键索引根据id查询。

### 聊聊事务的隔离级别。你们生产用的什么事务隔离级别？

### 如何保证数据库的主从一致性

### 分库分表后怎么保证主键仍然是递增的?
 TDDL的办法：有一张专门用于分配主键的表，每次用乐观锁的方式尝试去取一批主键过来分配，假如乐观锁失败就重试

### 瞬时写入量很大可能会打挂存储, 怎么保护?
断路器

### 分库分表的数据源中假如存在主键冲突要怎么解决？

### 怎么保证下游对 Binlog 的消费顺序？

### 如何在下游保证消费时的事务原子性？

### 分库分表后怎么查询分页?

### 需要支持深分页, 页码直接跳转, 怎么实现?

## Redis
[redis题目](https://blog.csdn.net/u010682330/article/details/81043419)

### redis的底层数据结构了解多少
最上层统一的头部
```c
struct RedisObject {
    int4 type;
    int4 encoding;
    int24 lru;
    int32 refcount;
    void *ptr;
}
```
ptr是个指针，指向具体的数据。

字符串采用动态字符串实现，数据结构为SDS，通过预先分配一个容量减少内存的频繁分配，也记录了字符串的长度大小。
```c
stuct SDS<T> {
    T capacity 
    T len
    byte flags
    byte[] content
}
```
而根据字符串的长度，采用了不同的方式存储字符串：
1. embstr，在字符串长度小于44字节时，使用这种格式存储。优点就是它只需要分配一次空间，redisObject和sds是连续的。缺点就是长度增加导致需要重新分配内存时，整个redisObject和sds都要重新分配。
   ![WX20190226-153702@2x](./images/WX20190226-161136@2x.png)
2. raw，字符串长度大于44字节时，使用这种格式。优点就是在重新分配内存的时候只需要分配sds的。缺点就是它需要分配两次内存空间，分别要为redisObject和sds分配空间。
   ![WX20190226-153702@2x](./images/WX20190226-161317@2x.png)

列表：
1. 在元素较少的时候（默认是512，redis.conf可配置），内部采用的是压缩列表（ziplist），通过分配一块连续的内存，将所有元素紧挨着一起存储。也因此导致了在扩容的时候，可能需要重新分配内存，将旧的内容拷贝到新的内存地址，在元素很多时就会损耗性能，所以ziplist不合适存储大型字符串。
    ```c
    stuct ziplist<T> {
        int32 zlbyts; // 整个压缩列表占用字节数
        int32 zltail_offset; // 最后一个元素距离列表起始位置的偏移量，用于快速定位最后一个节点，实现双向遍历。
        int16 zllength; // 元素个数
        T[]   entries; // 元素内容列表，挨个挨个紧凑存储
        int8  zlend;  // 标志压缩列表的结束，值恒为 0xFF
    }

    struct entry {
        int<var> prevlen; // 前一个 entry 的字节长度 当长度小于254字节，用1个字节存储；否则，用5个字节存储
        int<var> encoding; // 元素类型编码
        optional byte[] content; // 元素内容
    }
    ```
    ![ziplist](./images/WX20190226-163238@2x.png)
2. 在元素较多的时候，为了能够快速插入和删除，采用了链表（lintedlist）的方式实现。每个节点都需要存储前一个和后一个节点的地址，每一个指针需要占据8个字节（64位机器）。
    ```c
    struct listNode<T> {
        listNode* prev;
        listNode* next;
        T value;
    }
    struct list {
        listNode *head;
        listNode *tail;
        long length;
    }
    ```
    ![WX20190226-163333@2x](./images/WX20190226-163333@2x.png)
3. 由于链表中每个节点中占据的空间较大（16个字节的指针空间），因此在后续的版本中，实现上被改为了快速列表（quicklist）。将linkedlist按段切分，每一段使用ziplist来紧凑存储，多个ziplist之间使用双向指针串起来。而且还可以对ziplist使用LZF算法进行了压缩存储。
    ```c
    struct ziplist_compressed {
        int32 size;
        byte[] compressed_data;
    }
    struct quicklistNode {
        quicklistNode* prev;
        quicklistNode* next;
        ziplist* zl; // 指向压缩列表 每个ziplist长度为8k，超出后就需要新起一个ziplist，该配置由list-max-ziplist-size决定。
        int32 size; // ziplist 的字节总数
        int16 count; // ziplist 中的元素数量
        int2 encoding; // 存储形式 2bit，原生字节数组还是 LZF 压缩存储。
        ...
    }
    struct quicklist {
        quicklistNode* head;
        quicklistNode* tail;
        long count; // 元素总数
        int nodes; // ziplist 节点的个数
        int compressDepth; // LZF 算法压缩深度。默认为0，即不压缩。通过list-compress-depth配置。
        ...
    }
    ```

hash哈希：
1. 在元素个数小于512个（set-max-intset-entries配置）并且每个元素长度小于64字节，采用ziplist实现
    ```
    hset profile name "tome"
    hset profile age 25
    hset profile career "Programmer"
    ```
    ![WX20190226-162949@2x](./images/WX20190226-162949@2x.png)
1. 其他情况下采用hashtable实现。

字典中包含2个hashtable，通常其中一个是有值的，在扩容的时候需要分配新的hashtable，然后渐进式搬迁，一个hashtable用于存储旧值，一个用于存储新值。搬迁结束后，旧的hashtable会被删除。   
```c
struct dict {
    ...
    dictht ht[2];
}

struct dictht {
    dictEntry** table; // 二维
    long size; // 第一维数组的长度
    long used; // hash表中的元素个数
    ...
}

struct dictEntry {
    void* key;
    void* val;
    dictEntry* next; // 链接下一个entry
}
```
字典基于二维结构设计，其中第一维是数组，第二维是链表。数组中存储的是第二维链表中第一个元素的指针。  

关于字典的扩容：  
正常情况下，当hash表中元素的个数等于第一维数组的长度时，就会开始扩容，扩容的新数组是原数组大小的2倍。不过如果Redis正在做bgsave，为了减少内存页的过多分离 (Copy On Write)，Redis尽量不去扩容 (dict_can_resize)，但是如果hash表已经非常满了，元素的个数已经达到了第一维数组长度的5倍 (dict_force_resize_ratio)，说明hash表已经过于拥挤了，这个时候就会强制扩容。

关于字典的缩容：   
当 hash 表因为元素的逐渐删除变得越来越稀疏时，Redis 会对 hash 表进行缩容来减少 hash 表的第一维数组空间占用。缩容的条件是元素个数低于数组长度的 10%。缩容不会考虑 Redis 是否正在做 bgsave。

Set集合：
1. 当set集合容纳的元素都是整数并且元素较少的时候，内部采用intset来存储集合元素，它是一种紧凑的数组结构，同时支持16位、32位和64位整数
    ```c
    struct intset<T> {
        int32 encoding; // 决定整数位宽是16位、32位还是64位
        int32 length;  // 元素个数
        int<T> contents; // 整数数组，可以是16为、32位和64位
    }
    ```
    ![intset](./images/WX20190226-164057@2x.png)
2. 当set集合存储的不是整数时，采用了hash结构进行存储，只是里面的value都是NULL，其他特性和字典一模一样。

SortedSet：
1. 采用dict字典存储value和score值的映射关系
2. 元素数量小于128，元素长度都小于64字节，使用ziplist实现
    ![WX20190226-171214@2x](./images/WX20190226-171214@2x.png)
3. 采用跳跃表skiplist来作为存储score的数据结构。并且在skiplist的forward指针上增加了一个span属性，用于表示从前一个节点沿着当前层的forward指针跳到当前这个节点中间会跳过多少节点。
    ```c
    struct zset {
        dict *dict;
        zskiplist *zsl;
    }
    struct zslforward {
        zslnode* item;
        long span;  // 跨度
    }
    struct zslnode {
        String value;
        double score;
        zslforward*[] forwards;  // 多层连接指针
        zslnode* backward;  // 回溯指针
    }
    struct zskiplist {
        zslnode* header; // 跳跃列表头指针
        int maxLevel; // 跳跃列表当前的最高层
        map<string, zslnode*> ht; // hash 结构的所有键值对
    }
    ```
    ![skiplist](./images/WX20190226-153702@2x.png)

[redis的五大数据类型实现原理](https://www.cnblogs.com/ysocean/p/9102811.html#_label0_1)

### 知道动态字符串sds的优缺点么？
优点：
1. sds结构会直接存储字符串长度，不需要遍历字符串就能得到长度。
2. 由于申请字符串空间的时候，会通过capacity多申请一些冗余空间，因此在执行append的时候，如果字符串长度不大，可以直接在原数组上直接进行，不需要重新分配空间。
3. 可以根据当前字符串的大小，定义len和capacity的字段大小，对内存使用得到优化
4. 根据字符串的长短采用embstr和raw结构来存储，提高性能。

缺点：
1. 基于优点的第二点，如果append的字符串很大，就需要重新分配空间，并且做字符串复制迁移，这个开销非常大。

### redis有哪些数据结构，分别使用在什么场景？
string、list、hash、sort、sortedset   
bitmap、geo、Pub/Sub  
hyperloglog、bloomfilter

### 什么是缓存穿透？如何避免？什么是缓存雪崩？何如避免？
[缓存](https://github.com/zhengweikeng/blog/blob/master/posts/2018/redis/%E4%BC%98%E5%8C%96.md)

### redis分布式锁

### 简述Redis分布式锁的缺陷？

### 加锁机制，锁互斥机制，watch dog自动延期机制，可重入加锁机制，锁释放机制是什么？

### Redis 的Setnx命令是如何实现分布式锁的？

### 说说对Setnx 的实现锁的原理的理解？

### 如何避免死锁的出现？

### Redis里面有1亿个key，其中有10w个key是以某个固定的已知的前缀开头的，如何将它们全部找出来？
keys和scan

### 如何使用redis实现队列。又如何实现延时队列。
[队列和延时队列](https://juejin.im/book/5afc2e5f6fb9a07a9b362527/section/5afc3643518825672034404b)

### 如何实现持久化

### 主从间的同步机制

### 说说Redis哈希槽的概念？
Redis集群没有使用一致性hash,而是引入了哈希槽的概念，Redis集群有16384个哈希槽，每个key通过CRC16校验后对16384取模来决定放置哪个槽，集群的每个节点负责一部分hash槽。

### Redis集群的主从复制模型是怎样的？
为了使在部分节点失败或者大部分节点无法通信的情况下集群仍然可用，所以集群使用了主从复制模型,每个节点都会有N-1个复制品。

### Redis集群会有写操作丢失吗？为什么？
Redis并不能保证数据的强一致性，这意味这在实际中集群在特定的条件下可能会丢失写操作。

### Redis集群之间是如何复制的？
异步复制

### Redis集群方案什么情况下会导致整个集群不可用？
有A，B，C三个节点的集群,在没有复制模型的情况下,如果节点B失败了，那么整个集群就会以为缺少5501-11000这个范围的槽而不可用。

### 怎么理解Redis事务？

### Redis如何做内存优化，如何回收进程？
redis基于对象引用计数来进行垃圾回收管理。
```c
typedef struct redisObject{
    int4 type;
    int4 encoding;
    int24 lru;
    //引用计数
    int32 refcount;
    void *ptr;
}
```
这里有个refcount，用于计算对象的引用计数
1. 创建一个新对象，属性 refcount 初始化为1
2. 对象被一个新程序使用，属性 refcount 加1
3. 对象不再被一个程序使用，属性 refcount 减1
4. 当对象的引用计数值变为0时，对象所占用的内存就会被释放

对象引用计数容易造成内存泄漏（例如A引用B，B引用C，C引用A），redis提供了几种策略来进行内存的优化和释放  
[maxmemory-policy](https://github.com/zhengweikeng/blog/blob/master/posts/2018/redis/%E5%86%85%E5%AD%98%E7%AE%A1%E7%90%86.md#maxmemory-policy)

另外refcount也会用于优化内存。
```
set k1 100
set k2 100
```
此时，redis只会创建一个redisObject对象，此时refcount=2，达到共享的作用。  
![WX20190226-174426@2x](./images/WX20190226-174426@2x.png)

# 数据结构和算法
## 数据结构
### 链表和数组的优缺点？

### 解决hash冲突的方法有哪些？

## 算法
### 从无限的字符流中, 随机选出 10 个字符
[蓄水池采样算法](https://github.com/aylei/interview)

### M*N 横向纵向均递增的矩阵找指定数
[search-a-2d-matrix-ii](https://leetcode.com/problems/search-a-2d-matrix-ii/)

### 如何判断两个无环单链表有没有交叉点

### 如何判断两个有环单链表有没有交叉点

# 网络
## 网络基础
### http1.0和1.1的区别？

### 描述下http2.0有哪些特性

### 一个请求的过程（或者说从浏览器地址栏回车后发生的所有过程）

### http请求包含哪些数据结构？

### 什么是http的长连接和短连接？如何配置使用长连接

### TCP的keep-alive的作用？它和http的keep-alive有什么差别

### 谈谈tcp的三次握手和四次挥手。为什么建立连接需要三次，而不是两次？

### tcp有哪些状态，相应状态的含义。

### 如果服务端没有收到最后的ack包，客户端可以开始发数据么？

### 为什么接收方在FIN包后不能一次性发送ACK和FIN包给发送方，就像建立连接时一次性发送SYN和ACK包一样。

### 如果大量出现CLOSE_WAIT状态，说明什么？

### TIME_WAIT的作用？以及出现大量TIME_WAIT的原因。

### 如何优化time_wait状态的发生？

### 如果被断开的一方在收到FIN包后就跑路或者回复完ACK就跑路了，会怎么样？

### 如果出现大量的LAST_ACK状态，说明什么原因？

### 301和302有什么区别

### 504和500有什么区别

# 架构
## 微服务
### 微服务数据一致性问题，如何解决？

## 分布式
### 分布式系统的唯一id生成算法
[分布式系统的唯一id生成算法你了解吗？](https://juejin.im/post/5c6be4086fb9a04a060570df)

### 某一个业务中现在需要生成全局唯一的递增 ID, 并发量非常大, 怎么做
TDDL 那样一次取一个 ID 段，放在本地慢慢分配的策略

### 我现在要做一个限流功能, 怎么做?
令牌桶

### 这个限流要做成分布式的, 怎么做?
令牌桶维护到 Redis 里，每个实例起一个线程抢锁，抢到锁的负责定时放令牌

### 分布式锁设置超时后，有没可能在没有释放的情况下, 被人抢走锁。有的话，怎么解决？
有可能，单次处理时间过长，锁泄露。换zk，用心跳解决

### 不用zk的心跳, 可以怎么解决这个问题呢?
每次更新过期时间时，Redis用MULTI做check-and-set检查更新时间是否被其他线程修改了，假如被修改了，说明锁已经被抢走，放弃这把锁。

## 并发
### CAS
Compare and Swap，一种乐观锁的实现，简单来说就是不通过加锁的方式来解决并发情况下对共享变量的访问和修改。

# 语言
## golang
### 如何实现CAS。

### 关于golang for-range 里 goroutine 闭包捕获的代码
[go语言坑之for range](https://studygolang.com/articles/9701)

### goroutine 是怎么调度的？

### goroutine 和 kernel thread 之间是什么关系？

### golang 的 gc 算法

### Golang 里的逃逸分析是什么？怎么避免内存逃逸？

## node.js