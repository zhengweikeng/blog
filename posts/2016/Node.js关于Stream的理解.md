> 关于Node，在Stream这块一直是自己薄弱的地方，很多教程对这块也没有过多的介绍，朴灵大神的《深入浅出Node.js》更加把这块省略了（不过朴神要写2.0版了，希望这块能补上）。Stream从0.10开始已经很强大了，所以还是有必要来梳理一下这块的知识。  
> 本文不是翻译node.js的stream的API，而是对学习过程中遇到的问题做一些梳理。

为什么要使用流？  
想象一个场景，你在程序中，拷贝一个文件的内容到另外一个文件中。我们可能会这么做
```javascript
var source = fs.readFileSync('/path/to/source', {encoding: 'utf8'});
fs.writeFileSync('/path/to/dest', source);
```
这么做，就等于把源文件的内容全部读取到内存中，再将内存中的数据，写入到目标文件中。当然，如果是小文件，这种做法是完全可以的。然而，如果是像音频，视频这样动辄几G的文件，这种做法就不合适了，随时撑爆内存。  
面对这种大文件，我们应该是，读一些，再写一些，不断反复直到文件全部读取完毕，写入也结束。这种方式时间会长一点，但是是可取的。这也就是流的概念。  
对于流的理解，可看下图：  
![流的示意图](http://7xjw3r.com1.z0.glb.clouddn.com/image/e/b9/07d19321d2aa50a6853acbb543fbc.png)

目前node.js（4.0）的流分为4种，Readable,Writeable,Duplex,Transform，接下来将会对这几种流进行介绍。当然，阅读的前提是你看过Node.js关于Stream的api。没看过请移步[这里](https://nodejs.org/api/stream.html)。

##### Readable
可读流，这里需要重点理解Readable流有两种模式，流动模式（flow mode）和暂停模式（pause mode），理解这两种模式对该种流的理解将更透彻。  
流动模式，该模式下的流将会由底层系统自行调用，并尽可能快的提供给你的程序。简而言之，就是程序自动读取数据。  
暂停模式，该模式下的流，如果要获取数据，需要自己手动的调用流的read()方法。  
默认是暂停模式。  
根据官网所说的，有以下3种方式，将暂停模式的流转化为流动模式，  
1. 添加一个data事件的监听器来监听数据
2. 调用resume()方法来明确开启流动模式
3. 调用pipe()方法将数据导入一个可写流

new stream.Readable([options])  
options Object：
highWaterMark Number 在停止从底层资源读取之前，在内部缓冲中存储的最大字节数。默认为16kb，对于objectMode则是16。  
encoding String 如果被指定，那么缓冲将被利用指定编码解码为字符串，默认为null。  
objectMode Boolean 是否该流应该表现如一个对象的流。意思是说stream.read(n)返回一个单独的对象而不是一个大小为n的Buffer，即此时往read传递参数n也是无效的。stream默认把纯数据块送过来，否则会把数据放到一个 object中，默认为false。

在读取数据的过程中，根据不同情况将会触发一些事件，如readable，data，end，error等等。这里容易搞不清楚的就是readable和data事件的差别，在说明它们的区别前，我们先创建一个可读流。
```javascript
var Readable = require('stream').Readable;
var rs = new Readable();

// rs.push("hello");
// rs.push("world");
// rs.push(null);
```
注意到上述代码的注释的地方，一般不推荐使用这种方式为可读流填充数据，而是重写`_read()`方法，在该方法中实现数据的填充读取。需要注意的是`_read`方法是不能手动调用的，由流内部自动调用。如果是自己写流插件，继承了Readable后，需要重写该方法。
```javascript
var c = 97;
rs._read = function(){
  rs.push(String.fromCharCode(c++));
  if (c > 'd'.charCodeAt(0)) rs.push(null);
}
```
接下来，我们监听readable事件,
```javascript
// 1:
rs.on("readable", function(){
  data = rs.read();
  console.log("readable: " + data);
});

// 2:
// rs.on("readable", function(){
  // while (null !== (record = rs.read())) {
    // console.log('received: ' + JSON.stringify(record));
  // }
// });
```
可见，在该事件下，数据需要使用rs.read()来读取，所以此时还是暂停模式（pause mode）的流。  
只要监听了readable事件，可读流就会源源不断的去执行_read方法填充数据到缓存，直到填充的数据是空，即rs.push(null)，才会停止。  
而只要缓存中存在数据，就会不断触发readable事件，该事件中通过rs.read()来读取缓存中的数据，同时将数据从缓存中清除，直到rs.read()读取到了null，说明缓存中已经没有数据，则停止触发该事件。这就是上述代码1的意思。  
而2是什么意思呢。其实是差不多意思，只是数据的处理在第一次触发readable事件时，通过while循环读取数据，直到数据读取完成返回null。即在第一次触发readable就完成了数据的读取。后面几次触发的readable事件读取到的数据都是null。

那data事件又是如何，
```javascript
rs.on("data", function(chunk){
  console.log(chunk.toString());
});
```
很明显，在该事件下，无需我们自己调用read()就可以获取到数据chunk，注意这个数据是Buffer类型的。一旦触发该事件，可读流的模式将会转变为流动模式，数据的将会由程序自行读取，无需手动调用。只要一往流中填充数据，就会触发该事件，即push几次，就触发几次。每一次的chunk就是每次push的数据的Buffer化。  
和readable不同的地方不仅是模式的不同，还有读取到的数据的不同，readable事件中，read()读取到的可能是push了几遍后的数据。而data事件则不会。

到这里，也就基本清楚怎么理解Readable。那些error事件，end事件翻一下api就能懂了，就不说了。  
顺便提一下rs.push()这个方法。  
根据官网文档所说，该方法应该由流的实现者调用，如我们重写的_read()方法中就调用push方法。消费者则不应该调用push()方法，如我们不应该在data事件或者readable事件中调用该方法。
push进去的数据，可以通过read来进行获取。当push的数据是null时，即push(null)，则会发送一个数据结束信号，此时end事件将会触发。

在上面中，我们使用`_read`来往缓存中填充数据，使用该方法的好处就是，我们可以在需要消费数据时才往缓存中填充数据，这样可以使缓存不至于太快被填充满。再拿这段代码来说
```javascript
var Readable = require('stream').Readable;
var rs = new Readable();

rs.push("hello");
rs.push("world");
rs.push(null);
```
如果我们没有监听readable事件，或者可能要等待其他模块处理完后才进行数据处理，此时上面代码push的hello world还是会长期贮存在缓存中。如果我们能做到只在需要消费数据时才产生数据，这无疑能够减少缓存的压力。这也就是`_read`的好处。
```javascript
var c = 97;
rs._read = function(){
  rs.push(String.fromCharCode(c++));
  if (c > 'd'.charCodeAt(0)) rs.push(null);
}
```
这样即使无监听readable或者data事件，也不会产生数据。一旦监听，则会产生数据。

###### pipe
至今我们我们还为提及另外一个重量级功能，就是pipe。可以理解为管道，使用它可以实现各种不同流的连接，如将输入流和输出流连接，将输入流输入的数据，输出到输出流中。  
一般的做法就是监听data事件，然后创建一个输出流，然后往输出流写数据。然而使用pipe管道，直接将输出流接入到pipe中即可。
```javascript
var Readable = require('stream').Readable;
var rs = Readable();

var c = 97;
rs._read = function () {
  rs.push(String.fromCharCode(c++));
  if (c > 'z'.charCodeAt(0)) rs.push(null);
};

rs.pipe(process.stdout);
```
其实只是pipe内部替我们实现了data事件罢了。所以也会将流的模式转化为流动模式。  
当然，pipe如果只是这样就没什么吸引力了，这个方法会自动调节流量，所以当快速读取可读流时目标不会溢出。而且该方法返回流本身，因此可以不断调用pipe，形成链式调用。
```javascript
var r = fs.createReadStream('file.txt');
var z = zlib.createGzip();
var w = fs.createWriteStream('file.txt.gz');
r.pipe(z).pipe(w);
```
通过pipe我们可以在流中做各种处理，并只输出最终格式的流。我们所知的gulp就使用了这种方式进行文件处理。
```javascript
gulp.src("test.js")
.pipe(usemin())
.pipe(uglify())
.pipe(gulp.dest("dist"));
```
通过这种方式，文件的处理全部在流中处理，应该说在缓存中处理数据，最终输出到文件中，中间不形成任何临时文件。  
想想grunt怎么用的，usemin后产生一个处理后的文件到磁盘，从磁盘读取该文件后再进行uglify处理，再生成一个处理后的文件，最终再生成一个结果文件。如果需要做很多处理，则需要频繁进行磁盘的读取，这结果，你懂得。

OK。搞完了可读流，之后的流的理解就简单了。

##### Writable
可写流，理解起来是很简单的，无非就是调用write()方法。  
new stream.Writable([options])  
options Object：
highWaterMark Number write()方法开始返回false时的缓冲级别。默认为16kb，对于objectMode流则是16。  
decodeStrings Boolean 是否在传递给write()方法前将字符串解码成Buffer。默认为true。  
objectMode Boolean 是否write(anyObj)为一个合法操作。如果设置为true你可以写入任意数据而不仅是Buffer或字符串数据。可以是任何js对象。默认为false。当设置为true时，在调用write(data,encoding)时，encoding将会被忽略。  

这里有些地方是需要注意的。请看下面的例子
```javascript
var fs = require('fs');
var rs = fs.createReadStream('source/file');
var ws = fs.createWriteStream('dest/file');

rs.on('data', function(chunk){
  ws.write(chunk);
});

rs.on('end', function(){
  ws.end();
});
```
write()方法是有返回值的，该方法会返回true或者false，返回true说明已经将数据写入，返回false则说明数据尚未写完。  
因此上述代码，有个严重问题，就是当写入的速度低于读取的速度时，会造成数据的丢失。正确的做法应该是，写完一段后，再读取下一段，如果没有写完，就应该暂停读取，等待数据写入完毕后再进行读取。  
当write()返回false时，便会在合适的时机触发drain事件。

看以下代码:
```javascript
var fs = require('fs');
var rs = fs.createReadStream('source/file');
var ws = fs.createWriteStream('dest/file');

rs.on('data', function(chunk){
  if(ws.write(chunk) === false){ // 尚未写完，停止读取
    rs.pause();
  }
});

ws.on('drain', function(){
  rs.resume(); // 数据已经写完，继续读取
});

rs.on('end', function(){ // 已经没有跟多数据，关闭可写流
  ws.end();
});
```

其实更简洁的方式是使用pipe的，便无需像上面那么啰嗦了
```javascript
rs.pipe(ws)
```

接下来讲一下`_write()`这个方法。同Readable一样，Writable也有个私有方法_write  
所有的Writable流的实现都必须提供一个`_write()`方法来给底层资源传输数据   
_write(chunk, encoding, callback)
chunk：被写入的资源
encoding：如果数据块是一个字符串，那么这就是编码的类型。如果是一个buffer，那么则会忽略它
callback： 当你处理完给定的数据块后调用这个函数。回调函数使用标准的callback(error)模式来表示这个写操作成功或发生了错误。  
```javascript
var Writable = require('stream').Writable
var ws = new Writable({decodeStrings: false})

var ws._write = function(chunk, enc, cb){
  console.log chunk.toString()
  cb()
}

process.stdin.pipe(ws)
```

##### Duplex
讲完了可读流和可写流，接下来讲讲Duplex流。  
这是一种“双工流”，既是可读的，也是可写的。由于javascript不具备多重继承，所以该类是继承了Readable类，并寄生于Writable类。所以实现该类的时候，需要我们去重写`_read(n)`和`_write(chunk,encoding,cb)`方法。   

new stream.Duplex(options)
options Object 同时会传递给Writable和Readable构造函数。并且包含以下属性：  
allowHalfOpen Boolean 默认为true。如果设置为false，那么流的可读的一端结束时可写的一端也会自动结束，反之亦然。  
readableObjectMode Boolean 默认为false，为流的可读的一端设置objectMode。当objectMode为true时没有效果。  
writableObjectMode Boolean 默认为false，为流的可写的一端设置objectMode。当objectMode为true时没有效果。  

##### Transform
转换流也是一个双工流，用以处理输入输出是因果相关，位于管道中间层的 Transform 是即可读也可写的。  
![transform示意图](http://7xjw3r.com1.z0.glb.clouddn.com/image/0/3e/94a8ef75e04ff809dc7b9a74e6ec1.png)

该流不仅要实现`_read()`和`_write()`方法，还有实现`_transform()`方法，并且可选的实现`_flush()`方法。   

new stream.Transform([options])
options Object 同时传递给Writable和Readable构造函数。如`ObjectMode:true`

现在我们来看看`_transform()`方法的作用。  
`_transform` 方法在每次 stream 中有数据来了之后都会被执行  
```javascript
_transform = function(chunk, encoding, cb){...}
```
在该方法中，可以进行数据处理，例如小写字母变大写。  
调用transform.push()，则可以往输出流中写入数据。给后续的输入流使用。  
仅当目前的数据块被完全消费后，才会调用回调函数。  
需要注意的是如果将数据传入回调函数的第二个参数，那么数据将会被传递给push方法，也就等价于调用了push()。下面的两种情况是等价的：
```javascript
transform.prototype._transform = function (data, encoding, callback) {
  this.push(data);
  callback();
}

transform.prototype._transform = function (data, encoding, callback) {
  callback(null, data);
}
```
_flush(callback)
在所有的数据块都被 `_transform` 方法处理过后，才会调用 `_flush` 方法。所以它的作用就是处理残留数据的。  

关于transform，这里有一篇示例，[通过Node.js Stream API 实现逐行读取的实例](http://segmentfault.com/a/1190000000740718)

更多关于stream的理解，可以翻阅这篇[文档](https://github.com/substack/stream-handbook)
