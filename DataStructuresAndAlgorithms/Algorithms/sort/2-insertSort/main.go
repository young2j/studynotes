/*
 * File: main.go
 * Created Date: 2021-12-10 06:37:59
 * Author: ysj
 * Description:  插入排序 O(N^2)
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

/*
假设有n个数， 遍历0->n-1位置的数：
1.   只要1位置的数比0位置的数小，则交换1与0位置的数；
2.   只要2位置的数比1位置的数小，则交换2与1位置的数;
	   然后继续向前比较，只要1位置小于0位置的数，则交换1与0位置的数；
3.   ...
n-1. 只要n-1位置的数比前一个位置的数小，则交换n-1位置与前一个位置的数；
     然后继续向前比较，直到前一个位置的数比当前数小或到达0位置为止；
*/
func insertSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				utils.Swap(arr, j, j-1)
			}
		}
	}

	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = insertSort(arr)
	fmt.Println("sorted:", arr)
}
