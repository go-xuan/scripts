package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-xuan/quanx/types/stringx"

	"gen_code_tool/common"
)

type Table struct {
	Name     string   `json:"name"`     // 表名
	Schema   string   `json:"schema"`   // schema
	Comment  string   `json:"comment"`  // 表备注
	FieldLen int      `json:"fieldLen"` // 字段长度
	Fields   []*Field `json:"fields"`   // 字段列表
}

type Field struct {
	Name         string `json:"name" excel:"字段名"`       // 字段名
	Type         string `json:"type" excel:"数据类型"`     // 数据类型
	Precision    int    `json:"precision" excel:"长度"`    // 长度
	Scale        int    `json:"scale" excel:"小数点"`      // 小数点
	Nullable     bool   `json:"nullable" excel:"允许为空"` // 允许为空
	Default      string `json:"default" excel:"默认值"`    // 默认值
	Comment      string `json:"comment" excel:"注释"`      // 注释
	Database     string `json:"database"`                  // 数据库名
	Schema       string `json:"schema"`                    // schema名
	TableName    string `json:"tableName"`                 // 表名
	TableComment string `json:"tableComment"`              // 表注释
}

func (t *Table) JavaEntity() string {
	sb := strings.Builder{}
	for i, field := range t.Fields {
		lowerCamel := stringx.ToLowerCamel(field.Name)
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("\t")
		sb.WriteString("@ApiModelProperty(value = \"")
		sb.WriteString(field.Comment)
		sb.WriteString("\", name = \"")
		sb.WriteString(lowerCamel)
		sb.WriteString("\")\n\tprivate ")
		sb.WriteString(common.DB2JavaType(field.Type))
		sb.WriteString(" ")
		sb.WriteString(lowerCamel)
		sb.WriteString(";")
	}
	return sb.String()
}

func (t *Table) GormStruct() string {
	sb := strings.Builder{}
	for i, field := range t.Fields {
		uc, lc := stringx.ToUpperCamel(field.Name), stringx.ToLowerCamel(field.Name)
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("\t")
		sb.WriteString(uc)
		sb.WriteString(" ")
		sb.WriteString(common.DB2GoType(field.Type))
		sb.WriteString(" `json:\"")
		sb.WriteString(lc)
		sb.WriteString("\" gorm:\"type:")
		if field.Type == common.Varchar {
			sb.WriteString(common.Varchar)
			sb.WriteString("(")
			sb.WriteString(strconv.Itoa(field.Precision))
			sb.WriteString(")")
		} else {
			sb.WriteString(common.DB2GormType(field.Type))
		}
		if !field.Nullable {
			sb.WriteString("; not null")
		}
		if lc == "id" {
			sb.WriteString("; primary_key")
		} else if field.Default != "" {
			sb.WriteString("; default:")
			sb.WriteString(field.Default)
		}
		sb.WriteString("; comment:")
		sb.WriteString(field.Comment)
		sb.WriteString(";\"`")
	}
	return sb.String()
}

func (t *Table) GfModelDo() string {
	sb := strings.Builder{}
	for i, field := range t.Fields {
		uc, lc := stringx.ToUpperCamel(field.Name), stringx.ToLowerCamel(field.Name)
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("\t")
		sb.WriteString(uc)
		sb.WriteString(" ")
		sb.WriteString(common.DB2GoType(field.Type))
		sb.WriteString(" `json:\"")
		sb.WriteString(lc)
		sb.WriteString("\" gorm:\"type:")
		if field.Type == common.Varchar {
			sb.WriteString(common.Varchar)
			sb.WriteString("(")
			sb.WriteString(strconv.Itoa(field.Precision))
			sb.WriteString(")")
		} else {
			sb.WriteString(common.DB2GormType(field.Type))
		}
		if !field.Nullable {
			sb.WriteString("; not null")
		}
		if lc == "id" {
			sb.WriteString("; primary_key")
		} else if field.Default != "" {
			sb.WriteString("; default:")
			sb.WriteString(field.Default)
		}
		sb.WriteString("; comment:")
		sb.WriteString(field.Comment)
		sb.WriteString(";\"`")
	}
	return sb.String()
}

