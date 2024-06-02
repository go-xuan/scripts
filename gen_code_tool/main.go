package main

import (
	"flag"

	"github.com/go-xuan/quanx/utils/marshalx"

	"gen_code_tool/common"
)

func main() {
	var configIn = flag.String("config", "config.yaml", "输入配置文件，例如：-config=config.yaml")
	// 解析输入的参数
	flag.Parse()

	// 读取配置文件
	var err error
	if err = marshalx.UnmarshalFromFile(*configIn, common.Conf); err != nil {
		panic(err)
	}

	// 初始化应用数据库连接
	if common.Conf.Database != nil {
		common.Conf.Database.Enable = true
		if err = common.Conf.Database.Run(); err != nil {
			panic(err)
		}
	} else {
		panic("please check the [db] in the config.yaml")
	}

	// 代码生成
	if err = GenCode(); err != nil {
		panic(err)
	}
}

func GenCode() (err error) {
	// 1、查询应用数据库，获取需要生成代码的表结构
	//var tables = make(map[string]*model.Table)
	//if tables, err = GetTables(); err != nil {
	//	return
	//}
	//for _, table := range tables {
	//
	//}
	// 2、获取代码模板
	// 3、根据代码模板生成代码
	// 4、输出到指定目录
	return
}