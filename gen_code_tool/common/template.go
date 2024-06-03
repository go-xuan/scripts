package common

import (
	"bytes"
	"embed"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-xuan/quanx/os/filex"
	"github.com/go-xuan/quanx/types/stringx"

	"gen_code_tool/db"
)

//go:embed resource/template/*
var TemplateFs embed.FS

type Tmpls []*Tmpl

type Tmpl struct {
	Path      string           // 模板名称
	UseTables int              // 使用表数量（0/1/2）
	FuncMap   template.FuncMap // 自定义模板参数方法
}

func (t *Tmpl) OutputPath(table ...string) string {
	path := strings.TrimSuffix(t.Path, ".tmpl")
	split := strings.Split(path, "/")
	if l := len(split); l > 1 {
		path = filepath.Join(split...)
	}
	if len(table) > 0 {
		path = strings.Replace(path, "{{table}}", table[0], -1)
	}
	return path
}

func (t *Tmpl) GenCode(tables []*db.Table) (err error) {
	tt := template.Must(template.New(t.Path).Funcs(t.FuncMap).ParseFS(TemplateFs, t.Path))
	switch t.UseTables {
	case 1:
		for _, table := range tables {
			if err = GenCodeByTemplate(tt, table, t.OutputPath(table.Name)); err != nil {
				return
			}
		}
	case 2:
		var names []string
		for _, table := range tables {
			names = append(names, table.Name)
		}
		var allTables = struct{ Names []string }{Names: names}
		if err = GenCodeByTemplate(tt, allTables, t.OutputPath()); err != nil {
			return
		}
	default:
		if err = GenCodeByTemplate(tt, nil, t.OutputPath()); err != nil {
			return
		}
	}
	return
}

func GenCodeByTemplate(template *template.Template, data any, path string) (err error) {
	var buf = &bytes.Buffer{}
	if err = template.Execute(buf, data); err != nil {
		return
	}
	if err = filex.WriteFile(path, buf.String()); err != nil {
		return
	}
	return
}

func GoQuanxTemplates() []*Tmpl {
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/main.go.tmpl", UseTables: 2})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/main.go.tmpl", UseTables: 2})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/go.mod.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/common/consts.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/conf/config.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/internal/router/router.go.tmpl", UseTables: 2, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/internal/controller/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/internal/logic/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/internal/dao/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/internal/model/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-quanx/internal/model/entity/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	return tmpls
}

func GoFrameTemplates() []*Tmpl {
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/main.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/go.mod.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/consts/consts.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/conf/config.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/router/router.go.tmpl", UseTables: 2, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/controller/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/logic/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/dao/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/model/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "go-frame/model/entity/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	return tmpls
}

func SpringBootTemplates() []*Tmpl {
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/main.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/common/consts.go.tmpl", UseTables: 0})
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/controller/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/logic/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/dao/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/model/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Path: "spring-boot/model/entity/{{table}}.go.tmpl", UseTables: 1, FuncMap: funcMap})
	return tmpls
}
