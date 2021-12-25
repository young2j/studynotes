/*
 * File: main.go
 * Created Date: 2021-12-11 09:01:06
 * Author: ysj
 * Description:  线性表-顺序表 read/update: O(1) insert/delete:O(N)
 */

package main

import "fmt"

// 定义元素类型-可以是任意类型
type EleType interface{}

// 定义顺序表应该实现的接口
type IList interface {
	Len() int    // 长度
	Empty() bool // 是否为空
	Full() bool  // 是否已满

	GetEle(i int) EleType                // 获取位序i元素
	InsertEle(i int, value EleType) bool // 插入位序i元素
	UpdateEle(i int, value EleType) bool // 更新位序i元素
	DeleteEle(i int) EleType             // 删除位序i的元素
	DeleteValue(value EleType) int       // 删除值为value的元素
	IndexValue(value EleType) int        // 返回值为value的位序
	Reverse()                            // 逆序
	Clear()                              // 清空表元素
}

// 定义顺序表结构
const MAXSIZE = 10

type List struct {
	Length int
	Data   [MAXSIZE]EleType
}

// 初始化
func NewSequentialList(elements ...EleType) *List {
	var data [MAXSIZE]EleType
	length := 0
	for i, v := range elements {
		if i < MAXSIZE {
			data[i] = v
			length++
		}
	}
	return &List{
		Length: length,
		Data:   data,
	}
}

// 获取长度
func (l *List) Len() int {
	return l.Length
}

// 判断是否为空
func (l *List) Empty() bool {
	return l.Length == 0
}

// 判断是否已满
func (l *List) Full() bool {
	return l.Len() == MAXSIZE
}

// 获取位序i的元素
func (l *List) GetEle(i int) EleType {
	if l.Empty() {
		panic("no element exsists")
	}
	if i < 0 || i >= l.Length {
		panic("Index out of bounds")
	}
	return l.Data[i]
}

// 插入位序i元素
func (l *List) InsertEle(i int, value EleType) bool {
	if l.Full() {
		panic("length has reach max size")
	}
	if i < 0 || i > l.Length {
		panic("Index out of bounds")
	}
	for j := l.Length - 1; j >= i; j-- {
		l.Data[j+1] = l.Data[j]
	}
	l.Data[i] = value
	l.Length++
	return true
}

// 更新位序i元素
func (l *List) UpdateEle(i int, value EleType) bool {
	if l.Empty() {
		panic("no element exsists")
	}
	if i < 0 || i >= l.Length {
		panic("Index out of bounds")
	}
	l.Data[i] = value
	return true
}

// 删除位序i的元素
func (l *List) DeleteEle(i int) EleType {
	if l.Empty() {
		panic("no element exsists")
	}
	if i < 0 || i >= l.Length {
		panic("Index out of bounds")
	}
	delelteEle := l.Data[i]
	for j := i; j < l.Length; j++ {
		l.Data[j] = l.Data[j+1]
	}
	l.Data[l.Length-1] = nil
	l.Length--
	return delelteEle
}

// 删除值为value的元素
func (l *List) DeleteValue(value EleType) int {
	deleteCount := 0
	for i := 0; i < l.Length; i++ {
		if l.Data[i] == value {
			if ele := l.DeleteEle(i); ele != nil {
				deleteCount++
			}
		}
	}
	return deleteCount
}

// 返回值为value的位序
func (l *List) IndexValue(value EleType) int {
	i := -1
	for i := 0; i < l.Length; i++ {
		if l.Data[i] == value {
			return i
		}
	}
	return i
}

// 逆序
func (l *List) Reverse() {
	if l.Empty() {
		return
	}
	for i := 0; i < l.Length/2; i++ {
		temp := l.Data[i]
		l.Data[i] = l.Data[l.Length-1-i]
		l.Data[l.Length-1-i] = temp
	}
}

// 清空表元素
func (l *List) Clear() {
	if l.Empty() {
		return
	}
	// length := l.Length
	// for i := 0; i < length; i++ {
	// 	l.DeleteEle(0)
	// }
	l.Data = [MAXSIZE]EleType{}
	l.Length = 0
}

func main() {
	fmt.Println("初始化顺序表:")
	l := NewSequentialList(0, 1, 5, "5", 5)
	fmt.Println("字符串5的位序为:", l.IndexValue("5"))
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)
	l.UpdateEle(0, 1)
	fmt.Println("把第一个数更新为1:")
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)

	l.InsertEle(1, 2)
	fmt.Println("在第二个位置插入数字2:")
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)

	l.DeleteEle(0)
	fmt.Println("删除第一个数字:")
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)

	l.DeleteValue(5)
	fmt.Println("删除值为5的数字:")
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)

	fmt.Println("遍历:")
	for i := 0; i < l.Length; i++ {
		fmt.Println(l.GetEle(i))
	}
	fmt.Println("\n反转:")
	l.Reverse()
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)

	fmt.Println("清空:")
	l.Clear()
	fmt.Printf("data:%v length:%d \n\n", l.Data, l.Length)
}