func (t *Table) GfModelEntity() string {
	sb := strings.Builder{}
	for i, field := range t.Fields {
		uc, lc := stringx.ToUpperCamel(field.Name), stringx.ToLowerCamel(field.Name)
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("\t")
		sb.WriteString(uc)
		sb.WriteString(" ")
		sb.WriteString(common.DB2GoType(field.Type))
		sb.WriteString(" `json:\"")
		sb.WriteString(lc)
		sb.WriteString("\" gorm:\"type:")
		if field.Type == common.Varchar {
			sb.WriteString(common.Varchar)
			sb.WriteString("(")
			sb.WriteString(strconv.Itoa(field.Precision))
			sb.WriteString(")")
		} else {
			sb.WriteString(common.DB2GormType(field.Type))
		}
		if !field.Nullable {
			sb.WriteString("; not null")
		}
		if lc == "id" {
			sb.WriteString("; primary_key")
		} else if field.Default != "" {
			sb.WriteString("; default:")
			sb.WriteString(field.Default)
		}
		sb.WriteString("; comment:")
		sb.WriteString(field.Comment)
		sb.WriteString(";\"`")
	}
	return sb.String()
}

func (t *Table) SelectSql() string {
	ss := "%" + strconv.Itoa(-t.FieldLen) + "s as %s,"
	sb := strings.Builder{}
	sb.WriteString("select ")
	for i, field := range t.Fields {
		if i > 0 {
			sb.WriteString("\n")
			sb.WriteString("       ")
		}
		sb.WriteString(fmt.Sprintf(ss, field.Name, stringx.ToLowerCamel(field.Name)))
	}
	sb.WriteString(",\n  from ")
	sb.WriteString(t.Name)
	return strings.ReplaceAll(sb.String(), ",,", "")
}

func (t *Table) InsertSql() string {
	sb, iv := strings.Builder{}, strings.Builder{}
	sb.WriteString("insert into")
	sb.WriteString(" ")
	sb.WriteString(t.Name)
	sb.WriteString("\n(")
	var i int
	for _, field := range t.Fields {
		var fieldName = field.Name
		if common.IsBaseField(fieldName) && field.Default != "" {
			continue
		}
		if i > 0 {
			sb.WriteString("\n")
			iv.WriteString("\n")
		}
		sb.WriteString(fieldName)
		sb.WriteString(",")
		iv.WriteString("#{create.")
		iv.WriteString(stringx.ToLowerCamel(fieldName))
		iv.WriteString("},")
		i++
	}
	sb.WriteString(",)\nvalues \n(")
	sb.WriteString(iv.String())
	sb.WriteString(",)")
	return strings.ReplaceAll(sb.String(), ",,", "")
}

func (t *Table) UpdateSql() string {
	sb := strings.Builder{}
	sb.WriteString("update ")
	sb.WriteString(t.Name)
	sb.WriteString("\n<set>")
	for _, field := range t.Fields {
		var fieldName = field.Name
		if common.IsBaseField(fieldName) {
			continue
		}
		lc := "update." + stringx.ToLowerCamel(fieldName)
		sb.WriteString("\n\t")
		if field.Type == common.Varchar {
			sb.WriteString(fmt.Sprintf(`<if test="%s != null and %s != ''"> %s = #{%s}, </if>`, lc, lc, fieldName, lc))
		} else {
			sb.WriteString(fmt.Sprintf(`<if test="%s != null"> %s = #{%s}, </if>`, lc, fieldName, lc))
		}
	}
	sb.WriteString("\n</set>")
	sb.WriteString("\nwhere id = #{update.id}")
	return strings.ReplaceAll(sb.String(), ",,", "")
}
