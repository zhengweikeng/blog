# 慢查询
## 查询是否开启慢查询日志记录
```sql
-- 是否开启慢查询日志记录
show variables like 'slow_query_log'

-- 设置慢查询日志路径
set global slow_query_log_file='/home/mysql/sql_log/mysql-slow.log'

-- 开启查询没有使用索引的表时将记录日志的功能
set global log_queries_not_using_indexes=on

-- 设置慢查询时间
set global long_query_time=1
```

日志格式
```
# Time: 2017-05-02T02:36:44.948970Z
# User@Host: root[root] @ localhost [127.0.0.1]  Id:   196
# Query_time: 0.002199  Lock_time: 0.000060 Rows_sent: 599  Rows_examined: 599
SET timestamp=1493692604;
SELECT * FROM CUSTOMER;
```

## 慢查询工具
## mysql自带工具 —— mysqldumpslow 
```sql
-- 输出前10条最慢的日志
mysqldumpslow -t 10 /home/mysql/sql_log/mysql-slow.log
```
输出格式
```
Count: 1  Time=0.00s (0s)  Lock=0.00s (0s)  Rows=0.0 (0), 0users@0hosts
  Time: N-N-02T02:N:N.018845Z
  # User@Host: root[root] @ localhost [N.N.N.N]  Id:     N
  # Query_time: N.N  Lock_time: N.N Rows_sent: N  Rows_examined: N
  SET timestamp=N
```

其他工具还有`pt-query-digest`

## 分析SQL查询
```sql
explain select customer_id from customer;
```

|id | select_type | table | partitions | type | possible_keys | key | key_len | ref | rows | filtered | EXTRA
|------ | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | -------
1 | SIMPLE | customer | NULL | index | NULL | idx_fk_store_id | 1 | NULL | 599 | 100.00 | Using inex

字段解释：  
table: 显示这一行的数据是关于哪张表  
type: 显示连接使用了何种类型，从最好到最差的连接类型为const、eq_reg、ref、range、index和all  
possible_keys：显示可能应用在这张表中的索引。如果为空，则没有索引  
key：实际使用的索引  
key_len：使用的索引长度，索引越短越好  
ref：显示索引的哪一列被使用了，如果可能的话，是一个常数  
rows：MYSQL认为必须检查的用来返回请求数据的行数

关于EXTRA，如果看到该列的值为以下返回值，需要优化sql语句了  
Using filesort。mysql需要进行额外的步骤来发现如何对返回的行排序。  
一般来说，偶尔的查询并不会引起性能问题，但是频繁查询，就会有影响了。出现这种情况一般是使用了`ORDER BY`或者`GROUP BY`，其跟随的字段没有加索引，引发mysql排序的时候无法使用索引排序，造成性能低下

Using temporary。需要创建一个临时表来存储结果，这通常发生在对不同的列集进行ORDER BY和GROUP BY上  
如果order by的子句只引用了联接中的第一个表，MySQL会先对第一个表进行排序，然后进行联接。也就是expain中的Extra的Using Filesort.否则MySQL先把结果保存到临时表(Temporary Table),然后再对临时表的数据进行排序.此时expain中的Extra的显示Using temporary Using Filesort.

参考文档：
1. [详解 MySQL 中的 explain](http://www.cnblogs.com/zengkefu/p/5647206.html)
1. [MySql中explain的时候出现using filesort，优化之](http://blog.csdn.net/imzoer/article/details/8485680)
1. [MySQL Order By实现原理分析和Filesort优化](http://blog.csdn.net/hguisu/article/details/7161981)
1. [MySQL调优 —— LEFT JOIN](http://blog.csdn.net/wenniuwuren/article/details/44851819)

## 索引优化
建立索引的原则：
1. 查询频繁
1. 区分度高
1. 长度小
1. 尽量能覆盖常用查询字段

### 如何选择合适的列建立索引
1. 在where从句，group by从句，order by从句，on从句中出现的列建立索引
1. 索引字段越小越好
1. 离散度大的列放到联合索引的前面

```sql
SELECT * FROM payment WHERE staff_id AND customer_id=584
```
customer_id的离散度更大，应该使用`index(customer_id,staff_id)`的方式建立索引

参考资料
1. [mysql建索引的几大原则](http://blog.csdn.net/u013412790/article/details/51612304)

## 数据类型
很多时候我们都可以使用整型来存储数据
### 整型存储时间类型
可以以时间戳的方式将字符串类型的时间存储为整型
```sql
INSERT INTO test(timestr) VALUES(UNIX_TIMESTAMP('2017-01-01 13:12:00'))

SELECT FROM_UNIXTIME(timestr) FROM test
```

### 整型存储ip地址
```sql
INSERT INTO sessions(ipaddress) VALUES(INET_ATON('192.168.0.1'))

SELECT INET_NTOA(ipaddress) FROM sessions
```