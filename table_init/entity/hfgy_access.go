package entity

import "time"

// 出入校园-超速事件统计
type OverSpeed struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	Times      int       `json:"times" gorm:"type:bigint; not null; default:0; comment:超速次数"`
}

func (OverSpeed) TableName() string {
	return "edu_app_xycr_cssjtj"
}

// 出入校园-人数统计
type PeopleStat struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	Total      int       `json:"total" gorm:"type:bigint; not null; default:0; comment:总人数"`
	Student    int       `json:"student" gorm:"type:bigint; not null; default:0; comment:在校生总人数"`
	Visitor    int       `json:"visitor" gorm:"type:bigint; not null; default:0; comment:访客数"`
}

func (PeopleStat) TableName() string {
	return "edu_app_xycr_rstj"
}

// 出入校园-出入地点
type AccessPlace struct {
	PlaceName string `json:"placeName" gorm:"type:varchar(100); not null; comment:出入地点;"`
	Longitude string `json:"longitude" gorm:"type:varchar(100); not null; default:0; comment:坐标经度"`
	Latitude  string `json:"latitude" gorm:"type:varchar(100); not null; default:0; comment:坐标纬度"`
}

func (AccessPlace) TableName() string {
	return "edu_app_xycr_crdd"
}

// 出入校园-出入地点统计
type AccessPlaceStat struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	PlaceName  string    `json:"placeName" gorm:"type:varchar(100); not null; comment:出入地点;"`
	PeopleIn   int       `json:"peopleIn" gorm:"type:bigint; not null; default:0; comment:入校人数"`
	PeopleOut  int       `json:"peopleOut" gorm:"type:bigint; not null; default:0; comment:出校人数"`
	CarIn      int       `json:"carIn" gorm:"type:bigint; not null; default:0; comment:入校车辆数"`
	CarOut     int       `json:"carOut" gorm:"type:bigint; not null; default:0; comment:出校车辆数"`
}

func (AccessPlaceStat) TableName() string {
	return "edu_app_xycr_crddtj"
}

// 出入校园-出入人数统计
type AccessPeopleTimes struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	PlaceName  string    `json:"placeName" gorm:"type:varchar(100); not null; comment:出入地点;"`
	Total      int       `json:"total" gorm:"type:bigint; not null; default:0; comment:总人数"`
	Student    int       `json:"student" gorm:"type:bigint; not null; default:0; comment:学生人数"`
	Teacher    int       `json:"teacher" gorm:"type:bigint; not null; default:0; comment:教师人数"`
}

func (AccessPeopleTimes) TableName() string {
	return "edu_app_xycr_crrstj"
}

// 出入校园-出入车数统计
type AccessCarTimes struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	Total      int       `json:"total" gorm:"type:bigint; not null; default:0; comment:总车辆数"`
	CarIn      int       `json:"carIn" gorm:"type:bigint; not null; default:0; comment:入校车辆数"`
	CarOut     int       `json:"carOut" gorm:"type:bigint; not null; default:0; comment:出校车辆数"`
}

func (AccessCarTimes) TableName() string {
	return "edu_app_xycr_crcstj"
}

// 出入校园-学院人数统计
type CollegePeopleStat struct {
	Id          int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	InsertDate  time.Time `json:"insertDate" gorm:"type:date; not null; comment:数据写入日期;"`
	CollegeName string    `json:"collegeName" gorm:"type:varchar(100); not null; comment:学院名称;"`
	Total       int       `json:"total" gorm:"type:bigint; not null; default:0; comment:学院总人数"`
	PeopleIn    int       `json:"peopleIn" gorm:"type:bigint; not null; default:0; comment:在校人数"`
	PeopleNotIn int       `json:"peopleNotIn" gorm:"type:bigint; not null; default:0; comment:不在校人数"`
}

func (CollegePeopleStat) TableName() string {
	return "edu_app_xycr_xyrstj"
}
