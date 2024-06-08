package entity

import (
	"github.com/go-xuan/pinyin"
	"github.com/go-xuan/quanx/types/intx"
	"github.com/go-xuan/quanx/types/stringx"
	"github.com/go-xuan/quanx/utils/idx"
	"github.com/go-xuan/quanx/utils/randx"
	"strconv"
	"strings"
	"time"
)

type ScreenIndicesStats struct {
	Id          int64     `json:"id" gorm:"type:bigint; comment:主键ID;"`
	Class       string    `json:"class" gorm:"type:varchar(100); not null; comment:指标分类;"`
	Key         string    `json:"key" gorm:"type:varchar(100); not null; comment:指标KEY;"`
	Name        string    `json:"name" gorm:"type:varchar(100); not null; comment:指标名称;"`
	Value       float64   `json:"value" gorm:"type:float4; comment:指标值;"`
	Unit        string    `json:"unit" gorm:"type:varchar(20); comment:指标单位;"`
	Index       int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	UpdateTime  time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	Desensitize bool      `json:"desensitize" gorm:"type:bool; default:false; comment:是否脱敏;"`
}

func (s ScreenIndicesStats) TableName() string {
	return "screen.screen_indices_stats"
}

func (s ScreenIndicesStats) TableComment() string {
	return "大屏各个指标数据统计"
}

