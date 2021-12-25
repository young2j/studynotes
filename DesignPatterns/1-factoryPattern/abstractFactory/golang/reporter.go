/*
 * File: reporter.go
 * Created Date: 2021-11-29 06:25:02
 * Author: ysj
 * Description:  报告
 */

package main

import "fmt"

// 接口
type IReporter interface {
	exportReport()
}

type Reporter struct {
	ReportType string
}

func (r *Reporter) exportReport() {
	fmt.Printf("导出%s\n", r.ReportType)
}

// 具体实现
func NewTrialReporter() IReporter {
	return &Reporter{
		ReportType: "简要excel报告",
	}
}

// 具体实现
func NewBasicReporter() IReporter {
	return &Reporter{
		ReportType: "详细excel报告",
	}
}

// 具体实现
func NewPremiumReporter() IReporter {
	return &Reporter{
		ReportType: "精美pdf报告",
	}
}
