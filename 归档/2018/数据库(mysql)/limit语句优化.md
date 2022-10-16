# LIMIT语句
**LIMIT offset, length**  

采用LIMIT很容易实现分页功能

```sql
SELECT * FROM table LIMIT 10,100;
```

但是当我们的偏移量offset很大的时候，mysql需要废除掉前面的offset条记录，只返回length条记录，这个代价还是蛮高的

```sql
SELECT count(*) FROM article;
-- 243151 rows

SELECT * FROM article LIMIT 10,100;
-- 6.5ms

SELECT * FROM article LIMIT 1000,100;
-- 6.9ms

SELECT * FROM article LIMIT 10000,100;
-- 18.5ms

SELECT * FROM article LIMIT 20000,100;
-- 47.5ms
```
可见数据在查询10000行之后的数据已经要18毫秒了，到了20000已经需要47毫秒，当数据越多，这个时间会长，有没什么办法可以解决这个问题呢？

有一种方式是，业务场景上不要不要查询偏移量这个高的数据，如页面上不显示太多页码；

但是我们还是要来寻找技术上的解决方案。

## 子查询优化法
```sql
SELECT * FROM article LIMIT 10000,100;
-- 18.5ms

SELECT * FROM article WHERE id >= (
	SELECT id FROM article LIMIT 10000,1
) LIMIT 100
-- 7.6ms

SELECT * FROM article WHERE id >= (
	SELECT id FROM article LIMIT 20000,1
) LIMIT 100
-- 10.1ms
```
先找出第一条数据的id，然后查找大于这个id的数据。时间明显缩小了一半

这种方式的缺点需要where条件必须是连续的

## 取消offset，采用条件过滤
```sql
SELECT * FROM article WHERE id >= 20000 LIMIT 100;
```
跟之前的方式有点像，这里的20000可以让客户端传过来，这样我们也减少了子查询。

这种的好处就是无论是到多少页，性能都非常的好，跟查询从第10行还是20行开始效果是一样的。

这种方式的缺点依旧是需要where条件必须是连续的

## 通过连接查询来分页
上面的id中，要求id是连续的，并且不能中断，如果曾经删除过数据，那么id就不连续了。

这时可采用表连接的方式来做。
```sql
SELECT * FROM article AS t1 INNER JOIN (
	SELECT id FROM article LIMIT 20000,100
) AS t2 ON t1.id = t2.id;
-- 9.5ms
```
由于id是主键，通过主键索引将数据查找出来的速度会非常快，将根据id进行数据连接便可以查出数据。
