package common

import (
	"path/filepath"
	"strings"

	"gen_code_tool/adapter"

	"github.com/go-xuan/quanx/server/gormx"
)

var Conf = &Config{}

type Config struct {
	App      string          `json:"app" json:"app" default:"demo"`             // 应用名
	Frame    string          `json:"frame" yaml:"frame" default:"go-quanx"`     // 代码框架（go-quanx/go-frame/spring-boot）
	Adapter  string          `json:"adapter" yaml:"adapter" default:"template"` // 代码生成适配器（replace/template）
	Output   string          `json:"output" yaml:"output" default:"./"`         // 输出路径
	Tables   string          `json:"tables" yaml:"tables"`                      // 生成代码代码的表名（为空则获取全表）
	Database *gormx.Database `json:"db" yaml:"db"`                              // 应用数据库
}

type Generate struct {
	App      string `json:"app" json:"app" default:"demo"`         // 应用名
	Language string `json:"language" yaml:"language" default:"go"` // 生成代码语言
	Output   string `json:"output" yaml:"output" default:"./"`     // 代码生成路径
	Tables   string `json:"tables" yaml:"tables"`                  // 生成代码代码的表名（为空则获取全表）
}

// 生成适配器
func (c *Config) NewAdapter() adapter.Adapter {
	switch c.Adapter {
	case Replace:
		return adapter.NewReplaceAdapter(c, c.Frame)
	case Template:
		return adapter.NewTemplateAdapter(c, c.Frame)
	default:
		return nil
	}
}

func (c *Config) GetTableNames() []string {
	var tables []string
	if c.Tables != "" {
		tables = append(tables, strings.Split(c.Tables, ",")...)
	}
	return tables
}

func (c *Config) Root() string {
	if strings.HasSuffix(c.Output, c.App) {
		return c.Output
	} else {
		return filepath.Join(c.Output, c.App)
	}
}
