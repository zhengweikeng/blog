node.js原生支持了加密的库，例如md5加密
```javascript
var crypto = require('crypto')
var plainText = 'hello world'
var ciphertext ＝ crypto.createHash('md5').update(plainText).digest('hex')
// 5eb63bbbe01eeed093cb22bb8f5acdc3
```

这一串api相信很多人都还是不理解的，那我们根据官方api来学习加密和解密库crypto

## crypto.createHash(algorithm)
首先需要明确一点，这是一种使用指定散列算法将明文转换成不可逆的散列值，由于它无法被解密，因此不算是加密算法。也因此我们只会将它用于一些信息的认证。而不能用来消息传递的加密解密。

创建一个hash对象，该对象可用于根据指定的加密算法algorithm来生成密文。  
这里传入的参数algorithm的备选值，是根据程序所在平台的OpenSSL版本所决定的，如果OpenSSL支持该算法，则可以使用。

可以根据`openssl list-message-digest-algorithms`来查看支持的加密算法。  
也可以通过`crypto.getHashes()`来查看支持的加密算法。

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

## crypto.createHmac(algorithm, key)
该方法会将密钥和给定的key值进行哈希计算，因此也是一种不可逆的哈希值的生成，但是它可以防止彩虹表等的攻击。

使用方式和之前的createHash一样，因此就不解释了。放个官方的案例
```javascript
const crypto = require('crypto');
const hmac = crypto.createHmac('sha256', 'a secret');

hmac.update('some data to hash');
console.log(hmac.digest('hex'));
  // Prints:
  //   7fd04df92f636fd450bc841c9418e5825c17f33ad9c87c518115a45971f7f77e
```

## 加密和解密
像登陆密码这种是不需要解密的，因此hash算法就可以满足。但是很多时候，我们是将数据加密后在网络上传输，服务器获取后进行解密使用。

目前加密的方式可分为对称加密和非对称加密，具体差别可以自行网上搜索。

node.js的crypto支持对称加密，我们可以通过`crypto.getCiphers()`来查看支持的对称加密算法。

对称加密包含如下四个类，Cipher, Decipher,Sign, and Verify，即加密，解密，签名，验证。

## Cipher, Decipher
### 加密类Cipher
创建一个Cipher类的实例需要通过`crypto.createCipher()`或者`crypto.createCipheriv()`来实现。
```javascript
const cipher = crypto.createCipher('aes192', 'test1234')
```
其中createCipher的第二个参数用于生成加密密钥。

同理，也必须通过update添加待加密的数据。
#### [cipher.update(data[, input_encoding][, output_encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_cipher_update_data_input_encoding_output_encoding)
如果第一个参数data是一个字符串，那么第二个参数input_encoding必须是utf8，ascii或者binary中的一个。如果input_encoding不传，则data必须是Buffer对象。

第三个参数output_encoding用于该方法的返回值，若该参数值为binary，base64或者hex中的一个，则返回该值编码后的字符串值。若不传，则返回一个Buffer对象。
```javascript
var encrypted = cipher.update('hello world', 'utf8', 'hex')
```

#### [cipher.final([output_encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_cipher_final_output_encoding)
用于生成剩余的加密内容。若output_encoding参数值为binary，base64或者hex中的一个，则返回该值编码后的字符串值。若不传，则返回一个Buffer对象。
```javascript
encrypted += cipher.final('hex')
// e93081f70710bdee2b60e231e663566b
```

### 解密类Decipher
解密就是一个逆过程，拥有的方法和加密类Cipher一样。直接看代码。
```javascript
var encrypted = 'e93081f70710bdee2b60e231e663566b'
const decipher ＝ crypto.createDecipher('aes192', 'test1234')
var decrypted = decipher.update(encrypted, 'hex', 'utf8')
decrypted += decipher.final('utf8')
// hello world
```

## Sign, Verify
### 签名Sign
数据在网络传输时容易被篡改，因此我们获取到数据时，需要通过某种机制检验数据是否遭受篡改。  
我们可以模拟一个公钥和私钥加密的例子。
我们知道公钥和私钥是成对出现的，公钥需要根据私钥按照一定的规则产生。这里假设已经存在了私钥和公钥，privateKey和pubKey。
```javascript
var sign = crypto.createSign('RSA-SHA256')
```
同理，也是update方法，有两个参数，第一个参数是一个输入编码。
```javascript
var data = 'hello world'
sign.update(data)
```
sign方法用于生成签名，第一个参数是私钥的值，第二个参数也是输出编码。
```javascript
var sig = sign.sign(privateKey, 'hex')
```
上述过程简而言之就是使用私钥进行数据签名。传输是便可将公钥，原始数据，还有签名后的数据发送过去。

### 验证Verify
也即是一个逆过程，看代码。
```javascript
var data = 'hello world'
var verify = crypto.createVerify('RSA-SHA256')
verify.update(data)
verify.verify(pubKey, sig, 'utf8')
// 返回true or false
```
上述过程就是，通过传递过来的签名后的数据，使用对方的公钥进行解密，将解密后的数据和原始数据进行比对。

