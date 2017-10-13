## 什么是protobuf？
protobuf是google开源项目，是一种数据交换的格式，主要用于不同端（客户端和服务器端），不同平台（如java和node.js）的数据通信，它相比xml和json，对数据进行了压缩，使得传输更加高效。

举个例子，客户端和服务器端采用tcp连接，客户端传递数据时，需要将对象，序列化成二进制传递出去。  
服务器端接收后，将这段二进制解析出来。而这个序列化和反序列化的过程需要按照某种协议来做，protobuf就扮演了这样一个协议。

## 安装protobuf
### 方法一
到 [protobuf github](https://github.com/google/protobuf/releases)下载对应平台的release包，解压后安装

```bash
$ ./configure --prefix=/usr/local/protobuf
$ make
$ make check
$ make install
```

### 方法二
直接使用protobuf docker镜像，可以无需安装即可使用，更加简单的方便

```
docker pull znly/protoc
```

具体参考 [znly/protoc](https://hub.docker.com/r/znly/protoc/) 文档

## 使用protobuf
我们会使用Node.js来编写案例，并且采用docker镜像的方式来编译proto文件。

假设有这么一个案例，客户端和服务器端采用websocket建立连接，此时客户端用户填写了一个表单需要提交给服务器端，如果采用json的方式传递数据，会比较占流量，对用户也不好，因此我们将数据变成二进制，使得传输更加的轻量，提高了传输效率。

在node中，该二进制即为buffer。

首先我们需要定义一个数据格式，用于数据通信
```proto
syntax = "proto3";

message Person {
  string name=1;
  int32 age=2;
  string email=3;
}
```

我们定义了一个person的数据格式，接下来我们可以使用protoc编译这个文件，我们通过刚才拉取的proto镜像来编译
```bash
$ cd ./protos

$ docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --js_out=import_style=commonjs,binary:. -I. *.proto
```
编译成commonjs格式的js代码，供node.js使用。如果是前端可以采用Closure imports方式编译。
```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --js_out=library=base.js.,binary:. -I. *.proto
```

接下来我们就可以直接使用
```javascript
const base = require('./protos/base_pb')

const person = new base.Person()
person.setName('Tom')
person.setAge(10)
person.setEmail('test@outlook.com')

console.log('Name: ', person.getName())
console.log('Age: ', person.getAge())
console.log('Email: ', person.getEmail())
console.log(person.toObject())

// 构造一个buffer对象，即二进制数据流
const personBuff = person.serializeBinary()
console.log('serialize: ', personBuff)
clientSocket.write(personBuff)

// 解析客户端传递过来的二进制流，并且反序列化成一个对象
const deserPerson = base.Person.deserializeBinary(personBuff)
console.log('deserialize: ', deserPerson.toObject())
```

person实例即为proto文件中定义的message实例，该实例有一些方法供我们使用，如上述的序列化和反序列化，还有对应的getter和setter方法，具体参考 [文档](https://developers.google.com/protocol-buffers/docs/reference/javascript-generated)

