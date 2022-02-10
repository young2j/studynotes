/*
 * File: genApi.go
 * Created Date: 2022-02-09 02:19:07
 * Author: ysj
 * Description:  根据model定义快速生成curd api文件
 */

package main

import (
	"fmt"
	"go/ast"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const apiTpl = `
syntax = "v1"

info(
	title: "{{.ServiceName}}-api"
	desc: "xxx api"
	author: "xxx"
	email: "xxx@knownsec.com"
	version: 1.0
)

type (
	// 新增{{.Table}}信息
	Add{{.Table}}Req {
	{{- range $field := .Fields -}}
		{{if ne $field.Name "Id"}}
		{{$field.Name}} {{$field.Type}} {{$field.Tag}}
		{{- end}}
	{{- end}}
	}
	Add{{.Table}}Resp {
		Code int32  {{.Backticks}}json:"code" description:"返回码"{{.Backticks}}
		Msg  string {{.Backticks}}json:"msg" description:"消息说明"{{.Backticks}}
	}
	// 删除{{.Table}}信息
	Delete{{.Table}}Req {
		Id string {{.Backticks}}path:"id" description:"objectId"{{.Backticks}}
	}
	Delete{{.Table}}Resp {
		Code int32  {{.Backticks}}json:"code" description:"返回码"{{.Backticks}}
		Msg  string {{.Backticks}}json:"msg" description:"消息说明"{{.Backticks}}
	}

	// 修改{{.Table}}信息
	Change{{.Table}}Req {
	{{- range $field := .Fields}}
		{{$field.Name}} {{$field.Type}} {{$field.Tag}}
	{{- end}}
	}
	Change{{.Table}}Resp {
		Code int32  {{.Backticks}}json:"code" description:"返回码"{{.Backticks}}
		Msg  string {{.Backticks}}json:"msg" description:"消息说明"{{.Backticks}}
	}

	// 获取{{.Table}}信息
	{{.Table}}Info {
	{{- range $field := .Fields}}
		{{$field.Name}} {{$field.Type}} {{$field.Tag}}
	{{- end}}
	}
	// 获取{{.Table}}详情
	Get{{.Table}}Req {
		Id string {{.Backticks}}path:"id" description:"objectId"{{.Backticks}}
	}
	Get{{.Table}}Resp {
		Code int32  {{.Backticks}}json:"code" description:"返回码"{{.Backticks}}
		Msg  string {{.Backticks}}json:"msg" description:"消息说明"{{.Backticks}}
		Data *{{.Table}}Info {{.Backticks}}json:"data" description:"数据"{{.Backticks}}
	}

	// 获取{{.Table}}列表
	List{{.Table}}sReq {
		Page     int32 {{.Backticks}}form:"page,optional" description:"当前页"{{.Backticks}}
		PageSize int32 {{.Backticks}}form:"pageSize,optional" description:"每页数量"{{.Backticks}}
	}
	List{{.Table}}sResp {
		Code      int32             {{.Backticks}}json:"code" description:"返回码"{{.Backticks}}
		Msg       string            {{.Backticks}}json:"msg" description:"消息说明"{{.Backticks}}
		Data      []*{{.Table}}Info {{.Backticks}}json:"data" description:"数据"{{.Backticks}}
		Count     int32             {{.Backticks}}json:"count" description:"总数"{{.Backticks}}
		TotalPage int32             {{.Backticks}}json:"totalPage" description:"总页数"{{.Backticks}}
	}
)

service {{.ServiceName}}-api {
	@doc(
	  summary: 新增{{.Table}}
	)
	@handler Add{{.Table}}
	post /{{.ServiceName}}/add-{{ToLower .Table}}(Add{{.Table}}Req) returns(Add{{.Table}}Resp)
	
	@doc(
	  summary: 删除{{.Table}}
	)
	@handler Delete{{.Table}}
	delete /{{.ServiceName}}/delete-{{ToLower .Table}}/:id(Delete{{.Table}}Req) returns(Delete{{.Table}}Resp)
	
	@doc(
	  summary: 修改{{.Table}}
	)
	@handler Change{{.Table}}
	put /{{.ServiceName}}/change-{{ToLower .Table}}(Change{{.Table}}Req) returns(Change{{.Table}}Resp)
	
	@doc(
	  summary: 查询{{.Table}}
	)
	@handler Get{{.Table}}
	get /{{.ServiceName}}/get-{{ToLower .Table}}/:id(Get{{.Table}}Req) returns(Get{{.Table}}Resp)
	
	@doc(
	  summary: {{.Table}}列表
	)
	@handler List{{.Table}}s
	get /{{.ServiceName}}/list-{{ToLower .Table}}s(List{{.Table}}sReq) returns(List{{.Table}}sResp)
}
`

type ApiTplData struct {
	ServiceName string
	Table       string
	Fields      []*Field
	Backticks   string
}

func genApi(dir, modelName string, astFields []*ast.Field) {
	snakeName := strcase.ToSnake(modelName)
	snakeNameSplits := strings.Split(snakeName, "_")
	serviceName := snakeNameSplits[0]
	table := strcase.ToCamel(snakeNameSplits[1])

	tplData := &ApiTplData{
		ServiceName: serviceName,
		Table:       table,
		Fields:      make([]*Field, 0),
		Backticks:   "`",
	}

	for _, astField := range astFields {
		field := &Field{
			Name: astField.Names[0].Name,
		}

		lowerCamelFieldName := strcase.ToLowerCamel(field.Name)
		description := strings.TrimSpace(astField.Comment.Text())
		if astField.Tag != nil && description == "" {
			description = getDescTagValue(astField.Tag.Value)
		}

		switch astField.Type.(type) {
		case *ast.Ident:
			field.Type = astField.Type.(*ast.Ident).Name
			field.Tag = fmt.Sprintf("`json:\"%s\" description:\"%s\"`", lowerCamelFieldName, description)
		case *ast.SelectorExpr:
			expr, ok := astField.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			field.Type = expr.X.(*ast.Ident).Name + "." + expr.Sel.Name
			if field.Type == "bson.ObjectId" || field.Type == "time.Time" {
				field.Type = "string"
			}
			field.Tag = fmt.Sprintf("`json:\"%s\" description:\"%s\"`", lowerCamelFieldName, description)

		case *ast.ArrayType:
			elemType := astField.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
			field.Type = "[]" + elemType
			field.Tag = fmt.Sprintf("`json:\"%s\" description:\"%s\"`", lowerCamelFieldName, description)
		}
		tplData.Fields = append(tplData.Fields, field)
	}
	funcs := template.FuncMap{"ToLower": strings.ToLower}
	tpl, err := template.New(modelName).Funcs(funcs).Parse(apiTpl)
	if err != nil {
		log.Fatal(err)
	}

	target := filepath.Join(dir, serviceName+".api")
	t, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(t, tplData)
}