## Diffie–Hellman key exchange(迪菲－赫尔曼密钥交换)
它可以让双方在完全没有对方任何预先信息的条件下通过不安全信道创建起一个密钥。提出这种加密方式的两个人迪菲和赫尔曼也是图灵奖的获得者。

创建DiffieHellman类的方法有如下两个
```javascript
crypto.createDiffieHellman(prime[, prime_encoding][, generator][, generator_encoding])
crypto.createDiffieHellman(prime_length[, generator])
```
也即可以通过传入一个素数或者素数长度、结合生成器（不传则为2）因素来产生一个DiffieHellman类实例。
```javascript
var alice = crypto.createDiffieHellman(2048)
```
### [diffieHellman.generateKeys([encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_generatekeys_encoding)
用于生成一个私钥和公钥的键值对，根据指定的编码返回一个公钥。该公钥将会传递给另外一方，用于对称加密。
```javascript
const alice_key = alice.generateKeys()
```
### [diffieHellman.getPrime([encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_getprime_encoding)
获取创建diffieHellman实例的素数
```javascript
const alice_prime = alice.getPrime()
```
### [diffieHellman.getGenerator([encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_getgenerator_encoding)
获取创建diffieHellman实例的生成数
```javascript
const alice_generator = alice.getGenerator()
```
### [diffieHellman.getPrivateKey([encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_getprime_encoding) and [diffieHellman.getPublicKey([encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_getpublickey_encoding)
获取私钥和公钥
### [diffieHellman.setPrivateKey(private_key[, encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_setprivatekey_private_key_encoding) and [diffieHellman.setPublicKey(public_key[, encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_setpublickey_public_key_encoding)
设置私钥和公钥
### [diffieHellman.computeSecret(other_public_key[, input_encoding][, output_encoding])](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_diffiehellman_computesecret_other_public_key_input_encoding_output_encoding)
根据另外一个diffieHellman实例生成的公钥来计算生成一个密文。
```javascript
const bob = crypto.createDiffieHellman(alice_prime, alice_generator)
const bob_key = bob.generateKeys()

const alice_secret = alice.computeSecret(bob_key)
const bob_secret = alice.computeSecret(alice_key)

// alice_secret === bob_secret    
// true
```
### [crypto.getDiffieHellman(group_name)](https://nodejs.org/dist/latest-v6.x/docs/api/crypto.html#crypto_crypto_getdiffiehellman_group_name)
这是另外一种创建DiffieHellman实例的方式，不过有些不同。它可以根据‘modp1’, ‘modp2’, ‘modp5’ (defined in RFC 2412) and ‘modp14’, ‘modp15’, ‘modp16’, ‘modp17’, ‘modp18’ (defined in RFC 3526)来创建实例。

它创建后不能改变公钥和私钥。

使用它的优点是通常就是直接使用它不用交换生成key，只需要在握手前使用一样的group系数即可，节约了大家的处理时间和握手时间。
```javascript
var alice = crypto.getDiffieHellman('modp5')
var bob = crypto.getDiffieHellman('modp5')

alice.generateKeys()
bob.generateKeys()

var alice_secret = alice.computeSecret(bob.getPublicKey(), 'binary', 'hex')
var bob_secret = bob.computeSecret(alice.getPublicKey(), 'binary', 'hex')

// alice_secret === bob_secret    
// true
```
还有一种Elliptic Curve Diffie-Hellman(ECDH)的加密算法，用法都是一样的。

## crypto.pbkdf2()
我们知道md5这种哈希算法很容易通过彩虹表的方式破解，因此我们需要对我们的密码进行加盐后再加密。

上面我们知道node提供了我们hmac这种加盐api，但是我们可以更加简化操作，使用pbkdf2函数来做
```javascript
var text = 'hello world'
var salt = '123456789abcdefg'

crypto.pbkdf2(text, salt, 4096, 256, 'md5', (err, hash) => {
  if (err) { throw err; }
  console.log(hash.toString('hex'));
})
// bd99d067b97fc642c77e327909281c4ad1...
```
当然，如果我们每次加盐都是使用一样的salt，也是会有不安全的成分，所以我们可以生成一个随机salt。
```javascript
var text = 'hello world'
crypto.randomBytes(128, (err, salt) => {
  if (err) { throw err;}
  console.log(salt)
  
  crypto.pbkdf2(text, salt, 4096, 256, 'md5', (err, hash) => {
    if (err) { throw err; }
    console.log(hash.toString('hex'));
  })
})
```

## 参考文章
[Node.js加密算法库Crypto](http://blog.fens.me/nodejs-crypto/)  
[浅谈nodejs中的Crypto模块](https://cnodejs.org/topic/504061d7fef591855112bab5)
