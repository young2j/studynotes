/*
 * File: template.go
 * Created Date: 2023-03-13 02:44:39
 * Author: ysj
 * Description:  golang 模板方法
 */
package main

import "fmt"

type ForgeSwordTemplate interface {
	ForgeSword()        // 铸剑模板方法
	makeClayModel()     // 制做泥范, 每道工序20分
	dispenseMaterials() // 调剂材料, 每道工序20分
	smeltingMaterials() // 熔炼原料, 每道工序20分
	waterShaping()      // 浇灌成形, 每道工序20分
	repairProcessing()  // 修治加工, 每道工序20分
}

type ForgeSwordBase struct {
	Score int
}

func (f *ForgeSwordBase) ForgeSword() {
	f.makeClayModel()
	f.dispenseMaterials()
	f.smeltingMaterials()
	f.waterShaping()
	f.repairProcessing()
	f.result()
}

func (f *ForgeSwordBase) makeClayModel() {
	fmt.Println("制做泥范完美 +20分")
	f.Score += 20
}

func (f *ForgeSwordBase) dispenseMaterials() {
	fmt.Println("调剂材料完美 +20分")
	f.Score += 20
}

func (f *ForgeSwordBase) smeltingMaterials() {
	fmt.Println("熔炼原料完美 +20分")
	f.Score += 20
}

func (f *ForgeSwordBase) waterShaping() {
	fmt.Println("浇灌成形完美 +20分")
	f.Score += 20
}

func (f *ForgeSwordBase) repairProcessing() {
	fmt.Println("修治加工完美 +20分")
	f.Score += 20
}

func (f *ForgeSwordBase) result() {
	fmt.Println("总分:", f.Score)
	if f.Score == 100 {
		fmt.Println("获得绝世好剑!!!")
	} else if f.Score > 80 {
		fmt.Println("获得一把好剑!!")
	} else {
		fmt.Println("获得一把村好剑!")
	}
}
