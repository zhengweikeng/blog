> mongodb入门之增删改

```bash
./mongod

use test
```

### 新增
`db.foo.insert({bar: "baz"})`

批量插入
`db.foo.batchInsert([{f1: "b1"}, {f2: "b2"}, {f3: "b3"}])`

### 删除
删除文档所有数据  
`db.foo.remove()`

删除文档中指定的数据  
`db.foo.remove({bar: "baz"})` 

更快的清楚文档的数据  
`db.foo.drop()`

### 更新文档
update(condition, modifier)
```javascript
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "friends": 32,
  "enemies": 2
}

var joe = db.users.find({"name", "joe"})
joe.releationships = {"friends": joe.friends, "enemies": joe.enemies}

joe.username = joe.name
delete joe.name
delete joe.friends
delete joe.enemies

db.users.update({name: 'joe'}, joe)
```

#### mongo内置了一些修改器
1.$set
用于指定一个字段值，如果不存在则创建它，存在则更新它。
```javascript
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "age": 20,
  "sex": "male"
}

db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$set": {"favoriteBook": "war and piece"}
})

db.users.findOne()
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "age": 20,
  "sex": "male",
  "favoriteBook": "war and piece"
}
```
unset可以将键删除掉
db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$unset": {"favoriteBook": 1}
})

2.$inc
$inc用于增加数值
```javascript
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe"
}

db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$inc": {"score": 50}
})
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "score": 50
}

db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$inc": {"score": 100}
})
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "score": 150
}
```

3.$push
用于向已有的数组末尾添加一条记录，要是没有则新增一个数组。
```javascript
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "content": "test"
}

db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$push": {
    "comments": {
      name: "tim",
      content: "hello"
    }
  }
})
db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$push": {
    "comments": {
      name: "jim",
      content: "world"
    }
  }
})

{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "name": "joe",
  "comments": [{
    name: "tim",
    content: "hello" 
  }, {
    {
      name: "jim",
      content: "world"
    } 
  }]
}
```
配合$each一次添加多条记录
```javascript
db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$push": {
    "comments": {
      "$each": [{
        {
          name: "tom",
          content: "hello world"
        },
        {
          name: "marry",
          content: "sound's good"
        }
      }]
    }
  }
})
```
配合$slice可以控制数组的长度，只保留最后加入的n条记录。
```javascript
db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$push": {
    "comments": {
      "$each": [{
        {
          name: "tom",
          content: "hello world"
        },
        {
          name: "marry",
          content: "sound's good"
        }
      }],
      // $slice的值只能是负数，这里代表了数组只会保留最后加入的10条记录
      "$slice": -10
    }
  }
})
```
使用$order还可以将数组排序后再添加到库中，但是它必须和$each一同使用。
```javascript
db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$push": {
    "comments": {
      "$each": [{
        {
          name: "tom",
          content: "hello world",
          rate: 6.1
        },
        {
          name: "marry",
          content: "sound's good"
          rate: 3.2
        }
      }],
      // $slice的值只能是负数，这里代表了数组只会保留最后加入的10条记录
      "$slice": -10,
      "$order": {rate: -1}
    }
  }
})
```
4.添加不重复元素  
有时为了保证数组内的元素不重复，可以在添加的时候再做一个判断，而$ne可以做到。而使用$addToSet会更加直观好用一点。
```javascript
db.users.update({"favoriteSports": {"$ne": "backetball"}}, {
  "$push": {
    "favoriteSports": "backetball"
  }
})

db.users.update({"_id": ObjectId("571ae75b7d309173ab2c745c")}, {
  "$addToSet": {
    "favoriteSports": "backetball"
  }
})

// $addToSet还可以和$each联合使用，为数组添加多个值，并且数组中不会有重复值。
// 而ne不能这么做
db.user.update({"_id": ObjectId("571ae75b7d309173ab2c745c"}, {
  "$addToSet": {
    "favoriteSports": {
      "$each": ["backetball", "football"]
    }
  }
})
```
5.删除元素  
第一种方式是使用$pop删除头部或者尾部的元素
```javascirpt
db.lists.find({})
{
  "_id": ObjectId("571ae75b7d309173ab2c745c"),
  "todo": ["dishes", "laundry", "dry cleaning", "wash window"]
}

db.lists.update({}, {
  "todo": {
    "$pop": {"key": 1}
  }
})
db.lists.update({}, {
  "todo": {
    "$pop": {"key": -1}
  }
})
```
删除特定元素
```javascirpt
db.lists.update({}, {
  "$pull": {
    "todo": "laundry"
  }
})
```
6.upsert  
会先查询集合有没符合的文档，没有则会创建
```javascript
db.users.update({name: "tim"}, {
  score: 50
}, true)
```
update的第三个参数即为是否为upsert操作，如果集合中没有找到name为tim的记录，则会使用创建一个`{name: "tim", score: 50}`的记录。

另外使用$setOrInsert还可以做到只在第一次时创建值，之后的更新操作该值都不变。
```javascript
db.users.update({}, {
  "$setOrInsert": {"createdAt": new Date()}
}, true)
```
如果再次运行这个脚本，都不会改变createdAt的值

