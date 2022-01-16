/*
 * File: main.go
 * Created Date: 2022-01-17 02:36:35
 * Author: ysj
 * Description:  循环队列-数组实现
 */

package main

// 队列接口
type Queuer interface {
	Clear()                  // 清空队列
	IsEmpty() bool           // 队列是否为空
	Head() interface{}       // 获取队头元素
	Lpop() interface{}       // 出队
	Rpush(value interface{}) // 入队
	Len() int                // 链表元素个数
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

func main() {

}
