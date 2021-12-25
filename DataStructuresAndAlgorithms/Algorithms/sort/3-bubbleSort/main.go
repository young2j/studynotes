/*
 * File: main.go
 * Created Date: 2021-12-11 03:50:23
 * Author: ysj
 * Description: 冒泡排序 O(N^2)
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

/*
假设有n个数：
1. 0->n-1: 若0位置的数大于1位置的数，则互换两数；
				   继续比较，若1位置的数大于2位置的数，则互换两数;
				   ...
	 			   若n-1位置的数大于n位置的数，则互换两数。
				   此时，n位置的数为最大值；

2. 0->n-2: 若0位置的数大于1位置的数，则互换两数；
				   继续比较，若1位置的数大于2位置的数，则互换两数;
				   ...
	 			   若n-2位置的数大于n-1位置的数，则互换两数。
				   此时，n-1位置的数为第二大的数(或等于n位置的数)；
3.	 ...

n-1. 0->1: 若0位置的数大于1位置的数，则互换两数；
				   此时，0和1位置的数为最小的两个数；
*/
func bubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				utils.Swap(arr, j, j+1)
			}
		}
	}
	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = bubbleSort(arr)
	fmt.Println("sorted:", arr)
}
