setTimeout(function () {
  console.log(1)
})

setImmediate(function () {
  console.log(2)
}) 

process.nextTick(function () {
  console.log(3)
}) 

const p = new Promise(function (resolve, reject) {
  console.log(4)
  resolve()
})
p.then(function () {
  console.log(5)
})

console.log(6) 