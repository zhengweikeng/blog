# 内存和redo log
mysql在做更新操作的时候，会更新内存和写磁盘redo log，因此效率非常高，性能也非常好。

但是即便如此，也一样会有性能瓶颈，这里的瓶颈也依旧在内存和redo log上，成也萧何败萧何。

首先需要知道一个概念，当内存的数据和磁盘中的数据不一致时，我们称此时的内存页为“脏页”。当内存数据写入到磁盘后，此时内存和磁盘上的数据页也就一致了，此时的内存页称之为“干净页”。

这里也就引发出了一个问题，mysql什么时候将脏页刷回磁盘？   
先看如下两种刷盘的情况：
1. 系统比较空闲的时候
2. mysql正常关闭的时候

这两种情况，都是在系统比较正常的情况下发生的，因此也不会有什么性能问题。

而接下来分析的情况，则都会影响到mysql的读写性能。

## redo log问题
如果redo log被写满，这时候系统会停止所有更新操作，然后将日志中的checkpoint往前推进，将推进的区域指定的脏页数据刷回磁盘上。
```
-----------------------
|            |        |
write pos    cp1      cp2
```
这时write pos到cp2的区域就是可以写入redo log的区域。

出现这种情况，会比较严重，毕竟这种情况下数据库更新操作就都被阻塞住了，性能急剧下降。

## 内存问题
如果系统的内存不足，当需要新的内存页时，就需要淘汰内存的一些数据。如果淘汰的是脏页，那么就需要先把脏页写到磁盘。

InnoDB的管理内存是用buffer pool来管理的，有如下三种状态：
1. 还没有使用的
2. 使用了并且是干净页
3. 使用了并且是脏页

如果内存不足时，当要读取的数据又不在内存中时，只能先把最久的使用的数据页从内存中淘汰。

如果淘汰的是干净页，可以直接释放使用。  
如果淘汰的是脏页，那就必须先把脏页刷回磁盘，变成干净页才能被使用。

但是如果频繁刷脏页就可能导致性能问题，一个查询要淘汰的脏页太多了，查询的响应时间变长。

## 如何解决
日志问题和内存问题都会导致mysql性能的下降，有时可能只是一个简单的查询，也可能会异常的蛮。甚至更新操作直接被阻塞，系统写操作无法使用。

无论是内存还是日志问题，最终都是因为由于需要刷脏页到磁盘中导致的，如何控制刷脏页频率就决定了系统的吞吐能力。

InnoDB有几个很重要的参数
```
innodb_log_file_size
innodb_log_files_in_group
innodb_io_capacity
# 脏页比例上限，默认值75%
innodb_max_dirty_pages_pct
```

这里最重要的就是`innodb_io_capacity`，该参数用于告诉innodb你的磁盘能力。

之前我们说过mysql也会在系统比较空闲的时候进行内存的刷盘操作，mysql会根据`innodb_io_capacity`来判断更新的速度能不能在可以控制的范围内
1. 说过该值设置过大，就会导致mysql高估了系统的能力，导致脏页大量堆积，内存占用越积越大，而且redo log也会被写满。
2. 如果该值设置过低，机会导致mysql低估了系统的能力，导致频繁的刷盘，系统系统响应速度降低，换句话说就是数据库单位时间内提交的事务数(tps)降低。

一般来说，会将innodb_io_capacity设置成系统的IOPS，这个值可以用fio来测试
```
 fio -filename=$filename -direct=1 -iodepth 1 -thread -rw=randrw -ioengine=psync -bs=16k -size=500M -numjobs=10 -runtime=10 -group_reporting -name=mytest 
```

通过设置合理的`innodb_io_capacity`，并且关注脏页比例，让脏页比例不超过75%

脏页比例计算：`innodb_buffer_pool_pages_dirty/innodb_buffer_pool_pages_total`
```sql
mysql> select VARIABLE_VALUE into @a from global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_dirty';
select VARIABLE_VALUE into @b from global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_total';
select @a/@b;
```

另外还有个重要的问题，在mysql在刷脏页的时候，如果发现数据页旁边的数据页刚好也是脏页，那么也会一起刷掉。不仅如此，邻居数据页的数据页也是脏页，也会去刷，一直遍历下去。我们可以将`innodb_flush_neighbors`设置为0，关闭这个功能。而到了mysql 8.0时，该参数默认值已经是0了。