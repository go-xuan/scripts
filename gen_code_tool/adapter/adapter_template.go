package adapter

import (
	"gen_code_tool/common"
	"gen_code_tool/db"
)

type TemplateAdapter struct {
	App    string         // 应用名
	Root   string         // 代码生成路径
	Tables []*db.Table    // 表结构
	Tmpls  []*common.Tmpl // 模板
}

func NewTemplateAdapter(conf *common.Config, frame string) *TemplateAdapter {
	var adapter = &TemplateAdapter{App: conf.App, Root: conf.Root()}
	var err error
	if adapter.Tables, err = db.QueryTables(conf.Database, conf.GetTableNames()); err != nil {
		panic(err)
	}
	switch frame {
	case common.GoQuanx:
		adapter.Tmpls = common.GoQuanxTemplates()
	case common.GoFrame:
		adapter.Tmpls = common.GoFrameTemplates()
	case common.SpringBoot:
		adapter.Tmpls = common.SpringBootTemplates()
	}
	return adapter
}

func (a *TemplateAdapter) GenCode() (err error) {
	if len(a.Tmpls) > 0 {
		for _, tmpl := range a.Tmpls {
			switch tmpl.DataType {
			case common.NoData:
				if err = tmpl.WriteCodeToFile(a.Root, nil); err != nil {
					return
				}
			case common.TableData:
				for _, table := range a.Tables {
					if err = tmpl.WriteCodeToFile(a.Root, table, table.Name); err != nil {
						return
					}
				}
			case common.AdapterData:
				if err = tmpl.WriteCodeToFile(a.Root, a); err != nil {
					return
				}
			}
		}
	}
	return
}
