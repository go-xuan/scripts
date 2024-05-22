package entity

import (
	"github.com/go-xuan/quanx/utils/randx"
)

var zbmcs = []string{"指标一", "指标二", "指标三", "指标四"}

var cyxl = map[string]string{
	"新材料产业":         "战兴产业",
	"新一代信息技术产业":     "国家十四五产业",
	"高端设备制造产业":      "国家十四五产业",
	"新能源产业":         "战兴产业",
	"节能环保产业":        "国家十四五产业",
	"新能源(智能网联)汽车产业": "国家十四五产业",
	"生命健康产业":        "战兴产业",
	"高端纺织产业":        "战兴产业",
	"化工产业":          "战兴产业",
	"新兴数字产业":        "国家十四五产业",
	"航空航天产业":        "国家十四五产业",
	"集成电路产业":        "战兴产业",
}

var cyglsf map[string][]string

func getCysfys() map[string][]string {
	if cyglsf == nil {
		cyglsf = make(map[string][]string)
		for _, cymc := range cymcs {
			var n = randx.IntRange(5, 10)
			var m = make(map[string]struct{})
			for i := 0; i < n; i++ {
				m[randx.Enum(randx.ProvinceName, ",")] = struct{}{}
			}
			var list []string
			for k := range m {
				list = append(list, k)
			}
			cyglsf[cymc] = list
		}
	}
	return cyglsf
}

var cyglxkz map[string][]string

func getCyxkzys() map[string][]string {
	if cyglxkz == nil {
		cyglxkz = make(map[string][]string)
		for _, cymc := range cymcs {
			var n = randx.IntRange(5, 10)
			var m = make(map[string]struct{})
			for i := 0; i < n; i++ {
				xkzymc := "学科专业" + randx.Enum("一,二,三,四,五,六,七,八,九,十", ",")
				m[xkzymc] = struct{}{}
			}
			var list []string
			for k := range m {
				list = append(list, k)
			}
			cyglxkz[cymc] = list
		}
	}
	return cyglxkz
}

// 十四五产业-概览
type AppSswcyGl struct {
	Cymc  string  `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Cylx  string  `json:"cms" gorm:"type:varchar(100); not null; comment:产业类型;"`
	Ztzcd float64 `json:"ztzcd" gorm:"type:decimal(4,2); comment:整体支撑度(小数，保留两位小数);"`
	Pxxb  int     `json:"pxxb" gorm:"type:smallint; not null; default:1; comment:排序下标;"`
}

// 表名称
func (t *AppSswcyGl) TableName() string {
	var table = "app_sswcy_gl"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyGl) Comment() string {
	return "十四五产业-概览"
}

// 表数据
func (t *AppSswcyGl) InitData() any {
	var data []*AppSswcyGl
	for _, cymc := range cymcs {
		data = append(data, &AppSswcyGl{
			Cymc:  cymc,
			Cylx:  cyxl[cymc],
			Ztzcd: randx.Float64Range(0, 1, 2),
		})
	}
	return data
}

// 十四五产业-学科专业
type AppSswcyXkzy struct {
	Cymc   string `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Xkzymc string `json:"xkzymc" gorm:"type:varchar(100); not null; comment:学科专业名称;"`
}

// 表名称
func (t *AppSswcyXkzy) TableName() string {
	var table = "app_sswcy_xkzy"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyXkzy) Comment() string {
	return "十四五产业-学科专业"
}

// 表数据
func (t *AppSswcyXkzy) InitData() any {
	var data []*AppSswcyXkzy
	for _, cymc := range cymcs {
		for _, xkzymc := range getCyxkzys()[cymc] {
			data = append(data, &AppSswcyXkzy{
				Cymc:   cymc,
				Xkzymc: xkzymc,
			})
		}
	}
	return data
}

// 十四五产业-所在省份
type AppSswcySzsf struct {
	Cymc string  `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Sfmc string  `json:"sfmc" gorm:"type:varchar(100); not null; comment:省份名称;"`
	Zcd  float64 `json:"zcd" gorm:"type:decimal(4,2); comment:支撑度(小数，保留两位小数);"`
}

// 表名称
func (t *AppSswcySzsf) TableName() string {
	var table = "app_sswcy_szsf"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcySzsf) Comment() string {
	return "十四五产业-所在省份"
}

// 表数据
func (t *AppSswcySzsf) InitData() any {
	var data []*AppSswcySzsf
	for _, cymc := range cymcs {
		for _, sfmc := range getCysfys()[cymc] {
			data = append(data, &AppSswcySzsf{
				Cymc: cymc,
				Sfmc: sfmc,
				Zcd:  randx.Float64Range(0, 1, 2),
			})
		}
	}
	return data
}

// 十四五产业-研究生教育-支撑情况
type AppSswcyYjsjyZcqk struct {
	Cymc  string  `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Zcqk  string  `json:"zcqk" gorm:"type:varchar(100); not null; comment:支撑情况(支撑一般/基本支撑/支撑不足);"`
	Zcdzb float64 `json:"zcdzb" gorm:"type:decimal(4,2); comment:支撑度占比(百分比，保留两位小数);"`
}

