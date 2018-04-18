# 为什么redis采用单线程
redis的server是采用单线程来处理所有请求（执行持久化BGSAVE任务时会开启子进程），这么做的原因是什么呢？

首先redis是基于内存的操作，在内存中操作效率已经非常的高了。

其次如果是多线程或者多进程，还需要涉及到竞争和锁的问题。  
另外多线程或者多进程处理会增加线程或者进程间的切换，进而消耗cpu资源。

redis采用I/O多路复用来提高I/O的效率。

这里重点理解下I/O多路复用

## I/O多路复用（I/O multiplexing，又被称为“事件驱动”）
多路指的是多个多个描述符的I/O操作（也可以认为是多个socket连接），复用指的是复用同一个线程。

I/O多路复用常用的技术有select、poll、epoll。采用多路复用可以让单个线程高效的处理多个请求，同时redis是在内存中进行数据操作，速度非常快，综合这两点使得redis具备很高的吞吐量。

redis没有采用libevent，而是自己实现了一个io多路复用，根据不同系统对select、poll、epoll都进行了支持。在linux下采用了epoll。  

自己实现的原因是，libevent比较重，有很多功能redis用不上，为了减少redis的依赖，便自己实现了多路复用的功能。

### select和poll
单线程下怎么复用连接呢？

最简单的就是利用轮询。  
死循环的轮询每个流，如果有io事件就进行处理。虽然这样实现了复用线程功能，但是代价太大了，需要轮询每一个流，很多空闲的io也会被轮询到，导致cpu在空转，浪费资源。

看下select是如何解决这种情况的。

select会监听文件描述符（socket连接），如果描述符发生了变化，可能是可读了或者可写了，这时select会返回，通知线程去进行操作。如果没有io到达，就进入阻塞，直到select超时为止。

如果没有就绪的socket，select会一直阻塞直到超时，当然可以让其不超时一直阻塞。

select函数返回后，只是告诉系统已经有socket就绪了，但是却没说是哪个socket就绪了，因此还是需要去遍历一遍所有描述符找到就绪的那个socket。但是它解决了cpu空转的问题。

select有3个缺点：
1. 每次调用select，都需要把fd集合从用户态拷贝到内核态，这个开销在fd很多时会很大。
1. 每次调用select后，都需要在内核遍历传递进来的所有fd，这个开销在fd很多时也很大。
1. fd数量有限，默认1024。

为了解决文件描述符数量的限制，于是产生了poll。  
poll和select几乎是一样的，它解决了文件描述数量的限制，但是失去了select跨平台的特性。

从上述说明知道，select和poll都需要在返回后，通过遍历文件描述符来获取已经就绪的socket。但其实大部分情况下，同一时刻只有很少socket处于就绪状态，为了获取这个就绪的socket，我们要去遍历一遍列表就有点得不偿失了。  
这时候epoll出现了。

### epoll
为了解决select和poll中无效的遍历，只需要记录就绪的是哪些socket即可。

epoll使用了一个描述符来管理多个描述

参考下接口定义

定义一个epoll的句柄
```c
init epoll_create(int size)
```
这个size只是来建议内核分配一个多大的描述符个数。创建句柄结束后，会占用一个fd值，因此在使用完epoll后需要close()关闭它，否则会导致文件描述符被耗尽。

注册事件，主要是说明要监听哪些事件
```c
int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)
```
例如可以注册缓冲区非空事件，即有数据流入了；或者注册缓冲区非满事件，即流可以写入了。

最后等待IO事件直到上述注册的事件发生
```c
int epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout)
```


epoll先通过事件注册一个文件描述符，一旦该描述符就绪时，内核会采用类似callback的回调机制，迅速激活这个文件描述符，当进程调用epoll_wait()时便得到通知。


参考资料：  
[I/O多路复用技术（multiplexing）是什么？](https://www.zhihu.com/question/28594409)  
[IO多路复用深入浅出](https://www.jianshu.com/p/1020c11f016c)  
[我读过的最好的epoll讲解](http://blog.51cto.com/yaocoder/888374)