### save
如果文档不存在，使用它会自动创建文档。如果文档存在则更新它。  
它接受一个参数，文档，如果该文档存在"_id"键，则调用upsert，否则调用insert。
```javascript
var x = db.users.findOne()
x.age = 22
db.user.save(x)
```

### 更新多个文档
默认情况下，当有多个文档匹配时，update时只会更新第一个匹配的文档，若要更新多个，则需要将第四个参数置为true
```javascript
db.user.update({birthday: "07/31/2016"}, {"$set": {age: 22}}, false, true)
```

### find基本用法
find(condi)根据特定条件查询
```javascript
db.users.find({name: 'jim'})

{
  _id: ObjectId("..."),
  name: "jim",
  age: 22,
  sex: male
}
```
也可以只返回指定的键
```javascript
db.users.find({name: 'jim'}, {name: 1, age: 1})

{
  _id: ObjectId("..."),
  name: "jim",
  age: 22
}
```
默认情况下，_id一定会被返回。我们可以指定某个键不被返回，利用这个特性可以不返回\_id
```javascript
db.users.find({name: 'jim'}, {sex: 0, _id: 0})

{
  name: "jim",
  age: 22
}
```

### 查询条件
$lt、$lte、$gt、$gte分别代表了<、<=、>、>=。
```javascript
db.users.find({age: {"$gt": 20, "$lt": 30}})
```
还有一个$ne，代表不等于
```javascript
db.users.find({name: {"$ne": "tim"}})
```

#### or查询
$in用于查询单个键的多个值，与之相反的是$nin
```javascript
db.users.find({name: {"$in": ["tim", "tom"]}})
db.users.find({name: {"$nin": ["tim", "tom"]}})
```
$or用于查询多个键
```javascript
db.users.find({"$or": [{name: "tim"}, {age: {"$in": [20, 22]}}]})
```

#### not语句
```javascript
// $mod取模，第一个参数是除以给定值，第二个参数是余数
db.users.find({"id_num": {"$mod": [5, 1]}})
db.users.find({"id_num": {"$not": {"$mod": [5, 1]}}})
```

### 特定类型的查询
#### null
查询为null的记录

```javascript
db.users.find({male: null})
```
但是如果该指定的键在文档中不存在，文档会被查询出来。因此需要指定exists字段
```javascript
db.users.find({male: {"$in": [null], "$exists": true}})
```
#### 支持正则表达式
```javascript
db.users.find({name: /joe/i})
```

### 数组查询
当有这样一个文档
```javascript
db.food.insert({fruits: ['apple', 'banana', 'orange']})

db.food.find({fruits: 'banana'})
{
  _id: ObjectId(xxxx)
  fruits: ['apple', 'banana', 'orange']
}
```
这会查询到拥有banana的记录，想要绝对匹配，则需要指定为数组

```javascript
db.food.insert({fruits: ['apple', 'banana', 'orange']})
db.food.insert({fruits: ['apple', 'banana', 'cherry']})

db.food.find({fruits: ['apple', 'banana', 'orange']})
{
  _id: ObjectId(xxxx)
  fruits: ['apple', 'banana', 'orange']
}
```
可以指定从数组的某个下标开始查询
```javascript
db.food.insert({fruits: ['apple', 'banana', 'orange']})
db.food.insert({fruits: ['apple', 'banana', 'cherry']})

db.food.find({"fruits.2": "banana"})
```

#### $all
通过多个元素匹配数组
```javascript
db.food.insert({fruits: ['apple', 'banana', 'orange']})
db.food.insert({fruits: ['apple', 'kumquat', 'orange']})
db.food.insert({fruits: ['cherry', 'banana', 'apple']})

db.food.find({fruits: {"$all": ["apple", "banana"]}})
{
  _id: ObjectId(xxxx)
  fruits: ['apple', 'banana', 'orange']
}
{
  _id: ObjectId(xxxx)
  fruits: ['cherry', 'banana', 'banana']
}
```

#### $size
```javascript
db.food.insert({fruits: ['apple', 'banana', 'orange']})
db.food.insert({fruits: ['apple', 'banana']})

db.food.find({"fruits": {"$size": 3}})
{
  _id: ObjectId(xxxx)
  fruits: ['apple', 'banana', 'orange']
}
```

#### $slice
查询的第二个参数，可以返回指定的键，而使用slice可以返回数组中指定位置的值
```javascript
db.food.insert({fruits: ['apple', 'banana', 'orange', 'cherry']})

// 返回fruits的前2条记录
db.food.find({}, {"fruits": {"$slice": 2}})
{
  _id: ObjectId(xxxx)
  fruits: ['apple', 'banana']
}

// 返回fruits的后2条记录
db.food.find({}, {"fruits": {"$slice": -2}})
{
  _id: ObjectId(xxxx)
  fruits: ['orange', 'cherry']
}

// 返回fruits指定位置的记录
db.food.find({}, {"fruits": {"$slice": [1, 2]}})
{
  _id: ObjectId(xxxx)
  fruits: ['banana', 'orange']
}
```
