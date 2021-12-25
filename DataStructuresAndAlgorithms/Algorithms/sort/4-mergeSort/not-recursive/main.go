/*
 * File: main.go
 * Created Date: 2021-12-14 03:04:33
 * Author: ysj
 * Description:  归并排序-非递归实现
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
	for L0 <= M {
		help[i] = arr[L0]
		L0++
		i++
	}
	for R0 <= R {
		help[i] = arr[R0]
		R0++
		i++
	}
	for i, v := range help {
		arr[L+i] = v
	}
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mergeSize := 2
	N := len(arr)
	// 这里<2倍长度
	for mergeSize < 2*N {
		L := 0
		for L < N {
			// 先计算M
			M := L + (mergeSize >> 1) - 1
			// 边界条件
			if M > N-1 {
				break
			}
			R := utils.Min(M+(mergeSize>>1), N-1)
			merge(arr, L, M, R)
			L = R + 1
		}
		mergeSize <<= 1
	}
	return arr
}

func mergeSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mergeSize := 1
	N := len(arr)
	for mergeSize < N {
		L := 0
		for L < N {
			M := L + mergeSize - 1
			if M > N-1 {
				break
			}
			R := utils.Min(L+(mergeSize<<1)-1, N-1)
			merge(arr, L, M, R)
			L = R + 1
		}
		if mergeSize > (N >> 1) {
			break
		}
		mergeSize <<= 1
	}

	return arr
}

// 先计算右值R-要处理最后一次归并
func mergeSort3(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mergeSize := 2
	N := len(arr)
	// 这里<1倍长度
	for mergeSize < N {
		L := 0
		for L < N {
			// 先计算右值
			R := utils.Min(L+mergeSize-1, N-1)
			M := L + ((R - L) >> 1)
			merge(arr, L, M, R)
			L = R + 1
		}
		mergeSize <<= 1
	}
	// mergeSize>=N时，最后一次归并
	merge(arr, 0, (mergeSize>>1)-1, N-1)
	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = mergeSort(arr)
	fmt.Println("sorted:", arr)
	arr = mergeSort2(arr)
	fmt.Println("sorted:", arr)
	arr = mergeSort3(arr)
	fmt.Println("sorted:", arr)
}
