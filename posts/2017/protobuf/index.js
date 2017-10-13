const base = require('./protos/base_pb')

const person = new base.Person()
person.setName('Tom')
person.setAge(10)
person.setEmail('test@outlook.com')

console.log('Name: ', person.getName())
console.log('Age: ', person.getAge())
console.log('Email: ', person.getEmail())
console.log(person.toObject())

const personBuff = person.serializeBinary()
console.log('serialize: ', personBuff)

const deserPerson = base.Person.deserializeBinary(personBuff)
console.log('deserialize: ', deserPerson.toObject());

