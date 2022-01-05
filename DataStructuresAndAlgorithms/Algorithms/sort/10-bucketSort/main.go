/*
 * File: main.go
 * Created Date: 2022-01-05 05:40:09
 * Author: ysj
 * Description:  桶排序
 */
package main

import (
	"algorithms/utils"
	"fmt"
)

/*
1. 确定每个桶的大小size ->> 自定义大小，如果数组长度小于这个size，实质就退化为了对每个桶应用的排序算法。
2. 计算桶的数量(max-min)/size+1
3. 确定每个数入哪个桶，桶序号为(v-min)/size(桶从0号开始编号,这里不用特意+1,除非从1开始编号)
4. 对每个桶中的数进行排序(如插入排序)
5. 把每个桶中的数字顺序连接即可
示例:
arr = [3,4,2,2,1,5,3,7,9] size=3,max=9,min=1
桶数量为(max-min)/size=(9-1)/3+1=3
				[...]    [...]     [...]
				  0        1         2
每个数字入桶:
	3 -> (3-1)/3 = 0
	4 -> (4-1)/3 = 1
	2 -> (2-1)/3 = 0
	2 -> (2-1)/3 = 0
	1 -> (1-1)/3 = 0
	5 -> (5-1)/3 = 1
	3 -> (3-1)/3 = 0
	7 -> (7-1)/3 = 2
	9 -> (9-1)/3 = 2

	[3,2,2,1,3]  [4,5]  [7,9]
			0          1      2
对每个桶中的数做一次插入排序:
	[1,2,2,3,3]  [4,5]  [7,9]
最后，串起来完成排序：
	[1,2,2,3,3,4,5,7,9]
*/
func bucketSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	size := 3
	max := 0
	min := 0
	for _, v := range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	buckets := make([][]int, (max-min)/size+1)
	for _, v := range arr {
		n := (v - min) / size
		buckets[n] = append(buckets[n], v) //可以优化为入桶时马上做一次插入排序操作
	}
	arr = make([]int, 0) // 直接顶掉原arr，节约空间
	for _, bucket := range buckets {
		insertSort(bucket)
		arr = append(arr, bucket...)
	}
	return arr
}

// 插入排序
func insertSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				utils.Swap(arr, j, j-1)
			}
		}
	}
}

func main() {
	arr := utils.RandomIntArr(10, 100)
	fmt.Println("origin:", arr)
	arr = bucketSort(arr)
	fmt.Println("sorted:", arr)
}
