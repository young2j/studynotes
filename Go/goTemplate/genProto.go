/*
 * File: genProto.go
 * Created Date: 2022-02-09 05:11:17
 * Author: ysj
 * Description:  根据model定义快速生成curd proto文件
 */

package main

import (
	"go/ast"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const protoTpl = `
syntax = "proto3";

package {{.ServiceName}};
option go_package = "./{{.ServiceName}}pb";

// 新增{{.Table}}信息
message	Add{{.Table}}Req {
	{{- range $i,$field := .Fields -}}
		{{if ne $field.Name "id"}}
		{{$field.Type}} {{$field.Name}} = {{$i}};  //{{$field.Description}}
		{{- end}}
	{{- end}}
	}
message	Add{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
	}

// 删除{{.Table}}信息
message Delete{{.Table}}Req {
	string id = 1; //objectId
}
message Delete{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
}

// 修改{{.Table}}信息
message Change{{.Table}}Req {
{{- range $i,$field := .Fields}}
	{{$field.Type}} {{$field.Name}} = {{AddOne $i}}; //{{$field.Description}}
{{- end}}
}
message Change{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
}

// 获取{{.Table}}信息
message {{.Table}}Info {
{{- range $i,$field := .Fields}}
	{{$field.Type}} {{$field.Name}} = {{AddOne $i}}; //{{$field.Description}}
{{- end}}
}
// 获取{{.Table}}详情
message Get{{.Table}}Req {
	string id = 1; //objectId
}
message Get{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
	  {{.Table}}Info data = 3; //数据
}

// 获取{{.Table}}列表
message List{{.Table}}sReq {
	optional int32 page = 1;     //当前页
	optional int32 page_size = 2;  //每页数量
}
message List{{.Table}}sResp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
	  repeated {{.Table}}Info data = 3; //数据
	  int32 count = 4;  //总数
	  int32 total_page = 5; //总页数
}

service {{ToCamel .ServiceName}}Service {
  // 新增{{.Table}}
  rpc Add{{.Table}}(Add{{.Table}}Req) returns (Add{{.Table}}Resp);

  // 删除{{.Table}}
  rpc Delete{{.Table}}(Delete{{.Table}}Req) returns (Delete{{.Table}}Resp);

  // 修改{{.Table}}
  rpc Change{{.Table}}(Change{{.Table}}Req) returns (Change{{.Table}}Resp);

	// 查询{{.Table}}
	rpc Get{{.Table}}(Get{{.Table}}Req) returns(Get{{.Table}}Resp);

  // {{.Table}}列表
  rpc List{{.Table}}s(List{{.Table}}sReq) returns (List{{.Table}}sResp);
}
`

type ProtoTplData struct {
	ServiceName string
	Table       string
	Fields      []*Field
}

func genProto(dir, modelName string, astFields []*ast.Field) {
	snakeName := strcase.ToSnake(modelName)          // users_staff
	snakeNameSplits := strings.Split(snakeName, "_") // [users, staff]
	serviceName := snakeNameSplits[0]                // users
	table := strcase.ToCamel(snakeNameSplits[1])     // staff

	tplData := &ProtoTplData{
		ServiceName: serviceName,
		Table:       table,
		Fields:      make([]*Field, 0),
	}

	for _, astField := range astFields {
		description := strings.TrimSpace(astField.Comment.Text())
		if astField.Tag != nil && description == "" {
			description = getDescTagValue(astField.Tag.Value)
		}

		field := &Field{
			Name:        strcase.ToSnake(astField.Names[0].Name),
			Description: description,
		}
		switch astField.Type.(type) {
		case *ast.Ident:
			field.Type = astField.Type.(*ast.Ident).Name
		case *ast.SelectorExpr:
			expr, ok := astField.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			field.Type = expr.X.(*ast.Ident).Name + "." + expr.Sel.Name
			if field.Type == "bson.ObjectId" || field.Type == "time.Time" {
				field.Type = "string"
			}
		case *ast.ArrayType:
			elemType := astField.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
			field.Type = "repeated " + elemType
		}
		tplData.Fields = append(tplData.Fields, field)
	}

	funcs := template.FuncMap{"ToCamel": strcase.ToCamel, "AddOne": func(i int) int { return i + 1 }}

	tpl, err := template.New(modelName).Funcs(funcs).Parse(protoTpl)
	if err != nil {
		log.Fatal(err)
	}

	target := filepath.Join(dir, serviceName+".proto")
	t, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(t, tplData)
}
