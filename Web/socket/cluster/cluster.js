//nodejs 多进程
const cluster = require('cluster')
const http = require('http')
const numCPUS = require('os').cpus().length

if (cluster.isMaster){ //master process
    console.log(`主进程${process.pid}正在运行`)

    //衍生工作进程
    for (let i=0;i<numCPUS;i++){
        cluster.fork() //内部是child_process.fork()，可以通过IPC通信
    }

    cluster.on('fork',(worker)=>{
        console.log(`fork ${worker.process.pid} 已启动`)
    }) //同else中的console.log

    cluster.on('exit',(worker,code,signal)=>{
        console.log(`${worker.process.pid} 已退出`)
    })

    for (const id in cluster.workers) {
        // console.log(cluster.workers) //cluster.workers 是个对象
        cluster.workers[id].on('message',msg=>{
            console.log(`msg:${msg} from:${cluster.workers[id].process.pid}`)
        })
    }
} else { //worker process
    http.createServer((req,res)=>{
        res.writeHead(200)
        res.end('hello')
        console.log(`req from ${process.pid}`)
        //向主进程发送消息
        process.send('xxxx')
    }).listen(8000) //每个worker共享 8000 端口

    console.log(`${process.pid} 已启动`) 
}

