package adapter

import (
	"embed"
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

func NewTemplateAdapter(conf *common.Config, frame string, tmpls []string) *TemplateAdapter {
	if len(tmpls) > 0 {
		var adapter = &TemplateAdapter{App: conf.App, Root: conf.Root}
		var err error
		if adapter.Tables, err = db.GetTables(conf.Database, conf.GetTableNames()); err != nil {
			panic(err)
		}
		var funcMap = template.FuncMap{
			"UC":    stringx.ToUpperCamel,
			"LC":    stringx.ToLowerCamel,
			"DB2Go": common.DB2GoType,
		}
		var templates = make(map[string]*template.Template)
		for _, tmpl := range tmpls {
			path := filepath.Join(frame, tmpl)
			templates[tmpl] = template.Must(template.New(path).Funcs(funcMap).ParseFS(TemplateFs, path))
		}
		adapter.Templates = templates
		return adapter
	}
	return nil
}

func (a *TemplateAdapter) GenCommon() {

}

func (a *TemplateAdapter) GenController() {

}

func (a *TemplateAdapter) GenLogic() {

}

func (a *TemplateAdapter) GenModel() {

}

func (a *TemplateAdapter) GenDao() {

}
