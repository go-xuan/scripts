package adapter

import (
	"gen_code_tool/common"
	"gen_code_tool/db"
)

type ReplaceAdapter struct {
	App    string      // 应用名
	Root   string      // 代码生成跟路径
	Tables []*db.Table // 表结构
}

func NewReplaceAdapter(conf *common.Config, frame string, tmpls []string) *ReplaceAdapter {
	if len(tmpls) > 0 {
		var adapter = &ReplaceAdapter{App: conf.App, Root: conf.Root}
		var err error
		if adapter.Tables, err = db.GetTables(conf.Database, conf.GetTableNames()); err != nil {
			panic(err)
		}
		return adapter
	}
	return nil
}

func (a *ReplaceAdapter) GenCommon() {

}

func (a *ReplaceAdapter) GenController() {

}

func (a *ReplaceAdapter) GenLogic() {

}

func (a *ReplaceAdapter) GenModel() {

}

func (a *ReplaceAdapter) GenDao() {

}
