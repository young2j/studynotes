/*
 * File: main.go
 * Created Date: 2022-01-17 02:59:34
 * Author: ysj
 * Description:  链式队列
 */

package main

import "fmt"

// 队列接口
type Queuer interface {
	Clear()                  // 清空队列
	IsEmpty() bool           // 队列是否为空
	Len() int                // 链表元素个数
	Head() interface{}       // 获取队头元素
	Rpush(value interface{}) // 入队
	Lpop() interface{}       // 出队
}

// 定义队列节点
type Node struct {
	value interface{}
	next  *Node
}

// 定义队列
type LinkQueue struct {
	front *Node
	rear  *Node
	count int
}

// 初始化
func NewLinkQueue() Queuer {
	return &LinkQueue{}
}

// 清空队列
func (q *LinkQueue) Clear() {
	q.front = nil
	q.rear = nil
	q.count = 0
}

// 队列是否为空
func (q *LinkQueue) IsEmpty() bool {
	return q.front == nil && q.rear == nil
}

// 队列长度
func (q *LinkQueue) Len() int {
	return q.count
}

// 队头元素
func (q *LinkQueue) Head() interface{} {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.front.value
}

// 入队
func (q *LinkQueue) Rpush(value interface{}) {
	node := &Node{
		value: value,
		next:  nil,
	}
	if q.IsEmpty() {
		q.front = node
		q.rear = node
		q.count++
		return
	}
	q.rear.next = node
	q.rear = node
	q.count++
}

// 出队
func (q *LinkQueue) Lpop() interface{} {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	front := q.front
	q.front = q.front.next
	q.count--

	return front.value
}

func main() {
	q := NewLinkQueue()
	// q.Lpop() // panic: queue is empty
	values := []int{3, 5, 6, 7, 2}
	for _, v := range values {
		fmt.Println("queue push", v)
		q.Rpush(v)
	}
	fmt.Println("Len(q):", q.Len())
	fmt.Println("q head value is:", q.Head())
	fmt.Println("q pop:", q.Lpop())
	fmt.Println("q pop:", q.Lpop())
	fmt.Println("Len(q):", q.Len())
	q.Clear()
	fmt.Println("q clear...Len(q):", q.Len())
}
