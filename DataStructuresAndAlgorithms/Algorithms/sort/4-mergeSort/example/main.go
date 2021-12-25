/*
 * File: main.go
 * Created Date: 2021-12-15 03:22:18
 * Author: ysj
 * Description:  归并排序的扩展应用
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

// ========================= 小数和 =====================
func pick(arr []int, L, M, R int) (nums []int) {
	L0 := L
	R0 := M + 1
	i := 0
	help := make([]int, R-L+1)
	// --
	for L0 <= M && R0 <= R {
		if arr[L0] < arr[R0] {
			help[i] = arr[L0]
			// --
			for n := 0; n < R-R0+1; n++ {
				nums = append(nums, arr[L0])
			}
			//
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
		i++
		R0++
	}
	// --
	for i, v := range help {
		arr[L+i] = v
	}

	return nums
}

func pickSmallNumber(arr []int, L, R int) (nums []int) {
	// 终止条件
	if L == R {
		return nums
	}
	M := L + ((R - L) >> 1)
	letfNums := pickSmallNumber(arr, L, M)
	rightNums := pickSmallNumber(arr, M+1, R)
	currentNums := pick(arr, L, M, R)
	nums = append(nums, letfNums...)
	nums = append(nums, rightNums...)
	nums = append(nums, currentNums...)
	return nums
}

func smallNumberSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	nums := pickSmallNumber(arr, 0, len(arr)-1)
	sum := utils.Sum(nums)
	return sum
}

// ========================= 降序对 =====================
func pickPair(arr []int, L, M, R int) (pair [][]int) {
	L0 := L
	R0 := M + 1
	i := 0
	help := make([]int, R-L+1)

	for L0 <= M && R0 <= R {
		if arr[L0] > arr[R0] {
			help[i] = arr[L0]
			pair = append(pair, []int{arr[L0], arr[R0]})
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
	return pair
}

func pickHalfPair(arr []int, L, R int) (pair [][]int) {
	if L == R {
		return pair
	}
	M := L + ((R - L) >> 1)
	leftPair := pickHalfPair(arr, L, M)
	rightPair := pickHalfPair(arr, M+1, R)
	currentPair := pickPair(arr, L, M, R)
	pair = append(pair, leftPair...)
	pair = append(pair, rightPair...)
	pair = append(pair, currentPair...)
	return pair
}

func findDescPairs(arr []int) (pair [][]int) {
	if len(arr) < 2 {
		return pair
	}
	pair = pickHalfPair(arr, 0, len(arr)-1)
	return pair
}

func main() {
	// 1.求小数和
	arr := []int{3, 4, 1, 8, 6, 4, 2, 9}
	// 3 -> 0
	// 4 -> 3
	// 1 -> 0
	// 8 -> 3,4,1
	// 6 -> 3,4,1
	// 4 -> 3,1
	// 2 -> 1
	// 9 -> 3,4,1,8,6,4,2
	// sum = 0+3+0+3+4+1+3+4+1+3+1+1+3+4+1+8+6+4+2 = 52
	sum := smallNumberSum(arr)
	fmt.Println("sum:", sum)

	// 2.求降序对
	arr_ := []int{4, 3, 5, 1, 7, 8, 2}
	// 4 -> (4,3),(4,1),(4,2)
	// 3 -> (3,1),(3,2)
	// 5 -> (5,1),(5,2)
	// 1 -> ()
	// 7 -> (7,2)
	// 8 -> (8,2)
	// 2 -> ()
	pairs := findDescPairs(arr_)
	fmt.Println("desc pairs:", len(pairs), pairs)
}
