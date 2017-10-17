// 当采用json的方式时，这个包不包含解析器，包更小
const protobuf = require("protobufjs/light")
const bundle = require('./protos/bundle.json')

const root = protobuf.Root.fromJSON(bundle)
const Person = root.lookupType('Person')
const Address = root.lookupType('Address')
const Phone = root.lookupType('Phone')

const personObj = {
  name: 'Jack',
  age: 20,
  email: '123@163.com',
  sex: false
}

personObj.foo = Buffer.from('好')

const addressObj = {
  addr: 'shanghai',
  code: 1
}
const address = Address.create(addressObj)
personObj.address = address

const favorite = ['movie', 'music']
personObj.favorite = favorite

personObj.phone = {
  'workPhone': {
    phoneNum: 12345678901
  }
}

personObj.avatar = 'imageUrl'
personObj.imageUrl = 'this is image url'

personObj.pet = 1

const error = Person.verify(personObj)
if (error) return console.log(error)

const person = Person.create(personObj)

const buffer = Person.encode(person).finish()
console.log('Person Instance: ', person)
console.log('Buffer: ', buffer)

// >>>>>>>>>>>>> like client
const decodePerson = Person.decode(buffer)
console.log('decodePerson: ', Person.toObject(decodePerson, {longs: Number, enums: String, bytes: String, oneofs: true}))