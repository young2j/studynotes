/*
 * File: main.go
 * Created Date: 2021-12-19 01:31:08
 * Author: ysj
 * Description:  水王数(出现次数超过一半的数)
 */

package main

import "fmt"

func pickIrrigationNumber(arr []int) int {
	num := 0 // 候选
	hp := 0  // 血量
	for i := 0; i < len(arr); i++ {
		if hp == 0 { // 当前无候选，当前数为候选数，给1点hp
			num = arr[i]
			hp = 1
		} else if num == arr[i] { // 当前有候选，候选等于当前数，加1点hp
			hp++
		} else { //当前有候选，候选不等于当前数，减1点hp
			hp--
		}
	}
	if hp == 0 { // 最终，如果hp等于0，代表无候选
		return -1
	}
	c := 0
	for _, v := range arr { // 最终，如果hp大于0，统计候选数的出现次数
		if num == v {
			c++
		}
	}
	if c <= (len(arr) >> 1) { // 如果出现次数没有超过一半，则无水王数
		return -1
	}
	return num
}

func main() {
	arr := []int{2, 3, 2, 3, 2, 3, 3}
	num := pickIrrigationNumber(arr)
	fmt.Println("num:", num)
}
