/*
 * File: oldMaster.go
 * Created Date: 2023-03-13 02:57:46
 * Author: ysj
 * Description:  模板方式模式——老师傅铸剑
 */

package main

type OldMasterForgeSword struct {
	ForgeSwordBase
}

func NewOldMasterForgeSword() ForgeSwordTemplate {
	return &OldMasterForgeSword{}
}
