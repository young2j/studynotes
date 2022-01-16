/*
 * File: main.go
 * Created Date: 2022-01-17 12:29:47
 * Author: ysj
 * Description:  顺序栈—数组实现
 */
package main

//定义栈接口
type Stacker interface {
	Clear()                 // 清空栈
	IsEmpty() bool          // 栈是否为空
	GetTop() interface{}    // 获取栈顶元素
	Len() int               // 栈中元素个数
	Push(value interface{}) //元素入栈
	Pop() interface{}       // 元素出栈
}

const MAXSIZE = 20

// 定义顺序栈结构
type SequentialStack struct {
	top  int // 栈顶
	data [MAXSIZE]interface{}
}

func main() {

}
