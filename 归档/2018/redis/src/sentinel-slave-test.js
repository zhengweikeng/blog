const Redis = require('ioredis')

const preferredSlaves = [
  { ip: '127.0.0.1', port: '6479', prio: 1 },
];
const redis = new Redis({
  sentinels: [
    { host: '127.0.0.1', port: 26379 }, 
    { host: '127.0.0.1', port: 26380 },
    { host: '127.0.0.1', port: 26381 }
  ],
  name: 'mymaster',
  role: 'slave',
  preferredSlaves
});

setInterval(async () => {
  const val = await redis.get('hello')
  const infoVal = await redis.info('server')
  console.log(val)
  console.log(infoVal)
}, 3000)