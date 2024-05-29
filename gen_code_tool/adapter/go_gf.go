package adapter

import "gen_code_tool/db"

type GoQuanxAdapter struct {
	App    string      // 应用名
	Output string      // 代码生成路径
	Tables []*db.Table // 表结构
}

func (a *GoQuanxAdapter) GenCommon() {

}
func (a *GoQuanxAdapter) GenController() {

}

func (a *GoQuanxAdapter) GenLogic() {

}

func (a *GoQuanxAdapter) GenModel() {

}

func (a *GoQuanxAdapter) GenDao() {

}
