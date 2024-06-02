package adapter

import (
	"bytes"
	"embed"
	"github.com/go-xuan/quanx/os/filex"
	"path/filepath"
	"text/template"

	"github.com/go-xuan/quanx/types/stringx"

	"gen_code_tool/common"
	"gen_code_tool/db"
)

type TemplateAdapter struct {
	App       string                        // 应用名
	Root      string                        // 代码生成路径
	Tables    []*db.Table                   // 表结构
	Templates map[string]*template.Template // 模板
}

//go:embed resource/template/*
var TemplateFs embed.FS

func NewTemplateAdapter(conf *common.Config, frame string, tmplNames []string) *TemplateAdapter {
	if len(tmplNames) > 0 {
		var adapter = &TemplateAdapter{App: conf.App, Root: conf.Root()}
		var err error
		if adapter.Tables, err = db.QueryTables(conf.Database, conf.GetTableNames()); err != nil {
			panic(err)
		}
		var funcMap = template.FuncMap{
			"UC":  stringx.ToUpperCamel,
			"LC":  stringx.ToLowerCamel,
			"DTG": common.DB2GoType,
		}
		var templates = make(map[string]*template.Template)
		for _, tmplName := range tmplNames {
			path := filepath.Join(frame, tmplName)
			templates[tmplName] = template.Must(template.New(path).Funcs(funcMap).ParseFS(TemplateFs, path))
		}
		adapter.Templates = templates
		return adapter
	}
	return nil
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

func (a *TemplateAdapter) GenCommon() (err error) {
	if tmpl, ok := a.Templates[common.ConstsTmpl]; ok {
		if err = GenCodeByTemplate(tmpl, nil, filepath.Join(a.Root, "common", "consts.go")); err != nil {
			return
		}
	}
	return
}

func (a *TemplateAdapter) GenController() (err error) {
	if tmpl, ok := a.Templates[common.ControllerTmpl]; ok {
		for _, table := range a.Tables {
			if err = GenCodeByTemplate(tmpl, table, filepath.Join(a.Root, "controller", table.Name+".go")); err != nil {
				return
			}
		}
	}
	return

}

func (a *TemplateAdapter) GenLogic() (err error) {
	if tmpl, ok := a.Templates[common.LogicTmpl]; ok {
		for _, table := range a.Tables {
			if err = GenCodeByTemplate(tmpl, table, filepath.Join(a.Root, "logic", table.Name+".go")); err != nil {
				return
			}
		}
	}
	return

}

func (a *TemplateAdapter) GenModel() (err error) {
	if tmpl, ok := a.Templates[common.ModelTmpl]; ok {
		for _, table := range a.Tables {
			if err = GenCodeByTemplate(tmpl, table, filepath.Join(a.Root, "model", table.Name+".go")); err != nil {
				return
			}
		}
	}
	return
}

func (a *TemplateAdapter) GenDao() (err error) {
	if tmpl, ok := a.Templates[common.DaoTmpl]; ok {
		for _, table := range a.Tables {
			if err = GenCodeByTemplate(tmpl, table, filepath.Join(a.Root, "dao", table.Name+".go")); err != nil {
				return
			}
		}
	}
	return
}

func (a *TemplateAdapter) GenConfig() (err error) {
	if tmpl, ok := a.Templates[common.ConfigTmpl]; ok {
		if err = GenCodeByTemplate(tmpl, nil, filepath.Join(a.Root, "conf", "conf.yaml")); err != nil {
			return
		}
	}
	return
}
