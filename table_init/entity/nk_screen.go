package entity

import (
	"github.com/go-xuan/pinyin"
	"github.com/go-xuan/quanx/types/intx"
	"github.com/go-xuan/quanx/utils/idx"
	"github.com/go-xuan/quanx/utils/randx"
	"strings"
	"time"
)

type ScreenIndicesStats struct {
	Id          int64     `json:"id" gorm:"type:bigint; comment:主键ID;"`
	Class       string    `json:"class" gorm:"type:varchar(100); not null; comment:指标分类;"`
	Name        string    `json:"name" gorm:"type:varchar(100); not null; comment:指标名称;"`
	Key         string    `json:"key" gorm:"type:varchar(100); not null; comment:指标KEY;"`
	Value       float64   `json:"value" gorm:"type:float4; comment:指标值;"`
	Unit        string    `json:"unit" gorm:"type:varchar(20); comment:指标单位;"`
	Index       int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	UpdateTime  time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	Desensitize bool      `json:"desensitize" gorm:"type:bool; default:false; comment:是否脱敏;"`
}

func (s ScreenIndicesStats) TableName() string {
	return "nkdx.screen_indices_stats"
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
		py := pinyin.NewPinyin(item.Name).Fmt(pinyin.NoTone).Convert()
		pys := strings.Split(py, " ")
		var key string
		for _, p := range pys {
			if v, ok := stm[p]; ok {
				key = key + v
			} else if len(p) > 1 {
				key = key + p[:1]
			}
		}
		item.Key = key
	}
	return data
}

var stm = map[string]string{
	"ESI": "ESI",
	"1%":  "1%",
	"SCI": "SCI",
	"5":   "5",
	"/":   "/",
	"、":   "/",
	"‘":   "",
	"’":   "",
	"2.0": "2.0",
	"45":  "45",
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
	return "nkdx.enroll_students_stats"
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
	return "nkdx.subject_construction"
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
