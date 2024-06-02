package common

// tmplName
type TmplName string

const (
	ConstsTmpl      = "consts.tmpl"
	ConfigTmpl      = "config.tmpl"
	ControllerTmpl  = "controller.tmpl"
	LogicTmpl       = "logic.tmpl"
	DaoTmpl         = "dao.tmpl"
	ModelTmpl       = "model.tmpl"
	ModelDoTmpl     = "model_do.tmpl"
	ModelEntityTmpl = "model_entity.tmpl"
	RouterTmpl      = "router.tmpl"
	CmdTmpl         = "cmd.tmpl"
)

// 随机数生成器-数据类型
const (
	// go
	String  = "string"
	Int     = "int"
	Int64   = "int64"
	Float64 = "float64"
	Bool    = "bool"
	Time    = "time.Time"

	// database
	Date      = "date"
	Varchar   = "varchar"
	Char      = "char"
	Text      = "text"
	Int2      = "int2"
	Int4      = "int4"
	Int8      = "int8"
	Tinyint   = "tinyint"
	Smallint  = "smallint"
	Mediumint = "mediumint"
	Bigint    = "bigint"
	Float4    = "float4"
	Numeric   = "numeric" // 数字
	Numeric6  = "numeric(10,6)"
	Numeric2  = "numeric(10,2)"
	Decimal   = "decimal"
	Timestamp = "timestamp"
	Datetime  = "datetime"
	// java
	JavaString     = "String"
	JavaInteger    = "Integer"
	JavaInt        = "int"
	JavaLong       = "Long"
	JavaDate       = "Date"
	JavaBigDecimal = "BigDecimal"
	JavaFloat      = "Float"
	JavaBoolean    = "Boolean"
)

// BASE基础字段映射
func IsBaseField(t string) bool {
	switch t {
	case "id":
		return true
	case "create_time", "create_user_id", "create_by":
		return true
	case "update_time", "update_user_id", "update_by":
		return true
	default:
		return false
	}
}

// DB-java类型映射
func DB2JavaType(t string) string {
	switch t {
	case Char, Varchar, Text:
		return JavaString
	case Int2, Smallint, Mediumint:
		return JavaInteger
	case Tinyint:
		return JavaInt
	case Int, Int4, Int8, Bigint:
		return JavaLong
	case Float4, Numeric:
		return JavaFloat
	case Decimal:
		return JavaBigDecimal
	case Timestamp, Datetime, Date:
		return JavaDate
	case Bool:
		return JavaBoolean
	default:
		return JavaString
	}
}

// DB-Go类型映射
func DB2GoType(t string) string {
	switch t {
	case Char, Varchar, Text:
		return String
	case Int, Int2, Int4, Tinyint, Smallint, Mediumint:
		return Int
	case Int8, Bigint:
		return Int64
	case Float4, Numeric:
		return Float64
	case Timestamp, Datetime, Date:
		return Time
	case Bool:
		return Bool
	default:
		return String
	}
}

// DB-Gorm类型映射
func DB2GormType(t string) string {
	switch t {
	case Char:
		return Char
	case Text:
		return Text
	case Tinyint:
		return Tinyint
	case Smallint, Int2:
		return Smallint
	case Mediumint, Int4:
		return Mediumint
	case Bigint, Int8:
		return Bigint
	case Int:
		return Int
	case Float4:
		return Numeric6
	case Numeric:
		return Numeric2
	case Timestamp, Datetime:
		return Timestamp
	case Bool:
		return Bool
	case Date:
		return Date
	default:
		return t
	}
}
