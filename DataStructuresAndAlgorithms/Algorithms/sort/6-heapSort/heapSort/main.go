/*
 * File: main.go
 * Created Date: 2021-12-30 12:59:48
 * Author: ysj
 * Description:  堆排序
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

// 使用自顶向下的方式构建，自顶向下的方式最简单
// 大顶堆
func buildMaxHeapFromTop2Bottom(arr []int) {
	for i := 0; i < len(arr); i++ {
		for cur := i; arr[cur] > arr[(cur-1)/2]; cur = (cur - 1) / 2 {
			utils.Swap(arr, cur, (cur-1)/2)
		}
	}
}

// 小顶堆
func buildMinHeapFromTop2Bottom(arr []int) {
	for i := 0; i < len(arr); i++ {
		for cur := i; arr[cur] < arr[(cur-1)/2]; cur = (cur - 1) / 2 {
			utils.Swap(arr, cur, (cur-1)/2)
		}
	}
}

// adjust heap
func heapSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// use max-heap
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	buildMaxHeapFromTop2Bottom(arr[0 : i+1])
	// 	utils.Swap(arr, 0, i)
	// }

	// use min-heap
	for i := 0; i < len(arr); i++ {
		buildMinHeapFromTop2Bottom(arr[i:])
	}

	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = heapSort(arr)
	fmt.Println("sorted:", arr)
}
