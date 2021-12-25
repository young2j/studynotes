/*
 * File: xor.go
 * Created Date: 2021-12-10 04:59:10
 * Author: ysj
 * Description:  异或运算
 */

package main

import "fmt"

func main() {
	// 异或运算： 两个数相同为0，相异为1，且与顺序无关
	//1. 所有数字都出现了偶数次，仅有一个数字出现奇数次，找到并打印这个数字
	arr := []int{4, 5, 4, 6, 7, 8, 5, 8, 6}
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	fmt.Println("eor:", eor)

	//2. 所有数字都出现了偶数次，有两个数字出现奇数次，找到并打印这两个数字
	arr = []int{4, 5, 4, 6, 7, 8, 5, 8}
	eor = 0
	for _, v := range arr {
		eor ^= v
	}
	rightOne := eor & (^eor + 1)
	eor1 := 0
	eor2 := 0
	for _, v := range arr {
		if v&rightOne != 0 { // 将两个奇数分开
			eor1 ^= v
		} else {
			eor2 ^= v
		}
	}
	fmt.Println("eor1:", eor1)
	fmt.Println("eor2:", eor2)

	// 3. 十六进制状态管理
}