func (s ScreenIndicesStats) InitData() any {
	var data []*ScreenIndicesStats
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "党委数", Value: 39, Unit: "个", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "党支部数", Value: 688, Unit: "个", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "党员数", Value: 15251, Unit: "人", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "预备党员数", Value: 1757, Unit: "人", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "积极份子数", Value: 4078, Unit: "人", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "乡村工作站数", Value: randx.Float64Range(10, 100, 0), Unit: "个", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "团支部数", Value: randx.Float64Range(100, 1000, 0), Unit: "个", Index: idx.Sequence().NextVal("党团建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "党团建设", Name: "团员数", Value: randx.Float64Range(1000, 2000, 0), Unit: "人", Index: idx.Sequence().NextVal("党团建设")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "在校生情况", Name: "在籍学生数", Value: 34132, Unit: "人", Index: idx.Sequence().NextVal("在校生情况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "在校生情况", Name: "留学生数", Value: 391, Unit: "人", Index: idx.Sequence().NextVal("在校生情况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "在校生情况", Name: "赴境外交流学生数", Value: randx.Float64Range(100, 1000, 0), Unit: "人", Index: idx.Sequence().NextVal("在校生情况")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "专业学院", Value: 28, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "学科门类", Value: 11, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "本科专业", Value: 96, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "国家一流本科专业建设点", Value: 57, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "基础学科拔尖计划2.0基地", Value: 8, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "强基计划", Value: 7, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "教学概况", Name: "博士、硕士授权一级学科", Value: 32, Unit: "个", Index: idx.Sequence().NextVal("教学概况")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "近五年科研项目数", Value: 6164, Unit: "个", Index: idx.Sequence().NextVal("科研情况-理科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "近五年到账科研经费金额", Value: 338734, Unit: "万元", Index: idx.Sequence().NextVal("科研情况-理科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "全国/国家重点实验室数量", Value: 5, Unit: "个", Index: idx.Sequence().NextVal("科研情况-理科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "自然科学类省部级理科科研平台数量", Value: 80, Unit: "个", Index: idx.Sequence().NextVal("科研情况-理科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "近5年科研成果获国家奖、教育部奖、天津市科技奖数", Value: 55, Unit: "个", Index: idx.Sequence().NextVal("科研情况-理科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "近五年SCI收录论文数", Value: 10701, Unit: "篇", Index: idx.Sequence().NextVal("科研情况-理科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-理科科研", Name: "近5年获得专利授权数", Value: 1907, Unit: "个", Index: idx.Sequence().NextVal("科研情况-理科科研")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "国家级科研平台数量", Value: 2, Unit: "个", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "省部级科研平台数量", Value: 50, Unit: "个", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "近五年到账科研经费数量", Value: randx.Float64Range(1000, 10000, 2), Unit: "万元", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "近五年国家级科研项目数量", Value: randx.Float64Range(10, 100, 0), Unit: "个", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "近五年省部级科研获奖数量", Value: randx.Float64Range(10, 100, 0), Unit: "个", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "近三年一类高质量中文学术论文数量", Value: randx.Float64Range(10, 100, 0), Unit: "篇", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "近三年两报一刊理论版文章数量", Value: 51, Unit: "篇", Index: idx.Sequence().NextVal("科研情况-文科科研")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "科研情况-文科科研", Name: "近三年重要智库成果数量", Value: randx.Float64Range(10, 100, 0), Unit: "个", Index: idx.Sequence().NextVal("科研情况-文科科研")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学科建设", Name: "进入ESI前1%的学科数", Value: 17, Unit: "个", Index: idx.Sequence().NextVal("学科建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学科建设", Name: "ESI排名前1‰学科", Value: 3, Unit: "个", Index: idx.Sequence().NextVal("学科建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学科建设", Name: "一级学科", Value: 42, Unit: "个", Index: idx.Sequence().NextVal("学科建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学科建设", Name: "一级学科国家重点学科", Value: 6, Unit: "个", Index: idx.Sequence().NextVal("学科建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学科建设", Name: "二级学科国家重点学科", Value: 9, Unit: "个", Index: idx.Sequence().NextVal("学科建设")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学科建设", Name: "一级学科天津市重点学科", Value: 32, Unit: "个", Index: idx.Sequence().NextVal("学科建设")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "专任教师总人数", Value: randx.Float64Range(1000, 2000, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "外籍教师人数", Value: randx.Float64Range(10, 100, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "国家级人才人数", Value: randx.Float64Range(100, 200, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "天津市人才", Value: randx.Float64Range(200, 300, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "博士后总人数", Value: randx.Float64Range(300, 500, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "学校‘百青’人才人数", Value: randx.Float64Range(100, 200, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "博士生导师数量 ", Value: randx.Float64Range(100, 200, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "硕士生导师数量", Value: randx.Float64Range(100, 200, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "学术任职人数", Value: randx.Float64Range(100, 200, 0), Unit: "人", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "博士学位教师占比", Value: randx.Float64Range(10, 100, 2), Unit: "%", Index: idx.Sequence().NextVal("师资队伍")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "师资队伍", Name: "45周岁以下青年教师占比", Value: randx.Float64Range(10, 100, 2), Unit: "%", Index: idx.Sequence().NextVal("师资队伍")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "专任教师职称分布", Name: "正高级职称", Value: 963, Unit: "人", Index: idx.Sequence().NextVal("专任教师职称分布")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "专任教师职称分布", Name: "副高级职称", Value: 901, Unit: "人", Index: idx.Sequence().NextVal("专任教师职称分布")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "专任教师职称分布", Name: "中级职称", Value: 403, Unit: "人", Index: idx.Sequence().NextVal("专任教师职称分布")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "专任教师职称分布", Name: "初级职称", Value: 10, Unit: "人", Index: idx.Sequence().NextVal("专任教师职称分布")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "专任教师职称分布", Name: "未定职级", Value: 15, Unit: "人", Index: idx.Sequence().NextVal("专任教师职称分布")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "学校固定资产原值", Value: 1142973.4, Unit: "万元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "设备资产总值", Value: randx.Float64Range(1000000, 2000000, 2), Unit: "万元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "大型仪器设备数", Value: randx.Float64Range(10000, 20000, 0), Unit: "个", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "图书数", Value: 6241405, Unit: "册", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "学校占地总面积", Value: 505.5, Unit: "万平方米", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "公共教室数量", Value: randx.Float64Range(5000, 1000, 0), Unit: "间", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "校舍建筑面积", Value: 218.58, Unit: "万平方米", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "上一年校级预算收入", Value: 28.2, Unit: "亿元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "上一年校级实际收入", Value: 27.38, Unit: "亿元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "上一年校级预算支出", Value: 32.84, Unit: "亿元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "上一年校级实际支出", Value: 30.59, Unit: "亿元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "上一年收支预算赤字", Value: -4.64, Unit: "亿元", Index: idx.Sequence().NextVal("财务资产")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "财务资产", Name: "上一年收支实际赤字", Value: -3.21, Unit: "亿元", Index: idx.Sequence().NextVal("财务资产")})

	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学校占地总面积", Name: "学校占地", Value: 505.5, Unit: "万平方米", Index: idx.Sequence().NextVal("学校占地总面积")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学校占地总面积", Name: "八里台校区占地", Value: 121.6, Unit: "万平方米", Index: idx.Sequence().NextVal("学校占地总面积")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学校占地总面积", Name: "津南校区占地", Value: 245.9, Unit: "万平方米", Index: idx.Sequence().NextVal("学校占地总面积")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学校占地总面积", Name: "泰达校区占地", Value: 6.7, Unit: "万平方米", Index: idx.Sequence().NextVal("学校占地总面积")})
	data = append(data, &ScreenIndicesStats{Id: idx.SnowFlake().Int64(), Class: "学校占地总面积", Name: "滨海学院占地", Value: 59.5, Unit: "万平方米", Index: idx.Sequence().NextVal("学校占地总面积")})

	for _, item := range data {
		item.Key = getKey(item.Name)
	}
	return data
}

type EnrollStudentsStats struct {
	Year            int       `json:"year" gorm:"type:int4; primary_key; comment:年份;"`
	Undergraduate   int       `json:"undergraduate" gorm:"type:int4; not null; comment:本科生;"`
	MasterStudent   int       `json:"masterStudent" gorm:"type:int4; not null; comment:硕士生;"`
	DoctoralStudent int       `json:"doctoralStudent" gorm:"type:int4; not null; comment:博士生;"`
	OverseasStudent int       `json:"overseasStudent" gorm:"type:int4; not null; comment:留学生;"`
	UpdateTime      time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
}

func (e EnrollStudentsStats) TableName() string {
	return "screen.enroll_students_stats"
}

func (e EnrollStudentsStats) TableComment() string {
	return "招生录取统计"
}

func (e EnrollStudentsStats) InitData() any {
	var data []*EnrollStudentsStats
	year := time.Now().Year()
	for i := 0; i < 5; i++ {
		u := randx.IntRange(10000, 20000)
		m := randx.IntRange(2000, 5000)
		d := randx.IntRange(1000, 2000)
		o := randx.IntRange(1000, 2000)
		data = append(data, &EnrollStudentsStats{
			Year:            year - i,
			Undergraduate:   u,
			MasterStudent:   m,
			DoctoralStudent: d,
			OverseasStudent: o,
		})
	}
	return data
}

type SubjectConstruction struct {
	Id         int64     `json:"id" gorm:"type:bigint; comment:主键ID;"`
	Name       string    `json:"name" gorm:"type:varchar(100); not null; comment:学科名称;"`
	Unit       string    `json:"unit" gorm:"type:varchar(100); not null; comment:建设单位;"`
	Index      int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
}

func (e SubjectConstruction) TableName() string {
	return "screen.subject_construction"
}

func (e SubjectConstruction) TableComment() string {
	return "学科建设表"
}

func (e SubjectConstruction) InitData() any {
	var data []*SubjectConstruction
	for i := 1; i <= 10; i++ {
		no := intx.String(i)
		data = append(data, &SubjectConstruction{
			Id:    idx.SnowFlake().Int64(),
			Name:  "学科名称" + no,
			Unit:  "建设单位" + no,
			Index: idx.Sequence().NextVal("学科建设表"),
		})
	}
	return data
}

type IndicesValue struct {
	Id          int64     `json:"id" gorm:"type:bigint; comment:主键ID;"`
	ClassKey    string    `json:"class_key" gorm:"type:varchar(100); not null; comment:指标分类Key;"`
	ClassName   string    `json:"class_name" gorm:"type:varchar(100); not null; comment:指标分类名称;"`
	TitleKey    string    `json:"titleKey" gorm:"type:varchar(100); not null; comment:标题Key;"`
	TitleName   string    `json:"titleName" gorm:"type:varchar(100); not null; comment:标题名称;"`
	Key         string    `json:"key" gorm:"type:varchar(100); not null; comment:指标KEY;"`
	Name        string    `json:"name" gorm:"type:varchar(100); not null; comment:指标名称;"`
	FullName    string    `json:"FullName" gorm:"type:varchar(100); not null; comment:指标完整名称;"`
	Value       string    `json:"value" gorm:"type:varchar(100); comment:指标值;"`
	Unit        string    `json:"unit" gorm:"type:varchar(20); comment:指标单位;"`
	Index       int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	Desensitize bool      `json:"desensitize" gorm:"type:bool; default:false; comment:是否脱敏;"`
	UpdateTime  time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
}

func (IndicesValue) TableName() string {
	return "screen.indices_value"
}

func (IndicesValue) TableComment() string {
	return "教学指标值"
}

func (IndicesValue) InitData() any {
	var data []*IndicesValue
	// 教学科研
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学科研", Name: "专业学院", Value: "28", Unit: "个", Index: idx.Sequence().NextVal("教学科研")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学科研", Name: "生师比", Value: "18.69", Unit: "", Index: idx.Sequence().NextVal("教学科研")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学科研", Name: "博士学位教师占比", Value: "93.53", Unit: "%", Index: idx.Sequence().NextVal("教学科研")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学科研", Name: "本科专业", Value: "96", Unit: "个", Index: idx.Sequence().NextVal("教学科研")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学科研", Name: "专任教师人数", Value: "2292", Unit: "人", Index: idx.Sequence().NextVal("教学科研")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学科研", Name: "高级职称专任教师数量", Value: "1864", Unit: "人", Index: idx.Sequence().NextVal("教学科研")})

	// 高级职称专任教师数量
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "高级职称专任教师数量", Name: "正高级职称", Value: "963", Unit: "人", Index: idx.Sequence().NextVal("高级职称专任教师数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "高级职称专任教师数量", Name: "副高级职称", Value: "901", Unit: "人", Index: idx.Sequence().NextVal("高级职称专任教师数量")})

	// 教学资源
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "思政课专任教师与折合在校生比例（≥1:350）", Value: "1:616.89", Unit: "", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "专职辅导员与在校生比例（≥1:200）", Value: "1:171.36", Unit: "", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "专职从事心理健康教育教师与在校生比例（≥1:4000且至少2名）", Value: "1:5502.5", Unit: "", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "专职就业指导教师和专职就业工作人员与应届毕业生比例（≥1:500）", Value: "1:486.75", Unit: "", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "省部级以上一流本科专业建设点", Value: "89", Unit: "个", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "国家级一流本科专业建设点", Value: "56", Unit: "个", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "基础学科拔尖计划2.0基地", Value: "8", Unit: "个", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "强基计划", Value: "7", Unit: "个", Index: idx.Sequence().NextVal("教学资源")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学资源", Name: "国家级一流本科课程", Value: "59", Unit: "门", Index: idx.Sequence().NextVal("教学资源")})

	// 教学运行
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学运行", Name: "开课课程门次", Value: "3522", Unit: "门次", Index: idx.Sequence().NextVal("教学运行")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学运行", Name: "教授主讲本科课程人均学时", Value: "86.72", Unit: "学时", Index: idx.Sequence().NextVal("教学运行")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学运行", Name: "教授主讲课程授课比率", Value: "32.94", Unit: "%", Index: idx.Sequence().NextVal("教学运行")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学运行", Name: "教授为本科生授课比率", Value: "87.24", Unit: "%", Index: idx.Sequence().NextVal("教学运行")})

	// 资源分布
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "资源分布", Name: "资源总数", Value: "", Unit: "个", Index: idx.Sequence().NextVal("资源分布")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "资源分布", Name: "教室总数量", Value: "288", Unit: "间", Index: idx.Sequence().NextVal("资源分布")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "资源分布", Name: "本科生生均课程门次", Value: "0.21", Unit: "门次", Index: idx.Sequence().NextVal("资源分布")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "资源分布", Name: "小班授课比例", Value: "50.03", Unit: "%", Index: idx.Sequence().NextVal("资源分布")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "资源分布", Name: "校应使用马工程重点教材课程数量的比例", Value: "100%", Unit: "%", Index: idx.Sequence().NextVal("资源分布")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "资源分布", Name: "教室使用率", Value: "100", Unit: "%", Index: idx.Sequence().NextVal("资源分布")})

	// 学生学习
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "本科生数量", Value: "16902", Unit: "人", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "本科生占总全日制在校生比例", Value: "51.19", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "课程通过率", Value: "98.5", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "学生毕业必须修满公共艺术课程学分数（≥2学分）", Value: "100", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "劳动教育必修课或必修课程中劳动教育模块学时总数（≥32学时） ", Value: "100", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "实践教学学分占总学分（学时）比例（人文社科类专业≥15%，理工农医类专业≥25%）", Value: "98.23", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "以实验、实习、工程实践和社会调查等实践性工作为基础的毕业论文（设计）比例（≥50%）", Value: "68.07", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "本科生体质测试达标率", Value: "85.38", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "本科生竞赛获奖人数", Value: "2268", Unit: "人", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "国家奖学金人数", Value: "204", Unit: "人", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "学位授予率", Value: "98.05", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "创新创业获奖项目", Value: "83", Unit: "项", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "本科生以第一作者获批国家发明专利数", Value: "4个", Unit: "个", Index: idx.Sequence().NextVal("学生学习")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "学生学习", Name: "参加国家级大学生创新创业训练计划项目学生比例", Value: "2.41", Unit: "%", Index: idx.Sequence().NextVal("学生学习")})

	// 各学院学生数量
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "化学学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "材料科学与工程学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "数学科学学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "统计与数据科学学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "生命科学学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "物理科学学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "环境科学与工程学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "电子信息与光学工程学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "人工智能学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "计算机学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "网络空间安全学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "医学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "药学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "软件学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "历史学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "经济学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "金融学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "商学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "旅游与服务学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "马克思主义学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "文学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "哲学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "周恩来政府管理学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "法学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "外国语学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "汉语言文化学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "新闻与传播学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "各学院学生数量", Name: "社会学院", Value: strconv.Itoa(randx.IntRange(200, 500)), Unit: "人", Index: idx.Sequence().NextVal("各学院学生数量")})

	// 教学评价
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学评价", Name: "学生评教参与率", Value: "93.12", Unit: "%", Index: idx.Sequence().NextVal("教学评价")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学评价", Name: "学生评价教师覆盖率", Value: "99.79", Unit: "%", Index: idx.Sequence().NextVal("教学评价")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学评价", Name: "学生评价课程覆盖率", Value: "99.44", Unit: "%", Index: idx.Sequence().NextVal("教学评价")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学评价", Name: "教师评价平均分", Value: "95.43", Unit: "分", Index: idx.Sequence().NextVal("教学评价")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学评价", Name: "课程评价平均分", Value: "95.66", Unit: "分", Index: idx.Sequence().NextVal("教学评价")})

	// 教学成果
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "教材建设数量", Value: "209", Unit: "本", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "教材获奖数量", Value: "2", Unit: "本", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "国家级教学成果奖数量", Value: "53", Unit: "项", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "省部级教学成果奖", Value: "163", Unit: "项", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "天津市教学成果奖", Value: "163", Unit: "项", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "霍英东青年教师基金与青年教师奖", Value: "9", Unit: "项", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "宝钢优秀教师奖", Value: "28", Unit: "项", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "天津市教学名师奖", Value: "46", Unit: "项", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "黄大年式教师团队", Value: "2", Unit: "个", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "天津市教学团队", Value: "22", Unit: "个", Index: idx.Sequence().NextVal("教学成果")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "教学成果", Name: "天津市优秀教师", Value: "11", Unit: "人", Index: idx.Sequence().NextVal("教学成果")})

	// 省部级以上学科数量获奖人次
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "省部级以上学科数量获奖人次", Key: "2021", Name: "2021年", Value: "1635", Unit: "人次", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "省部级以上学科数量获奖人次", Key: "2022", Name: "2022年", Value: "1131", Unit: "人次", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "省部级以上学科数量获奖人次", Key: "2023", Name: "2023年", Value: "1953", Unit: "人次", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})

	// 近五年建设教材数
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "近五年建设教材数", Key: "2019", Name: "2019", Value: strconv.Itoa(randx.IntRange(100, 200)), Unit: "本", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "近五年建设教材数", Key: "2020", Name: "2020", Value: strconv.Itoa(randx.IntRange(100, 200)), Unit: "本", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "近五年建设教材数", Key: "2021", Name: "2021", Value: strconv.Itoa(randx.IntRange(100, 200)), Unit: "本", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "近五年建设教材数", Key: "2022", Name: "2022", Value: strconv.Itoa(randx.IntRange(100, 200)), Unit: "本", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})
	data = append(data, &IndicesValue{Id: idx.SnowFlake().Int64(), TitleName: "近五年建设教材数", Key: "2023", Name: "2023", Value: strconv.Itoa(randx.IntRange(100, 200)), Unit: "本", Index: idx.Sequence().NextVal("省部级以上学科数量获奖人次")})

	for _, item := range data {
		item.ClassKey = "undergraduate"
		item.ClassName = "本科教学"
		if item.TitleKey == "" {
			item.TitleKey = getKey(item.TitleName)
		}
		if item.Key == "" {
			item.Key = getKey(item.Name)
		}
	}

	return data
}

