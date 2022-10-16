## 加入Redis里有1亿个Key，其中10w个key是以某个固定的已知的前缀开头，如何将他们全部找出来。
最原始也是最慢的方式就是采用`keys prefiex*`的方式，这种方式不仅慢，而且会阻塞redis线程，造成服务不可用，对线上服务会有一定的影响。

此时可以采用`SCAN`命令

```
SCAN cursor [MATCH pattern] [COUNT count]
```
scan每次执行都只返回少量的元素，是一个基于游标的迭代器，每次命令的调用都需要使用上一次调用所返回的游标作为该次调用的游标参数，以此来延续之前的迭代过程

当游标参数cursor为0时，则开始新的一次迭代，当redis返回值为0时，表示迭代结束

```redis
redis> SCAN 0
0
  hello
  user:3
  user:1
  user:6
  user:7
  user:2
  user:10
  user:9
  class
  user:4

redis> SCAN 5
0
  user:8
  user:5
```

每次迭代，返回的元素个数是不定的，对于大数据集来说，每次迭代可能会返回数十个元素，但是对于一个足够小的数据集来说，可能会一次性返回都有的Key

可以使用count指定返回的元素个数，但是它只是提示redis返回的个数，redis并不会严格遵循它，只是一个大致的约束
```redis
redis> SCAN 0 COUNT 20
0
  hello
  user:3
  user:1
  user:6
  user:7
  user:2
  user:10
  user:9
  class
  user:4
  user:8
  user:5

redis> SCAN 0 COUNT 5
2
  hello
  user:3
  user:1
  user:6
  user:7

redis> SCAN 2 COUNT 5
5
  user:2
  user:10
  user:9
  class
  user:4

redis> SCAN 5 COUNT 5
0
  user:8
  user:5
```

如果只想匹配具备某个特定格式的key，则需使用MATCH
```redis
redis> SCAN 0 MATCH user:*
5
  user:3
  user:1
  user:6
  user:7
  user:2
  user:10
  user:9
  user:4

redis> SCAN 5 MATCH user:*
0
  user:8
  user:5
```

SCAN的缺点：  
1. 同一个元素可能会被返回多次。 处理重复元素的工作交由应用程序负责
1. 如果一个元素是在迭代过程中被添加到数据集的，又或者是在迭代过程中从数据集中被删除的，那么这个元素可能会被返回，也可能不会，这是未定义的

除了SCAN外，还有`SSCAN`、`HSCAN`、`ZSCAN`，区别在于：
1. SCAN是用于迭代redis中所有的键
2. SSCAN用于迭代集合键中的元素
3. 用于迭代哈希键中的键值对
4. 用于迭代有序集合中的元素（包括元素成员和元素分值）
5. SSCAN 命令、 HSCAN 命令和 ZSCAN 命令的第一个参数总是一个数据库键

```redis
redis> HSCAN class 0
0
  c1
  New Value
  c2
  bbb
```
