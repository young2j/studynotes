const io = require('socket.io-client')

const socket = io('http://localhost:8000')


socket.on('connect',()=>{
    console.log('连接服务器成功')
    socket.on('message',msg=>{
        console.log(`message ${msg} <from server>`)
    })
})

socket.on('news',data=>{
    console.log(`receive:${data} <from server>`)
    socket.emit('customEvent','hello')
})

socket.on('volatileEvent',data=>{
    console.log(`receive: ${data} <from server>`)
})

