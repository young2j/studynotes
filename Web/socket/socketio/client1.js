const io = require('socket.io-client')

const socket = io('http://localhost:8000')


socket.on('connect',()=>{
    console.log('连接服务器成功')
    socket.emit('acknowledge','message',(data)=>{
        console.log(`handle ${data} in server`)
    }) //消息确认机制
})

socket.on('news',data=>{
    console.log(`receive:${data} <from server>`)
    socket.emit('customEvent','hello')
})

socket.on('broadcast',data=>{
    console.log(`receive: ${data} <from server>`)
})