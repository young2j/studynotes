/*
 * File: main.go
 * Created Date: 2022-01-04 07:08:57
 * Author: ysj
 * Description:  希尔排序
 */

package main

import (
	"algorithms/utils"
	"fmt"
)

// 对指定下标的元素做插入排序
// arr-数组， idxs-下标数组
func insertSort(arr []int, idxs []int) {
	for i := 0; i < len(idxs); i++ {
		for j := i; j > 0; j-- {
			if arr[idxs[j]] < arr[idxs[j-1]] {
				utils.Swap(arr, idxs[j], idxs[j-1])
			}
		}
	}
}

/*
希尔排序:
先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
待整个序列中的记录“基本有序”时，再对全体记录进行依次直接插入排序。
 arr = [4, 6, 7, 8, 8, 4 , 2, 9, 8, 5, 2, 1, 7]
        0  1  2  3  4  5   6  7  8  9  10 11 12
 gap = len(arr) = 13
 gap = gap/2 = 13/2 = 6  [0,6)上：下标差6取数
										     [0,6,12] [1,7] [2,8] [3,9] [4,10] [5,11] 对每一组做插入排序
 gap = gap/2 = 6/2  = 3  [0,3)上：下标差3取数
                         [0,3,6,9,12] [1,4,7,10] [2,5,8,11] 对每一组做插入排序
 gap = gap/2 = 3/2  = 1  [0,1)上：下标差1取数，就是整个数组
												 [0,1,2,3,4,5,6,7,8,9,10,11,12] 最后做一次插入排序
*/
func shellSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	gap := len(arr) / 2
	for gap != 0 { // 等于0时停止，运行完1时结束
		for i := 0; i < gap; i++ {
			idxs := make([]int, (len(arr)/gap)+1) // 准备一个存放下标的slice，容量和长度最多(len(arr)/gap)+1
			idx := i
			for idx < len(arr) {
				idxs = append(idxs, idx)
				idx += gap
			}
			insertSort(arr, idxs)
		}
		gap = gap / 2
	}

	return arr
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = shellSort(arr)
	fmt.Println("sorted:", arr)
}
