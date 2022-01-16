/*
 * File: main.go
 * Created Date: 2021-12-13 04:20:45
 * Author: ysj
 * Description:  单链表
 */

package main

import (
	"errors"
	"fmt"
)

//单链表接口
type SingleLinkLister interface {
	Len() int
	Get(i int) (interface{}, error)
	Insert(i int, value interface{})
	Delete(i int)
	Reverse()
}

// 定义链表节点
type Node struct {
	value interface{}
	next  *Node
}

type SingleLinkList Node

/*
初始化单链表
head.next -> nil
*/
func NewSingleLinkList() SingleLinkLister {
	return &SingleLinkList{}
}

// 获取链表长度
func (s *SingleLinkList) Len() int {
	length := 0
	node := s.next
	for node != nil {
		node = node.next
		length++
	}
	return length
}

//获取i元素的值
func (s *SingleLinkList) Get(i int) (interface{}, error) {
	// 查找i次
	node := s.next
	for node != nil && i > 0 {
		node = node.next
		i--
	}
	if node == nil {
		return nil, errors.New("i: index out of range")
	}

	return node.value, nil
}

/*
在i位置插入值value
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
    i					   0            1            2            3            4
假如要在i=2的位置插入一个节点:
head.next -> Node.next -> Node.next    Node.next -> Node.next -> Node.next -> nil
    i					   0            1  ↘️     ↗️  3(2')        4            5
														     Node.next
																	   2
1. 判断插入位置是否为第一个，如果是，则新的节点next指向头节点的next，头节点的next指向新的节点；
2. 如果不是，则
3. 找到i-1节点，将新加入的节点next指向i-1节点的next
4. 再将i-1节点的next指向新的节点
*/

func (s *SingleLinkList) Insert(i int, value interface{}) {

	// 新插入的节点
	node := &Node{
		value: value,
	}
	if i == 0 {
		node.next = s.next
		s.next = node
		return
	}

	// 移动到i-1节点
	preNode := s.next
	for preNode != nil && i > 1 {
		preNode = preNode.next
		i--
	}
	// 到达最后一个节点了，i超出了最大长度
	if preNode == nil {
		panic("i: index out of range")
	}
	node.next = preNode.next
	preNode.next = node
}

/*
删除i位置的值
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
    i					   0            1            2            3            4
假如要删除i=2的位置的节点:
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
    i					   0            1            3            4
														     Node.next
																	   2(X)
1. 判断删除位置是否为第一个，如果是，则头节点的next指向待删除的节点next；
2. 如果不是，则
3. 找到i-1节点，将i-1节点的next指向待删除的节点next；
*/
func (s *SingleLinkList) Delete(i int) {
	// 删除第一个节点
	if i == 0 {
		if s.next == nil {
			return
		}
		s.next = s.next.next
		return
	}
	node := s.next
	for node != nil && i > 1 {
		node = node.next // node是要删除节点的前一个节点
		i--
	}
	if node == nil || node.next == nil {
		panic("i: index out of range")
	}
	node.next = node.next.next
}

/*
链表反转
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
    i					   0            1            2            3            4
    nil  <-  Node.next <- Node.next <- Node.next <- Node.next <- Node.next <- head.next
    i  			     0            1            2            3            4
*/
func (s *SingleLinkList) Reverse() {
	var preNode *Node
	curNode := s.next

	for curNode != nil {
		nextNode := curNode.next
		curNode.next = preNode // 这行一定要先行
		if nextNode == nil { // 最后一个节点
			s.next = curNode
			return
		}
		preNode = curNode
		curNode = nextNode
	}
}

func main() {
	linkList := NewSingleLinkList()
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		linkList.Insert(i, v)
	}
	fmt.Println("len:", linkList.Len())
	for i := 0; i < linkList.Len(); i++ {
		v, _ := linkList.Get(i)
		fmt.Println(i, ":", v)
	}
	fmt.Println("-----删除i=3的数-----")
	linkList.Delete(3)
	fmt.Println("len:", linkList.Len())
	for i := 0; i < linkList.Len(); i++ {
		v, _ := linkList.Get(i)
		fmt.Println(i, ":", v)
	}

	fmt.Println("-----链表反转-----")
	linkList.Reverse()
	fmt.Println("len:", linkList.Len())
	for i := 0; i < linkList.Len(); i++ {
		v, _ := linkList.Get(i)
		fmt.Println(i, ":", v)
	}
}
