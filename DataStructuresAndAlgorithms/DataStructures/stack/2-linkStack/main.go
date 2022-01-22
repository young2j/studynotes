/*
 * File: main.go
 * Created Date: 2022-01-17 12:30:07
 * Author: ysj
 * Description:  链栈
 */

package main

import "fmt"

//定义栈接口
type Stacker interface {
	Clear()                 // 清空栈
	IsEmpty() bool          // 栈是否为空
	GetTop() interface{}    // 获取栈顶元素
	Len() int               // 栈中元素个数
	Push(value interface{}) //元素入栈
	Pop() interface{}       // 元素出栈
}

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

// 初始化链栈，top==nil代表空栈，始终以链表第一个元素为栈顶
func NewLinkStack() Stacker {
	return &LinkStack{}
}

// 清空链栈
func (s *LinkStack) Clear() {
	s.top = nil
	s.count = 0
}

// 栈是否为空
func (s *LinkStack) IsEmpty() bool {
	return s.top == nil
}

// 获取栈顶元素
func (s *LinkStack) GetTop() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.top.value
}

// 获取栈元素个数
func (s *LinkStack) Len() int {
	return s.count
}

// 入栈
func (s *LinkStack) Push(value interface{}) {
	node := &Node{
		value: value,
		next:  s.top,
	}
	s.top = node
	s.count++
}

// 出栈
func (s *LinkStack) Pop() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	top := s.top
	s.top = s.top.next
	s.count--
	return top.value
}

func main() {
	linkStack := NewLinkStack()
	// linkStack.Pop() // panic: stack is empty
	values := []int{3, 5, 6, 7, 2}
	for _, v := range values {
		fmt.Println("linkStack push", v)
		linkStack.Push(v)
	}
	fmt.Println("Len(linkStack):", linkStack.Len())
	fmt.Println("linkStack top value is:", linkStack.GetTop())
	fmt.Println("linkStack pop:", linkStack.Pop())
	fmt.Println("linkStack pop:", linkStack.Pop())
	fmt.Println("Len(linkStack):", linkStack.Len())
	linkStack.Clear()
	fmt.Println("linkStack clear...Len(linkStack):", linkStack.Len())
}
