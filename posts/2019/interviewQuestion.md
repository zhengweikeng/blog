- [Database](#database)
  - [Mysql](#mysql)
    - [知道mysql的索引算法么？](#%E7%9F%A5%E9%81%93mysql%E7%9A%84%E7%B4%A2%E5%BC%95%E7%AE%97%E6%B3%95%E4%B9%88)
    - [为什么mysql要用b+树而不是b树或者其他树？](#%E4%B8%BA%E4%BB%80%E4%B9%88mysql%E8%A6%81%E7%94%A8b%E6%A0%91%E8%80%8C%E4%B8%8D%E6%98%AFb%E6%A0%91%E6%88%96%E8%80%85%E5%85%B6%E4%BB%96%E6%A0%91)
    - [聊聊事务的隔离级别。你们生产用的什么事务隔离级别？](#%E8%81%8A%E8%81%8A%E4%BA%8B%E5%8A%A1%E7%9A%84%E9%9A%94%E7%A6%BB%E7%BA%A7%E5%88%AB%E4%BD%A0%E4%BB%AC%E7%94%9F%E4%BA%A7%E7%94%A8%E7%9A%84%E4%BB%80%E4%B9%88%E4%BA%8B%E5%8A%A1%E9%9A%94%E7%A6%BB%E7%BA%A7%E5%88%AB)
    - [如何保证数据库的主从一致性](#%E5%A6%82%E4%BD%95%E4%BF%9D%E8%AF%81%E6%95%B0%E6%8D%AE%E5%BA%93%E7%9A%84%E4%B8%BB%E4%BB%8E%E4%B8%80%E8%87%B4%E6%80%A7)
    - [数据库的高可用结构](#%E6%95%B0%E6%8D%AE%E5%BA%93%E7%9A%84%E9%AB%98%E5%8F%AF%E7%94%A8%E7%BB%93%E6%9E%84)
  - [Redis](#redis)
    - [redis的底层数据结构了解多少](#redis%E7%9A%84%E5%BA%95%E5%B1%82%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%BA%86%E8%A7%A3%E5%A4%9A%E5%B0%91)
    - [知道动态字符串sds的优缺点么？（sds是redis底层数据结构之一）](#%E7%9F%A5%E9%81%93%E5%8A%A8%E6%80%81%E5%AD%97%E7%AC%A6%E4%B8%B2sds%E7%9A%84%E4%BC%98%E7%BC%BA%E7%82%B9%E4%B9%88sds%E6%98%AFredis%E5%BA%95%E5%B1%82%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B9%8B%E4%B8%80)
    - [redis单线程特性有什么优缺点？](#redis%E5%8D%95%E7%BA%BF%E7%A8%8B%E7%89%B9%E6%80%A7%E6%9C%89%E4%BB%80%E4%B9%88%E4%BC%98%E7%BC%BA%E7%82%B9)
    - [redis有哪些数据结构，分别使用在什么场景？](#redis%E6%9C%89%E5%93%AA%E4%BA%9B%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%88%86%E5%88%AB%E4%BD%BF%E7%94%A8%E5%9C%A8%E4%BB%80%E4%B9%88%E5%9C%BA%E6%99%AF)
    - [什么是缓存穿透？如何避免？什么是缓存雪崩？何如避免？](#%E4%BB%80%E4%B9%88%E6%98%AF%E7%BC%93%E5%AD%98%E7%A9%BF%E9%80%8F%E5%A6%82%E4%BD%95%E9%81%BF%E5%85%8D%E4%BB%80%E4%B9%88%E6%98%AF%E7%BC%93%E5%AD%98%E9%9B%AA%E5%B4%A9%E4%BD%95%E5%A6%82%E9%81%BF%E5%85%8D)
    - [redis分布式锁](#redis%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81)
    - [简述Redis分布式锁的缺陷？](#%E7%AE%80%E8%BF%B0redis%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E7%9A%84%E7%BC%BA%E9%99%B7)
    - [加锁机制，锁互斥机制，watch dog自动延期机制，可重入加锁机制，锁释放机制是什么？](#%E5%8A%A0%E9%94%81%E6%9C%BA%E5%88%B6%E9%94%81%E4%BA%92%E6%96%A5%E6%9C%BA%E5%88%B6watch-dog%E8%87%AA%E5%8A%A8%E5%BB%B6%E6%9C%9F%E6%9C%BA%E5%88%B6%E5%8F%AF%E9%87%8D%E5%85%A5%E5%8A%A0%E9%94%81%E6%9C%BA%E5%88%B6%E9%94%81%E9%87%8A%E6%94%BE%E6%9C%BA%E5%88%B6%E6%98%AF%E4%BB%80%E4%B9%88)
    - [Redis 的 Setnx 命令是如何实现分布式锁的？](#redis-%E7%9A%84-setnx-%E5%91%BD%E4%BB%A4%E6%98%AF%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E7%9A%84)
    - [说说对Setnx 的实现锁的原理的理解？](#%E8%AF%B4%E8%AF%B4%E5%AF%B9setnx-%E7%9A%84%E5%AE%9E%E7%8E%B0%E9%94%81%E7%9A%84%E5%8E%9F%E7%90%86%E7%9A%84%E7%90%86%E8%A7%A3)
    - [如何避免死锁的出现？](#%E5%A6%82%E4%BD%95%E9%81%BF%E5%85%8D%E6%AD%BB%E9%94%81%E7%9A%84%E5%87%BA%E7%8E%B0)
    - [Redis里面有1亿个key，其中有10w个key是以某个固定的已知的前缀开头的，如何将它们全部找出来？](#redis%E9%87%8C%E9%9D%A2%E6%9C%891%E4%BA%BF%E4%B8%AAkey%E5%85%B6%E4%B8%AD%E6%9C%8910w%E4%B8%AAkey%E6%98%AF%E4%BB%A5%E6%9F%90%E4%B8%AA%E5%9B%BA%E5%AE%9A%E7%9A%84%E5%B7%B2%E7%9F%A5%E7%9A%84%E5%89%8D%E7%BC%80%E5%BC%80%E5%A4%B4%E7%9A%84%E5%A6%82%E4%BD%95%E5%B0%86%E5%AE%83%E4%BB%AC%E5%85%A8%E9%83%A8%E6%89%BE%E5%87%BA%E6%9D%A5)
    - [如何使用redis实现队列。又如何实现延时队列。](#%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8redis%E5%AE%9E%E7%8E%B0%E9%98%9F%E5%88%97%E5%8F%88%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E5%BB%B6%E6%97%B6%E9%98%9F%E5%88%97)
    - [如何实现持久化](#%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E6%8C%81%E4%B9%85%E5%8C%96)
    - [主从间的同步机制](#%E4%B8%BB%E4%BB%8E%E9%97%B4%E7%9A%84%E5%90%8C%E6%AD%A5%E6%9C%BA%E5%88%B6)
    - [聊聊redis sentinal和redis cluster](#%E8%81%8A%E8%81%8Aredis-sentinal%E5%92%8Credis-cluster)
    - [Redis主要消耗什么物理资源？](#redis%E4%B8%BB%E8%A6%81%E6%B6%88%E8%80%97%E4%BB%80%E4%B9%88%E7%89%A9%E7%90%86%E8%B5%84%E6%BA%90)
    - [Redis有哪几种数据淘汰策略？](#redis%E6%9C%89%E5%93%AA%E5%87%A0%E7%A7%8D%E6%95%B0%E6%8D%AE%E6%B7%98%E6%B1%B0%E7%AD%96%E7%95%A5)
    - [为什么Redis需要把所有数据放到内存中？](#%E4%B8%BA%E4%BB%80%E4%B9%88redis%E9%9C%80%E8%A6%81%E6%8A%8A%E6%89%80%E6%9C%89%E6%95%B0%E6%8D%AE%E6%94%BE%E5%88%B0%E5%86%85%E5%AD%98%E4%B8%AD)
    - [说说Redis哈希槽的概念？](#%E8%AF%B4%E8%AF%B4redis%E5%93%88%E5%B8%8C%E6%A7%BD%E7%9A%84%E6%A6%82%E5%BF%B5)
    - [Redis集群的主从复制模型是怎样的？](#redis%E9%9B%86%E7%BE%A4%E7%9A%84%E4%B8%BB%E4%BB%8E%E5%A4%8D%E5%88%B6%E6%A8%A1%E5%9E%8B%E6%98%AF%E6%80%8E%E6%A0%B7%E7%9A%84)
    - [Redis集群会有写操作丢失吗？为什么？](#redis%E9%9B%86%E7%BE%A4%E4%BC%9A%E6%9C%89%E5%86%99%E6%93%8D%E4%BD%9C%E4%B8%A2%E5%A4%B1%E5%90%97%E4%B8%BA%E4%BB%80%E4%B9%88)
    - [Redis集群之间是如何复制的？](#redis%E9%9B%86%E7%BE%A4%E4%B9%8B%E9%97%B4%E6%98%AF%E5%A6%82%E4%BD%95%E5%A4%8D%E5%88%B6%E7%9A%84)
    - [Redis集群最大节点个数是多少？](#redis%E9%9B%86%E7%BE%A4%E6%9C%80%E5%A4%A7%E8%8A%82%E7%82%B9%E4%B8%AA%E6%95%B0%E6%98%AF%E5%A4%9A%E5%B0%91)
    - [Redis集群如何选择数据库？](#redis%E9%9B%86%E7%BE%A4%E5%A6%82%E4%BD%95%E9%80%89%E6%8B%A9%E6%95%B0%E6%8D%AE%E5%BA%93)
    - [Redis集群方案应该怎么做？都有哪些方案？](#redis%E9%9B%86%E7%BE%A4%E6%96%B9%E6%A1%88%E5%BA%94%E8%AF%A5%E6%80%8E%E4%B9%88%E5%81%9A%E9%83%BD%E6%9C%89%E5%93%AA%E4%BA%9B%E6%96%B9%E6%A1%88)
    - [Redis集群方案什么情况下会导致整个集群不可用？](#redis%E9%9B%86%E7%BE%A4%E6%96%B9%E6%A1%88%E4%BB%80%E4%B9%88%E6%83%85%E5%86%B5%E4%B8%8B%E4%BC%9A%E5%AF%BC%E8%87%B4%E6%95%B4%E4%B8%AA%E9%9B%86%E7%BE%A4%E4%B8%8D%E5%8F%AF%E7%94%A8)
    - [怎么理解Redis事务？](#%E6%80%8E%E4%B9%88%E7%90%86%E8%A7%A3redis%E4%BA%8B%E5%8A%A1)
    - [Redis事务相关的命令有哪几个？](#redis%E4%BA%8B%E5%8A%A1%E7%9B%B8%E5%85%B3%E7%9A%84%E5%91%BD%E4%BB%A4%E6%9C%89%E5%93%AA%E5%87%A0%E4%B8%AA)
    - [Redis key的过期时间和永久有效分别怎么设置？](#redis-key%E7%9A%84%E8%BF%87%E6%9C%9F%E6%97%B6%E9%97%B4%E5%92%8C%E6%B0%B8%E4%B9%85%E6%9C%89%E6%95%88%E5%88%86%E5%88%AB%E6%80%8E%E4%B9%88%E8%AE%BE%E7%BD%AE)
    - [Redis如何做内存优化？](#redis%E5%A6%82%E4%BD%95%E5%81%9A%E5%86%85%E5%AD%98%E4%BC%98%E5%8C%96)
    - [Redis回收进程如何工作的？](#redis%E5%9B%9E%E6%94%B6%E8%BF%9B%E7%A8%8B%E5%A6%82%E4%BD%95%E5%B7%A5%E4%BD%9C%E7%9A%84)
- [数据结构和算法](#%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%92%8C%E7%AE%97%E6%B3%95)
  - [数据结构](#%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)
    - [链表和数组的优缺点？](#%E9%93%BE%E8%A1%A8%E5%92%8C%E6%95%B0%E7%BB%84%E7%9A%84%E4%BC%98%E7%BC%BA%E7%82%B9)
    - [解决hash冲突的方法有哪些？](#%E8%A7%A3%E5%86%B3hash%E5%86%B2%E7%AA%81%E7%9A%84%E6%96%B9%E6%B3%95%E6%9C%89%E5%93%AA%E4%BA%9B)
  - [算法](#%E7%AE%97%E6%B3%95)
- [网络](#%E7%BD%91%E7%BB%9C)
  - [网络基础](#%E7%BD%91%E7%BB%9C%E5%9F%BA%E7%A1%80)
    - [http1.0和1.1的区别？](#http10%E5%92%8C11%E7%9A%84%E5%8C%BA%E5%88%AB)
    - [http请求包含哪些数据结构？](#http%E8%AF%B7%E6%B1%82%E5%8C%85%E5%90%AB%E5%93%AA%E4%BA%9B%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84)
    - [什么是http的长连接和短连接？](#%E4%BB%80%E4%B9%88%E6%98%AFhttp%E7%9A%84%E9%95%BF%E8%BF%9E%E6%8E%A5%E5%92%8C%E7%9F%AD%E8%BF%9E%E6%8E%A5)
    - [谈谈tcp的三次握手和四次挥手。为什么建立连接需要三次，而不是两次？](#%E8%B0%88%E8%B0%88tcp%E7%9A%84%E4%B8%89%E6%AC%A1%E6%8F%A1%E6%89%8B%E5%92%8C%E5%9B%9B%E6%AC%A1%E6%8C%A5%E6%89%8B%E4%B8%BA%E4%BB%80%E4%B9%88%E5%BB%BA%E7%AB%8B%E8%BF%9E%E6%8E%A5%E9%9C%80%E8%A6%81%E4%B8%89%E6%AC%A1%E8%80%8C%E4%B8%8D%E6%98%AF%E4%B8%A4%E6%AC%A1)
    - [tcp有哪些状态，相应状态的含义。](#tcp%E6%9C%89%E5%93%AA%E4%BA%9B%E7%8A%B6%E6%80%81%E7%9B%B8%E5%BA%94%E7%8A%B6%E6%80%81%E7%9A%84%E5%90%AB%E4%B9%89)
    - [如果服务端没有收到最后的ack包，客户端可以开始发数据么？](#%E5%A6%82%E6%9E%9C%E6%9C%8D%E5%8A%A1%E7%AB%AF%E6%B2%A1%E6%9C%89%E6%94%B6%E5%88%B0%E6%9C%80%E5%90%8E%E7%9A%84ack%E5%8C%85%E5%AE%A2%E6%88%B7%E7%AB%AF%E5%8F%AF%E4%BB%A5%E5%BC%80%E5%A7%8B%E5%8F%91%E6%95%B0%E6%8D%AE%E4%B9%88)
    - [为什么接收方在FIN包后不能一次性发送ACK和FIN包给发送方，就像建立连接时一次性发送SYN和ACK包一样。](#%E4%B8%BA%E4%BB%80%E4%B9%88%E6%8E%A5%E6%94%B6%E6%96%B9%E5%9C%A8fin%E5%8C%85%E5%90%8E%E4%B8%8D%E8%83%BD%E4%B8%80%E6%AC%A1%E6%80%A7%E5%8F%91%E9%80%81ack%E5%92%8Cfin%E5%8C%85%E7%BB%99%E5%8F%91%E9%80%81%E6%96%B9%E5%B0%B1%E5%83%8F%E5%BB%BA%E7%AB%8B%E8%BF%9E%E6%8E%A5%E6%97%B6%E4%B8%80%E6%AC%A1%E6%80%A7%E5%8F%91%E9%80%81syn%E5%92%8Cack%E5%8C%85%E4%B8%80%E6%A0%B7)
    - [如果大量出现CLOSE_WAIT状态，说明什么？](#%E5%A6%82%E6%9E%9C%E5%A4%A7%E9%87%8F%E5%87%BA%E7%8E%B0closewait%E7%8A%B6%E6%80%81%E8%AF%B4%E6%98%8E%E4%BB%80%E4%B9%88)
    - [TIME_WAIT的作用？以及出现大量TIME_WAIT的原因。](#timewait%E7%9A%84%E4%BD%9C%E7%94%A8%E4%BB%A5%E5%8F%8A%E5%87%BA%E7%8E%B0%E5%A4%A7%E9%87%8Ftimewait%E7%9A%84%E5%8E%9F%E5%9B%A0)
    - [如何优化time_wait状态的发生？](#%E5%A6%82%E4%BD%95%E4%BC%98%E5%8C%96timewait%E7%8A%B6%E6%80%81%E7%9A%84%E5%8F%91%E7%94%9F)
    - [如果被断开的一方在收到FIN包后就跑路或者回复完ACK就跑路了，会怎么样？](#%E5%A6%82%E6%9E%9C%E8%A2%AB%E6%96%AD%E5%BC%80%E7%9A%84%E4%B8%80%E6%96%B9%E5%9C%A8%E6%94%B6%E5%88%B0fin%E5%8C%85%E5%90%8E%E5%B0%B1%E8%B7%91%E8%B7%AF%E6%88%96%E8%80%85%E5%9B%9E%E5%A4%8D%E5%AE%8Cack%E5%B0%B1%E8%B7%91%E8%B7%AF%E4%BA%86%E4%BC%9A%E6%80%8E%E4%B9%88%E6%A0%B7)
    - [如果出现大量的LAST_ACK状态，说明什么原因？](#%E5%A6%82%E6%9E%9C%E5%87%BA%E7%8E%B0%E5%A4%A7%E9%87%8F%E7%9A%84lastack%E7%8A%B6%E6%80%81%E8%AF%B4%E6%98%8E%E4%BB%80%E4%B9%88%E5%8E%9F%E5%9B%A0)
- [架构](#%E6%9E%B6%E6%9E%84)
  - [微服务](#%E5%BE%AE%E6%9C%8D%E5%8A%A1)
    - [微服务数据一致性问题，如何解决？](#%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%95%B0%E6%8D%AE%E4%B8%80%E8%87%B4%E6%80%A7%E9%97%AE%E9%A2%98%E5%A6%82%E4%BD%95%E8%A7%A3%E5%86%B3)

# Database
## Mysql
### 知道mysql的索引算法么？

### 为什么mysql要用b+树而不是b树或者其他树？

### 聊聊事务的隔离级别。你们生产用的什么事务隔离级别？

### 如何保证数据库的主从一致性

### 数据库的高可用结构

## Redis
### redis的底层数据结构了解多少
字符串采用动态字符串实现，数据结构为SDS，通过预先分配一个容量减少内存的频繁分配，也记录了字符串的长度大小。
```c
stuct SDS<T> {
    T capacity 
    T len
    byte flags
    byte[] content
}
```

列表：
1. 在元素较少的时候，内部采用的是压缩列表（ziplist），通过分配一块连续的内存，将所有元素紧挨着一起存储。
    ```c
    stuct ziplist<T> {
        int32 zlbyts; // 整个压缩列表占用字节数
        int32 zltail_offset; // 最后一个元素距离列表起始位置的偏移量，用于快速定位最后一个节点，实现双向遍历。
        int16 zllength; // 元素个数
        T[]   entries; // 元素内容列表，挨个挨个紧凑存储
        int8  zlend;  // 标志压缩列表的结束，值恒为 0xFF
    }
    ```
2. 在元素较多的时候，会采用快速列表（quicklist），因为普通的链表需要附加指针的空间太大，会浪费空间，加重内存的碎片化

### 知道动态字符串sds的优缺点么？（sds是redis底层数据结构之一）

### redis单线程特性有什么优缺点？

### redis有哪些数据结构，分别使用在什么场景？
string、list、hash、sort、sortedset   
hyperloglog、geo、Pub/Sub  
bloomfilter

### 什么是缓存穿透？如何避免？什么是缓存雪崩？何如避免？
[缓存](https://github.com/zhengweikeng/blog/blob/master/posts/2018/redis/%E4%BC%98%E5%8C%96.md)

### redis分布式锁

### 简述Redis分布式锁的缺陷？

### 加锁机制，锁互斥机制，watch dog自动延期机制，可重入加锁机制，锁释放机制是什么？

### Redis 的 Setnx 命令是如何实现分布式锁的？

### 说说对Setnx 的实现锁的原理的理解？

### 如何避免死锁的出现？

### Redis里面有1亿个key，其中有10w个key是以某个固定的已知的前缀开头的，如何将它们全部找出来？
keys和scan

### 如何使用redis实现队列。又如何实现延时队列。

### 如何实现持久化

### 主从间的同步机制

### 聊聊redis sentinal和redis cluster

### Redis主要消耗什么物理资源？

### Redis有哪几种数据淘汰策略？

### 为什么Redis需要把所有数据放到内存中？

### 说说Redis哈希槽的概念？

### Redis集群的主从复制模型是怎样的？

### Redis集群会有写操作丢失吗？为什么？

### Redis集群之间是如何复制的？

### Redis集群最大节点个数是多少？

### Redis集群如何选择数据库？

### Redis集群方案应该怎么做？都有哪些方案？

### Redis集群方案什么情况下会导致整个集群不可用？

### 怎么理解Redis事务？

### Redis事务相关的命令有哪几个？

### Redis key的过期时间和永久有效分别怎么设置？

### Redis如何做内存优化？

### Redis回收进程如何工作的？

# 数据结构和算法
## 数据结构
### 链表和数组的优缺点？

### 解决hash冲突的方法有哪些？

## 算法

# 网络
## 网络基础
### http1.0和1.1的区别？

### http请求包含哪些数据结构？

### 什么是http的长连接和短连接？

### 谈谈tcp的三次握手和四次挥手。为什么建立连接需要三次，而不是两次？

### tcp有哪些状态，相应状态的含义。

### 如果服务端没有收到最后的ack包，客户端可以开始发数据么？

### 为什么接收方在FIN包后不能一次性发送ACK和FIN包给发送方，就像建立连接时一次性发送SYN和ACK包一样。

### 如果大量出现CLOSE_WAIT状态，说明什么？

### TIME_WAIT的作用？以及出现大量TIME_WAIT的原因。

### 如何优化time_wait状态的发生？

### 如果被断开的一方在收到FIN包后就跑路或者回复完ACK就跑路了，会怎么样？

### 如果出现大量的LAST_ACK状态，说明什么原因？

# 架构
## 微服务
### 微服务数据一致性问题，如何解决？
