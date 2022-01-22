/*
 * File: main.go
 * Created Date: 2022-01-17 02:36:35
 * Author: ysj
 * Description:  循环队列-数组实现
 */

package main

import "fmt"

// 队列接口
type Queuer interface {
	Clear()                 // 清空队列
	IsEmpty() bool          // 队列是否为空
	IsFull() bool           // 队列是否已满
	Len() int               // 链表元素个数
	Head() interface{}      // 获取队头元素
	Push(value interface{}) // 入队
	Pop() interface{}       // 出队
}

const MAXSIZE = 20

// 定义循环队列
type LoopQueue struct {
	front int                  // 队头
	rear  int                  // 队尾
	data  [MAXSIZE]interface{} // 队列元素
}

/*
1. 队列为空：front==rear
2. 队列已满：(rear+1)%queueSize == front
3. 队列长度: (rear-front+queueSize)%queueSize
*/

/*
初始化循环队列
front=rear=0
q [ - , - , - , - , - , ]
i   0   1   2   3   4
*/
func NewLoopQueue() Queuer {
	return &LoopQueue{}
}

// 清空队列
func (q *LoopQueue) Clear() {
	q.front = 0
	q.rear = 0
	for i := 0; i < MAXSIZE; i++ {
		q.data[i] = nil
	}
}

// 队列是否为空
func (q *LoopQueue) IsEmpty() bool {
	return q.front == q.rear
}

// 队列是否已满
func (q *LoopQueue) IsFull() bool {
	return (q.rear+1)%MAXSIZE == q.front
}

// 队列长度
func (q *LoopQueue) Len() int {
	return (q.rear - q.front + MAXSIZE) % MAXSIZE
}

// 获取对头元素
func (q *LoopQueue) Head() interface{} {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.data[q.front]
}

/*
入队 rear = (rear+1)%MAXSIZE

front = 0
rear = 2
q [ 3 , 5 , - , - , - , ]
i   0   1   2   3   4

push 6
rear = (2+1)%20 = 3

front = 0
rear = 3
q [ 3 , 5 , 6 , - , - , ]
i   0   1   2   3   4
*/
func (q *LoopQueue) Push(value interface{}) {
	if q.IsFull() {
		panic("queue is full")
	}

	q.data[q.rear] = value
	q.rear = (q.rear + 1) % MAXSIZE
}

//出队 front 往后移动一位
func (q *LoopQueue) Pop() interface{} {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % MAXSIZE
	return value
}

func main() {
	q := NewLoopQueue()
	// q.Pop() // panic: queue is empty
	values := []int{3, 5, 6, 7, 2}
	for _, v := range values {
		fmt.Println("queue push", v)
		q.Push(v)
	}
	fmt.Println("Len(q):", q.Len())
	fmt.Println("q head value is:", q.Head())
	fmt.Println("q pop:", q.Pop())
	fmt.Println("q pop:", q.Pop())
	fmt.Println("Len(q):", q.Len())
	q.Clear()
	fmt.Println("q clear...Len(q):", q.Len())
}
