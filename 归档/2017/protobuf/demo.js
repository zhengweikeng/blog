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

