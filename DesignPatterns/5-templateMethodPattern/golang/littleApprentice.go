/*
 * File: littleApprentice.go
 * Created Date: 2023-03-13 03:00:07
 * Author: ysj
 * Description:  模板方式模式——小徒弟铸剑
 */

package main

import "fmt"

type LittleApprentice struct {
	ForgeSwordBase
}

// 初始化
func NewLittleApprentice() ForgeSwordTemplate {
	return &LittleApprentice{}
}

func (la *LittleApprentice) ForgeSword() {
	la.makeClayModel()
	la.dispenseMaterials()
	la.smeltingMaterials()
	la.waterShaping()
	la.repairProcessing()

	la.result()
}

func (la *LittleApprentice) dispenseMaterials() {
	fmt.Println("调剂材料比例不均 +10分")
	la.Score += 10
}

func (la *LittleApprentice) smeltingMaterials() {
	fmt.Println("熔炼原料火候不够 +10分")
	la.Score += 10
}
