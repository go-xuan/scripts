package gen

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/go-xuan/quanx/os/file/filex"
	"github.com/go-xuan/quanx/types/stringx"

	"gen_code_tool/common"
	"gen_code_tool/db"
)

type Template struct {
	Frame     string
	Multiplex []string
	Tables    []*db.Table // 表结构
}

func (t *Template) GenController() {
	tmpl, err := template.New("controller").Funcs(template.FuncMap{
		"UC":    stringx.ToUpperCamel,
		"LC":    stringx.ToLowerCamel,
		"DB2Go": common.DB2Go,
	}).ParseFS(common.Templates, filepath.Join("resource/template", t.Frame, "controller.tmpl"))
	if err != nil {
		panic(err)
	}

	for _, table := range t.Tables {
		// 执行模板并将结果写入buffer
		var buf = &bytes.Buffer{}
		if err = tmpl.Execute(buf, table); err != nil {
			panic(err)
		}
		if err = filex.WriteFile("template/entity/quanchao.go", buf.String()); err != nil {
			panic(err)
		}
	}
}
