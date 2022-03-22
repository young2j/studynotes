/*
 * File: genProto.go
 * Created Date: 2022-02-09 05:11:17
 * Author: ysj
 * Description:  根据model定义快速生成curd proto文件
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

const protoTpl = `
syntax = "proto3";

package {{ToLower .ServiceName}};
option go_package = "./{{ToLower .ServiceName}}pb";

// {{.Table}}信息
message {{.Table}}Info {
{{- range $i,$field := .Fields}}
	{{$field.Type}} {{$field.Name}} = {{AddOne $i}}; //{{$field.Description}}
{{- end}}
}

// Upsert{{.Table}}信息
message	Upsert{{.Table}}Req {
  bytes query = 1; // upsert约束字段
  {{.Table}}Info data = 2;  // upsert操作字段
	}
message	Upsert{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
		string data = 3; // UpsertedID
	}

// 新增{{.Table}}信息
message	Add{{.Table}}Req {
	 repeated {{.Table}}Info {{ToSnake .Table}}= 1; // []*{{.Table}}Info
	}
message	Add{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
		repeated string data = 3; // objectId
	}

// 删除{{.Table}}信息
message Delete{{.Table}}Req {
	string id = 1; // objectId
	bytes query = 2; // 删除条件
}
message Delete{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
}

// 修改{{.Table}}信息
message Change{{.Table}}Req {
  string id = 1;                // objectId
  bytes query = 2;              // 修改条件
  {{.Table}}Info {{ToSnake .Table}} = 3; // 更新数据
}
message Change{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
}

// 获取{{.Table}}详情
message Get{{.Table}}Req {
	string id = 1; //objectId
	bytes query = 2; // 查询条件
}
message Get{{.Table}}Resp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
	  {{.Table}}Info data = 3; //数据
}

// 获取{{.Table}}列表
message List{{.Table}}sReq {
	int32  page = 1;       //当前页
	int32  page_size = 2;  //每页数量
	string search = 3;     // 搜索
	string sort_keys = 4;       //排序键
	repeated string project_fields = 5; // 需要的字段
  repeated string exclude_fields = 6; // 排除的字段
	bytes  query = 7; // 查询条件
}
message List{{.Table}}sResp {
		int32 code = 1; //返回码
		string msg = 2; //消息说明
	  repeated {{.Table}}Info data = 3; //数据
	  int32 count = 4;  //总数
	  int32 total_page = 5; //总页数
}

service {{ToCamel .ServiceName}}Service {
  // Upsert{{.Table}}
  rpc Upsert{{.Table}}(Upsert{{.Table}}Req) returns (Upsert{{.Table}}Resp);
  
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

func genProto(dir, modelName, serviceName string, astFields []*ast.Field) {
	snakeName := strcase.ToSnake(modelName) // users_staff
	table := ""
	if serviceName == "" {
		snakeNameSplits := strings.SplitN(snakeName, "_", 2) // [users, staff]
		serviceName = snakeNameSplits[0]                     // users
		table = strcase.ToCamel(snakeNameSplits[1])          // staff
	} else {
		table = strings.TrimPrefix(modelName, strcase.ToCamel(serviceName))
	}
	fmt.Printf("serviceName: %v\n", serviceName)
	fmt.Printf("modelName: %v\n", modelName)
	fmt.Printf("tableName: %v\n", table)

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

		switch fieldtt := astField.Type.(type) {
		case *ast.Ident:
			field.Type = fieldtt.Name
		case *ast.SelectorExpr:
			field.Type = fieldtt.X.(*ast.Ident).Name + "." + fieldtt.Sel.Name
			if field.Type == "bson.ObjectId" || field.Type == "time.Time" {
				field.Type = "string"
			}
		case *ast.ArrayType:
			elt := fieldtt.Elt
			eltype := ""
			switch eltt := elt.(type) {
			case *ast.Ident:
				eltype = eltt.Name
			case *ast.StarExpr:
				eltype = eltt.X.(*ast.Ident).Name
			case *ast.SelectorExpr:
				eltype = eltt.X.(*ast.Ident).Name + "." + eltt.Sel.Name
				if eltype == "bson.ObjectId" || eltype == "time.Time" {
					eltype = "string"
				}
			}
			field.Type = "repeated " + eltype
		}
		tplData.Fields = append(tplData.Fields, field)
	}

	funcs := template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
		"ToSnake": strcase.ToSnake,
		"AddOne":  func(i int) int { return i + 1 },
	}

	tpl, err := template.New(modelName).Funcs(funcs).Parse(protoTpl)
	if err != nil {
		log.Fatal(err)
	}

	target := filepath.Join(dir, strings.ToLower(serviceName)+".proto")
	t, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(t, tplData)
}
