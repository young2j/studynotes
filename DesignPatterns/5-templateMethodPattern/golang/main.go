/*
 * File: main.go
 * Created Date: 2023-03-07 06:59:11
 * Author: ysj
 * Description: 模板方式模式——客户端调用
 */

package main

import "fmt"

func main() {
	fmt.Println("=============老师傅铸剑============")
	oldMaster := NewOldMasterForgeSword()
	oldMaster.ForgeSword()

	fmt.Println("=============小徒弟铸剑============")
	littleApprentice := NewLittleApprentice()
	littleApprentice.ForgeSword()
}
