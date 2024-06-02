package adapter

import (
	"embed"
	"gen_code_tool/common"
	"gen_code_tool/db"
)

type ReplaceAdapter struct {
	App    string      // 应用名
	Root   string      // 代码生成跟路径
	Tables []*db.Table // 表结构
}

//go:embed resource/replace/*
var ReplaceFs embed.FS

func NewReplaceAdapter(conf *common.Config, frame string, files []string) *ReplaceAdapter {
	if len(files) > 0 {
		var adapter = &ReplaceAdapter{App: conf.App, Root: conf.Root()}
		var err error
		if adapter.Tables, err = db.QueryTables(conf.Database, conf.GetTableNames()); err != nil {
			panic(err)
		}
		return adapter
	}
	return nil
}

func (a *ReplaceAdapter) GenCommon() error {

	return nil
}

func (a *ReplaceAdapter) GenController() error {
	return nil
}

func (a *ReplaceAdapter) GenLogic() error {
	return nil
}

func (a *ReplaceAdapter) GenModel() error {
	return nil
}

func (a *ReplaceAdapter) GenDao() error {
	return nil
}

func (a *ReplaceAdapter) GenConfig() error {
	return nil
}
