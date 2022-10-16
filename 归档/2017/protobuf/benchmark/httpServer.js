const express = require('express')
const app = express()

app.get('/', function (req, res) {
  const params = req.query
  res.send(params)
})

app.listen(3001)
