package main

import (
	"flag"
	"fmt"
	"github.com/go-xuan/quanx/os/filex"

	"github.com/go-xuan/quanx/utils/sqlx"
)

func main() {
	var filePath = flag.String("file", "input.sql", "请输入需要格式化化的sql文件，例如：-file=input.sql")
	// 解析输入的参数
	flag.Parse()
	fmt.Println("file:", *filePath)

	var err error
	var bytes []byte
	if bytes, err = filex.ReadFile(*filePath); err != nil {
		panic(err)
	}
	// 将sql转为select对象
	//fmtSql := sqlx.SqlFormat(string(bytes))
	s := sqlx.Format(string(bytes))
	fmtSql := s.String()
	fmt.Println(fmtSql)
	// 覆盖原SQL
	if err = filex.WriteFile("output.sql", fmtSql); err != nil {
		panic(err)
	}
}