type EvaluationCourseScore struct {
	Id         int64  `json:"id" gorm:"type:bigint; comment:主键ID;"`
	Course     string `json:"course" gorm:"type:varchar(100); not null; comment:课程名称;"`
	Department string `json:"college" gorm:"type:varchar(100); not null; comment:学院;"`
	Teacher    string `json:"teacher" gorm:"type:varchar(100); not null; comment:任课教师;"`
	Rate       string `json:"rate" gorm:"type:varchar(100); comment:参评率;"`
	Score      string `json:"score" gorm:"type:varchar(100); comment:得分;"`
	Rank       int    `json:"rank" gorm:"type:smallint; not null; comment:排名;"`
}

func (EvaluationCourseScore) TableName() string {
	return "screen.evaluation_course_score"
}

func (EvaluationCourseScore) TableComment() string {
	return "评价-课程得分"
}

func (EvaluationCourseScore) InitData() any {
	var data []*EvaluationCourseScore
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "新生研讨课", Department: "教务部", Teacher: "刘双喜", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "新生研讨课", Department: "教务部", Teacher: "阎萌", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "新生研讨课", Department: "教务部", Teacher: "张洁", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "新生研讨课", Department: "教务部", Teacher: "郑琳", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "新生研讨课", Department: "教务部", Teacher: "周培源", Rate: "91.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "文科高等数学", Department: "高等数学教学部", Teacher: "安桂梅", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高等数学（A类）I", Department: "高等数学教学部", Teacher: "张阳", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "运筹学", Department: "商学院", Teacher: "邢金刚", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "物流与供应链管理（全英文）", Department: "商学院", Teacher: "崔连广", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "实践教学", Department: "商学院", Teacher: "王晶", Rate: "80.49%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "计算机网络与通信技术", Department: "商学院", Teacher: "外聘教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "航运管理", Department: "商学院", Teacher: "杨静蕾", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "创新管理", Department: "商学院", Teacher: "杨坤", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "目录学", Department: "商学院", Teacher: "柯平", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "图书馆与档案馆管理", Department: "商学院", Teacher: "李樵", Rate: "80.95%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "技术经济", Department: "商学院", Teacher: "隋静", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "企业战略管理综合案例分析", Department: "商学院", Teacher: "柳茂平", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "变革管理", Department: "商学院", Teacher: "田莉", Rate: "77.78%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "营销与创新保护", Department: "商学院", Teacher: "田莉", Rate: "85.71%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "认识实习", Department: "化学学院", Teacher: "指导教师", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "分子科学与工程信息学", Department: "化学学院", Teacher: "指导教师", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "膜科学与技术", Department: "化学学院", Teacher: "指导教师", Rate: "33.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生物化工导论", Department: "化学学院", Teacher: "指导教师", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高分子化学", Department: "化学学院", Teacher: "指导教师", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生产实习", Department: "化学学院", Teacher: "指导教师", Rate: "37.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "综合化学化工实验（2）", Department: "化学学院", Teacher: "指导教师", Rate: "37.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "化工机械基础", Department: "化学学院", Teacher: "指导教师", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "化工热力学", Department: "化学学院", Teacher: "指导教师", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "绿色化学工艺学", Department: "化学学院", Teacher: "指导教师", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "化工传质与分离过程课程设计", Department: "化学学院", Teacher: "指导教师", Rate: "37.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高等有机化学", Department: "化学学院", Teacher: "李庆山", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "微纳米材料与技术导论", Department: "化学学院", Teacher: "指导教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "新型分离技术", Department: "化学学院", Teacher: "指导教师", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "分子识别与组装", Department: "化学学院", Teacher: "蔡康", Rate: "80.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "分子识别与组装", Department: "化学学院", Teacher: "张瀛溟", Rate: "77.78%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "精细化学品与高新技术", Department: "化学学院", Teacher: "指导教师", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "化工传质与分离过程", Department: "化学学院", Teacher: "指导教师", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "电气工程学概论（1）", Department: "化学学院", Teacher: "指导教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "管理概论", Department: "化学学院", Teacher: "指导教师", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "海外科研实习", Department: "化学学院", Teacher: "王佰全", Rate: "33.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "农药生物学", Department: "化学学院", Teacher: "范志金", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "有机合成化学", Department: "化学学院", Teacher: "王忠文", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "聚合物胶体", Department: "化学学院", Teacher: "赵汉英", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生物医用材料导论", Department: "化学学院", Teacher: "袁直", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代高分子化学", Department: "化学学院", Teacher: "刘丽", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "博士生文献报告", Department: "化学学院", Teacher: "张弛", Rate: "90.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "有机结构分析(全英文)", Department: "化学学院", Teacher: "孔祥蕾", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "有机结构分析(全英文)", Department: "化学学院", Teacher: "王志宏", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "化学反应动力学（全英文）", Department: "化学学院", Teacher: "孔祥蕾", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "人工智能在化学化工中的应用", Department: "南开天大合办专业", Teacher: "指导教师", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "多尺度化学工程与技术", Department: "南开天大合办专业", Teacher: "指导教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "计算催化反应工程（全英文）", Department: "南开天大合办专业", Teacher: "指导教师", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级汉语口语2-1", Department: "汉语言文化学院", Teacher: "史建伟", Rate: "20.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中国概况2-2", Department: "汉语言文化学院", Teacher: "白宏钟", Rate: "20.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "Internet应用技术", Department: "汉语言文化学院", Teacher: "刘晓红", Rate: "20.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级综合汉语2-1", Department: "汉语言文化学院", Teacher: "吴星云", Rate: "20.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "学年论文", Department: "汉语言文化学院", Teacher: "吴星云", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "报刊选读2-2", Department: "汉语言文化学院", Teacher: "王红厂", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "虚词研究", Department: "汉语言文化学院", Teacher: "郭继懋", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级汉语听力2-1", Department: "汉语言文化学院", Teacher: "史建伟", Rate: "20.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "韩国语（二外）2-1", Department: "汉语言文化学院", Teacher: "桂香", Rate: "63.64%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "教育心理学", Department: "汉语言文化学院", Teacher: "夏全胜", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "论文写作", Department: "汉语言文化学院", Teacher: "郭继懋", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "第二语言习得研究", Department: "汉语言文化学院", Teacher: "温宝莹", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "语义学", Department: "汉语言文化学院", Teacher: "王毓钧", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "社会语言学", Department: "汉语言文化学院", Teacher: "王吉辉", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "语言对比与偏误分析", Department: "汉语言文化学院", Teacher: "梁磊", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "劳动法", Department: "网络空间安全学院", Teacher: "柯振兴", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "公共选择经济学", Department: "经济学院", Teacher: "高嵩", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中国经济史2-2", Department: "经济学院", Teacher: "龚关", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "国际金融实务", Department: "经济学院", Teacher: "梁琪", Rate: "53.85%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代经济学前沿专题（2-1）", Department: "经济学院", Teacher: "金威", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代经济学前沿专题（2-1）", Department: "经济学院", Teacher: "李俊青", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代经济学前沿专题（2-1）", Department: "经济学院", Teacher: "毛其淋", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "网络经济学", Department: "经济学院", Teacher: "王金杰", Rate: "77.78%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "R语言", Department: "经济学院", Teacher: "外聘教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "计量经济学2", Department: "经济学院", Teacher: "王丽莉", Rate: "76.47%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光谱学", Department: "电子信息与光学工程学院", Teacher: "王湘晖", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "PLC原理与实验", Department: "电子信息与光学工程学院", Teacher: "张颖", Rate: "90.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光通信实验", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "工程光学实验（1）", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "83.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "电子工艺实习", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "81.82%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生物医学光子学（双语）", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光纤光学", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光学工艺与检测", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "22.22%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光纤通信及系统", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光学系统设计", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "机械工程训练", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光电子激光实验", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光电信息专业实验", Department: "电子信息与光学工程学院", Teacher: "刘晓颀", Rate: "60.53%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光纤传感技术及应用", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "固体物理", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "半导体物理", Department: "电子信息与光学工程学院", Teacher: "天大教师", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生态地学", Department: "环境科学与工程学院", Teacher: "鲍艳宇", Rate: "70.59%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境影响评价与环境规划", Department: "环境科学与工程学院", Teacher: "徐鹤", Rate: "80.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境生物学实验", Department: "环境科学与工程学院", Teacher: "李尧", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境生物学实验", Department: "环境科学与工程学院", Teacher: "张寅清", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境化学", Department: "环境科学与工程学院", Teacher: "王婷", Rate: "63.64%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境化学", Department: "环境科学与工程学院", Teacher: "姚义鸣", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "陈威", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "冯银厂", Rate: "70.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "刘东方", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "刘璐", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "毛洪钧", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "王玉秋", Rate: "70.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "张彤", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "周启星", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境研究", Department: "环境科学与工程学院", Teacher: "祝凌燕", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境数学模型", Department: "环境科学与工程学院", Teacher: "张丽", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境工程学", Department: "环境科学与工程学院", Teacher: "郭晓燕", Rate: "70.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境工程学", Department: "环境科学与工程学院", Teacher: "李凤祥", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境工程学", Department: "环境科学与工程学院", Teacher: "王鑫", Rate: "53.85%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境工程学实验", Department: "环境科学与工程学院", Teacher: "郭晓燕", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境工程学实验", Department: "环境科学与工程学院", Teacher: "李凤祥", Rate: "62.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境工程学实验", Department: "环境科学与工程学院", Teacher: "翟利芳", Rate: "70.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "实践教学", Department: "环境科学与工程学院", Teacher: "卢会霞", Rate: "60.87%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "专业实践", Department: "环境科学与工程学院", Teacher: "齐宇", Rate: "57.14%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "专业实践", Department: "环境科学与工程学院", Teacher: "王军锋", Rate: "57.14%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "专业实践", Department: "环境科学与工程学院", Teacher: "张墨", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "电工与电子基础", Department: "环境科学与工程学院", Teacher: "展思辉", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代生命科学导论", Department: "环境科学与工程学院", Teacher: "张雷", Rate: "90.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境流行病学", Department: "环境科学与工程学院", Teacher: "张敏英", Rate: "83.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境生态工程", Department: "环境科学与工程学院", Teacher: "卢学强", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生态规划与设计", Department: "环境科学与工程学院", Teacher: "李洪远", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "毕业实习", Department: "环境科学与工程学院", Teacher: "刘春光", Rate: "44.44%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境生态工程实验", Department: "环境科学与工程学院", Teacher: "卢会霞", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "环境污染与修复", Department: "环境科学与工程学院", Teacher: "姜传佳", Rate: "76.47%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "仪器分析实验", Department: "环境科学与工程学院", Teacher: "卢媛", Rate: "33.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "仪器分析实验", Department: "环境科学与工程学院", Teacher: "杨楠", Rate: "33.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "恢复生态学", Department: "环境科学与工程学院", Teacher: "鲍艳宇", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "恢复生态学", Department: "环境科学与工程学院", Teacher: "李田", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "污染场地调查与修复", Department: "环境科学与工程学院", Teacher: "李田", Rate: "80.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "污染场地调查与修复", Department: "环境科学与工程学院", Teacher: "刘维涛", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "大气环境化学", Department: "环境科学与工程学院", Teacher: "彭剑飞", Rate: "62.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "大气环境化学", Department: "环境科学与工程学院", Teacher: "宋少洁", Rate: "62.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "实 验 投 资 学", Department: "金融学院", Teacher: "刘晓峰", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "海上保险（双语）", Department: "金融学院", Teacher: "刘玮", Rate: "93.75%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "汉西口译", Department: "外国语学院", Teacher: "宓田", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "经贸西班牙语", Department: "外国语学院", Teacher: "RAYDIS MOISES FRANCO GUERRA", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中国与西语国家关系", Department: "外国语学院", Teacher: "李兴华", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "旅游西班牙语", Department: "外国语学院", Teacher: "李兴华", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "语言学导论", Department: "外国语学院", Teacher: "李玮琦", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "西班牙语写作2-1", Department: "外国语学院", Teacher: "OLGARELYS DUENAS MENDOZA", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "翻译理论与实践（汉译法）", Department: "外国语学院", Teacher: "李澜雪", Rate: "78.95%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级西班牙语3-1", Department: "外国语学院", Teacher: "梁嘉艺", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级俄语视听说3-1", Department: "外国语学院", Teacher: "BELIKOV SERGEI", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级法语精读3-1", Department: "外国语学院", Teacher: "贺梦莹", Rate: "84.62%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "营销学原理（英）", Department: "外国语学院", Teacher: "李晶", Rate: "79.17%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "德语国家概况2-1", Department: "外国语学院", Teacher: "薄一荻", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "英语演讲与辩论", Department: "外国语学院", Teacher: "回春萍", Rate: "91.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级俄语3-1", Department: "外国语学院", Teacher: "博尔多诺娃", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级俄语3-1", Department: "外国语学院", Teacher: "王丽丹", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "德语泛读7-4", Department: "外国语学院", Teacher: "托马斯", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "俄罗斯文学史", Department: "外国语学院", Teacher: "姜敏", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "翻译(俄译汉)", Department: "外国语学院", Teacher: "叶芳芳", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "创新研究与训练", Department: "外国语学院", Teacher: "李玮琦", Rate: "91.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "创新研究与训练", Department: "外国语学院", Teacher: "任明丽", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "基础俄语4-3", Department: "外国语学院", Teacher: "赵春梅", Rate: "88.24%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "英语文学赏析", Department: "外国语学院", Teacher: "张强", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "当代俄罗斯社会", Department: "外国语学院", Teacher: "安妮索娃", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "初级德语精读2-1", Department: "外国语学院", Teacher: "马迎彬", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "初级德语精读2-1", Department: "外国语学院", Teacher: "远思", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中级德语精读2-1", Department: "外国语学院", Teacher: "韩捷", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "西班牙语口语3-2", Department: "外国语学院", Teacher: "OLGARELYS DUENAS MENDOZA", Rate: "85.71%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "西班牙语视听说3-3", Department: "外国语学院", Teacher: "OLGARELYS DUENAS MENDOZA", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "日语精读4-1", Department: "外国语学院", Teacher: "东孝拓", Rate: "96.55%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "日语精读4-1", Department: "外国语学院", Teacher: "王灿娟", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "日语写作2-1", Department: "外国语学院", Teacher: "MATSUOKA YUYA", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "德语综合训练1", Department: "外国语学院", Teacher: "雷海花", Rate: "95.24%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "德语综合训练3", Department: "外国语学院", Teacher: "雷海花", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中葡口译", Department: "外国语学院", Teacher: "王程序", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "基础意大利语4-1", Department: "外国语学院", Teacher: "乐小悦", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "基础意大利语4-1", Department: "外国语学院", Teacher: "杨琳", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "基础意大利语4-3", Department: "外国语学院", Teacher: "杨琳", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "意大利语国家概况2-1", Department: "外国语学院", Teacher: "石豆", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "意大利语报刊选读", Department: "外国语学院", Teacher: "乐小悦", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "意大利语基础语法", Department: "外国语学院", Teacher: "倪杨", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "经贸意大利语", Department: "外国语学院", Teacher: "阮玉凤", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "西方文明概览2-1", Department: "外国语学院", Teacher: "刘英", Rate: "83.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "英美当代文学初读2-1", Department: "外国语学院", Teacher: "黄宗贤", Rate: "84.62%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "意大利语口语3-2", Department: "外国语学院", Teacher: "乐小悦", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "意大利语阅读4-1", Department: "外国语学院", Teacher: "彭倩", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "跨文化交际", Department: "外国语学院", Teacher: "乐小悦", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级葡萄牙语精读3-3", Department: "外国语学院", Teacher: "Diana Maia", Rate: "72.73%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "二外法语4—3", Department: "外国语学院", Teacher: "贺梦莹", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "二外俄语4-3", Department: "外国语学院", Teacher: "许力", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "二外西班牙语4-3", Department: "外国语学院", Teacher: "肖音", Rate: "50.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "二外意大利语4-3", Department: "外国语学院", Teacher: "阮玉凤", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "二外阿拉伯语4-1", Department: "外国语学院", Teacher: "AHMED EBRAHIM ELDESOUKY ABDELRAOUF ELSAYED", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "德语翻译研究", Department: "外国语学院", Teacher: "施大山", Rate: "78.57%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "阿拉伯语视听说高级2-1", Department: "外国语学院", Teacher: "AHMED EBRAHIM ELDESOUKY ABDELRAOUF ELSAYED", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级阿拉伯语3-3", Department: "外国语学院", Teacher: "苏楚婷", Rate: "42.86%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "经贸阿拉伯语", Department: "外国语学院", Teacher: "苏楚婷", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "外文戏剧", Department: "外国语学院", Teacher: "谷佳维", Rate: "85.71%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "意大利文学作品选读2-1", Department: "外国语学院", Teacher: "ANDREA SARTORI", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "电影和文学中的当代意大利", Department: "外国语学院", Teacher: "ANDREA SARTORI", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "老年学概论", Department: "周恩来政府管理学院", Teacher: "胡雯", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "健康心理学", Department: "周恩来政府管理学院", Teacher: "王崇颖", Rate: "86.96%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "发展人类学", Department: "周恩来政府管理学院", Teacher: "朱健刚", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "希腊语", Department: "历史学院", Teacher: "武鹏", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中国古代陶瓷器", Department: "历史学院", Teacher: "刘毅", Rate: "90.48%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "明史", Department: "历史学院", Teacher: "马子木", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "拉丁语", Department: "历史学院", Teacher: "叶民", Rate: "72.73%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "世界上古中古史2-1", Department: "历史学院", Teacher: "邵大路", Rate: "87.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "科技考古概论", Department: "历史学院", Teacher: "张国文", Rate: "93.75%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "宋代制度史", Department: "历史学院", Teacher: "曹杰", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "出版物设计与制作", Department: "新闻与传播学院", Teacher: "张树楠", Rate: "68.42%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "数字出版创意与策划", Department: "新闻与传播学院", Teacher: "梁小建", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "竞争法", Department: "经济学院", Teacher: "孙炜", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "免疫学", Department: "生命科学学院", Teacher: "吴震州", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级英文生物化学", Department: "生命科学学院", Teacher: "Adam Midgley", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高级英文生物化学", Department: "生命科学学院", Teacher: "周卫红", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "科研训练3-1", Department: "生命科学学院", Teacher: "贡红日", Rate: "88.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "科研训练3-3", Department: "生命科学学院", Teacher: "赵绍伟", Rate: "35.71%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "遗传学基础研讨", Department: "生命科学学院", Teacher: "朱正茂", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生命科学前沿研讨", Department: "生命科学学院", Teacher: "王恺", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "山水临摹2-1", Department: "文学院", Teacher: "柳潜", Rate: "90.91%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "山水临摹2-2", Department: "文学院", Teacher: "杜森", Rate: "87.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "山水临摹2-2", Department: "文学院", Teacher: "柳潜", Rate: "83.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代中国文学2-1", Department: "文学院", Teacher: "罗维斯", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中国文化通识2-1（留学生）", Department: "文学院", Teacher: "田桂民", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "书法4-3", Department: "文学院", Teacher: "王劼", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "工笔人物写生2-2", Department: "文学院", Teacher: "李晓石", Rate: "80.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "工笔人物写生2-2", Department: "文学院", Teacher: "张旺", Rate: "87.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "汉语写作2-1（留学生）", Department: "文学院", Teacher: "孙易", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "现代汉语2-1（留学生）", Department: "文学院", Teacher: "孙倩", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "汉语精读2-1（留学生）", Department: "文学院", Teacher: "马尚", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "说唱文学研究与创作", Department: "文学院", Teacher: "鲍震培", Rate: "85.71%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "清代小说史研究", Department: "文学院", Teacher: "陈宏", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "语言科学前沿研究赏析", Department: "文学院", Teacher: "冉启斌", Rate: "87.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "西方政治思想史", Department: "马克思主义学院", Teacher: "付洪", Rate: "80.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "中国共产党与中国现代化", Department: "马克思主义学院", Teacher: "纪亚光", Rate: "87.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "教育哲学", Department: "马克思主义学院", Teacher: "任铃", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "抽象函数与巴拿赫代数", Department: "数学科学学院", Teacher: "刘锐", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "组合学中的代数方法", Department: "数学科学学院", Teacher: "王星炜", Rate: "55.56%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "图同态与其应用", Department: "数学科学学院", Teacher: "王周宁馨", Rate: "80.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "专业实习", Department: "医学院", Teacher: "杨宝成", Rate: "41.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "临床视光学", Department: "医学院", Teacher: "王雁", Rate: "81.82%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "体育理论与训练7", Department: "体育部", Teacher: "指导教师", Rate: "41.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生物药物研究与应用", Department: "药学院", Teacher: "王玲", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生化分析与分子诊断", Department: "药学院", Teacher: "孟萌", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "干细胞与遗传药学研究", Department: "药学院", Teacher: "帅领", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "生物统计基础", Department: "药学院", Teacher: "刘祥", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "马克思主义经典作家论逻辑", Department: "哲学院", Teacher: "郎需瑞", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "西方哲学原著选读（全英文）", Department: "哲学院", Teacher: "张俊国", Rate: "71.43%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "海外中国哲学研究（全英文）", Department: "哲学院", Teacher: "Misha Andrew Tadd", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "材料制备技术", Department: "物理科学学院", Teacher: "余华", Rate: "81.82%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "高等物理实验", Department: "物理科学学院", Teacher: "陈靖", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "光子学与技术", Department: "物理科学学院", Teacher: "李玉栋", Rate: "42.86%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "固体物理(二)", Department: "物理科学学院", Teacher: "曹学伟", Rate: "60.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "粒子物理", Department: "物理科学学院", Teacher: "杨茂志", Rate: "66.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "非参数统计", Department: "统计与数据科学学院", Teacher: "刘进", Rate: "75.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "会展媒体与传播", Department: "旅游与服务学院", Teacher: "苏醒", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "服务学习：中华诗教的传承与实践", Department: "文学院", Teacher: "张静", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "服务学习：营养中国", Department: "医学院", Teacher: "刘寅", Rate: "86.67%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "描述集合论", Department: "数学科学学院", Teacher: "丁龙云", Rate: "83.33%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "马克思主义哲学前沿问题研究", Department: "哲学院", Teacher: "侯振武", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "马克思主义哲学前沿问题研究", Department: "哲学院", Teacher: "王南湜", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "马克思主义哲学前沿问题研究", Department: "哲学院", Teacher: "夏钊", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "专利基础与情报分析", Department: "教务处", Teacher: "李佳", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "专利基础与情报分析", Department: "教务处", Teacher: "周静", Rate: "88.89%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "初级葡萄牙语2-1", Department: "公共外语教学部", Teacher: "王程序", Rate: "87.50%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "拟阵基础", Department: "数学科学学院", Teacher: "杨立波", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	data = append(data, &EvaluationCourseScore{Id: idx.SnowFlake().Int64(), Course: "服务学习：社区、文化与发展", Department: "周恩来政府管理学院", Teacher: "朱健刚", Rate: "100.00%", Score: "100", Rank: idx.Sequence().NextVal("评价分TOP10课程列表")})
	return data
}

