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

func NewReplaceAdapter(conf *common.Config, frame string) *ReplaceAdapter {
	var adapter = &ReplaceAdapter{App: conf.App, Root: conf.Root()}
	return adapter
}

func (a *ReplaceAdapter) GenCode() error {

	return nil
}
