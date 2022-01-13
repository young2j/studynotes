/*
 * File: main.go
 * Created Date: 2022-01-13 05:20:45
 * Author: ysj
 * Description:  十六进制标签管理
 */

package main

import "fmt"

const (
	LABEL_1 = 0x0001 // 正常1
	LABEL_2 = 0x0002 // 正常2
	LABEL_3 = 0x0004 // 违规1
	LABEL_4 = 0x0008 // 违规2
	LABEL_5 = 0x0010 // 无法判断1
	LABEL_6 = 0x0020 // 无法判断2
)

// 标签组:当需要添加标签结果时，使用或运算|
const (
	NORMAL_LABEL  = LABEL_1 | LABEL_2 // 正常
	EVIL_LABEL    = LABEL_3 | LABEL_4 // 违规
	UNKNOWN_LABEL = LABEL_5 | LABEL_6 // 无法判断
)

// 当需要判断是否包含某种标签时使用与运算&,结果为0则代表不包含指定标签结果
const (
	CONTAINS_LABEL_1 = (NORMAL_LABEL&LABEL_1 != 0)
	CONTAINS_LABEL_3 = (NORMAL_LABEL&LABEL_3 != 0)
)

// 当需要排除标签时，使用与运算、取反运算&^(注意取反操作在go中为^，其他语言通常为~)
const (
	NORMAL_LABEL_1  = NORMAL_LABEL & ^LABEL_2  // == LABEL_1
	EVIL_LABEL_3    = EVIL_LABEL & ^LABEL_4    // == LABEL_3
	UNKNOWN_LABEL_5 = UNKNOWN_LABEL & ^LABEL_6 // == LABEL_5
)

func main() {
	fmt.Println(NORMAL_LABEL, EVIL_LABEL, UNKNOWN_LABEL)
	fmt.Println(CONTAINS_LABEL_1, CONTAINS_LABEL_3)
	fmt.Println(
		NORMAL_LABEL_1 == LABEL_1,
		EVIL_LABEL_3 == LABEL_3,
		UNKNOWN_LABEL_5 == LABEL_5,
	)
}
