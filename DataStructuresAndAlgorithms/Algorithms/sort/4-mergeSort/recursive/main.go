/*
 * File: main.go
 * Created Date: 2021-12-13 05:09:44
 * Author: ysj
 * Description:  归并排序-递归实现
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

func merge(arr []int, L, M, R int) {
	L0 := L
	R0 := M + 1
	i := 0
	help := make([]int, R-L+1)
	//
	for L0 <= M && R0 <= R {
		if arr[L0] < arr[R0] {
			help[i] = arr[L0]
			L0++
		} else {
			help[i] = arr[R0]
			R0++
		}
		i++
	}
	//
	for L0 <= M {
		help[i] = arr[L0]
		L0++
		i++
	}
	//
	for R0 <= R {
		help[i] = arr[R0]
		R0++
		i++
	}

	for i, v := range help {
		arr[L+i] = v
	}
}

func processHalf(arr []int, L, R int) []int {
	// 终止条件
	if L == R {
		return arr
	}

	M := L + ((R - L) >> 1)
	processHalf(arr, L, M)
	processHalf(arr, M+1, R)
	merge(arr, L, M, R)

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	processHalf(arr, 0, len(arr)-1)
	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = mergeSort(arr)
	fmt.Println("sorted:", arr)
}
