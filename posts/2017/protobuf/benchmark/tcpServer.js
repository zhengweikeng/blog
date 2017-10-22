const net = require('net')

const server = net.createServer(socket => {
  socket.on('data', chunk => {
    console.log(chunk);
    socket.write(chunk)
  })
})

server.listen('3000')
