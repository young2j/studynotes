package template

import (
	"fmt"
	"xiaolan/constants"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// GenSheets 生成模板工作表
func GenSheets(resultBookName string) *excelize.File {
	f := excelize.NewFile()
	_ = f.NewSheet(constants.PlanSheetName)
	_ = f.NewSheet(constants.ExecSheetName)

	err := f.SetSheetFormatPr(constants.PlanSheetName, excelize.DefaultColWidth(24), excelize.DefaultRowHeight(15))
	if err != nil {
		fmt.Println(err)
	}

	numStyle, err := f.NewStyle(`{"number_format": 2}`)
	err = f.SetColStyle(constants.PlanSheetName, "F:N", numStyle)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetColWidth(constants.PlanSheetName, "F", "N", 24)
	if err != nil {
		fmt.Println(err)
	}

	err = f.SetSheetFormatPr(constants.ExecSheetName, excelize.DefaultColWidth(22), excelize.DefaultRowHeight(15))
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetColStyle(constants.ExecSheetName, "F:K", numStyle)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetColWidth(constants.ExecSheetName, "F", "K", 22)
	if err != nil {
		fmt.Println(err)
	}


	err = f.SetSheetRow(constants.PlanSheetName, "A1", &[]string{
		"公司名称",
		"项目名称",
		"收支类型",
		"单位",
		"款项内容",
		"上月末帐列应付账款",
		"本月预计发生应付账款",
		"应付账款合计",
		"本月计划收付款",
		"其中1：计划支付商业承兑",
		"其中2：计划支付银行承兑",
		"其中3：计划工程抵款",
		"其中4：计划支付现金",
		"集团批复金额",
		"备注",
	})

	err = f.SetSheetRow(constants.ExecSheetName, "A1", &[]string{
		"公司名称",
		"项目名称",
		"年度",
		"月度",
		"项目",
		"批复金额",
		"追加金额",
		"收支计划合计",
		"实际收支金额（NC取数）",
		"调整数",
		"实际收支金额（实际数）",
		"调整原因",
		"所属片区",
	})

	if err != nil {
		fmt.Println(err)
	}

	style, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#22a5f1"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle(constants.PlanSheetName, "A1", "O1", style)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle(constants.ExecSheetName, "A1", "M1", style)
	if err != nil {
		fmt.Println(err)
	}

	if err = f.SaveAs(resultBookName); err != nil {
		fmt.Println(err)
	}
	return f
}
