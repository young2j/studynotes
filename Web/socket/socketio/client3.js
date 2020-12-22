const io = require('socket.io-client')

const users = io.connect('http://localhost:8000/users')
const others = io.connect('http://localhost:8000/others')

users.on('connect',()=>{
    console.log('连接服务器成功')
})

users.on('users',data=>{
    console.log(`receive: ${data} <from server>`)
})



others.on('connect',()=>{
    console.log('连接服务器成功')
})
others.on('others',data=>{
    console.log(`receive: ${data} <from server>`)
})