const WebSocket = require('ws')

const server = new WebSocket.Server({port:8080},()=>{
    console.log('socket server running on port 8080.')
})


server.on('connection',(socket)=>{
    console.log('客户端连接成功.')
    socket.on('message',msg => { //接收来自客户端的消息
        console.log(`message from client:${msg}`)
    })
    socket.on('close', () => {
        console.log('客户端断开连接')
    })
    socket.send('hello<from server>') //向客户端发送消息
})

