> 之前梳理了一次关于node.js中关于流的用法的知识点，感觉还是少了些东西。所以就这阵子再梳理一些东西，帮助自己多点理解stream。

##### readable
我们知道，可读流可以读取数据到缓存中，当缓存中的数据可用时，便会触发`readable`事件。在该事件中我们便可以去使用read方法去读取缓存中的数据。
```javascript
var fs = require("fs")
var rs = fs.createReadStream("./test.md")
rs.on("readable", function log() {
  console.log("in readable")
  console.log(rs.read())
})
```
readable事件被触发，则表明了出现如下情况之一：
1. 缓存中有新的数据可以被读取，此时调用read读取数据的话则会返回数据
2. 已经到了流的结尾，此时调用read读取数据的话会返回null

当所有数据都已经读取结束，则会触发end事件
```javascript
rs.on("end", () => {
  console.log("there will no more data.")
})
```

这里我们有必要说一下read这个方法。  
read([size]) size参数是可选的，表示每次从缓存中读取多少字节的数据。  
如不加上这个参数，则读取的是缓存中的所有数据。若指定了每次读取的字节数，则有一点要注意一下。
```javascript
rs.on("readable", function log() {
  console.log("in readable")
  console.log(rs.read(3))
})
```
上面的例子中每次读取3个字节，会发现我们的读取的文件中的数据并没有读完，而且end事件也没有触发。  
原因就是readable事件并没有一直被触发。根据上面所述我们知道readable的触发机制必须为缓存中有新数据或者已经没有数据了。  
现在的情况是我们每次只读取3个字节，而缓存中还存有剩余的数据，因此readable便不会被触发了。  

要解决这种情况，我们需要在readable事件中不断去读取数据，直到把缓存的数据读完为止。
```javascript
rs.on("readable", function log() {
  console.log("in readable")
  while(null !== (chunk = rs.read(3))) {
    console.log(chunk.toString())
  }
})
```  
当缓存中的数据已经被读完时，再调用read方法则会返回null，此时执行完readable事件的回调后便会去文件中读取新的数据到缓存中，于是readable便源源不断的被触发了。
