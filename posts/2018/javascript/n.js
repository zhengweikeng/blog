var fs = require('fs')

fs.readFile('./n.js', () => {
  setTimeout(function () {
   console.log(1)
  }, 0)

  setImmediate(function () {
    console.log(2)
  })

  process.nextTick(function () {
    console.log(3)
  })
})