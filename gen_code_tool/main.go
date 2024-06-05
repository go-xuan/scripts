package main

import (
	"flag"

	"github.com/go-xuan/quanx/utils/marshalx"
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

	// 代码生成
	if err = config.Generator().GenCode(); err != nil {
		panic(err)
	}
}
