/*
 * File: main.go
 * Created Date: 2021-12-16 02:51:53
 * Author: ysj
 * Description:  荷兰国旗问题
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

// 以某个值为基准(如中点值)，将数组的值分开为两部分
func NetherlandsFlag(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	L := 0
	R := len(arr) - 1
	M := L + ((R - L) >> 1)
	base := arr[M]
	fmt.Println("base:", base)
	less := -1
	greater := len(arr)
	for L < greater {
		if arr[L] < base {
			utils.Swap(arr, L, less+1)
			less++
			L++
		} else if arr[L] > base {
			utils.Swap(arr, L, greater-1)
			greater--
		} else {
			L++
		}
	}
	return arr
}

func main() {
	// arr := []int{4, 6, 8, 4, 1, 4, 2, 9}
	arr := utils.RandomIntArr(10, 100)
	fmt.Println(arr)

	arr = NetherlandsFlag(arr)
	fmt.Println(arr)
}
