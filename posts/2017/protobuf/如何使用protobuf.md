> 阅读该文需要你对node.js的Buffer有一定的基础

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

## 使用Google官方protobuf
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
  bytes foo=4; 
}
```

我们定义了一个person的数据格式，接下来我们可以使用protoc编译这个文件，我们通过刚才拉取的proto镜像来编译。因为安装protoc需要你的机器满足安装protoc需要的依赖库，会比较麻烦，对于初学者肯定想快速使用，而不是把时间浪费在安装上，那么docker就是最好的选择了。
```bash
$ cd ./protos

$ docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --js_out=import_style=commonjs,binary:. -I. *.proto
```
编译成commonjs格式的js代码，供node.js使用。

如果是前端可以采用Closure imports方式编译。
```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --js_out=library=base.js.,binary:. -I. *.proto
```

由于我们是在服务端使用的，所以我们采取了第一种commonjs的编译方式，编译之后即可看到base_pb.js文件  
打开文件我们可以看到第一行代码引入了谷歌的protobuf的js库，因此这里我们需要自己手动install下这个库

```bash
$ npm install google-protobuf
```

接下来我们就可以直接使用，创建demo.js
```javascript
// demo.js
const {Person, Address, Phone} = require('./protos/person_pb')

const person = new Person()
person.setName('Tom')
person.setAge(10)
person.setEmail('test@outlook.com')
const buf = Buffer.from('好')
person.setFoo(buf.toString('base64'))

const address = new Address()
address.setAddr('shanghai')
address.setCode(1)
person.setAddress(address)

person.setFavoriteList(['movie', 'music'])

const phone = new Phone()
phone.setPhoneNum(12345678901)
const phoneMap = person.getPhoneMap()
phoneMap.set('workPhone', phone)

person.setSex(false)

person.setImageUrl('www.baidu.com')

person.setPet(1)

console.log('Name: ', person.getName())
console.log('Age: ', person.getAge())
console.log('Email: ', person.getEmail())
console.log('Foo:', person.getFoo())
console.log('Foo as: base-64', person.getFoo_asB64())
console.log('Foo as: Uint8Array', person.getFoo_asU8())
console.log('Address: ', person.getAddress().toObject().addr)
console.log('Favorite: ', person.getFavoriteList())
console.log('workPhone', person.getPhoneMap().get('workPhone').getPhoneNum())
console.log('Sex: ', person.getSex() ? 'male' : 'female')
console.log('Avatar: ', person.getImageUrl())
console.log('Pet: ', person.getPet())
console.log(person.toObject())

// 构造成一个buffer对象
// const personBuff = person.serializeBinary()
// console.log('serialize: ', personBuff)

// 解析客户端传递过来的二进制流，并且反序列化成一个对象
// const deserPerson = Person.deserializeBinary(personBuff)
// console.log('deserialize: ', deserPerson.toObject())
// console.log('Foo: ', Buffer.from(person.getFoo_asB64(), 'base64').toString())
```

person实例即为proto文件中定义的message实例，该实例有一些方法供我们使用，主要是一些获取字段的方法：
1. getter和setter方法，用于获取字段和为字段赋值，如果字段类型为bytes，则还会有`getXxx_asB64`和`getXxx_asU8`两个方法用于获取base64和Uint8Array格式的数据
1. toObject，会返回proto中的message所定义的object对象
1. closeMessage，对message实例和它的属性进行深拷贝

我们在对proto文件编译的时候，有这么一段指令：`--js_out=import_style=commonjs,binary:.`  
其中有个binary的定义，这是protoc提供给我们的编译的备选项，binary这个备选项可以使得生成的js文件能够将对象编译成proto，也可以将proto二进制数据反序列化成对象，即会提供如下几个方法：
1. deserializeBinary，反序列化一个protocol buffer二进制数据成一个object对象
1. serializeBinary，序列化一个对象成protocol buffer二进制数据

具体参考 [文档](https://developers.google.com/protocol-buffers/docs/reference/javascript-generated) 中关于Messages的描述。

使用时发现一个问题，如果message中有个字段类型是bytes的话，在js中使用setter时需要先变成将buffer转成base64再set，不然serializeBinary的时候会报错。
```javascript
const buf = Buffer.from('好')
person.setFoo(buf)
// 此时会报错
person.serializeBinary()
```
