/*
 * File: main.go
 * Created Date: 2022-01-17 12:30:07
 * Author: ysj
 * Description:  链栈
 */

package main

// 定义栈节点
type Node struct {
	value interface{}
	next  *Node
}

// 定义链栈
type LinkStack struct {
	top   *Node // 栈顶节点
	count int   // 栈元素个数
}

//定义栈接口
type Stacker interface {
	Clear()                 // 清空栈
	IsEmpty() bool          // 栈是否为空
	GetTop() interface{}    // 获取栈顶元素
	Len() int               // 栈中元素个数
	Push(value interface{}) //元素入栈
	Pop() interface{}       // 元素出栈
}

func main() {

}
