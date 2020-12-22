const net = require('net')
//net.connect()是net.createConnection()的别名
const client = net.createConnection({port:8124},()=>{
    //连接到服务器时，执行下面内容
    console.log('已连接到服务器')

    //向服务端发送消息
    client.write('你好<from client>') 
}) //Socket实例


//接收来自服务端的消息
client.on('data',data=>{ 
    console.log(data.toString())
    // client.end()
})

//断开连接
client.on('end',()=>{
    console.log('已从服务器断开')
})
