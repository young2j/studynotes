/*
 * File: main.go
 * Created Date: 2021-12-20 01:49:56
 * Author: ysj
 * Description:  堆排序
 */
package main

import (
	"algorithms/utils"
	"fmt"
)

/*
 1. 完全二叉树: 最下层叶子结点从左到右依次排列；其余节点都是满节点。
 2. 堆：
 		大根堆：所有父节点都大于其子节点的完全二叉树；(子节点：子树上的所有节点)
 		小根堆：所有父节点都小于其子节点的完全二叉树；(子节点：子树上的所有节点)
 3. 可以用一位数组模拟完全二叉树，自上而下、从左到右按顺序编号：
    1). 序号即为数组下标；
		2). 下标为i的数，对应树中序号为i的节点；
		3). 节点i的父节点的序号为(i-1)/2，左子节点的序号为(2i+1), 右子节点的序号为(2i+2);
4. 举例:
	arr [ 3, 9, 1, 6, 7, 1, 9, 5 ]
	 i    0  1  2  3  4  5  6  7

	 i = 0, 父=(0-1)/2=-0=0, 左=(2*0+1)=1, 右=(2*0+2)=2;
	 i = 1, 父=(1-1)/2=0,    左=(2*1+1)=3, 右=(2*1+2)=4;
	 i = 2, 父=(2-1)/2=0,    左=(2*2+1)=5, 右=(2*2+2)=6;
	 i = 3, 父=(3-1)/2=1,    左=(2*3+1)=7, 右=(2*3+2)=8>7=无;
	 i = 4, 父=(4-1)/2=1,    左=(2*4+1)>7=无, 右=(2*4+2)>7=无;
	 i = 5, 父=(5-1)/2=2,    左=(2*5+1)>7=无, 右=(2*5+2)>7=无;
	 i = 6, 父=(6-1)/2=2,    左=(2*6+1)>7=无, 右=(2*6+2)>7=无;
	 i = 7, 父=(7-1)/2=3,    左=(2*7+1)>7=无, 右=(2*7+2)>7=无;
*/

//===================================构建堆===============================
// 构建大根堆：
// 自顶向下，大数不断上浮 O(N*logN)
func buildMaxHeapFromTop2Bottom(arr []int) {
	// 自顶向下
	for i := 0; i < len(arr); i++ {
		// 1. 将当前数与其父节点比较，若大于父节点则交换位置；
		// 2. 将父节点置为当前数，重复步骤1；
		// 3. 当前数<=父节点时，停止。

		for cur := i; arr[cur] > arr[(cur-1)/2]; cur = (cur - 1) / 2 {
			utils.Swap(arr, cur, (cur-1)/2)
		}
	}
}

// 构建小根堆:
// 自顶向下，小数不断上浮 O(N*logN)
func buildMinHeapFromTop2Bottom(arr []int) {
	// 自顶向下
	for i := 0; i < len(arr); i++ {
		// 1. 将当前数与其父节点比较，若小于父节点则交换位置；
		// 2. 将父节点置为当前数，重复步骤1；
		// 3. 当前数>=父节点时，停止。

		for cur := i; arr[cur] < arr[(cur-1)/2]; cur = (cur - 1) / 2 {
			utils.Swap(arr, cur, (cur-1)/2)
		}

		// cur := i
		// for arr[cur] < arr[(cur-1)/2] {
		// 	utils.Swap(arr, cur, (cur-1)/2)
		// 	cur = (cur - 1) / 2
		// }
	}
}

/*
构建大根堆：
自底向上，大数不断上浮,小数不断下沉 O(N)
*/
func buildMaxHeapFromBottom2Top(arr []int) {
	// 自底向上
	for i := len(arr) - 1; i >= 0; i-- {
		// 1. 以当前节点为父节点，往下看
		// 2. 找最大：找出左右子节点、父节点三者中数最大者，记最大数下标为largest
		// 3. 停  止：如果最大数为父节点，即largest==i，停止；
		// 4. 上  浮：如果最大数为子节点之一，则子节点上浮，与父节点交换；
		// 5. 下  沉：继续以最大数(子节点较大者)为父节点，重复上述步骤。
		p := i
		left := (2*p + 1)
		for left < len(arr) {
			// 找左右子节点最大者：存在右子节点、且右子节点较大
			largest := left
			if (left+1) < len(arr) && arr[left] < arr[left+1] {
				largest = left + 1
			}
			// 找较大子节点与父节点中的最大者
			if arr[p] > arr[largest] {
				largest = p
			}

			// 如果最大数为父节点, 则停止，无需上浮下沉
			if largest == p {
				break
			}
			// 如果最大数为子节点之一，则子节点上浮
			utils.Swap(arr, p, largest)
			// 以largest为下次循环的父节点，往下游走，让小数继续下沉
			p = largest
			// 计算其左子节点，继续循环
			left = p*2 + 1
		}
	}
}

/*
构建小根堆：
自底向上，小数不断上浮，大数不断下沉 O(N)
*/
func buildMinHeapFromBottom2Top(arr []int) {
	// 自底向上
	for i := len(arr) - 1; i >= 0; i-- {
		// 1. 以当前节点为父节点，往下看
		// 2. 找最小：找出左右子节点、父节点三者中数最小者，记最小数下标为smallest
		// 3. 停  止：如果最小数为父节点，即smallest==i，停止；
		// 4. 上  浮：如果最小数为子节点之一，则子节点上浮，与父节点交换；
		// 5. 下  沉：继续以最小数(子节点较小者)为父节点，重复上述步骤。
		p := i
		left := (2*p + 1)
		for left < len(arr) {
			// 找左右子节点最小者：存在右子节点、且右子节点较小
			smallest := left
			if (left+1) < len(arr) && arr[left] > arr[left+1] {
				smallest = left + 1
			}
			// 找较小子节点与父节点中的最小者
			if arr[p] < arr[smallest] {
				smallest = p
			}

			// 如果最小数为父节点, 则停止，无需上浮下沉
			if smallest == p {
				break
			}
			// 如果最小数为子节点之一，则子节点上浮
			utils.Swap(arr, p, smallest)
			// 以smallest为下次循环的父节点，往下游走，让大数下沉到底
			p = smallest
			// 计算其左子节点，继续循环
			left = p*2 + 1
		}
	}
}

func main() {
	arr1 := []int{3, 9, 1, 6, 7, 1, 9, 5}
	fmt.Println("原数组:", arr1)

	// buildMaxHeapFromTop2Bottom(arr1)
	buildMaxHeapFromBottom2Top(arr1)
	fmt.Println("大顶堆:", arr1)

	arr2 := []int{3, 9, 1, 6, 7, 1, 9, 5}
	// buildMinHeapFromTop2Bottom(arr2)
	buildMinHeapFromBottom2Top(arr2)
	fmt.Println("小顶堆:", arr2)
}
