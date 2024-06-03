package common

import (
	"bytes"
	"embed"
	"os"
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
	Frame     string           // 框架名
	Path      string           // 模板名称
	NeedTable string           // 模板用表（0不用表/1生成文件需要一张表/2生成文件需要所有表）
	FuncMap   template.FuncMap // 自定义模板参数方法
}

func (t *Tmpl) OutputPath(table ...string) string {
	path := strings.TrimSuffix(t.Path, ".tmpl")
	split := strings.Split(path, string(os.PathSeparator))
	if split[0] == t.Frame {
		split = split[1:]
	}
	if len(table) > 0 {
		i := len(split) - 1
		split[i] = strings.Replace(split[i], "{{table}}", table[0], -1)
	}
	return filepath.Join(split...)
}

func (t *Tmpl) GenCode(tables []*db.Table) (err error) {
	// 初始化template
	tt := template.Must(template.New(t.Path).Funcs(t.FuncMap).ParseFS(TemplateFs, t.Path))
	switch t.NeedTable {
	case NoTable:
		if err = GenCodeByTemplate(tt, nil, t.OutputPath()); err != nil {
			return
		}
	case OneTable:
		for _, table := range tables {
			if err = GenCodeByTemplate(tt, table, t.OutputPath(table.Name)); err != nil {
				return
			}
		}
	case AllTable:
		var names []string
		for _, table := range tables {
			names = append(names, table.Name)
		}
		var allTables = struct{ Names []string }{Names: names}
		if err = GenCodeByTemplate(tt, allTables, t.OutputPath()); err != nil {
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
	var frame = GoQuanx
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "main.go.tmpl"), NeedTable: AllTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "go.mod.tmpl"), NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "common", "consts.go.tmpl"), NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "conf", "config.yaml.tmpl"), NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "router", "router.go.tmpl"), NeedTable: AllTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "controller", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "logic", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "dao", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "entity", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	return tmpls
}

func GoFrameTemplates() []*Tmpl {
	var frame = GoFrame
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "main.go.tmpl"), NeedTable: AllTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "go.mod.tmpl"), NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "consts", "consts.go.tmpl"), NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "manifest", "config", "global_conf", "nacos.yaml.tmpl"), NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "router", "router.go.tmpl"), NeedTable: AllTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "controller", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "logic", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "dao", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "do", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "entity", "{{table}}.go.tmpl"), NeedTable: OneTable, FuncMap: funcMap})
	return tmpls
}

func SpringBootTemplates() []*Tmpl {
	var frame = SpringBoot
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/main.go.tmpl", NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/common/consts.go.tmpl", NeedTable: NoTable})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/controller/{{table}}.go.tmpl", NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/logic/{{table}}.go.tmpl", NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/dao/{{table}}.go.tmpl", NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/model/{{table}}.go.tmpl", NeedTable: OneTable, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/model/entity/{{table}}.go.tmpl", NeedTable: OneTable, FuncMap: funcMap})
	return tmpls
}
