package adapter

type Adapter interface {
	GenCommon()     // 生成通用代码
	GenController() // 生成Controller代码
	GenLogic()      // 生成Logic代码
	GenModel()      // 生成Model代码
	GenDao()        // 生成Dao代码
}
