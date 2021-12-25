/*
 * File: main.go
 * Created Date: 2021-12-13 04:20:45
 * Author: ysj
 * Description:  单链表
 */

package main

import "fmt"

// 节点值类型
type Value interface{}

// 链表节点
type Node struct {
	Value Value // 节点值
	Next  *Node // 下一个节点
}

// 链表结构
type SingleLinkedList struct {
	Length int     // 长度=元素个数
	Data   []*Node // 链表数据
}

// 链表接口
type ISingleLinkedList interface {
	Len() int    // 元素个数
	Empty() bool // 链表是否为空

	GetValue(i int) Value       // 获取第i个值
	UpdateValue(i int, v Value) // 更新第i个值
	DeleteValue(i int)          // 删除第i个值
	IndexValue(v Value) int     // 返回第一个值为Value的位序
	InsertValue(i int, v Value) // 在i位置插入值v

	Reverse() // 逆序
}

// 链表节点数
func (s *SingleLinkedList) Len() int {
	return s.Length
}

// 链表是否为空
func (s *SingleLinkedList) Empty() bool {
	return s.Len() == 0
}

//

func main() {
	var v []int
	fmt.Println(v)
}