// 表名称
func (t *AppSswcyYjsjyZcqk) TableName() string {
	var table = "app_sswcy_yjsjy_zcqk"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyYjsjyZcqk) Comment() string {
	return "十四五产业-研究生教育-支撑情况"
}

// 表数据
func (t *AppSswcyYjsjyZcqk) InitData() any {
	var data []*AppSswcyYjsjyZcqk
	for _, cymc := range cymcs {
		for _, zcqk := range zcqks {
			data = append(data, &AppSswcyYjsjyZcqk{
				Cymc:  cymc,
				Zcqk:  zcqk,
				Zcdzb: randx.Float64Range(0, 1, 2),
			})
		}
	}
	return data
}

// 十四五产业-学科评估
type AppSswcyXkpg struct {
	Cymc   string  `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Xkzymc string  `json:"xkzymc" gorm:"type:varchar(100); not null; comment:学科专业名称;"`
	Zbmc   string  `json:"zbmc" gorm:"type:varchar(100); not null; comment:指标名称;"`
	Zxz    float64 `json:"zxz" gorm:"type:decimal(4,2); comment:最小值(保留两位小数);"`
	Dysfws float64 `json:"dysfws" gorm:"type:decimal(4,2); comment:第一四分位数(保留两位小数);"`
	Zws    float64 `json:"zws" gorm:"type:decimal(4,2); comment:中位数(保留两位小数);"`
	Dssfws float64 `json:"dssfws" gorm:"type:decimal(4,2); comment:第三四分位数(保留两位小数);"`
	Zdz    float64 `json:"zdz" gorm:"type:decimal(4,2); comment:最大值(保留两位小数);"`
	Pxxb   int     `json:"pxxb" gorm:"type:smallint; not null; default:1; comment:排序下标;"`
}

// 表名称
func (t *AppSswcyXkpg) TableName() string {
	var table = "app_sswcy_xkpg"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyXkpg) Comment() string {
	return "十四五产业-学科评估"
}

// 表数据
func (t *AppSswcyXkpg) InitData() any {
	var data []*AppSswcyXkpg
	for _, cymc := range cymcs {
		for _, xkzymc := range getCyxkzys()[cymc] {
			for j, zbmc := range zbmcs {
				data = append(data, &AppSswcyXkpg{
					Cymc:   cymc,
					Xkzymc: xkzymc,
					Zbmc:   zbmc,
					Zxz:    randx.Float64Range(1, 50, 2),
					Dysfws: randx.Float64Range(20, 40, 2),
					Zws:    randx.Float64Range(40, 60, 2),
					Dssfws: randx.Float64Range(60, 80, 2),
					Zdz:    randx.Float64Range(80, 100, 2),
					Pxxb:   j + 1,
				})
			}
		}
	}
	return data
}

// 十四五产业-强相关学科-专业建设情况
type AppSswcyQxgxkZyjsqk struct {
	Cymc    string  `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Sfmc    string  `json:"sfmc" gorm:"type:varchar(100); not null; comment:省份名称（分为全国和省份）;"`
	Gxsl    int     `json:"gxsl" gorm:"type:int; not null; comment:高校数量;"`
	Sylgxzb float64 `json:"sylgxzb" gorm:"type:decimal(4,2); not null; comment:双一流高校占比(百分比，保留两位小数);"`
	Bsdsl   int     `json:"bsdsl" gorm:"type:int; not null; comment:博士点数量;"`
	Bsdzb   float64 `json:"bsdzb" gorm:"type:decimal(4,2); not null; comment:博士点占比(百分比，保留两位小数);"`
	Ssdsl   int     `json:"ssdsl" gorm:"type:int; not null; comment:硕士点数量;"`
	Ssdzb   float64 `json:"ssdzb" gorm:"type:decimal(4,2); not null; comment:硕士点占比(百分比，保留两位小数);"`
}

// 表名称
func (t *AppSswcyQxgxkZyjsqk) TableName() string {
	var table = "app_sswcy_qxgxk_zyjsqk"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyQxgxkZyjsqk) Comment() string {
	return "十四五产业-强相关学科-专业建设情况"
}

