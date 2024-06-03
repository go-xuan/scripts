package adapter

import "gen_code_tool/common"

type Adapter interface {
	GenCode() error // 生成通用代码
}

// 生成适配器
func NewAdapter(c *common.Config) Adapter {
	switch c.Adapter {
	case common.Replace:
		return NewReplaceAdapter(c, c.Frame)
	case common.Template:
		return NewTemplateAdapter(c, c.Frame)
	default:
		return nil
	}
}
