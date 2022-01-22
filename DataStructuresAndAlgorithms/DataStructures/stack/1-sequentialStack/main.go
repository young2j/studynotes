/*
 * File: main.go
 * Created Date: 2022-01-17 12:29:47
 * Author: ysj
 * Description:  顺序栈—数组实现
 */
package main

import "fmt"

//定义栈接口
type Stacker interface {
	Clear()                 // 清空栈
	IsEmpty() bool          // 栈是否为空
	IsFull() bool           // 栈是否已满
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

// 初始化顺序栈，空栈，栈顶top为-1
func NewSequentialStack() Stacker {
	return &SequentialStack{
		top:  -1,
		data: [MAXSIZE]interface{}{},
	}
}

// 清空栈
func (s *SequentialStack) Clear() {
	s.top = -1
	for i := 0; i < MAXSIZE; i++ {
		s.data[i] = nil
	}
}

// 是否为空栈
func (s *SequentialStack) IsEmpty() bool {
	return s.top == -1
}

// 栈是否已满
func (s *SequentialStack) IsFull() bool {
	return s.top >= MAXSIZE -1
}



// 获取栈顶元素
func (s *SequentialStack) GetTop() interface{} {
	return s.data[s.top]
}

// 获取栈中元素个数
func (s *SequentialStack) Len() int {
	return s.top + 1
}

// 元素入栈
func (s *SequentialStack) Push(value interface{}) {
	if s.IsFull() {
		panic("stack is full")
	}

	s.data[s.top+1] = value
	s.top++
}

// 元素出栈
func (s *SequentialStack) Pop() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	value := s.data[s.top]
	s.top--
	return value
}

func main() {
	sqStack := NewSequentialStack()
	// sqStack.Pop() // panic: stack is empty
	values := []int{3, 5, 6, 7, 2}
	for _, v := range values {
		fmt.Println("sqStack push", v)
		sqStack.Push(v)
	}
	fmt.Println("Len(sqStack):", sqStack.Len())
	fmt.Println("sqStack top value is:", sqStack.GetTop())
	fmt.Println("sqStack pop:", sqStack.Pop())
	fmt.Println("sqStack pop:", sqStack.Pop())
	fmt.Println("Len(sqStack):", sqStack.Len())
	sqStack.Clear()
	fmt.Println("sqStack clear...Len(sqStack):", sqStack.Len())
}
