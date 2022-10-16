const {Person, Address, Phone} = require('./protos/person_pb2')

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

const err = Person.verify(personObj)
if (err) return console.error(err)

const person = Person.create(personObj)

const buffer = Person.encode(person).finish()
console.log('Person Instance: ', person)
console.log('Buffer: ', buffer)

// >>>>>>>>>>>>> like client
const decodePerson = Person.decode(buffer)
console.log('decodePerson: ', Person.toObject(decodePerson, {longs: Number}))

