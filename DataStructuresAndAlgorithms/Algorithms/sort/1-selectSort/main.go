/*
 * File: main.go
 * Created Date: 2021-12-10 05:35:59
 * Author: ysj
 * Description:  选择排序 O(N^2)
 */

package main

import (
	"algorithms/utils"
	"fmt"
)


/*
假设有n个数：
1. 遍历0->n-1位置的数，只要存在某个数比第0位置的数小，则交换位置；
2. 遍历1->n-1位置的数，只要存在某个数比第1位置的数小，则交换位置；
3. 遍历2->n-1位置的数，只要存在某个数比第2位置的数小，则交换位置；
4. ...
n. 遍历n-2->n-1位置的数，只要存在某个数比第n-2位置的数小，则交换位置；
*/
func selectSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				utils.Swap(arr, i, j)
			}
		}
	}
	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = selectSort(arr)
	fmt.Println("sorted:", arr)
}
