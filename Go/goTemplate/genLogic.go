/*
 * File: completeLogic.go
 * Created Date: 2022-03-15 10:23:37
 * Author: ysj
 * Description:  快速生成rpc logic部分的代码
 */

package main

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	parser "go-template/protoparser"

	"github.com/iancoleman/strcase"
)

const (
	upsertTpl = `
	if len(in.Query) <= 0 {
		return nil, errors.New("请指定upsert的约束条件")
	}

	var err error
	data := &model.{{.ModelName}}{}
	gocopy.CopyWithOption(data, in.Data, &gocopy.Option{
		IgnoreFields: []string{"Id"},
		Converters: utils.CONVERTER_TO_COMPLEX,
	})

	q := bson.M{}
	err = bson.Unmarshal(in.Query, &q)
	if err != nil {
		return nil, err
	}
	changeInfo, err := l.svcCtx.{{.ModelName}}Model.Upsert(l.ctx, q, data)
	if err != nil {
		return nil, err
	}

	id := ""
	objectId, ok := changeInfo.UpsertedId.(bson.ObjectId)
	if ok {
		id = objectId.Hex()
	}

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code: http.StatusOK,
		Msg:  "操作成功",
		Data: id,
	}, nil
	`
	addTpl = `	
	data := make([]interface{}, 0)
	for _, v := range in.{{.TableName}} {
		doc := &model.{{ .ModelName }}{}
		gocopy.CopyWithOption(doc, v, &gocopy.Option{
			Converters: utils.CONVERTER_TO_COMPLEX,
		})
		data = append(data, doc)
	}

	oids, err := l.svcCtx.{{ .ModelName }}Model.Insert(l.ctx, data...)
	if err != nil {
		return nil, err
	}
	oidsHex := make([]string, len(oids))
	for i, oid := range oids {
		oidsHex[i] = oid.Hex()
	}

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code: http.StatusOK,
		Msg:  "添加成功",
		Data: oidsHex,
	}, nil
`
	changeTpl = `
	var err error
	update := bson.M{}
	if in.{{.TableName}} != nil {
		data := bson.M{}
		gocopy.CopyWithOption(&data, in.{{.TableName}}, &gocopy.Option{
			IgnoreZero:   true,
			Converters:   utils.CONVERTER_TO_COMPLEX,
			NameFromTo:   map[string]string{"Id": "_id"},
		})
		update["$set"] = data
	}

	if len(in.Update) > 0 {
		data := bson.M{}
		err = bson.Unmarshal(in.Update, &data)
		if err != nil {
			return nil, err
		}
		gocopy.CopyWithOption(&update, data, &gocopy.Option{
			Append: true,
		})
	}

	if in.Id != "" {
		err = l.svcCtx.{{.ModelName}}Model.UpdateOneId(l.ctx, in.Id, update)
		if err != nil {
			return nil, err
		}
	} else if len(in.Query) > 0 {
		q := bson.M{}
		err = bson.Unmarshal(in.Query, &q)
		if err != nil {
			return nil, err
		}
		_, err = l.svcCtx.{{.ModelName}}Model.UpdateAll(l.ctx, q, update)
		if err != nil {
			return nil, err
		}
	}

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code: http.StatusOK,
		Msg:  "修改成功",
	}, nil
`
	deleteTpl = `
	if in.Id != "" {
		err := l.svcCtx.{{ .ModelName }}Model.RemoveId(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
	} else if len(in.Query) > 0{
		q := bson.M{}
		if err := bson.Unmarshal(in.Query, &q); err != nil {
			return nil, err
		}
		_, err := l.svcCtx.{{ .ModelName }}Model.RemoveAll(l.ctx, q)
		if err != nil {
			return nil, err
		}
	}

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil
`
	getTpl = `
	var (
		{{LowerCamel .TableName}} *model.{{ .ModelName }}
		err       error
	)
	if in.Id != "" {
		{{LowerCamel .TableName}}, err = l.svcCtx.{{ .ModelName }}Model.FindOneId(l.ctx, in.Id, in.Nocache)
		if err != nil {
			return nil, err
		}
	} else {
		q := bson.M{}
		if err := bson.Unmarshal(in.Query, &q); err != nil {
			return nil, err
		}
		{{LowerCamel .TableName}}, err = l.svcCtx.{{ .ModelName }}Model.FindOne(l.ctx, q, in.Nocache)
		if err != nil {
			return nil, err
		}
	}

	data := &{{ .PbPackage }}.{{ .TableName }}Info{}
	gocopy.CopyWithOption(data, {{LowerCamel .TableName}}, &gocopy.Option{
		NameFromTo: map[string]string{"_id": "Id"},
		Converters: utils.CONVERTER_TO_STRING,
	})

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: data,
	}, nil
`
	listTpl = `
	opts := utils.GetQueryOption(&utils.QueryOption{
		Page:          in.Page,
		PageSize:      in.PageSize,
		SortKeys:      in.SortKeys,
		ProjectFields: in.ProjectFields,
		ExcludeFields: in.ExcludeFields,
	})
	q := bson.M{}
	if len(in.Query) > 0 {
		err := bson.Unmarshal(in.Query, &q)
		if err != nil {
			return nil, err
		}
	}
	
	count, err := l.svcCtx.{{ .ModelName }}Model.Count(l.ctx, q)
	if err != nil {
		return nil, err
	}
	{{LowerCamel .TableName}}, err := l.svcCtx.{{ .ModelName }}Model.FindAll(l.ctx, q, opts)
	if err != nil {
		return nil, err
	}

	data := []*{{ .PbPackage }}.{{ .TableName }}Info{}
	gocopy.CopyWithOption(&data, {{LowerCamel .TableName}}, &gocopy.Option{
		NameFromTo: map[string]string{"_id": "Id"},
		Converters: utils.CONVERTER_TO_STRING,
	})

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code:      http.StatusOK,
		Msg:       "查询成功",
		Data:      data,
		Count:     count,
		TotalPage: utils.GetTotalPage(count, in.PageSize),
	}, nil
	`
	countTpl = `
	q := bson.M{}
	err := bson.Unmarshal(in.Query, &q)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.{{ .ModelName }}Model.Count(l.ctx, q)
	if err != nil {
		return nil, err
	}

	return &{{ .PbPackage }}.{{ .ReturnsType }}{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: count,
	}, nil
	
	`
)

