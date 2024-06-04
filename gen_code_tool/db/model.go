package db

type Table struct {
	Name     string   `json:"name"`     // 表名
	Schema   string   `json:"schema"`   // schema
	Comment  string   `json:"comment"`  // 表备注
	FieldLen int      `json:"fieldLen"` // 字段最大长度
	Fields   []*Field `json:"fields"`   // 字段列表
}

type Field struct {
	Name         string `json:"name" excel:"字段名"`      // 字段名
	Type         string `json:"type" excel:"数据类型"`     // 数据类型
	Precision    int    `json:"precision" excel:"长度"`  // 长度
	Scale        int    `json:"scale" excel:"小数点"`     // 小数点
	Nullable     bool   `json:"nullable" excel:"允许为空"` // 允许为空
	Default      string `json:"default" excel:"默认值"`   // 默认值
	Comment      string `json:"comment" excel:"注释"`    // 注释
	Database     string `json:"database"`              // 数据库名
	Schema       string `json:"schema"`                // schema名
	TableName    string `json:"tableName"`             // 表名
	TableComment string `json:"tableComment"`          // 表注释
}
