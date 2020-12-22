const net = require('net')
const server = net.createServer(c=>{ //c: Socket
    //客户端连接成功后，执行下面的内容
    console.log('客户端已连接')

    //向客户端发送消息
    c.write('你好<from server>') 

    //接收来自客户端的消息
    c.on('data',data=>{
        console.log(data.toString())
    })
    
    //将客户端发来的消息写入一个可写流对象
    c.pipe(c) 

    //客户端断开连接
    c.on('end',()=>{
        console.log('客户端已断开连接')
    })
})//server实例


server.on('error',err=>{
    throw err
})

server.listen(8124,()=>{
    console.log('服务器已启动')
})