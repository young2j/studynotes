/*
 * File: main.go
 * Created Date: 2022-01-17 02:59:34
 * Author: ysj
 * Description:  链式队列
 */


package main


// 队列接口
type Queuer interface {
	Clear()                  // 清空队列
	IsEmpty() bool           // 队列是否为空
	Head() interface{}       // 获取队头元素
	Lpop() interface{}       // 出队
	Rpush(value interface{}) // 入队
	Len() int                // 链表元素个数
}

// 定义队列节点
type Node struct {
	value interface{}
	next *Node
}


// 定义队列
type LinkQueue struct {
	front *Node
	rear *Node
}




func main() {
	
}



