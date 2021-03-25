package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"
	"xiaolan/constants"
	"xiaolan/extract"
	"xiaolan/templates"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// var wg sync.WaitGroup
var wg sync.WaitGroup

func rwPlanSheetResult(p Params) {
	defer wg.Done()
	compAndProjName := extract.GetCompAndProjName(p.sourceFile)
	results := extract.GetPlanDetail(p.sourceFile)

	for i, result := range results {
		row := append(compAndProjName, result...)
		err := p.targetFile.SetSheetRow(constants.PlanSheetName, fmt.Sprintf("A%v", 2+i+*p.count), &row)
		if err != nil {
			log.Fatalf("提取资金计划失败【%v】, ERROR:%v \n", p.wb, err)
		}
	}
	*p.count += len(results)
	log.Printf("提取资金计划完成【%v】\n", p.wb)
}

func rwExecSheetResult(p Params) {
	defer wg.Done()
	compAndProjName := extract.GetCompAndProjName(p.sourceFile)
	results := extract.GetExecDetail(p.sourceFile, p.dateStart, p.datEnd)
	for i, result := range results {
		row := append(compAndProjName, result...)
		err := p.targetFile.SetSheetRow(constants.ExecSheetName, fmt.Sprintf("A%v", 2+i+*p.count), &row)
		if err != nil {
			log.Fatalf("提取资金执行失败【%v】, ERROR:%v \n", p.wb, err)
		}
	}
	*p.count += len(results)
	log.Printf("提取资金执行完成【%v】\n", p.wb)
}

type Params struct {
	sourceFile *excelize.File
	targetFile *excelize.File
	count      *int
	wb         string
	dateStart  time.Time
	datEnd     time.Time
}

func main() {
	str := `
#-----------------------------------------------------------#
|                BelongsTo: 陈肖兰                          |
#-----------------------------------------------------------#
  使用说明:                                                  
    1. 输入数据文件所在目录路径                                
    2. 若文件目录为空，则默认为程序文件所在目录下的data文件夹     
    3. 输入时间范围                                           
    4. 没有了                                                
-------------------------------------------------------------
`

	fmt.Println(str)

	dateStart, dateEnd, dataPath := extract.InputDateRange()
	start := time.Now()
	
	
	log.Println("开始提取数据......")
	targetFile := template.GenSheets()
	log.Println("生成模板文件......")

	if dataPath == "" {
		pwd, _ := os.Getwd()
		dataPath = path.Join(pwd, "/data")
	}

	workbooks, err := ioutil.ReadDir(dataPath)
	if err != nil {
		fmt.Println(err)
	}

	count1 := 0
	count2 := 0
	for _, file := range workbooks {
		wb := path.Join(dataPath, file.Name())
		log.Printf("开始处理【%v】\n", wb)
		sourceFile := extract.GetWb(wb)

		p1 := Params{
			sourceFile: sourceFile,
			targetFile: targetFile,
			count:      &count1,
			wb:         wb,
		}
		p2 := Params{
			sourceFile: sourceFile,
			targetFile: targetFile,
			count:      &count2,
			wb:         wb,
			dateStart:  dateStart,
			datEnd:     dateEnd,
		}

		wg.Add(2)
		go rwPlanSheetResult(p1)
		go rwExecSheetResult(p2)
	}

	wg.Wait()

	err = targetFile.SaveAs(constants.ResultFileName)
	if err != nil {
		log.Fatalln(err)
	}

	end := time.Now()
	elapsed := end.Sub(start).Seconds()

	log.Printf("task done. elpased: %vs\n", elapsed)
}
