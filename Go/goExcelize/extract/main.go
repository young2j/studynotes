package extract

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"xiaolan/constants"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// GetWb return workbook
func GetWb(wb string) *excelize.File {
	f, err := excelize.OpenFile(wb)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

// GetCompAndProjName 获取公司名称和项目名称
func GetCompAndProjName(f *excelize.File) []interface{} {
	companyName, err := f.GetCellValue(constants.CashFlowSheetName, "B2")
	if err != nil {
		fmt.Println(err)
	}

	projName, err := f.GetCellValue(constants.CashFlowSheetName, "B3")
	if err != nil {
		fmt.Println(err)
	}
	return []interface{}{companyName, projName}
}

// GetPlanDetail 获取资金计划明细
func GetPlanDetail(f *excelize.File) [][]interface{} {
	rows, err := f.GetRows(constants.PlanDetailSheetName)
	if err != nil {
		fmt.Println(err)
	}

	results := make([][]interface{}, 0, len(rows))
	for i, row := range rows {
		// 跳过标题行
		if i == 0 {
			continue
		}

		result := make([]interface{}, 0, len(row))

		for j, colCell := range row {
			// 跳过前4列
			if j < 4 {
				continue
			}

			// 跳过收支类型列为空的行
			if j == 4 && colCell == "" {
				break
			}
			if j >= 7 {
				colCellFloat, err := strconv.ParseFloat(colCell, 32)
				if err != nil {
					result = append(result, colCell)
				} else {
					result = append(result, colCellFloat)
				}
			} else {
				result = append(result, colCell)
			}
		}
		if len(result) > 0 {
			results = append(results, result)
		}
	}
	return results
}

// InputDateRange 输入时间范围
func InputDateRange() (time.Time, time.Time, string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请输入数据文件目录(例如D:/files/data):")
		dataPath, _ := reader.ReadString('\n')
		dataPath = strings.ReplaceAll(strings.TrimSpace(dataPath), "\\", "/")
		dataPath = strings.Trim(dataPath, "\"")

		fmt.Println("请输入起始年月(以-分隔,例如2020.1-2021.12):")
		dateRange, _ := reader.ReadString('\n')
		dateRangeSlice := strings.Split(strings.TrimSpace(dateRange), "-")
		dateStart, err := time.Parse("2006.1", dateRangeSlice[0])
		if err != nil {
			fmt.Printf("时间格式输入错误! %v\n", err)
			continue
		}
		dateEnd, err := time.Parse("2006.1", dateRangeSlice[1])
		if err != nil {
			fmt.Printf("时间格式输入错误! %v\n", err)
			continue
		}
		return dateStart, dateEnd, dataPath
	}
}

// GetExecDetail 获取资金执行明细
func GetExecDetail(f *excelize.File, dateStart time.Time, dateEnd time.Time) [][]interface{} {
	rows, err := f.GetRows(constants.ExecDetailSheetName)
	if err != nil {
		fmt.Println(err)
	}

	results := make([][]interface{}, 0, len(rows))
	for i, row := range rows {
		// 跳过标题行
		if i == 0 {
			continue
		}
		result := make([]interface{}, 0, len(row))
		timeValue, _ := time.Parse("2006年1月", row[1]+row[2])
		if (timeValue.Before(dateEnd) || timeValue.Equal(dateEnd)) && (timeValue.After(dateStart) || timeValue.Equal(dateStart)) {
			for j, colCell := range row {
				if j == 0 || j == 3 {
					continue
				}
				if j >= 5 && j <= 10 {
					colCellFloat, err := strconv.ParseFloat(colCell, 32)
					if err != nil {
						result = append(result, colCell)
					} else {
						result = append(result, colCellFloat)
					}
				} else {
					result = append(result, colCell)
				}
			}
		}
		if len(result) > 0 {
			results = append(results, result)
		}
	}
	return results
}