type LogicTplData struct {
	PbPackage   string // eg. accesspb
	ModelName   string // eg. AccessResourcePerms
	TableName   string // eg. ResourcePerms
	ReturnsType string // eg. AddResourcePermsResp
}

// 将模版生成的代码存入buffer
func getTplBuf(tpl *template.Template, tplData *LogicTplData) *bytes.Buffer {
	stmt := make([]byte, 0)
	buf := bytes.NewBuffer(stmt)
	buf.WriteString("package main\n")
	buf.WriteString("func main(){\n")
	err := tpl.Execute(buf, tplData)
	if err != nil {
		log.Panicln("getTplBuf:", err)
		return buf
	}
	buf.WriteString("}")
	return buf
}

func genLogic(inputFile string) {
	protoPath, err := filepath.Abs(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("protoPath: %v\n", protoPath)

	protoParser := parser.NewDefaultProtoParser()
	p, err := protoParser.Parse(protoPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, rpc := range p.Service.RPC {
		if strings.HasPrefix(rpc.Name, "Upsert") {
			action := "Upsert"
			table := strings.TrimPrefix(rpc.Name, action)
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}
			tpl, err := template.New(rpc.Name).Parse(upsertTpl)
			if err != nil {
				log.Println(err)
				continue
			}
			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)

		} else if strings.HasPrefix(rpc.Name, "Add") {
			action := "Add"
			table := strings.TrimPrefix(rpc.Name, action)
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}
			tpl, err := template.New(rpc.Name).Parse(addTpl)
			if err != nil {
				log.Println(err)
				continue
			}
			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)

		} else if strings.HasPrefix(rpc.Name, "Change") {
			action := "Change"
			table := strings.TrimPrefix(rpc.Name, action)
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}
			tpl, err := template.New(rpc.Name).Parse(changeTpl)
			if err != nil {
				log.Println(err)
				continue
			}

			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)

		} else if strings.HasPrefix(rpc.Name, "Delete") {
			action := "Delete"
			table := strings.TrimPrefix(rpc.Name, action)
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}
			tpl, err := template.New(rpc.Name).Parse(deleteTpl)
			if err != nil {
				log.Println(err)
				continue
			}

			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)

		} else if strings.HasPrefix(rpc.Name, "Get") {
			action := "Get"
			table := strings.TrimPrefix(rpc.Name, action)
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}
			funcs := template.FuncMap{
				"LowerCamel": strcase.ToLowerCamel,
			}

			tpl, err := template.New(rpc.Name).Funcs(funcs).Parse(getTpl)
			if err != nil {
				log.Println(err)
				continue
			}

			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)

		} else if strings.HasPrefix(rpc.Name, "List") {
			action := "List"
			table := strings.TrimPrefix(rpc.Name, action)
			table = strings.TrimSuffix(table, "s")
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}
			funcs := template.FuncMap{
				"LowerCamel": strcase.ToLowerCamel,
			}

			tpl, err := template.New(rpc.Name).Funcs(funcs).Parse(listTpl)
			if err != nil {
				log.Println(err)
				continue
			}
			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)

		} else if strings.HasPrefix(rpc.Name, "Count") {
			action := "Count"
			table := strings.TrimPrefix(rpc.Name, action)
			table = strings.TrimSuffix(table, "s")
			tplData := &LogicTplData{
				PbPackage:   p.PbPackage,
				ModelName:   parser.CamelCase(p.Package.Name) + table,
				TableName:   table,
				ReturnsType: rpc.ReturnsType,
			}

			tpl, err := template.New(rpc.Name).Parse(countTpl)
			if err != nil {
				log.Println(err)
				continue
			}
			buf := getTplBuf(tpl, tplData)
			writeLogic(protoPath, rpc.Name, action, buf)
		}
	}
}