// 表数据
func (t *AppSswcyQxgxkZyjsqk) InitData() any {
	var data []*AppSswcyQxgxkZyjsqk
	for _, cymc := range cymcs {
		data = append(data, &AppSswcyQxgxkZyjsqk{
			Cymc:    cymc,
			Sfmc:    "全国",
			Gxsl:    randx.IntRange(100, 1000),
			Sylgxzb: randx.Float64Range(10, 99, 2),
			Bsdsl:   randx.IntRange(100, 1000),
			Bsdzb:   randx.Float64Range(10, 99, 2),
			Ssdsl:   randx.IntRange(100, 1000),
			Ssdzb:   randx.Float64Range(10, 99, 2),
		})
		for _, sfmc := range getCysfys()[cymc] {
			data = append(data, &AppSswcyQxgxkZyjsqk{
				Cymc:    cymc,
				Sfmc:    sfmc,
				Gxsl:    randx.IntRange(10, 100),
				Sylgxzb: randx.Float64Range(10, 99, 2),
				Bsdsl:   randx.IntRange(10, 100),
				Bsdzb:   randx.Float64Range(10, 99, 2),
				Ssdsl:   randx.IntRange(10, 100),
				Ssdzb:   randx.Float64Range(10, 99, 2),
			})
		}
	}
	return data
}

// 十四五产业-强相关学科-学生数量情况
type AppSswcyQxgxkXsslqk struct {
	Cymc      string  `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Sfmc      string  `json:"sfmc" gorm:"type:varchar(100); not null; comment:省份名称（分为全国和省份）;"`
	Bssl      int     `json:"bssl" gorm:"type:int; not null; comment:博士数量;"`
	Sylgxbszb float64 `json:"sylgxbszb" gorm:"type:decimal(4,2); not null; comment:双一流高校博士占比(百分比，保留两位小数);"`
	Sssl      int     `json:"sssl" gorm:"type:int; not null; comment:硕士数量;"`
	Sylgxsszb float64 `json:"sylgxsszb" gorm:"type:decimal(4,2); not null; comment:双一流高校硕士占比(百分比，保留两位小数);"`
}

// 表名称
func (t *AppSswcyQxgxkXsslqk) TableName() string {
	var table = "app_sswcy_qxgxk_xsslqk"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyQxgxkXsslqk) Comment() string {
	return "十四五产业-强相关学科-学生数量情况"
}

// 表数据
func (t *AppSswcyQxgxkXsslqk) InitData() any {
	var data []*AppSswcyQxgxkXsslqk
	for _, cymc := range cymcs {
		for _, sfmc := range getCysfys()[cymc] {
			data = append(data, &AppSswcyQxgxkXsslqk{
				Cymc:      cymc,
				Sfmc:      sfmc,
				Bssl:      randx.IntRange(1000, 2000),
				Sylgxbszb: randx.Float64Range(10, 99, 2),
				Sssl:      randx.IntRange(2000, 4000),
				Sylgxsszb: randx.Float64Range(10, 99, 2),
			})
		}
	}
	return data
}

// 十四五产业-强相关学科-就业情况
type AppSswcyQxgxkJyqk struct {
	Cymc        string `json:"cymc" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Sfmc        string `json:"sfmc" gorm:"type:varchar(100); not null; comment:省份名称（分为全国和省份）;"`
	Bsbyrs      int    `json:"bsbyrs" gorm:"type:int; not null; comment:本省毕业人数;"`
	Bsjyrs      int    `json:"bsjyrs" gorm:"type:int; not null; comment:本省就业人数;"`
	Bsbysbsjyrs int    `json:"bsbysbsjyrs" gorm:"type:int; not null; comment:本省毕业生本省就业人数;"`
	Wabysbsjyrs int    `json:"wabysbsjyrs" gorm:"type:int; not null; comment:外省毕业生本省就业人数;"`
}

// 表名称
func (t *AppSswcyQxgxkJyqk) TableName() string {
	var table = "app_sswcy_qxgxk_jyqk"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcyQxgxkJyqk) Comment() string {
	return "十四五产业-强相关学科-就业情况"
}

// 表数据
func (t *AppSswcyQxgxkJyqk) InitData() any {
	var data []*AppSswcyQxgxkJyqk
	for _, cymc := range cymcs {
		data = append(data, &AppSswcyQxgxkJyqk{
			Cymc:        cymc,
			Sfmc:        "全国",
			Bsbyrs:      randx.IntRange(1000, 2000),
			Bsjyrs:      randx.IntRange(1000, 2000),
			Bsbysbsjyrs: randx.IntRange(1000, 2000),
			Wabysbsjyrs: randx.IntRange(1000, 2000),
		})
		for _, sfmc := range getCysfys()[cymc] {
			data = append(data, &AppSswcyQxgxkJyqk{
				Cymc:        cymc,
				Sfmc:        sfmc,
				Bsbyrs:      randx.IntRange(1000, 2000),
				Bsjyrs:      randx.IntRange(1000, 2000),
				Bsbysbsjyrs: randx.IntRange(1000, 2000),
				Wabysbsjyrs: randx.IntRange(1000, 2000),
			})
		}
	}
	return data
}
