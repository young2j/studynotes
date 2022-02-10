/*
 * File: genTag.go
 * Created Date: 2022-02-09 11:52:43
 * Author: ysj
 * Description: 根据结构体字段补全json和bson标签
 */
package main

import (
	"fmt"
	"go/ast"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const modelTpl = `
type {{.ModelName}} struct {
{{range $field := .Fields}}  {{$field.Name}} {{$field.Type}} {{$field.Tag}}
{{end}}
}
`

type Field struct {
	Name        string
	Type        string
	Tag         string
	Description string
}

type ModelTplData struct {
	ModelName string
	Fields    []*Field
}

// 根据结构体字段快速生成json和bson标签
func genFieldTag(modelName string, astFields []*ast.Field) {
	tpl, err := template.New(modelName).Parse(modelTpl)
	if err != nil {
		log.Fatal(err)
	}
	tplData := &ModelTplData{
		ModelName: modelName,
		Fields:    make([]*Field, 0),
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
			field.Tag = fmt.Sprintf("`json:\"%s\" bson:\"%s\" description:\"%s\"`", lowerCamelFieldName, lowerCamelFieldName, description)
		case *ast.SelectorExpr:
			expr, ok := astField.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}
			field.Type = expr.X.(*ast.Ident).Name + "." + expr.Sel.Name
			if lowerCamelFieldName == "id" {
				field.Tag = fmt.Sprintf("`json:\"%s\" bson:\"%s\" description:\"%s\"`", lowerCamelFieldName, "_id,omitempty", description)
			} else {
				field.Tag = fmt.Sprintf("`json:\"%s\" bson:\"%s\" description:\"%s\"`", lowerCamelFieldName, lowerCamelFieldName, description)
			}
		case *ast.ArrayType:
			elemType := astField.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
			field.Type = "[]" + elemType
			field.Tag = fmt.Sprintf("`json:\"%s\" bson:\"%s\" description:\"%s\"`", lowerCamelFieldName, lowerCamelFieldName, description)
		}
		tplData.Fields = append(tplData.Fields, field)
	}
	tpl.Execute(os.Stdout, tplData)
}
