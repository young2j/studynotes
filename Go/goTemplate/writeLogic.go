/*
 * File: writeLogic.go
 * Created Date: 2022-03-15 04:24:59
 * Author: ysj
 * Description:
 */

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

func writeLogic(protoPath, rpcName, action string, buf *bytes.Buffer) {
	dir, protoFile := filepath.Split(protoPath)
	logicFile := fmt.Sprintf("internal/logic/%sLogic.go", strcase.ToLowerCamel(rpcName))
	logicFile = filepath.Join(dir, logicFile)

	fset := token.NewFileSet()
	astf, err := parser.ParseFile(fset, logicFile, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}
	// delete todo comment
	todoIdx := 0
	for i, comment := range astf.Comments {
		if comment.Text() == "todo: add your logic here and delete this line\n" {
			todoIdx = i
			break
		}
	}
	if len(astf.Comments) > 0 {
		astf.Comments = append(
			astf.Comments[0:todoIdx],
			astf.Comments[todoIdx:len(astf.Comments)-1]...,
		)
	}

	for _, dec := range astf.Decls {
		switch decl := dec.(type) {
		case *ast.GenDecl:
			originImspecs := make([]ast.Spec, 0)
			imspecs := getImportSpec(action, protoFile)
			isImportSpec := false
			for _, spec := range decl.Specs {
				_, isImportSpec = spec.(*ast.ImportSpec)
				if isImportSpec && !containsSameImportSpec(imspecs, spec) {
					originImspecs = append(originImspecs, spec)
				}
			}
			if isImportSpec {
				decl.Specs = append(imspecs, originImspecs...)
			}

		case *ast.FuncDecl:
			if decl.Name.String() == rpcName {
				bufAst, err := parser.ParseFile(fset, "buf", buf.Bytes(), parser.ParseComments)
				if err != nil {
					log.Fatalln(err)
					continue
				}
				decl.Body.List = bufAst.Decls[0].(*ast.FuncDecl).Body.List
			}
		}
	}
	logicFileCode := make([]byte, 0)
	logicBuf := bytes.NewBuffer(logicFileCode)
	printer.Fprint(logicBuf, fset, astf)
	ioutil.WriteFile(logicFile, logicBuf.Bytes(), 0644)
	exec.Command("gofmt", "-l", "-w", logicFile).Run()
	fmt.Printf("complete %s\n", logicFile)
}

func getImportSpec(action, protoFile string) []ast.Spec {
	ret := make([]ast.Spec, 0)
	serviceName := strings.TrimSuffix(protoFile, ".proto") //access.proto
	switch action {
	case "Upsert":
		imports := []string{
			"\"errors\"",
			"\"net/http\"",
			"\"github.com/globalsign/mgo/bson\"",
			"\"scana/common/utils\"",
			fmt.Sprintf("\"scana/services/%s/model\"", serviceName),
			"\"scana/common/gocopy\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	case "Add":
		imports := []string{
			"\"net/http\"",
			"\"scana/common/utils\"",
			fmt.Sprintf("\"scana/services/%s/model\"", serviceName),
			"\"scana/common/gocopy\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	case "Change":
		imports := []string{
			"\"net/http\"",
			"\"scana/common/utils\"",
			"\"github.com/globalsign/mgo/bson\"",
			"\"scana/common/gocopy\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	case "Delete":
		imports := []string{
			"\"net/http\"",
			"\"github.com/globalsign/mgo/bson\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	case "Get":
		imports := []string{
			"\"net/http\"",
			"\"scana/common/utils\"",
			fmt.Sprintf("\"scana/services/%s/model\"", serviceName),
			"\"github.com/globalsign/mgo/bson\"",
			"\"scana/common/gocopy\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	case "List":
		imports := []string{
			"\"net/http\"",
			"\"scana/common/utils\"",
			"\"github.com/globalsign/mgo/bson\"",
			"\"scana/common/gocopy\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	case "Count":
		imports := []string{
			"\"net/http\"",
			"\"github.com/globalsign/mgo/bson\"",
		}
		for _, v := range imports {
			imp := &ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: v,
				},
			}
			ret = append(ret, imp)
		}
	}
	return ret
}

func containsSameImportSpec(specs []ast.Spec, spec ast.Spec) bool {
	imspec, ok := spec.(*ast.ImportSpec)
	if !ok {
		return false
	}

	for _, v := range specs {
		imv, ok := v.(*ast.ImportSpec)
		if !ok {
			continue
		}
		if imv.Path.Value == imspec.Path.Value {
			return true
		}
	}
	return false
}
