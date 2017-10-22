const superagent = require('superagent')
const {personObj} = require('../decodeIoDemo3')

let count = 0
const MAX_COUNT = 5000
const resTimes = []

const request = () => {
  let t1 = Date.now()
  superagent.get('http://127.0.0.1:3001')
  .query(personObj)
  .end((err, res) => {
    // console.log('t1', t1);
    const resTime = Date.now() - t1
    // console.log(resTime);
    resTimes.push(resTime)
    // console.log(resTime);

    if (count++ <= MAX_COUNT) {
      request()
    } else {
      // resTimes.shift()
      const result = resTimes.reduce((result, time) => {
        result += time
        return result
      }, 0)
      console.log(result / MAX_COUNT)
      process.exit(0)
    }
  })
}

request()
