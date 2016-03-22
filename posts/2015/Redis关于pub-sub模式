> 做的项目中很频繁的使用了redis，而且使用了redis中的pub/sub功能，之前使用redis时没用过，故学习后，写下这篇博文。文中除了redis的代码，还会穿插Node.js的代码。


###### Redis发布/订阅模式
redis中实现了设计模式中的发布/订阅模式，与传统的pub/sub不一样的是，redis称之为Publish/Subscribe messaging paradigm，即发布与订阅信息泛型。  
简单来说就是将发布和订阅解耦，在发布和订阅的中间设立频道（channel）,客户端订阅（subscribe）的是频道，另一客户端将消息发布（publish）到频道上，之后频道则会转发该消息给对这个频道感兴趣的客户端。  
这种解耦，使得系统结构更加清晰，订阅者和发布者只需要关心各自的东西即可，也方便维护，可扩展性大大提高了。  
Redis中使用如下几个命令实现pub/sub：subscribe、unsubscribe、publish、psubscribe、punsubscribe。下面一一来解释命令的用法

###### 订阅（SUBSCRIBE）
客户端可以订阅任何自己感兴趣的频道，redis的命令如下`SUBSCRIBE foo bar`，SUBSCRIBE后的参数即为频道的名称。调用命令后返回的数据分三个部分
```
1. "subscribe"
2. "foo"
3. (integer) 1

1. "subscribe"
2. "bar"
3. (integer) 2

第一部分：此次动作的类型，这次为订阅动作，因此返回subscribe
第二部分：订阅的频道的名称，这里订阅了两个频道，foo和bar
第三部分：该频道的订阅数量。

node.js code
redisClient1.subscribe("foo", "bar")

redisClient1.on("subscribe", function(channel, count){
  console.log("client1 subscribed to " + channel + ", " + count + " total subscriptions");
})
```

###### 发布（PUBLISH）
另一个客户端通过publish发布消息到频道,`PUBLISH foo hello`，PUBLISH的第一个参数为频道名称，第二个消息为消息，此时订阅了该频道的客户端将收到同样分为三部分的格式的消息：
```
1. "message"
2. "foo"
3. "hello"

第一部分：此次动作的类型，这里是收到频道转发的消息，因此返回了message
第二部分：频道的名称
第三部分：收到的消息

node.js code:
redisClient2.publish("foo", "hello")

redisClient1.on("message", function(channel, message){
  console.log("client1 channel " + channel + ": " + message);
})
```

###### 退订（UNSUBSCRIBE）
客户端可以通过UNSUBSCRIBE退订订阅的频道。该命令可以接参数也可以不接参数，参数即位为要退订频道的名称。
```
UNSUBSCRIBE
# 会返回如下数据
1. "unsubscribe"
2. "foo"
3. (integer) 0

1. "unsubscribe"
2. "bar"
3. (integer) 1

第一部分：此次动作的类型，这里是退订频道，因此返回unsubscribe
第二部分：频道名称
第三部分：订阅该频道的客户端数量（退订后的数量）

node.js code:
redisClient1.unsubscribe()

redisClient1.on("unsubscribe", function(channel, count){
  console.log("client1 unsubscribed from " + channel + ", " + count + " total subscriptions");
})
```

###### 模式匹配（PSUBSCRIBE和PUNSUBSCRIBE）
redis的发布与订阅模式支持模式匹配，客户端可以订阅一个带*的频道，如果某个/某些频道的名字和这个模式匹配，那么当有信息发送给这个/这些频道的时候，客户端也会收到这个/这些频道的信息。  
例如：`psubscribe news.*`  
订阅后，客户端可以接收来自类似news.art和news.music等等频道的消息。返回的数据类型如下所示：  
```
1. "psubscribe"
2. "news.*"
3. (integer) 1
```
而`PUNSUBSCRIBE news.*`则可以退订任何以news.开头的频道。  
接收方面，与上述message不同的是，接收模式匹配的消息，动作类型为pmessage，即
```
1. "pmessage"
2. "art"
3. "good"

node.js code
redisClient1.psubscribe("news.*");

redisClient1.on("psubscribe", function(pattern, count){
  console.log("client1 psubscribed to " + pattern + ", " + count + " total subscriptions");
});

redisClient2.publish("news.art", "a good art channel");

redisClient1.on("pmessage", function(pattern, channel, message){
  console.log("("+  pattern +")" + " client1 received message on " + channel + ": " + message);
});

redisClient1.punsubscribe();

redisClient1.on("punsubscribe", function (pattern, count) {
  console.log("client1 punsubscribed from " + pattern + ", " + count + " total subscriptions");
});
```
