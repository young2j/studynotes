/*
 * File: main.go
 * Created Date: 2022-01-16 02:46:57
 * Author: ysj
 * Description:  循环链表-和单链表基本一致，尾节点指向头节点
 */

package main

import (
	"errors"
	"fmt"
)

// 定义链表节点
type Node struct {
	value interface{}
	next  *Node
}

// 定义循环链表
type LoopLinkList Node

// 定义循环链表接口
type LoopLinkLister interface {
	Len() int
	Get(i int) (interface{}, error)
	Insert(i int, value interface{})
	Delete(i int)
	Reverse()
}

// 初始化循环链表
// head.next -> head
func NewLoopLinkList() LoopLinkLister {
	head := &LoopLinkList{}
	head.next = (*Node)(head)
	return head
}

// 获取链表长度
func (l *LoopLinkList) Len() int {
	node := l.next
	length := 0
	for node != (*Node)(l) {
		node = node.next
		length++
	}
	return length
}

//获取i位置的值
func (l *LoopLinkList) Get(i int) (interface{}, error) {
	// 查找i次
	node := l.next
	for node != (*Node)(l) && i > 0 {
		node = node.next
		i--
	}
	if node == (*Node)(l) {
		return nil, errors.New("i: index out of range")
	}
	return node.value, nil
}

/*
在i位置插入值value
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> head
    i					   0            1            2            3            4
假如要在i=2的位置插入一个节点:
head.next -> Node.next -> Node.next    Node.next -> Node.next -> Node.next -> head
    i					   0            1  ↘️     ↗️  3(2')        4            5
														     Node.next
																	   2
1. 判断插入位置是否为第一个，如果是，则新的节点next指向头节点的next，头节点的next指向新的节点；
2. 如果不是，则
3. 找到i-1节点，将新加入的节点next指向i-1节点的next
4. 再将i-1节点的next指向新的节点
*/
func (l *LoopLinkList) Insert(i int, value interface{}) {
	// 新节点
	node := &Node{
		value: value,
	}
	// 是否是第一个位置
	if i == 0 {
		node.next = l.next
		l.next = node
		return
	}

	// 移动到i-1节点处
	preNode := l.next
	for preNode != (*Node)(l) && i > 1 {
		preNode = preNode.next
		i--
	}
	if preNode == (*Node)(l) {
		panic("i: index out of range")
	}

	node.next = preNode.next
	preNode.next = node
}

/*
删除i位置的值
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> head
    i					   0            1            2            3            4
假如要删除i=2的位置的节点:
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> head
    i					   0            1            3            4
														     Node.next
																	   2(X)
1. 判断删除位置是否为第一个，如果是，则头节点的next指向待删除的节点next；
2. 如果不是，则
3. 找到i-1节点，将i-1节点的next指向待删除的节点next；
*/
func (l *LoopLinkList) Delete(i int) {
	if i == 0 {
		if l.next == (*Node)(l) {
			return
		}
		l.next = l.next.next
		return
	}

	node := l.next
	for node != (*Node)(l) && i > 1 {
		node = node.next
		i--
	}
	if node == (*Node)(l) || node.next == (*Node)(l) {
		panic("i: index out of range")
	}
	node.next = node.next.next
}

/*
链表反转
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> head
    i					   0            1            2            3            4
   head  <-  Node.next <- Node.next <- Node.next <- Node.next <- Node.next <- head.next
    i  			     0            1            2            3            4
*/
func (l *LoopLinkList) Reverse() {
	preNode := (*Node)(l)
	curNode := l.next
	for curNode != (*Node)(l) {
		nextNode := curNode.next
		curNode.next = preNode
		if nextNode == (*Node)(l) { // 最后一个节点
			l.next = curNode
			return
		}
		preNode = curNode
		curNode = nextNode
	}
}

func main() {
	loopList := NewLoopLinkList()
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		loopList.Insert(i, v)
	}
	fmt.Println("len:", loopList.Len())
	for i := 0; i < loopList.Len(); i++ {
		v, _ := loopList.Get(i)
		fmt.Println(i, ":", v)
	}
	fmt.Println("-----删除i=3的数-----")
	loopList.Delete(3)
	fmt.Println("len:", loopList.Len())
	for i := 0; i < loopList.Len(); i++ {
		v, _ := loopList.Get(i)
		fmt.Println(i, ":", v)
	}

	fmt.Println("-----链表反转-----")
	loopList.Reverse()
	fmt.Println("len:", loopList.Len())
	for i := 0; i < loopList.Len(); i++ {
		v, _ := loopList.Get(i)
		fmt.Println(i, ":", v)
	}
}