type EvaluationDepartmentRate struct {
	Id         int64  `json:"id" gorm:"type:bigint; comment:主键ID;"`
	Department string `json:"name" gorm:"type:varchar(100); not null; comment:院系名称;"`
	Rate       string `json:"value" gorm:"type:varchar(100); comment:参评率;"`
	Rank       int    `json:"rank" gorm:"type:smallint; not null; comment:排名;"`
}

func (EvaluationDepartmentRate) TableName() string {
	return "screen.evaluation_department_rate"
}

func (EvaluationDepartmentRate) TableComment() string {
	return "评价-各院系参评率"
}

func (EvaluationDepartmentRate) InitData() any {
	var data []*EvaluationDepartmentRate

	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "工商管理类", Rate: "99.40%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "中国语言文学类", Rate: "98.42%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "工科试验班", Rate: "98.24%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "理科试验班（物理与光电信息技术工程）", Rate: "97.67%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "历史学院", Rate: "97.22%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "理科试验班（数理科学与大数据）", Rate: "97.13%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "理科试验班（物质科学与可持续发展）", Rate: "97.08%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "商学院", Rate: "96.82%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "法学院", Rate: "96.51%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "生命科学学院", Rate: "95.61%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "周恩来政府管理学院", Rate: "95.38%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "哲学院", Rate: "95.17%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "国际教育学院", Rate: "94.92%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "计算机学院", Rate: "94.87%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "医学院", Rate: "94.34%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "统计与数据科学学院", Rate: "94.23%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "网络空间安全学院", Rate: "94.03%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "外国语学院", Rate: "93.94%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "新闻与传播学院", Rate: "93.53%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "金融学院", Rate: "92.90%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "人工智能学院", Rate: "92.48%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "数学科学学院", Rate: "92.23%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "软件学院", Rate: "91.94%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "马克思主义学院", Rate: "91.56%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "文学院", Rate: "91.56%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "经济学院", Rate: "91.30%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "材料科学与工程学院", Rate: "89.73%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "化学学院", Rate: "89.53%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "电子信息与光学工程学院", Rate: "88.00%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "物理科学学院", Rate: "85.56%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "药学院", Rate: "84.10%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "汉语言文化学院", Rate: "82.09%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "旅游与服务学院", Rate: "82.08%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "环境科学与工程学院", Rate: "81.89%", Rank: idx.Sequence().NextVal("各院系参评率")})
	data = append(data, &EvaluationDepartmentRate{Id: idx.SnowFlake().Int64(), Department: "交换生", Rate: "21.43%", Rank: idx.Sequence().NextVal("各院系参评率")})
	return data
}

