package main

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/go-xuan/quanx/server/gormx"
	"github.com/go-xuan/quanx/utils/marshalx"

	"gen_code_tool/common"
	"gen_code_tool/db"
	"gen_code_tool/template"
)

var config = &Config{}

func main() {
	var configIn = flag.String("config", "config.yaml", "输入配置文件，例如：-config=config.yaml")
	// 解析输入的参数
	flag.Parse()

	// 读取配置文件
	var err error
	if err = marshalx.UnmarshalFromFile(*configIn, config); err != nil {
		panic(err)
	}

	// 初始化数据库连接
	if config.DB != nil {
		if db.DB, err = config.DB.NewGormDB(); err != nil {
			panic(err)
		}
	}

	// 代码生成
	if err = config.Generator().GenCode(); err != nil {
		panic(err)
	}
}

type Config struct {
	App    string          `json:"app" json:"app" default:"demo"`         // 应用名
	Frame  string          `json:"frame" yaml:"frame" default:"go-quanx"` // 代码框架（go-quanx/go-frame/spring-boot）
	Output string          `json:"output" yaml:"output"`                  // 输出路径
	Table  string          `json:"table" yaml:"table"`                    // 生成代码代码的表名（为空则获取全表,多个以“,”分割）
	DB     *gormx.Database `json:"db" yaml:"db"`                          // 应用数据库
}

func (c *Config) Tables() []string {
	var tables []string
	if c.Table != "" {
		tables = append(tables, strings.Split(c.Table, ",")...)
	}
	return tables
}

func (c *Config) Tmpls() []*template.Tmpl {
	switch config.Frame {
	case common.GoQuanx:
		return template.GoQuanxTemplates()
	case common.GoFrame:
		return template.GoFrameTemplates()
	case common.SpringBoot:
		return template.SpringBootTemplates()
	}
	return nil
}

func (c *Config) Root() string {
	if strings.HasSuffix(c.Output, c.App) {
		return c.Output
	} else {
		return filepath.Join(c.Output, c.App)
	}
}

func (c *Config) Generator() *Generator {
	generator := &Generator{App: config.App, Root: config.Root(), DB: config.DB, Tmpls: config.Tmpls()}
	// 查询数据表
	if tables, err := db.QueryTables(config.DB, config.Tables()); err == nil {
		generator.Tables = tables
	}
	return generator
}

type Generator struct {
	App    string           // 应用名
	Root   string           // 代码生成路径
	DB     *gormx.Database  // 应用数据库
	Tables []*db.Table      // 表结构
	Tmpls  []*template.Tmpl // 模板
}

func (gen *Generator) GenCode() (err error) {
	if len(gen.Tmpls) > 0 {
		for _, tmpl := range gen.Tmpls {
			switch tmpl.DataType {
			case common.NoData:
				if err = tmpl.WriteCodeToFile(gen.Root, nil); err != nil {
					return
				}
			case common.TableData:
				for _, table := range gen.Tables {
					if err = tmpl.WriteCodeToFile(gen.Root, table, table.Name); err != nil {
						return
					}
				}
			case common.AppData:
				if err = tmpl.WriteCodeToFile(gen.Root, gen); err != nil {
					return
				}
			}
		}
	}
	return
}
