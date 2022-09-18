/*
 * File: genPython.go
 * Created Date: 2022-09-09 10:27:46
 * Author: ysj
 * Description:  根据model定义快速生成fastapi的基础代码
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

const (
	pyApiTpl = `
import json

from fastapi import APIRouter, Query, Request, Path
from fastapi import status as HttpStatus    # noqa

from app.api.utils import send_center_request, is_valid_objectid
from app.schemas.common import Str2IntList, Str2List, CountResponse
from app.schemas.{{ToLower .ServiceName}}.{{ToLower .Table}} import (
    Add{{.Table}}Request,
    Add{{.Table}}Response,
    Change{{.Table}}Request,
    Change{{.Table}}Response,
    Get{{.Table}}Response,
    List{{.Table}}sResponse,
)

router = APIRouter()


@router.post("/add-{{ToLower .Table}}",
             name="xxx:新增{{.Table}}",
             response_model=Add{{.Table}}Response,
             response_model_by_alias=True)
async def create_{{ToLower .Table}}(request: Request, data: Add{{.Table}}Request):
    {{ToLower .Table}} = data.dict(by_alias=True, exclude_none=True)
    resp = await send_center_request(request, "{{ToLower .ServiceName}}", "post", json={{ToLower .Table}})
    return resp.json()


@router.put("/change-{{ToLower .Table}}",
            name="xxx:修改{{.Table}}",
            response_model=Change{{.Table}}Response,
            response_model_by_alias=True)
async def change_{{ToLower .Table}}(request: Request, data: Change{{.Table}}Request):
    {{ToLower .Table}} = data.dict(by_alias=True, exclude_none=True)
    resp = await send_center_request(request, "{{ToLower .ServiceName}}", "put", json={{ToLower .Table}})
    return resp.json()


@router.get("/get-{{ToLower .Table}}/{id}",
            name="xxx:{{.Table}}详情",
            response_model=Get{{.Table}}Response,
            response_model_by_alias=True)
async def get_{{ToLower .Table}}(request: Request, id: str = Path(...)):
    ok = is_valid_objectid(id)
    if not ok:
        return {"code": HttpStatus.HTTP_400_BAD_REQUEST, "msg": "{{.Table}}ID错误"}

    resp = await send_center_request(request, "{{ToLower .ServiceName}}", "get")
    return resp.json()


@router.get("/list-{{ToLower .Table}}s",
            name="xxx:{{.Table}}列表",
            response_model=List{{.Table}}sResponse,
            response_model_by_alias=True)
async def list_{{ToLower .Table}}s(
    request: Request,
    page: int = Query(1, description="页数"),
    page_size: int = Query(50, alias="pageSize", description="每页数"),
    search: str | None = Query(None, description="搜索内容"),
    sort: str | None = Query("-_id", description="排序项"),
		project_fields: Str2List | None = Query([], alias="only", description="仅包含字段,逗号分隔,与exclude互斥"),
		exclude_fields: Str2List | None = Query([], alias="exclude", description="仅排除字段,逗号分隔,与only互斥"),
    operator_names: Str2List = Query([], alias="operatorNames", description="操作人"),
    create_time: str | None = Query(None, alias="createTime", description="创建时间"),
):
    # 请求参数
    q = {"page": page, "pageSize": page_size}
    if search:
        q["search"] = search
    if sort:
        q["sort"] = sort
    if project_fields:
        q["projectFields"] = json.dumps(project_fields)
    if exclude_fields:
        q["excludeFields"] = json.dumps(exclude_fields)
    if operator_names:
        q["operatorNames"] = json.dumps(operator_names)
    if create_time:
        q["createTime"] = create_time

    # 发送请求
    resp = await send_center_request(request, "{{ToLower .ServiceName}}", "get", params=q)
    return resp.json()


@router.get("/count-{{ToLower .Table}}s", name="xxx:常用筛选项数量", response_model=CountResponse)
async def count_{{ToLower .Table}}s(request: Request,
                        search: str | None = Query(None, description="搜索内容"),
                       ):
    # 请求参数
    q = {}
    if search:
        q["search"] = search
   
    # 发送请求
    resp = await send_center_request(request, "{{ToLower .ServiceName}}", "get", params=q)
    return resp.json()
	`
	pySchemaTpl = `
from typing import List, Mapping, Any #noqa
from pydantic import BaseModel, Field
from app.schemas.common import ListMixin, OperatorNameID, DetailMixin

# base
class {{.Table}}(BaseModel):
  {{- range $field := .Fields}}
		{{- if ne $field.Name "id" }}
    {{$field.Name}}: {{$field.Type}} | None = Field(None, alias="{{ToCamel $field.Name}}", description="{{$field.Description}}")
		{{- end}}
  {{- end}}

    class Config:
        anystr_strip_whitespace = True

# add
class Add{{.Table}}Request({{.Table}}):
	{{- range $field := .Fields}}
	 {{- if NotAutoField $field.Name }}
    {{$field.Name}}: {{$field.Type}} | None = Field(..., alias="{{ToCamel $field.Name}}", description="{{$field.Description}}")
   {{- end}}
	{{- end}}

    class Config:
        anystr_strip_whitespace = True

class Add{{.Table}}Response(DetailMixin):
    pass


# change
class Change{{.Table}}Request({{.Table}}):
    id: str = Field(..., description="数据id")
   
    class Config:
        anystr_strip_whitespace = True

class Change{{.Table}}Response(DetailMixin):
    pass


# get-list
class {{.Table}}Data({{.Table}}):
    id: str = Field(..., description="数据id")
    creator: OperatorNameID | None = None
    create_time: str | None = Field(None, alias="createTime")
    operator: OperatorNameID | None = None
    update_time: str | None = Field(None, alias="updateTime")


# get
class Get{{.Table}}Response(DetailMixin):
    data: {{.Table}}Data | None = None


# list
class List{{.Table}}sResponse(ListMixin):
    data: List[{{.Table}}Data] | None = None
	`
)

type PyTplData struct {
	ServiceName string
	Table       string
	Fields      []*Field
}

func genPython(dir, modelName, serviceName string, astFields []*ast.Field) {
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

	genPyApi(dir, serviceName, table)
	genPySchema(dir, serviceName, table, astFields)
}

func genPyApi(dir, serviceName, table string) {
	tplData := &PyTplData{
		ServiceName: serviceName,
		Table:       table,
	}
	funcs := template.FuncMap{
		"ToCamel": strcase.ToLowerCamel,
		"ToLower": strings.ToLower,
		"ToSnake": strcase.ToSnake,
	}
	tpl, err := template.New(serviceName + "-pyapi").Funcs(funcs).Parse(pyApiTpl)
	if err != nil {
		log.Fatal(err)
	}

	target := filepath.Join(dir, strings.ToLower(serviceName)+"_api.py")
	t, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(t, tplData)
}

func genPySchema(dir, serviceName, table string, astFields []*ast.Field) {
	tplData := &PyTplData{
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

		// 解析go字段类型
		switch fieldtt := astField.Type.(type) {
		case *ast.Ident:
			field.Type = fieldtt.Name
		case *ast.SelectorExpr:
			field.Type = fieldtt.X.(*ast.Ident).Name + "." + fieldtt.Sel.Name
			if field.Type == "bson.ObjectId" || field.Type == "time.Time" {
				field.Type = "string"
			}
		case *ast.MapType:
			field.Type = fmt.Sprintf("map<%v, %v>", fieldtt.Key, fieldtt.Value)
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
		// 对应python字段类型
		fieldType := field.Type
		switch {
		case fieldType == "string":
			field.Type = "str"
		case fieldType == "int32":
			field.Type = "int"
		case strings.HasPrefix(fieldType, "repeated"):
			fieldType = strings.Split(fieldType, " ")[1]
			if fieldType == "string" {
				field.Type = "List[str]"
			} else if fieldType == "int32" {
				field.Type = "List[int]"
			} else {
				field.Type = "List[Any]"
			}
		case strings.HasPrefix(fieldType, "map"):
			fieldType = strings.NewReplacer("map", "", "<", "", ">", "").Replace(fieldType)
			kv := strings.Split(fieldType, ",")
			for i := 0; i < len(kv); i++ {
				if kv[i] == "string" {
					kv[i] = "str"
				} else if kv[i] == "int32" {
					kv[i] = "int"
				} else {
					kv[i] = "Any"
				}
			}
			field.Type = "Mapping[" + kv[0] + ", " + kv[1] + "]"
		}
		tplData.Fields = append(tplData.Fields, field)
	}
	funcs := template.FuncMap{
		"ToCamel": strcase.ToLowerCamel,
		"ToLower": strings.ToLower,
		"ToSnake": strcase.ToSnake,
		"NotAutoField": func(fieldName string) bool {
			autoFields := []string{
				"id",
				"creator",
				"create_time",
				"operator",
				"update_time",
			}
			for _, v := range autoFields {
				if fieldName == v {
					return false
				}
			}
			return true
		},
	}

	tpl, err := template.New(serviceName + "-pyschema").Funcs(funcs).Parse(pySchemaTpl)
	if err != nil {
		log.Fatal(err)
	}

	target := filepath.Join(dir, strings.ToLower(serviceName)+"_schema.py")
	t, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(t, tplData)
}
