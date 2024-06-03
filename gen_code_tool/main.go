package main

import (
	"embed"
	"flag"

	"github.com/go-xuan/quanx/utils/marshalx"

	"gen_code_tool/adapter"
	"gen_code_tool/common"
)

//go:embed template/*
var TemplateFs embed.FS

func main() {
	var configIn = flag.String("config", "config.yaml", "输入配置文件，例如：-config=config.yaml")
	// 解析输入的参数
	flag.Parse()

	// 读取配置文件
	var err error
	if err = marshalx.UnmarshalFromFile(*configIn, common.Conf); err != nil {
		panic(err)
	}
	common.TemplateFs = TemplateFs

	// 初始化应用数据库连接
	if common.Conf.DB != nil {
		common.Conf.DB.Enable = true
		if err = common.Conf.DB.Run(); err != nil {
			panic(err)
		}
	} else {
		panic("please check the [db] in the config.yaml")
	}

	// 代码生成
	if newAdapter := adapter.NewAdapter(common.Conf); newAdapter != nil {
		if err = newAdapter.GenCode(); err != nil {
			panic(err)
		}
	} else {
		panic("please check the [adapter] in the config.yaml")
	}
}
