/*
 * File: main.go
 * Created Date: 2022-01-11 02:12:09
 * Author: ysj
 * Description:  静态链表-用数组模拟链表操作进行增删改查
 */

package main

import (
	"fmt"
)

/*
1. 静态链表使用数组来模拟
2. 头尾位置通常不存储值，头元素位置可以存储元素个数
3. 未存储有数据值的区域称为空闲区
4. 使用cur变量(cursor)来模拟指针移动；
   每个数据节点node.cur=存储下一个数据节点的下标；
	 如果下一个数据节点为空闲区，其cur指向head位置，即node.cur=0，代表下一位置无数据；
	 头元素位置head.cur=存储空闲区的第一个下标(第一个可插入数据的位置)；
	 尾元素位置tail.cur=存储数据区的第一个下标(第一个数据值的下标),初始为0；

假如将容量为12的数组arr=[12]int表示为静态链表：
      head                     |---------空闲区--------|  tail
index   0   1   2   3   4   5   6    7    8    9    10   11
arr		[nil, 3,  5,  6,  8,  9,  - ,  - ,  - ,  - ,  - ,  nil ]
cur     6   2   3   4   5   0   7    8    9    10   11   1
*/

// 定义静态链表节点
type Node struct {
	value interface{} // 存储任意值
	cur   int         // 游标-模拟指针
}

// 定义静态链表--容量最大12
const MAX = 12

type StaticList [MAX]*Node

// 声明静态链表接口
type StaticLister interface {
	Insert(i int, value interface{}) // 插入值
	Delete(i int)                    // 删除i位置的值
	Len()                            // 长度，元素个数
}

/*
初始化一个静态链表
      head |--------------------空闲区------------------| tail
index   0   1   2   3   4   5   6    7    8    9    10   11
arr		[nil, -,  -,  -,  -,  -,  - ,  - ,  - ,  - ,  - ,  nil ]
cur     1   2   3   4   5   6   7    8    9    10   11   0
*/
func NewStaticList() StaticList {
	staticList := StaticList{}
	for i := 0; i < MAX-1; i++ {
		node := &Node{
			cur: i + 1,
		}
		staticList[i] = node
	}
	staticList[MAX-1] = &Node{cur: 0}
	return staticList
}

/*
假如要在3,5,6,8,9中第三个位置插入一个值2，即Insert(3,666)后: 3,5,666,6,8,9
需要先在空闲位插入值，然后调整cur的值：
1. 首先查找是否具有空位，有则6号位先插入值666，head.cur记录更新为下一个空位7；
2. 看尾节点，尾节点记录的是数据域第一个位置，循环(i-1)次，拿到第i个数的cur
   (通过多次增删，可能变得不会如示例一样看起来那么有序，需要通过尾节点的cur往下遍历)；
	 循环3-1=2次，1.cur=2.cur=3, 则更新6号位的cur为3，表示新插入值的下一个值为原3号位；
3. 原3号位的上一位2号位的cur=3更新为6，表示下一个值为新插入的6号位；


      head                     |---------空闲区--------|  tail
index   0   1   2   3   4   5   6    7    8    9    10   11
arr		[nil, 3,  5,  6,  8,  9,  - ,  - ,  - ,  - ,  - ,  nil ]
cur     6   2   3   4   5   0   7    8    9    10   11   1

      head                     |---------空闲区--------|  tail
index   0   1   2   3   4   5   6    7    8    9    10   11
arr		[nil, 3,  5,  6,  8,  9, 666,  - ,  - ,  - ,  - ,  nil ]
cur     7   2   6   4   5   0   3    8    9    10   11   1
*/
// 获取空闲区cur, 并更新head.cur
func (s *StaticList) headCur() int {
	cur := s[0].cur
	if cur != 0 {
		s[0].cur = s[cur].cur
	}
	return cur
}

// 插入值
func (s *StaticList) Insert(i int, value interface{}) {
	// 边界条件, 掐头去尾
	if i < 1 || i >= MAX-1 {
		panic("i: index range error, should be [1,98]")
	}
	cur := s.headCur() // 6, head.cur=7
	if cur != 0 {
		s[cur].value = value

		k := MAX - 1 // 99
		for ii := 1; ii <= i-1; ii++ {
			k = s[k].cur // 1->2
		}
		s[cur].cur = s[k].cur //s[6].cur=s[2].cur=3
		s[k].cur = cur        // s[2].cur=6
	}
}

// 获取链表数据个数
func (s *StaticList) Len() int {
	k := MAX - 1
	length := 0
	for s[k].cur != 0 {
		k = s[k].cur
		length++
	}
	return length
}

/*
删除链表元素
1. 删除时，下标范围不能超过元素个数；
2. 删除时，删除位置要回收到空闲区，即作为空闲区第一个位置，更新到head.cur
3. 删除位置的cur值更新到前一个节点的cur值
假如要删除第4个数(注意不是8，而是3号位置的数6):
1. 判断4不越界；
2. 从尾节点开始遍历4-1=3次，cur=1->2->6;
3. 6号位的cur=3，即要删除的数的位置；将6号位置的cur值更新为要删除位置的cur=4；
4. 将要删除位置的cur更新为下一个空闲区，即head.cur
5. 更新head.cur的值为要删除的位号3，表示空闲区第一个位置为3；

      head                     |---------空闲区--------|  tail
index   0   1   2   3   4   5   6    7    8    9    10   11
arr		[nil, 3,  5,  6,  8,  9, 666,  - ,  - ,  - ,  - ,  nil ]
cur     7   2   6   4   5   0   3    8    9    10   11   1

      head                     |---------空闲区--------|  tail
index   0   1   2   3   4   5   6    7    8    9    10   11
arr		[nil, 3,  5,  -,  8,  9, 666,  - ,  - ,  - ,  - ,  nil ]
cur     3   2   6   7   5   0   4    8    9    10   11   1
*/
func (s *StaticList) Delete(i int) {
	if i < 1 || i > s.Len() {
		panic("i: index range error")
	}
	k := MAX - 1
	for ii := 1; ii <= i-1; ii++ {
		k = s[k].cur // 1->2->6
	}
	j := s[k].cur       // 待删除的位号 3
	s[k].cur = s[j].cur // s[6].cur = s[3].cur = 4
	s[j].cur = s[0].cur // s[3].cur = s[0].cur = 7
	s[j].value = nil    // 值置为nil
	s[0].cur = k        // 3
}

func main() {
	staticList := NewStaticList()
	arr := []int{3, 5, 6, 8, 9}
	// 顺序插入5个数
	for i, v := range arr {
		staticList.Insert(i+1, v)
	}
	fmt.Println("len(staticList):", staticList.Len())
	for n, node := range staticList {
		fmt.Printf("index:%d value:%v cur:%d \n", n, node.value, node.cur)
	}

	// 往4位置插入666
	fmt.Println(".....往第4个位置插入一个数666.....")
	staticList.Insert(4, 666)
	fmt.Println("len(staticList):", staticList.Len())

	for n, node := range staticList {
		fmt.Printf("index:%d value:%v cur:%d \n", n, node.value, node.cur)
	}

	// 删除第4个数
	fmt.Println("....删除第4个数666后....")
	staticList.Delete(4)
	fmt.Println("len(staticList):", staticList.Len())
	for n, node := range staticList {
		fmt.Printf("index:%d value:%v cur:%d \n", n, node.value, node.cur)
	}
}
