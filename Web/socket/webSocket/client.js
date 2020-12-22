const WebSocket = require('ws')

const socket = new WebSocket('ws://localhost:8080')

socket.onopen = ()=>{
    console.log('连接服务端成功')
    socket.send('hello<from client>')
}

// socket.onmessage = e=>{
//     console.log('Message from server: ',e.data)
// }
//等同于
socket.on('message',data=>{
    console.log('Message from server: ',data)
})


socket.onclose = ()=>{
    console.log('服务器关闭')
}