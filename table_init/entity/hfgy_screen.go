package entity

import "time"

// 物业监管-巡检类型统计
type InspectionTypeStat struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	TypeName   string    `json:"typeName" gorm:"type:varchar(100); not null; comment:巡检类型名称;"`
	Times      int       `json:"times" gorm:"type:bigint; not null; default:0; comment:巡检次数"`
}

func (InspectionTypeStat) TableName() string {
	return "edu_app_wyjg_xjlxtj"
}

// 物业监管-巡检处理统计
type InspectionDealStat struct {
	Id          int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate  time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	XjTotal     int       `json:"xjTotal" gorm:"type:bigint; not null; default:0; comment:巡检总数;"`
	XjReplyNum  int       `json:"xjReplyNum" gorm:"type:bigint; not null; default:0; comment:巡检回复数;"`
	XjAdoptNum  int       `json:"XjAdoptNum" gorm:"type:bigint; not null; default:0; comment:巡检采纳数"`
	XjRejectNum int       `json:"XjRejectNum" gorm:"type:bigint; not null; default:0; comment:巡检驳回数"`
	XcTotal     int       `json:"xcTotal" gorm:"type:bigint; not null; default:0; comment:巡查总数;"`
	XcReplyNum  int       `json:"xcReplyNum" gorm:"type:bigint; not null; default:0; comment:巡查回复数;"`
	XcAdoptNum  int       `json:"xcAdoptNum" gorm:"type:bigint; not null; default:0; comment:巡查采纳数"`
	XcRejectNum int       `json:"xcRejectNum" gorm:"type:bigint; not null; default:0; comment:巡查驳回数"`
}

func (InspectionDealStat) TableName() string {
	return "edu_app_wyjg_xjcltj"
}

// 物业监管-巡检地点统计
type InspectionPlaceStat struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	PlaceName  string    `json:"placeName" gorm:"type:varchar(100); not null; comment:巡检地点名称;"`
	Times      int       `json:"times" gorm:"type:bigint; not null; default:0; comment:巡检次数"`
}

func (InspectionPlaceStat) TableName() string {
	return "edu_app_wyjg_xjddtj"
}

// 物业监管-巡检扣分详情
type InspectionDeductDetail struct {
	Id          int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate  time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	PlaceName   string    `json:"placeName" gorm:"type:varchar(100); not null; comment:巡检地点名称;"`
	DeductPoint float64   `json:"deductPoint" gorm:"type:decimal(10,1); not null; comment:扣分;"`
	DeductTme   time.Time `json:"deductTme" gorm:"type:timestamp(0); default:now(); comment:扣分时间;"`
}

func (InspectionDeductDetail) TableName() string {
	return "edu_app_wyjg_xjkfxq"
}

// 物业监管-巡检次数统计
type InspectionTimesStat struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	Year       int       `json:"year" gorm:"type:smallint; not null; default:2023; comment:年;"`
	Month      int       `json:"month" gorm:"type:smallint; not null; default:1; comment:月;"`
	Day        int       `json:"day" gorm:"type:smallint; not null; default:1; comment:日;"`
	XjTimes    int       `json:"xjTimes" gorm:"type:bigint; not null; default:0; comment:巡检次数;"`
	XcTimes    int       `json:"xcTimes" gorm:"type:bigint; not null; default:0; comment:巡查次数;"`
}

func (InspectionTimesStat) TableName() string {
	return "edu_app_wyjg_xjcstj"
}
