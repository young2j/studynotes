/** 
//使用简单http
const server = require('http').createServer((req,res)=>{
    res.writeHead(200)
    res.end('socket server')
}) 
*/

//使用express
const app = require('express')()
const server = require('http').Server(app)
const io = require('socket.io')(server)


server.listen(8000,()=>{
    console.log('server running on 8000')
})

app.get('/',(req,res)=>{
    res.send('socket server')
})

io.on('connection',socket=>{
    console.log(`客户端${socket.id}连接成功`)
    
    io.emit('allClient',{}) //向所有客户端发起事件
    socket.broadcast.emit('broadcast','message to everyone') //广播:向所有客户端发送消息
    
    socket.emit('news','hello') //向所有监听了news事件的客户端发送消息
    socket.volatile.emit('volatileEvent','volatile message') //发送一个不确定能否准确到达的消息

    socket.on('customEvent',data=>{
        console.log(`receive:${data} <from client>`)
    })

    socket.on('acknowledge',(data,fn)=>{
        fn(data)
    }) //消息确认机制：客户端会需要确认向服务器端发送的事件是否在服务器端正确执行了。-就是传递一个回调作为消息在服务器端执行

    socket.on('disconnect',()=>{ //客户端调用socket.disconnect()触发
        io.emit('user disconnect')
    })
})

//自定义路由，或叫命名空间，默认是根路径/
const users = io
    .of('/users')
    .on('connection',socket=>{
        console.log(`${socket.id} 连接成功`)
        socket.emit('users','message only /users will get')
        users.emit('users','message all /users will get')
    })

const others = io
    .of('/others')
    .on('connection',socket=>{
        console.log(`${socket.id} 连接成功`)
        socket.emit('others','message from /others')
    })


/**
// 客户端
io.connect(url) //客户端连接上服务器端
socket.on('eventName', msg => {}) //客户端监听服务器端事件
socket.emit('eventName', msg) //客户端向服务器端发送数据
socket.disconnect()    //客户端断开链接
// 服务端
socket.on('eventName', msg => {}) //服务器端监听客户端emit的事件，事件名称可以和客户端是重复的，但是并没有任何关联。socket.io内置了一些事件比如connection，disconnect，exit事件，业务中错误处理需要用到。
socket.emit('eventName', msg) //服务端各自的socket向各自的客户端发送数据
socket.broadcast('eventName', msg) //服务端向其他客户端发送消息，不包括自己的客户端
socket.join(channel) //创建一个频道（非常有用，尤其做分频道的时候，比如斗地主这种实时棋牌游戏）
io.sockets.in(channel) //加入一个频道
socket.broadcast.to(channel).emit('eventName', msg) //向一个频道发送消息，不包括自己
io.sockets.adapter.rooms //获取所有的频道
*/