const net = require('net')
const protobuf = require("protobufjs/light")
const {personBuf} = require('../decodeIoDemo3')
const bundle = require('../protos/bundle.json')

const root = protobuf.Root.fromJSON(bundle)
const Person = root.lookupType('Person')

const client = net.connect(3000, '127.0.0.1')
let count = 0
const MAX_COUNT = 5000
const resTimes = []

client.on('connect', () => {
  client.reqTime = Date.now()
  // console.log(client.reqTime)
  client.write(personBuf)
})
client.on('data', chunk => {
  const decodePerson = Person.decode(chunk)
  const resTime = Date.now() - client.reqTime
  // console.log(resTime);
  resTimes.push(resTime)
  // console.log('decodePerson: ', Person.toObject(decodePerson, {longs: Number, enums: String, bytes: String, oneofs: true}))

  if (count++ <= MAX_COUNT) {
    client.reqTime = Date.now()
    // console.log(client.reqTime)
    client.write(personBuf)
  } else {
    const result = resTimes.reduce((result, time) => {
      result += time
      return result
    }, 0)
    // console.log(result)
    console.log(result / MAX_COUNT)
    process.exit(0)
  }
})
