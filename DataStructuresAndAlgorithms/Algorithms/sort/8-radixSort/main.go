/*
 * File: main.go
 * Created Date: 2022-01-04 05:11:06
 * Author: ysj
 * Description:  基数排序 O(N)
 */
package main

import (
	"algorithms/utils"
	"fmt"
)

func getDigit(num int, bit int) int {
	// bit = 0,1,2,3,...
	for i := 0; i < bit; i++ {
		num = num / 10
	}
	return num % 10
}

/*
仅适用于非负整数
1. 找出最大元素值，计算其位数
2. 遍历位数，按位上的数字进行排序，准备一个0-9的二维数组(桶)：
	 a. 从个位开始，将个位上的数字放入对应的桶中，0放入0号桶，1放入1号桶，2放入2号桶...然后将每个桶中的数字顺序拼接；
	 b. 在a的结果基础上，继续排十位上的数字，将十位上的数字放入对应桶中，0放入0号桶，1放入1号桶，2放入2号桶...然后将每个桶中的数字顺序拼接；
	 c. 在b的结果基础上，继续排百分位上的数字...
*/
func radixSort(arr []int) []int {
	// 边界
	if len(arr) < 2 {
		return arr
	}

	// 获取最大位数
	maxValue := 0
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	bits := 1 // 至少一位
	for (maxValue / 10) != 0 {
		bits++
		maxValue = maxValue / 10
	}

	// 从个位开始对每一位数字进行排序
	for i := 0; i < bits; i++ {
		buckets := make([][]int, 10)
		for _, v := range arr {
			digit := getDigit(v, i)
			buckets[digit] = append(buckets[digit], v)
		}
		arr = []int{}
		for _, bucket := range buckets {
			arr = append(arr, bucket...)
		}
	}

	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 1000)
	fmt.Println("origin:", arr)
	arr = radixSort(arr)
	fmt.Println("sorted:", arr)
}
