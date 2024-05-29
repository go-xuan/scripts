package adapter

import "gen_code_tool/db"

type GoFrameAdapter struct {
	App    string      // 应用名
	Output string      // 代码生成路径
	Tables []*db.Table // 表结构
}

func (a *GoFrameAdapter) GenCommon() {

}
func (a *GoFrameAdapter) GenController() {

}

func (a *GoFrameAdapter) GenLogic() {

}

func (a *GoFrameAdapter) GenModel() {

}

func (a *GoFrameAdapter) GenDao() {

}
