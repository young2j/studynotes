/*
 * File: main.go
 * Created Date: 2022-01-16 04:49:56
 * Author: ysj
 * Description:  双向链表
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
	pre   *Node
}

// 定义双向链表(头节点)
type DoubleLinkList Node

//双向链表接口
type DoubleLinkLister interface {
	Len() int
	Get(i int) (interface{}, error)
	Insert(i int, value interface{})
	Delete(i int)
	Reverse()
}

/*
初始化双向链表
head.next -> nil
head.pre -> nil
*/
func NewDoubleLinkList() DoubleLinkLister {
	return &DoubleLinkList{}
}

// 获取链表长度
func (d *DoubleLinkList) Len() int {
	length := 0
	node := d.next
	for node != nil {
		node = node.next
		length++
	}
	return length
}

//获取i元素的值
func (d *DoubleLinkList) Get(i int) (interface{}, error) {
	// 查找i次
	node := d.next
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
 nil      <- Node.pre  <- Node.pre  <- Node.pre  <- Node.pre  <- Node.pre
  i  					   0            1            2            3            4

假如要在i=2的位置插入一个节点:
head.next -> Node.next -> Node.next    Node.next -> Node.next -> Node.next -> nil
    i					   0            1  ↘️     ↗️  3(2')        4            5
														     Node.next
																	   2
 nil      <- Node.pre  <- Node.pre     Node.pre  <- Node.pre  <- Node.pre
  i 					   0            1  ↖     ↙  3(2')        4            5
														     Node.pre
																	   2
1. 判断插入位置是否为第一个，如果是，则新的节点next指向头节点的next，头节点的next指向新的节点；新节点的pre指向头节点；
2. 如果不是，则
3. 找到i-1节点，将新加入的节点next指向i-1节点的next；将i-1节点的next.pre指向新的节点；
4. 再将i-1节点的next指向新的节点；新节点的pre指向i-1节点。
*/

func (d *DoubleLinkList) Insert(i int, value interface{}) {

	// 新插入的节点
	node := &Node{
		value: value,
	}
	if i == 0 {
		node.next = d.next
		if d.next != nil {
			d.next.pre = node
		}
		d.next = node
		node.pre = nil
		return
	}

	// 移动到i-1节点
	preNode := d.next
	for preNode != nil && i > 1 {
		preNode = preNode.next
		i--
	}
	// 到达最后一个节点了，i超出了最大长度
	if preNode == nil {
		panic("i: index out of range")
	}
	node.next = preNode.next
	if preNode.next != nil {
		preNode.next.pre = node
	}
	preNode.next = node
	node.pre = preNode
}

/*
删除i位置的值
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
 nil      <- Node.pre  <- Node.pre  <- Node.pre  <- Node.pre  <- Node.pre
  i  					   0            1            2            3            4

假如要删除i=2的位置的节点:
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
    i					   0            1            3            4
														     Node.next
																	   2(X)
   nil    <- Node.pre <- Node.pre <- Node.pre <- Node.pre
    i					   0           1           3           4
														     Node.pre
																	   2(X)
1. 判断删除位置是否为第一个，如果是，则头节点的next指向待删除的节点next；删除节点的next.pre指向nil；
2. 如果不是，则
3. 找到i-1节点，将i-1节点的next指向待删除的节点next；待删除节点的next.pre指向i-1节点;
*/
func (d *DoubleLinkList) Delete(i int) {
	// 删除第一个节点
	if i == 0 {
		if d.next == nil {
			return
		}
		if d.next.next != nil {
			d.next.next.pre = nil
		}
		d.next = d.next.next
		return
	}
	node := d.next
	for node != nil && i > 1 {
		node = node.next // node是要删除节点的前一个节点
		i--
	}
	if node == nil || node.next == nil {
		panic("i: index out of range")
	}
	if node.next.next != nil {
		node.next.next.pre = node
	}
	node.next = node.next.next
}

/*
链表反转
head.next -> Node.next -> Node.next -> Node.next -> Node.next -> Node.next -> nil
    i					   0            1            2            3            4
    nil   <- Node.next <- Node.next <- Node.next <- Node.next <- Node.next <- head.next [反转next]
    i  			     0            1            2            3            4

 		nil   <- Node.pre  <- Node.pre  <- Node.pre  <- Node.pre  <- Node.pre
  	i  				   0            1            2            3            4
 head.pre -> Node.pre  -> Node.pre  -> Node.pre  -> Node.pre  -> Node.pre -> nil  [反转pre]
  	i  				   0            1            2            3            4

将每个节点的next和pre交换
head.next要指向最后一个节点
*/
func (d *DoubleLinkList) Reverse() {
	curNode := d.next
	d.next = d.pre
	d.pre = curNode
	for curNode != nil {
		nextNode := curNode.next
		// 先交换
		curNode.next = curNode.pre
		curNode.pre = nextNode

		if nextNode == nil { // 最后一个节点
			d.next = curNode
		}
		curNode = nextNode
	}
}

func main() {
	linkList := NewDoubleLinkList()
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
