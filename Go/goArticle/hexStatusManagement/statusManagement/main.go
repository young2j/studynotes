/*
 * File: main.go
 * Created Date: 2022-01-13 04:47:00
 * Author: ysj
 * Description: 十六进制状态管理
 */

package main

import "fmt"

const (
	STATUS_1 = 0x0001 // 初始状态1
	STATUS_2 = 0x0002 // 初始状态2
	STATUS_3 = 0x0004 // 中间状态1
	STATUS_4 = 0x0008 // 中间状态2
	STATUS_5 = 0x0010 // 最终状态1
	STATUS_6 = 0x0020 // 最终状态2
)

// 状态组:当需要添加状态时，使用或运算|
const (
	INIT_STATUS   = STATUS_1 | STATUS_2 // 初始状态
	MIDDLE_STATUS = STATUS_3 | STATUS_4 // 中间状态
	FINAL_STATUS  = STATUS_5 | STATUS_6 // 最终状态
)

// 当需要判断是否包含某种状态时使用与运算&,结果为0则代表不包含指定状态
const (
	CONTAINS_STATUS_1 = (INIT_STATUS&STATUS_1 != 0)
	CONTAINS_STATUS_3 = (INIT_STATUS&STATUS_3 != 0)
)

// 当需要排除状态时，使用与运算、取反运算&^(注意取反操作在go中为^，其他语言通常为~)
const (
	INIT_STATUS_1   = INIT_STATUS & ^STATUS_2   // == STATUS_1
	MIDDLE_STATUS_3 = MIDDLE_STATUS & ^STATUS_4 // == STATUS_3
	FINAL_STATUS_5  = FINAL_STATUS & ^STATUS_6  // == STATUS_5
)

func main() {
	fmt.Println(INIT_STATUS, MIDDLE_STATUS, FINAL_STATUS)
	fmt.Println("初始状态包含状态1:", CONTAINS_STATUS_1)
	fmt.Println("初始状态包含状态3:", CONTAINS_STATUS_3)
	fmt.Println(
		INIT_STATUS_1 == STATUS_1,
		MIDDLE_STATUS_3 == STATUS_3,
		FINAL_STATUS_5 == STATUS_5,
	)
}
