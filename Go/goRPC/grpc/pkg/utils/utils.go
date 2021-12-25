/*
 * File: utils.go
 * Created Date: 2021-12-25 12:49:43
 * Author: ysj
 * Description:  工具类
 */

package utils

import (
	"time"
)

// 判断某个元素是否在切片中
func EleInStrArray(arr []string, Ele string) bool {
	for _, item := range arr {
		if Ele == item {
			return true
		}
	}
	return false
}

// 判断切片某个元素是否在另一个切片中
func EleAnyInStrArray(arr1 []string, arr2 []string) bool {
	for _, item := range arr2 {
		if EleInStrArray(arr1, item) {
			return true
		}
	}
	return false
}

// 切片去重
func RemoveDuplicatedEle(arr []string) []string {
	temp := map[string]struct{}{}
	for _, ele := range arr {
		if _, ok := temp[ele]; !ok {
			temp[ele] = struct{}{}
		}
	}
	result := []string{}
	for k := range temp {
		result = append(result, k)
	}
	return result
}

// 获取当前时间
func GetTimeNow() time.Time {
	cstSh, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Now()
	}
	return time.Now().In(cstSh)
}
