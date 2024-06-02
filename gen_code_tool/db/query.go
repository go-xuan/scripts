package db

import (
	"strings"

	"github.com/go-xuan/quanx/server/gormx"
	log "github.com/sirupsen/logrus"
)

func GetTables(database *gormx.Database, tableNames []string) (tables []*Table, err error) {
	var fields []*Field
	if fields, err = GetFields(database, tableNames); err != nil {
		return
	}
	if len(fields) > 0 {
		var tableMap = make(map[string]*Table)
		for _, field := range fields {
			if _, ok := tableMap[field.TableName]; ok {
				tableMap[field.TableName].Fields = append(tableMap[field.TableName].Fields, field)
				if tableMap[field.TableName].FieldLen < len(field.Name) {
					tableMap[field.TableName].FieldLen = len(field.Name)
				}
			} else {
				tableMap[field.TableName] = &Table{
					Name:     field.TableName,
					Schema:   field.Schema,
					Comment:  field.TableComment,
					FieldLen: len(field.Name),
					Fields:   []*Field{field},
				}
			}
		}
		for _, table := range tableMap {
			tables = append(tables, table)
		}
	}
	return
}

func GetFields(database *gormx.Database, tableNames []string) (fields []*Field, err error) {
	switch database.Type {
	case gormx.Mysql:
		if fields, err = mysqlTableFields(database.Source, database.Database, tableNames...); err != nil {
			return
		}
	case gormx.Postgres:
		if fields, err = pgsqlTableFields(database.Source, database.Database, database.Schema, tableNames...); err != nil {
			return
		}
	}
	return
}

// 查询表字段列表
func mysqlTableFields(source, database string, table ...string) (fields []*Field, err error) {
	sql := strings.Builder{}
	sql.WriteString(`
select t1.column_name as "name",
       t1.table_name as "table_name",
       t1.table_schema as "database",
       t1.data_type as "type",
       if(t1.column_default is null, t1.extra, t1.column_default) as "default",
       t1.column_comment as "comment",
       t2.table_comment as "table_comment",
       if(t1.data_type = 'varchar', t1.character_maximum_length, t1.numeric_precision) as "precision",
       t1.numeric_scale as "scale",
       t1.is_nullable = 'YES' as "nullable"
  from information_schema.columns t1
  left join information_schema.tables t2
    on t1.table_schema = t2.table_schema
   and t1.table_name = t2.table_name`)
	sql.WriteString(` where t1.table_schema = '`)
	sql.WriteString(database)
	sql.WriteString(`'`)
	if len(table) > 0 {
		sql.WriteString(` and t1.table_name in ('`)
		sql.WriteString(strings.Join(table, `','`))
		sql.WriteString(`')`)
	}
	sql.WriteString(` order by t1.table_name, t1.ordinal_position`)
	if err = gormx.DB(source).Raw(sql.String()).Scan(&fields).Error; err != nil {
		log.Errorf("查询表字段列表 失败：%s\n ", err)
		return
	}
	return
}

// 查询表字段列表
func pgsqlTableFields(source, database, schema string, table ...string) (fields []*Field, err error) {
	sql := strings.Builder{}
	sql.WriteString(`
select t1.column_name as name,
       t1.table_name as table_name,
       t1.table_schema as schema,
       t1.table_catalog as database,
       t1.udt_name as type,
       t1.column_default as default,
       obj_description(t3.oid) as tableComment,
       t5.description as comment,
       case when t1.numeric_precision is null then t1.character_maximum_length else t1.numeric_precision end as precision,
       t1.numeric_scale as scale,
       t1.is_nullable = 'YES' as nullable
  from information_schema.columns t1
  left join pg_namespace t2 on t1.table_schema = t2.nspname
  left join pg_class t3 on t3.relname = t1.table_name and t3.relnamespace = t2.oid
  left join pg_attribute t4 on t4.attname = t1.column_name and t4.attrelid = t3.oid
  left join pg_description t5 on t5.objoid = t4.attrelid and t5.objsubid = t4.attnum`)
	sql.WriteString(` where t1.table_catalog = '`)
	sql.WriteString(database)
	sql.WriteString(`'`)
	if schema != "" {
		sql.WriteString(` and t1.table_schema = '`)
		sql.WriteString(schema)
		sql.WriteString(`'`)
	}
	if len(table) > 0 {
		sql.WriteString(` and t1.table_name in ('`)
		sql.WriteString(strings.Join(table, `','`))
		sql.WriteString(`')`)
	}
	sql.WriteString(` order by t1.table_schema, t1.table_name, t1.ordinal_position`)
	if err = gormx.DB(source).Raw(sql.String()).Scan(&fields).Error; err != nil {
		log.Errorf("查询表字段列表 失败：%s\n ", err)
		return
	}
	return
}
