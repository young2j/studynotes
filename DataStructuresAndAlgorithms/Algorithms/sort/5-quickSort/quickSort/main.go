/*
 * File: main.go
 * Created Date: 2021-12-16 02:33:17
 * Author: ysj
 * Description:  快速排序
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

func netherlands(arr []int, L, R int) []int {
	M := L + ((R - L) >> 1)
	base := arr[M]
	less := L - 1
	greater := R + 1
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
	return []int{less, greater}
}

func process(arr []int, L, R int) {
	if L >= R {
		return
	}
	leftRight := netherlands(arr, L, R)
	// fmt.Println("leftRight:", leftRight)
	process(arr, L, leftRight[0])
	process(arr, leftRight[1], R)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	process(arr, 0, len(arr)-1)
	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	// arr := []int{4, 6, 8, 4, 1, 4, 2, 9}
	// arr := []int{2, 4, 4}
	// arr := []int{4,4,4,4}
	fmt.Println("origin:", arr)
	arr = quickSort(arr)
	fmt.Println("sorted:", arr)
}
