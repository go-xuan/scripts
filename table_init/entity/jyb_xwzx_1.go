package entity

import (
	"strings"

	"github.com/go-xuan/quanx/utils/randx"
)

var zcqks = []string{"支撑一般", "基本支撑", "支撑不足"}
var cymcs = []string{
	"新材料产业", "新一代信息技术产业", "高端设备制造产业", "新能源产业",
	"节能环保产业", "新能源(智能网联)汽车产业", "生命健康产业", "高端纺织产业",
	"化工产业", "新兴数字产业", "航空航天产业", "集成电路产业",
}

var schema string

var xkzys = []string{
	"地理学", "生态学", "中药",
	"信息与通信工程", "计算机科学与技术", "电子科学与技术",
	"生物医学工程", "材料科学与工程", "石油与天然气工程",
	"纺织科学与工程", "轻工技术与工程", "船舶与海洋工程",
	"食品科学与工程", "畜牧学", "公共卫生与预防医学",
}

// 十四五产业-省域分布
type AppSswcySyfb struct {
	Cymc string `json:"cymc" echart:"lgd_data" gorm:"type:varchar(100); not null; comment:产业名称;"`
	Sfsl int    `json:"sfsl" echart:"data" gorm:"type:int; not null; default:0; comment:省份数量;"`
	Pxxb int    `json:"pxxb" gorm:"type:smallint; not null; default:1; comment:排序下标;"`
}

