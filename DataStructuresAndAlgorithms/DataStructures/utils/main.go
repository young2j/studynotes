/*
 * File: main.go
 * Created Date: 2021-12-10 06:08:50
 * Author: ysj
 * Description:  工具人
 */

package utils

import (
	"math/rand"
)

func Swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func RandomIntArr(eleNum int, maxValue int) []int {
	rand.Seed(100)
	arr := make([]int, eleNum)
	if maxValue == 0 {
		maxValue = 100
	}
	for i := 0; i < eleNum; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
