package adapter

import (
	"github.com/go-xuan/quanx/db/gormx"

	"gen_code_tool/common"
	"gen_code_tool/db"
)

type Adapter interface {
	GenCommon()     // 生成通用代码
	GenController() // 生成Controller代码
	GenLogic()      // 生成Logic代码
	GenModel()      // 生成Model代码
	GenDao()        // 生成Dao代码
}

func NewAdapter(conf *common.Config) Adapter {
	// 查询所有表
	var err error
	var fields []*db.Field
	if database := conf.Database; database != nil {
		var tableNames = conf.GetTableNames()
		switch database.Type {
		case gormx.Mysql:
			if fields, err = db.MysqlTableFieldList(database.Source, database.Database, tableNames...); err != nil {
				panic(err)
			}
		case gormx.Postgres:
			if fields, err = db.PgTableFieldList(database.Source, database.Database, database.Schema, tableNames...); err != nil {
				panic(err)
			}
		}
	}
	var tables []*db.Table
	if len(fields) > 0 {
		var tableMap = make(map[string]*db.Table)
		for _, field := range fields {
			if _, ok := tableMap[field.TableName]; ok {
				tableMap[field.TableName].Fields = append(tableMap[field.TableName].Fields, field)
				if tableMap[field.TableName].FieldLen < len(field.Name) {
					tableMap[field.TableName].FieldLen = len(field.Name)
				}
			} else {
				tableMap[field.TableName] = &db.Table{
					Name:     field.TableName,
					Schema:   field.Schema,
					Comment:  field.TableComment,
					FieldLen: len(field.Name),
					Fields:   []*db.Field{field},
				}
			}
		}
		for _, table := range tableMap {
			tables = append(tables, table)
		}
	}
	switch conf.Language {
	case "go":
		return &GoQuanxAdapter{App: conf.App, Output: conf.Output, Tables: tables}
	}
	return nil
}
