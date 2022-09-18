/*
 * File: main.go
 * Created Date: 2022-02-09 10:35:52
 * Author: ysj
 * Description: gozero 辅助生成api和proto工具
 */

package main

import (
	"flag"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

var descRegx = regexp.MustCompile(`description:"(.*)"`)

func getDescTagValue(tagValue string) string {
	res := descRegx.FindStringSubmatch(tagValue)
	if len(res) > 1 {
		return res[1]
	}
	return ""
}

func parseInputFileName(fileName string) (string, string) {
	dir, file := filepath.Split(fileName)
	file = strings.TrimSuffix(file, "Model.go")
	modelName := strcase.ToCamel(file)
	return dir, modelName
}

func main() {
	target := flag.String("t", "", "生成的目标, tag-补全模型tag, api-生成.api文件, rpc-生成.proto文件, logic-补全rpc逻辑代码, python-生成python代码")
	inputFile := flag.String("f", "", "传入goctl生成的模型文件,如 usersCustomerModel.go\n传入补全rpc代码时的.proto文件,如 access.proto")
	serviceName := flag.String("s", "", "模块的名称，默认提取模型文件名第一个单词")

	flag.Parse()

	dir, fileName := parseInputFileName(*inputFile)

	switch *target {
	case "logic":
		genLogic(*inputFile)
	case "tag":
		// 补全模型tag
		fields := parseModelDefinition(*inputFile, fileName)
		genFieldTag(fileName, fields)
	case "api":
		// 生成API文件
		fields := parseModelDefinition(*inputFile, fileName)
		genApi(dir, fileName, *serviceName, fields)
	case "rpc":
		// 生成Proto文件
		fields := parseModelDefinition(*inputFile, fileName)
		genProto(dir, fileName, *serviceName, fields)
	case "python":
		// 生成fastapi代码
		fields := parseModelDefinition(*inputFile, fileName)
		genPython(dir, fileName, *serviceName, fields)

	default:
		fields := parseModelDefinition(*inputFile, fileName)
		// 补全模型tag
		genFieldTag(fileName, fields)
		// 生成API文件
		genApi(dir, fileName, *serviceName, fields)
		// 生成Proto文件
		genProto(dir, fileName, *serviceName, fields)
	}
}
