package adapter

type Adapter interface {
	GenCommon() error     // 生成通用代码
	GenController() error // 生成Controller代码
	GenLogic() error      // 生成Logic代码
	GenModel() error      // 生成Model代码
	GenDao() error        // 生成Dao代码
	GenConfig() error
}

// 根据适配器生成代码
func GenCodes(adapter Adapter) (err error) {
	if err = adapter.GenCommon(); err != nil {
		return
	}
	if err = adapter.GenController(); err != nil {
		return
	}
	if err = adapter.GenLogic(); err != nil {
		return
	}
	if err = adapter.GenModel(); err != nil {
		return
	}
	if err = adapter.GenDao(); err != nil {
		return
	}
	if err = adapter.GenConfig(); err != nil {
		return
	}
	return
}
