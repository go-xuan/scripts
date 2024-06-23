package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-xuan/quanx/os/filex"
	"github.com/go-xuan/quanx/os/filex/excelx"
	"github.com/tealeg/xlsx"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	Mysql    = "mysql"
	Postgres = "postgres"
)

var config *Config

type Config struct {
	Type     string `json:"type" yaml:"type"`         // 数据库类型
	Host     string `json:"host" yaml:"host"`         // 数据库Host
	Port     int    `json:"port" yaml:"port"`         // 数据库端口
	Database string `json:"database" yaml:"database"` // 数据库名
	Schema   string `json:"schema" yaml:"schema"`     // 模式名
	Username string `json:"username" yaml:"username"` // 用户名
	Password string `json:"password" yaml:"password"` // 密码
}

func main() {
	var err error
	var configIn = flag.String("config", "config.yaml", "输入配置文件，例如：-config=config.yaml")
	//var table = flag.String("table", "edu_app_rlzy_zzmm_lx,edu_app_rlzy_zjxyqk_lx,edu_app_rlzy_zdrw_lx,edu_app_rlzy_zdkyxm_lx,edu_app_rlzy_tq_lx,edu_app_rlzy_sjcqk_lx,edu_app_rlzy_rwqjl_lx,edu_app_rlzy_rwhjl_lx,edu_app_rlzy_rcqk_lx,edu_app_rlzy_qnbtff_lx,edu_app_rlzy_poqk_lx,edu_app_rlzy_kjjl_lx,edu_app_rlzy_khqk_lx,edu_app_rlzy_jypxjl_lx,edu_app_rlzy_jtcyxx_lx,edu_app_rlzy_jszyjy_lx,edu_app_rlzy_jbxx_lx,edu_app_rlzy_hzjzgqk_lx,edu_app_rlzy_gzdc_lx,edu_app_rlzy_grsjzh_lx,edu_app_rlzy_dfjzqk_lx,edu_app_rlzy_cltxzsq_lx,edu_app_rlzy_cjttqk_lx,edu_app_rlzy_cgjqk_lx", "输入表名，例如：-table=t_abc_def")
	//var table = flag.String("table", "hr_confirm_task,hr_confirm_task_person,hr_confirm_task_record,hr_confirm_task_workflow,hr_scoring,hr_scoring_indicator,hr_scoring_person", "输入表名，例如：-table=t_abc_def")
	var table = flag.String("table", "edu_app_rlzy_jbxx_lx,edu_app_rlzy_jbxx_sjb,edu_app_rlzy_jbxx_sjb_xzjl,edu_app_rlzy_jbxx_xzmb,edu_app_rlzy_jbxx_xzzd", "输入表名，例如：-table=t_abc_def")
	// 解析输入的参数
	flag.Parse()
	var bytes []byte
	if bytes, err = os.ReadFile(*configIn); err != nil {
		fmt.Println(err)
		panic(err)
	}
	config = &Config{}
	if err = yaml.Unmarshal(bytes, config); err != nil {
		fmt.Println(err)
		panic(err)
	}
	var tables []string
	if table != nil {
		tables = strings.Split(*table, ",")
	}
	var allFields []*TableField
	if allFields, err = config.TableFieldList(tables); err != nil {
		fmt.Println(err)
		panic(err)
	}

	//导出
	if err = Export(allFields); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Export(fields []*TableField) (err error) {
	var tableMap = make(map[string][]*TableField)
	for _, field := range fields {
		tableMap[field.Table] = append(tableMap[field.Table], field)
	}
	var xlsxFile = xlsx.NewFile()
	var sheet *xlsx.Sheet
	if sheet, err = xlsxFile.AddSheet("Sheet1"); err != nil {
		fmt.Printf("创建excel文件失败：%s", err)
		return
	}
	headers := excelx.GetHeadersByReflect(TableField{})
	var headerStyle = excelx.HeaderStyle()
	var defaultStyle = excelx.DefaultStyle()
	var headerCells []*xlsx.Cell
	for _, header := range headers {
		cell := &xlsx.Cell{Value: header.Name}
		cell.SetStyle(headerStyle)
		headerCells = append(headerCells, cell)
	}

	hMerge := len(headers) - 1
	for tableName, tableFields := range tableMap {
		// 添加表注释行
		tableCommentRow := sheet.AddRow()
		tableCommentCell := tableCommentRow.AddCell()
		tableCommentCell.Value = tableFields[0].TableComment
		tableCommentCell.HMerge = hMerge
		tableCommentCell.SetStyle(headerStyle)
		for i := 0; i < hMerge; i++ {
			tableCommentRow.AddCell().SetStyle(headerStyle)
		}
		// 添加表名行
		tableRow := sheet.AddRow()
		tableCell := tableRow.AddCell()
		tableCell.Value = tableName
		tableCell.HMerge = hMerge
		tableCell.SetStyle(headerStyle)
		for i := 0; i < hMerge; i++ {
			tableRow.AddCell().SetStyle(headerStyle)
		}

		// 添加表头行
		sheet.AddRow().Cells = headerCells
		for _, field := range tableFields {
			b, _ := json.Marshal(field)
			data := gjson.ParseBytes(b).Map()
			row := sheet.AddRow()
			for _, header := range headers {
				cell := row.AddCell()
				cell.SetStyle(defaultStyle)
				cell.Value = data[header.Key].String()
			}
		}
		// 添加间隔行
		sheet.AddRow()
	}
	// 这里重新生成
	dir := filepath.Join(config.Type, config.Database)
	filex.CreateDir(dir)
	excelPath := filepath.Join(dir, config.Database+"_export_"+time.Now().Format("20060102150405")+".xlsx")
	if err = xlsxFile.Save(excelPath); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (conf *Config) TableFieldList(tables []string) (result []*TableField, err error) {
	var db *gorm.DB
	db, err = conf.NewGormDB()
	if err != nil {
		fmt.Println("数据库连接失败" + conf.Format())
		return
	}
	switch conf.Type {
	case Mysql:
		result, err = MysqlTableFieldList(db, conf.Database, tables)
	case Postgres:
		result, err = PgTableFieldList(db, conf.Database, conf.Schema, tables)
	}
	if err != nil {
		fmt.Println("数据库查询失败" + conf.Format())
		return
	}
	return
}

// 查询表字段列表
func MysqlTableFieldList(db *gorm.DB, database string, tables []string) (resultList []*TableField, err error) {
	querySql := strings.Builder{}
	querySql.WriteString(`
select t1.column_name as "name",
       t1.table_name as "table",
       t1.table_schema as "database",
       t1.data_type as "type",
       t1.column_default as "default",
       t1.column_comment as "comment",
       t2.table_comment as "table_comment",
       if(t1.data_type = 'varchar', t1.character_maximum_length, t1.numeric_precision) as "precision",
       t1.numeric_scale as "scale",
       t1.is_nullable = 'YES' as "nullable"
  from information_schema.columns t1
  left join information_schema.tables t2
    on t1.table_schema = t2.table_schema
   and t1.table_name = t2.table_name`)
	querySql.WriteString(` where t1.table_schema = '`)
	querySql.WriteString(database)
	querySql.WriteString(`'`)
	if len(tables) > 0 {
		querySql.WriteString(` and t1.table_name in ('`)
		querySql.WriteString(strings.Join(tables, `','`))
		querySql.WriteString(`')`)
	}
	querySql.WriteString(` order by t1.table_name, t1.ordinal_position`)
	err = db.Raw(querySql.String()).Scan(&resultList).Error
	if err != nil {
		fmt.Printf("查询表字段列表 失败：%s\n ", err)
		return
	}
	return
}

// 查询表字段列表
func PgTableFieldList(db *gorm.DB, database, schema string, tables []string) (resultList []*TableField, err error) {
	querySql := strings.Builder{}
	querySql.WriteString(`
select t1.column_name as name,
       t1.table_name as table,
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
	querySql.WriteString(` where t1.table_catalog = '`)
	querySql.WriteString(database)
	querySql.WriteString(`'`)
	if schema != "" {
		querySql.WriteString(` and t1.table_schema = '`)
		querySql.WriteString(schema)
		querySql.WriteString(`'`)
	}
	if len(tables) > 0 {
		querySql.WriteString(` and t1.table_name in ('`)
		querySql.WriteString(strings.Join(tables, `','`))
		querySql.WriteString(`')`)
	}
	querySql.WriteString(` order by t1.table_schema, t1.table_name, t1.ordinal_position`)
	err = db.Raw(querySql.String()).Scan(&resultList).Error
	if err != nil {
		fmt.Printf("查询表字段列表 失败：%s\n ", err)
		return
	}
	return
}

// 配置信息格式化
func (conf *Config) Format() string {
	return fmt.Sprintf(
		"type=%s host=%s port=%d database=%s schema=%s username=%v password=%s",
		conf.Type, conf.Host, conf.Port, conf.Database, conf.Schema, conf.Username, conf.Password)
}

// 创建数据库连接
func (conf *Config) NewGormDB() (gormDB *gorm.DB, err error) {
	gormDB, err = GetGormDB(conf.GetDSN(), conf.Type)
	if err != nil {
		return
	}
	var sqlDB *sql.DB
	sqlDB, err = gormDB.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(60)
	sqlDB.SetConnMaxLifetime(time.Hour)
	gormDB = gormDB.Debug()
	return
}

// 获取数据库连接DSN
func (conf *Config) GetDSN() (dsn string) {
	switch strings.ToLower(conf.Type) {
	case Mysql:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf.Username, conf.Password, conf.Host, conf.Port, conf.Database) +
			"?clientFoundRows=false&parseTime=true&timeout=1800s&charset=utf8&collation=utf8_general_ci&loc=Local"
	case Postgres:
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.Host, conf.Port, conf.Username, conf.Password, conf.Database)
	}
	return
}

// 根据dsn生成gormDB
func GetGormDB(dsn, dialect string) (gormDb *gorm.DB, err error) {
	switch strings.ToLower(dialect) {
	case Mysql:
		return gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	case Postgres:
		return gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	}
	return
}

// 表字段信息
type TableField struct {
	Name         string `json:"name" excel:"字段名"`    // 字段名
	Table        string `json:"table"`               // 表名
	TableComment string `json:"tableComment"`        // 表注释
	Database     string `json:"database"`            // 数据库名
	Schema       string `json:"schema"`              // schema名
	Type         string `json:"type" excel:"数据类型"`   // 数据类型
	Default      string `json:"default" excel:"默认值"` // 默认值
	//Precision    int    `json:"precision" excel:"长度"`    // 长度
	//Scale        int    `json:"scale" excel:"小数点"`       // 小数点
	//Nullable     bool   `json:"nullable" excel:"是否允许为空"` // 是否允许为空
	Comment string `json:"comment" excel:"注释"` // 注释
}
