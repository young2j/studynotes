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

func parseModelFileName(fileName string) (string, string) {
	dir, file := filepath.Split(fileName)
	file = strings.TrimSuffix(file, "Model.go")
	modelName := strcase.ToCamel(file)
	return dir, modelName
}

func main() {
	target := flag.String("t", "", "生成的目标, tag-补全模型tag, api-生成.api文件, rpc-生成.proto文件, 默认全部")
	fileName := flag.String("f", "", "传入goctl生成的模型文件,如usersCustomerModel.go")
	flag.Parse()

	dir, modelName := parseModelFileName(*fileName)
	fields := parseModelDefinition(*fileName, modelName)

	switch *target {
	case "tag":
		// 补全模型tag
		genFieldTag(modelName, fields)
	case "api":
		// 生成API文件
		genApi(dir, modelName, fields)
	case "rpc":
		// 生成Proto文件
		genProto(dir, modelName, fields)
	default:
		// 补全模型tag
		genFieldTag(modelName, fields)
		// 生成API文件
		genApi(dir, modelName, fields)
		// 生成Proto文件
		genProto(dir, modelName, fields)
	}
}
