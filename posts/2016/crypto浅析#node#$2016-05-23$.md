node.js原生支持了加密的库，例如md5加密
```javascript
var crypto = require('crypto')
var plainText = 'hello world'
var ciphertext ＝ crypto.createHash('md5').update(plainText).digest('hex')
// 5eb63bbbe01eeed093cb22bb8f5acdc3
```

这一串api相信很多人都还是不理解的，那我们根据官方api来学习加密和解密库crypto

## crypto.createHash(algorithm)
首先需要明确一点，这是一种使用指定散列算法将明文转换成不可逆的散列值，因此不算是加密算法。

创建一个hash对象，该对象可用于根据指定的加密算法algorithm来生成密文。  
这里传入的参数algorithm的备选值，是根据程序所在平台的OpenSSL版本所决定的，如果OpenSSL支持该算法，则可以使用。

可以根据`openssl list-message-digest-algorithms`来查看支持的加密算法。

该方法返回一个[Hash类](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_class_hash)的实例。   
该类继承了stream，因此可以通过流的方式来加密数据。如官网的例子
```javascript
const crypto = require('crypto');
const hash = crypto.createHash('sha256');

hash.on('readable', () => {
  var data = hash.read();
  if (data)
    console.log(data.toString('hex'));
    // Prints:
    //   6a2da20943931e9834fc12cfe5bb47bbd9ae43489a30726962b576f4e3993e50
});

hash.write('some data to hash');
hash.end();

// 或者使用管道pipe的方式
const crypto = require('crypto');
const fs = require('fs');
const hash = crypto.createHash('sha256');

const input = fs.createReadStream('test.js');
input.pipe(hash).pipe(process.stdout);
```

另外也可以通过调用update和digest方法的方式来加密数据。

### [hash.update(data[, input_encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_hash_update_data_input_encoding)
需要先调用update方法来添加待加密的数据

该方法会返回hash对象自身，因此update可以被调用多次，传入的加密数据最终会被连接起来。
```javascript
const crypto = require('crypto');
const hash = crypto.createHash('sha256');

hash.update('hello');
hash.update('world');
// 和调用 hash.update('hello').update('world') 一样
// 和调用 hash.update('helloworld') 一样
```
同时update还有第二个参数，用于指定待加密数据的编码，如果数据是字符串，则可以是utf8、ascii或者binary，默认是utf8。如果数据是buffer对象，则该参数被忽略。

### [hash.digest([encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_hash_digest_encoding)
将以上update进去的数据进行加密计算。接受一个编码参数。备选值有hex、binary或者base64。如果该参数没有传，则返回一个Buffer对象。  
该方法只能调用一次，否则会报错。

文章开头的例子就是加密后返回一串用16进制编码后的密文。

正常来说，我们已经不会简单的使用hash算法来加密数据，建立庞大的彩虹表已经是有可能进行破解。因此我们必须加密前进行加盐。
