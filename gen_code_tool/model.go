package main

import (
	"errors"
	"gorm.io/gorm"
	"path/filepath"
	"strings"

	"github.com/go-xuan/quanx/server/gormx"
	"github.com/go-xuan/quanx/types/stringx"
	log "github.com/sirupsen/logrus"

	"gen_code_tool/common"
	"gen_code_tool/template"
)

type Config struct {
	App         string          `json:"app" json:"app" default:"demo"`         // 应用名
	Frame       string          `json:"frame" yaml:"frame" default:"go-quanx"` // 代码框架（go-quanx/go-frame/spring-boot）
	Output      string          `json:"output" yaml:"output"`                  // 输出路径
	Include     string          `json:"include" yaml:"include"`                // 生成代码的表名（为空则获取全表,多个以“,”分割）
	Exclude     string          `json:"exclude" yaml:"exclude"`                // 排除表名（多个以“,”分割）
	TablePrefix string          `json:"tablePrefix" yaml:"tablePrefix"`        // 表前缀
	TableSuffix string          `json:"tableSuffix" yaml:"tableSuffix"`        // 表后缀缀
	DB          *gormx.Database `json:"db" yaml:"db"`                          // 应用数据库
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
	// 查询应用数据库表对应模型
	if models, err := config.GetModels(); err == nil {
		generator.Models = models
	}
	return generator
}

func (c *Config) Trim(table string) string {
	if c.TablePrefix != "" {
		table = strings.TrimPrefix(table, c.TablePrefix)
	}
	if c.TableSuffix != "" {
		table = strings.TrimSuffix(table, c.TableSuffix)
	}
	return table
}

func (c *Config) GetModels() (result []*Model, err error) {
	var fields []*Field
	var gormDB *gorm.DB
	if gormDB, err = c.DB.NewGormDB(); err == nil && gormDB != nil {
		var sql string
		switch c.DB.Type {
		case gormx.Mysql:
			sql = mysqlTableFieldQuerySql(c.DB.Database, c.Include, c.Exclude)
		case gormx.Postgres:
			sql = pgsqlTableFieldQuerySql(c.DB.Database, c.DB.Schema, c.Include, c.Exclude)
		default:
			err = errors.New("db.type can only be mysql or postgres")
			return
		}
		if err = gormDB.Raw(sql).Scan(&fields).Error; err != nil {
			log.Errorf("查询表字段列表 失败：%s\n ", err)
			return
		}
	}
	if len(fields) > 0 {
		var models = make(map[string]*Model)
		for _, field := range fields {
			field.GoName = stringx.ToUpperCamel(field.Name)
			field.GoType = common.DB2GoType(field.Type)
			field.GormType = common.DB2GormType(field.Type)
			field.JavaType = common.DB2JavaType(field.Type)
			var nameLen, typeLen = len(field.GoName), len(field.GoType)
			var table = field.TableName
			if _, ok := models[table]; ok {
				models[table].Fields = append(models[table].Fields, field)
				if models[table].FiledNameLen < nameLen {
					models[table].FiledNameLen = nameLen
				}
				if models[table].FiledTypeLen < typeLen {
					models[table].FiledTypeLen = typeLen
				}
			} else {
				models[table] = &Model{
					Table:        table,
					Name:         c.Trim(table),
					Schema:       field.Schema,
					Comment:      field.TableComment,
					FiledNameLen: nameLen,
					FiledTypeLen: typeLen,
					Fields:       []*Field{field},
				}
			}
		}
		for _, model := range models {
			for _, field := range model.Fields {
				field.GoName = stringx.Fill(field.GoName, " ", model.FiledNameLen, true)
				field.GoType = stringx.Fill(field.GoType, " ", model.FiledTypeLen, true)
			}
			result = append(result, model)
		}
	}
	return
}

// 代码生成器
type Generator struct {
	App    string           // 应用名
	Root   string           // 代码生成路径
	DB     *gormx.Database  // 应用数据库
	Models []*Model         // 模型列表
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
				for _, model := range gen.Models {
					if err = tmpl.WriteCodeToFile(gen.Root, model, model.Name); err != nil {
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

type Model struct {
	Table        string   `json:"table"`        // 表名
	Name         string   `json:"name"`         // 模型名
	Schema       string   `json:"schema"`       // schema
	Comment      string   `json:"comment"`      // 表备注
	FiledNameLen int      `json:"filedNameLen"` // 字段名称长度
	FiledTypeLen int      `json:"filedTypeLen"` // 字段类型长度
	Fields       []*Field `json:"fields"`       // 字段列表
}

type Field struct {
	Name         string `json:"name"`         // 字段名
	Type         string `json:"type"`         // 数据类型
	Precision    int    `json:"precision"`    // 长度
	Scale        int    `json:"scale"`        // 小数点
	Nullable     bool   `json:"nullable"`     // 允许为空
	Default      string `json:"default"`      // 默认值
	Comment      string `json:"comment"`      // 注释
	Database     string `json:"database"`     // 数据库名
	Schema       string `json:"schema"`       // schema名
	TableName    string `json:"tableName"`    // 表名
	TableComment string `json:"tableComment"` // 表注释
	GoName       string `json:"goName"`       // go字段名
	GoType       string `json:"goType"`       // go数据类型
	GormType     string `json:"ormType"`      // orm字段类型
	JavaType     string `json:"javaType"`     // Java字段类型
}

// 查询表字段列表
func mysqlTableFieldQuerySql(database, include, exclude string) string {
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
	if include != "" {
		include = strings.ReplaceAll(include, " ", "")
		include = strings.ReplaceAll(include, ",", "','")
		sql.WriteString(` and t1.table_name in ('`)
		sql.WriteString(include)
		sql.WriteString(`')`)
	} else if exclude != "" {
		exclude = strings.ReplaceAll(exclude, " ", "")
		exclude = strings.ReplaceAll(exclude, ",", "','")
		sql.WriteString(` and t1.table_name not in ('`)
		sql.WriteString(exclude)
		sql.WriteString(`')`)
	}
	sql.WriteString(` order by t1.table_name, t1.ordinal_position`)
	return sql.String()
}

// 查询表字段列表
func pgsqlTableFieldQuerySql(database, schema, include, exclude string) string {
	sql := strings.Builder{}
	sql.WriteString(`
select t1.column_name as name,
       t1.table_name as table_name,
       t1.table_schema as schema,
       t1.table_catalog as database,
       t1.udt_name as type,
       t1.column_default as default,
       obj_description(t3.oid) as table_comment,
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
	if include != "" {
		include = strings.ReplaceAll(include, " ", "")
		include = strings.ReplaceAll(include, ",", "','")
		sql.WriteString(` and t1.table_name in ('`)
		sql.WriteString(include)
		sql.WriteString(`')`)
	} else if exclude != "" {
		exclude = strings.ReplaceAll(exclude, " ", "")
		exclude = strings.ReplaceAll(exclude, ",", "','")
		sql.WriteString(` and t1.table_name not in ('`)
		sql.WriteString(exclude)
		sql.WriteString(`')`)
	}
	sql.WriteString(` order by t1.table_name, t1.ordinal_position`)
	return sql.String()
}
