/*
 * File: main.go
 * Created Date: 2022-01-04 07:08:57
 * Author: ysj
 * Description:  希尔排序
 */

package main

/*
先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
待整个序列中的记录“基本有序”时，再对全体记录进行依次直接插入排序。
*/

func hillSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	gap := len(arr)/2

	for gap != 1 {
				
		gap = gap / 2
	}

	return arr
}

func main() {

}
