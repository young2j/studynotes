/*
 * File: parseModel.go
 * Created Date: 2022-02-09 11:55:33
 * Author: ysj
 * Description: 解析模型定义
 */

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func parseModelDefinition(fileName, modelName string) []*ast.Field {
	fset := token.NewFileSet()
	astf, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, dec := range astf.Decls {
		gdec, ok := dec.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gdec.Specs {
			tspec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if tspec.Name.String() != modelName {
				continue
			}
			st, ok := tspec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			return st.Fields.List
		}
	}
	return []*ast.Field{}
}