type EvaluationTeacherRatio struct {
	Id         int64  `json:"id" gorm:"type:bigint; comment:主键ID;"`
	ScoreRange string `json:"name" gorm:"type:varchar(100); not null; comment:得分范围;"`
	Number     int    `json:"number" gorm:"type:int; comment:人数;"`
	Ratio      string `json:"value" gorm:"type:varchar(100); comment:占比;"`
	Rank       int    `json:"rank" gorm:"type:smallint; not null; comment:排名;"`
}

func (EvaluationTeacherRatio) TableName() string {
	return "screen.evaluation_teacher_ratio"
}

func (EvaluationTeacherRatio) TableComment() string {
	return "评价-各分数段教师占比"
}

func (EvaluationTeacherRatio) InitData() any {
	var data []*EvaluationTeacherRatio
	data = append(data, &EvaluationTeacherRatio{Id: idx.SnowFlake().Int64(), ScoreRange: "95-100", Ratio: "67.07%", Rank: idx.Sequence().NextVal("各分数段教师占比")})
	data = append(data, &EvaluationTeacherRatio{Id: idx.SnowFlake().Int64(), ScoreRange: "90-95", Ratio: "25.79%", Rank: idx.Sequence().NextVal("各分数段教师占比")})
	data = append(data, &EvaluationTeacherRatio{Id: idx.SnowFlake().Int64(), ScoreRange: "80-90", Ratio: "6.38%", Rank: idx.Sequence().NextVal("各分数段教师占比")})
	data = append(data, &EvaluationTeacherRatio{Id: idx.SnowFlake().Int64(), ScoreRange: "60-80", Ratio: "0.70%", Rank: idx.Sequence().NextVal("各分数段教师占比")})
	data = append(data, &EvaluationTeacherRatio{Id: idx.SnowFlake().Int64(), ScoreRange: "60以下", Ratio: "0.06%", Rank: idx.Sequence().NextVal("各分数段教师占比")})
	return data
}

func getKey(name string) string {
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, "、", "")
	if i := stringx.Index(name, "比例"); i > 0 {
		name = name[:i+len("比例")]
	}
	if i := stringx.Index(name, "（≥"); i > 0 {
		name = name[:i]
	}
	if j, k := stringx.Between(name, "（", "）"); j > 0 {
		name = name[:j-len("（")] + name[j:k] + name[k+len("）"):]
	}
	if j, k := stringx.Between(name, "‘", "’"); j > 0 {
		name = name[:j-len("‘")] + name[j:k] + name[k+len("’"):]
	}
	py := pinyin.NewPinyin(name).Fmt(pinyin.NoTone).Convert()
	pys := strings.Split(py, " ")
	var stm = map[string]string{
		"ESI":   "ESI",
		"1%":    "bfzy",
		"1‰":    "qfzy",
		"SCI":   "SCI",
		"5":     "5",
		"2.0":   "2d0",
		"45":    "45",
		"TOP10": "TOP10",
	}
	var key string
	for _, p := range pys {
		if v, ok := stm[p]; ok {
			key = key + v
		} else if len(p) > 1 {
			key = key + p[:1]
		}
	}
	return key
}
