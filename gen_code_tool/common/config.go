package common

import (
	"embed"
	"github.com/go-xuan/quanx/db/gormx"
	"strings"
)

//go:embed resource/*
var Templates embed.FS

var Conf = &Config{}

type Config struct {
	App      string          `json:"app" json:"app" default:"demo"`         // 应用名
	Language string          `json:"language" yaml:"language" default:"go"` // 生成代码语言
	Output   string          `json:"output" yaml:"output" default:"./"`     // 代码生成路径
	Tables   string          `json:"tables" yaml:"tables"`                  // 生成代码代码的表名（为空则获取全表）
	Database *gormx.Database `json:"db" yaml:"db"`                          // 应用数据库
}

type Generate struct {
	App      string `json:"app" json:"app" default:"demo"`         // 应用名
	Language string `json:"language" yaml:"language" default:"go"` // 生成代码语言
	Output   string `json:"output" yaml:"output" default:"./"`     // 代码生成路径
	Tables   string `json:"tables" yaml:"tables"`                  // 生成代码代码的表名（为空则获取全表）
}

type Module struct {
	Name   string `json:"name" yaml:"name"`     // 模块名
	Tables string `json:"tables" yaml:"tables"` // 模块所属表名
}

func (c *Config) GetTableNames() []string {
	var tables []string
	if c.Tables != "" {
		tables = append(tables, strings.Split(c.Tables, ",")...)
	}
	return tables
}