// 表名称
func (t *AppSswcySyfb) TableName() string {
	var table = "app_sswcy_syfb"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (t *AppSswcySyfb) Comment() string {
	return "十四五产业-省域分布"

}

// 表数据
func (t *AppSswcySyfb) InitData() any {
	var data []*AppSswcySyfb
	for i, cymc := range cymcs {
		data = append(data, &AppSswcySyfb{
			Cymc: cymc,
			Sfsl: randx.IntRange(1, 30),
			Pxxb: i + 1,
		})
	}
	return data
}

// 十四五产业-省域产业数量统计
type AppSswcySysltj struct {
	Sycyzs int `json:"sycyzs" gorm:"type:int; not null; default:0; comment:省域产业数;"`
	Glcys  int `json:"glcys" gorm:"type:int; not null; default:0; comment:归类产业数;"`
}

// 表名称
func (AppSswcySysltj) TableName() string {
	var table = "app_sswcy_sysltj"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (AppSswcySysltj) Comment() string {
	return "十四五产业-省域产业数量统计"
}

// 表数据
func (AppSswcySysltj) InitData() any {
	var data []*AppSswcySysltj
	data = append(data, &AppSswcySysltj{
		Sycyzs: randx.IntRange(1, 300),
		Glcys:  randx.IntRange(1, 100),
	})
	return data
}

// 十四五产业-省域产业数量分布
type AppSswcySyslfb struct {
	Sfmc string `json:"sfmc" gorm:"type:varchar(100); not null; comment:省份名称;"`
	Cysl int    `json:"cysl" gorm:"type:int; not null; default:0; comment:产业数量;"`
}

// 表名称
func (AppSswcySyslfb) TableName() string {
	var table = "app_sswcy_syslfb"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (AppSswcySyslfb) Comment() string {
	return "十四五产业-省域产业数量分布"
}

// 表数据
func (AppSswcySyslfb) InitData() any {
	var data []*AppSswcySyslfb
	var sfmcs = strings.Split(randx.ProvinceName, ",")
	for _, sfmc := range sfmcs {
		data = append(data, &AppSswcySyslfb{
			Sfmc: sfmc,
			Cysl: randx.IntRange(1, 100),
		})
	}
	return data
}

// 十四五产业-支撑情况-各省支撑度
type AppSswcyZcqkGszcd struct {
	Sfmc string  `json:"sfmc" gorm:"type:varchar(100); not null; comment:省份名称;"`
	Zcd  float64 `json:"zcd" gorm:"type:decimal(4,2); comment:支撑度(保留两位小数);"`
	Cysl int     `json:"cysl" gorm:"type:int; not null; default:0; comment:产业数量;"`
}

// 表名称
func (AppSswcyZcqkGszcd) TableName() string {
	var table = "app_sswcy_zcqk_gszcd"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (AppSswcyZcqkGszcd) Comment() string {
	return "十四五产业-支撑情况-各省支撑度"
}

// 表数据
func (AppSswcyZcqkGszcd) InitData() any {
	var data []*AppSswcyZcqkGszcd
	var sfmcs = strings.Split(randx.ProvinceName, ",")
	for _, sfmc := range sfmcs {
		data = append(data, &AppSswcyZcqkGszcd{
			Sfmc: sfmc,
			Zcd:  randx.Float64Range(0, 1, 2),
			Cysl: randx.IntRange(1, 100),
		})
	}
	return data
}

// 十四五产业-研究生教育-综合支撑度
type AppSswcyZcqkZhzcd struct {
	Zhzcd float64 `json:"zhzcd" gorm:"type:decimal(4,2); comment:综合支撑度(保留两位小数);"`
}

// 表名称
func (AppSswcyZcqkZhzcd) TableName() string {
	var table = "app_sswcy_zcqk_zhzcd"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (AppSswcyZcqkZhzcd) Comment() string {
	return "十四五产业-支撑情况-综合支撑度"
}

// 表数据
func (AppSswcyZcqkZhzcd) InitData() any {
	var data []*AppSswcyZcqkZhzcd
	data = append(data, &AppSswcyZcqkZhzcd{Zhzcd: randx.Float64Range(0.5, 1, 2)})
	return data
}

// 十四五产业-支撑情况-统计
type AppSswcyZcqkZcqktj struct {
	Zcqk  string  `json:"zcqk" gorm:"type:varchar(100); not null; comment:支撑情况(支撑一般/基本支撑/支撑不足);"`
	Sfsl  int     `json:"sfsl" gorm:"type:int; not null; default:0; comment:省份数量;"`
	Gdpzb float64 `json:"gdpzb" gorm:"type:decimal(4,2); comment:GDP占比(保留两位小数);"`
	Zcdzb float64 `json:"zcdzb" gorm:"type:decimal(4,2); comment:支撑度占比(保留两位小数);"`
}

// 表名称
func (AppSswcyZcqkZcqktj) TableName() string {
	var table = "app_sswcy_zcqk_tj"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (AppSswcyZcqkZcqktj) Comment() string {
	return "十四五产业-支撑情况-统计"
}

// 表数据
func (AppSswcyZcqkZcqktj) InitData() any {
	var data []*AppSswcyZcqkZcqktj
	for _, zcqk := range zcqks {
		data = append(data, &AppSswcyZcqkZcqktj{
			Zcqk:  zcqk,
			Sfsl:  randx.IntRange(1, 100),
			Gdpzb: randx.Float64Range(0, 1, 2),
			Zcdzb: randx.Float64Range(0, 1, 2),
		})
	}
	return data
}

// 十四五产业-核心支撑学科专业
type AppSswcyHxzcxkzy struct {
	Xkzymc string  `json:"xkzymc" gorm:"type:varchar(100); not null; comment:学科专业名称;"`
	Zccysl int     `json:"zccysl" gorm:"type:int; not null; default:0; comment:支撑产业数量;"`
	Rd     float64 `json:"rd" gorm:"type:decimal(4,2); comment:热度(保留两位小数);"`
}

// 表名称
func (AppSswcyHxzcxkzy) TableName() string {
	var table = "app_sswcy_hxzcxkzy"
	if schema != "" {
		table = schema + "." + table
	}
	return table
}

// 表备注
func (AppSswcyHxzcxkzy) Comment() string {
	return "十四五产业-核心支撑学科专业"
}

// 表数据
func (AppSswcyHxzcxkzy) InitData() any {
	var data []*AppSswcyHxzcxkzy
	for _, xkzy := range xkzys {
		data = append(data, &AppSswcyHxzcxkzy{
			Xkzymc: xkzy,
			Zccysl: randx.IntRange(1, 100),
			Rd:     randx.Float64Range(1, 100, 2),
		})
	}
	return data
}
