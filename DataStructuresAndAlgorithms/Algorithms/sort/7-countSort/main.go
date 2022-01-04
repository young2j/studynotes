/*
 * File: main.go
 * Created Date: 2022-01-04 02:48:48
 * Author: ysj
 * Description:  计数排序 O(N)
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

/*
计数排序要求输入的元素必须是有限个非负整数
1. 准备一个最大元素长度的辅助空间；
2. 遍历每个数组元素，在辅助空间中对应元素大小下标处计数值+1
3. 遍历辅助空间，按顺序输出非0值个数的下标，即为排序数组
*/

func countSort(arr []int) []int {
	// 边界
	if len(arr) < 2 {
		return arr
	}
	// 最大值
	maxValue := 0
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	// 一般情况
	help := make([]int, maxValue+1)
	for _, v := range arr {
		help[v]++
	}

	c := 0
	for i, v := range help {
		if v == 0 {
			continue
		}
		for j := 0; j < v; j++ {
			arr[c] = i
			c++
		}
	}
	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = countSort(arr)
	fmt.Println("sorted:", arr)
}
