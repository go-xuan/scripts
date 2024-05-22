package entity

import "time"

type SchoolData struct {
	Teacher             int     `json:"teacher" gorm:"type:int; not null; comment:教职工数;"`
	Student             int     `json:"student" gorm:"type:int; not null; comment:学生数;"`
	Major               int     `json:"major" gorm:"type:int; not null; comment:开设专业数;"`
	Textbook            int     `json:"textbook" gorm:"type:int; not null; comment:选用教材总数;"`
	Class               int     `json:"class" gorm:"type:int; not null; comment:班级数;"`
	StudentTeacherRatio float64 `json:"studentTeacherRatio" gorm:"type:decimal(10,2); not null; comment:生师比;"`
}

func (d SchoolData) TableComment() string {
	return "学校数据表"
}

func (d SchoolData) InitData() any {
	return nil
}

func (SchoolData) TableName() string {
	return "screen_new.t_school_data"
}

type SchoolAssetData struct {
	SchoolArea                float64 `json:"schoolArea" gorm:"type:decimal(10,2); not null; comment:校园占地面积;"`
	SchoolAreaAvg             float64 `json:"schoolAreaAvg" gorm:"type:decimal(10,2); not null; comment:生均用地;"`
	SchoolHouseArea           float64 `json:"schoolHouseArea" gorm:"type:decimal(10,2); not null; comment:校舍建筑面积;"`
	SchoolHouseAreaAvg        float64 `json:"schoolHouseAreaAvg" gorm:"type:decimal(10,2); not null; comment:生均校舍建筑面积;"`
	TeachingEquipmentValue    float64 `json:"teachingEquipmentValue" gorm:"type:decimal(10,2); not null; comment:教学仪器设备总值;"`
	TeachingEquipmentValueAvg float64 `json:"teachingEquipmentValueAvg" gorm:"type:decimal(10,2); not null; comment:生均教学仪器设备值;"`
	Books                     int     `json:"books" gorm:"type:int; not null; comment:图书册数;"`
	BooksAvg                  float64 `json:"booksAvg" gorm:"type:decimal(10,2); not null; comment:生均图书;"`
}

func (d SchoolAssetData) TableComment() string {
	return "学校资产表"
}

func (d SchoolAssetData) InitData() any {
	return nil
}

func (SchoolAssetData) TableName() string {
	return "screen_new.t_school_asset_data"
}

type TeachingData struct {
	Course                          int     `json:"course" gorm:"type:int; not null; comment:开设课程数;"`
	OnlineCourse                    int     `json:"onlineCourse" gorm:"type:int; not null; comment:上网课程数;"`
	SchoolTeachingEquipmentValue    float64 `json:"schoolTeachingEquipmentValue" gorm:"type:decimal(10,2); not null; comment:校内实训仪器设备值;"`
	SchoolTeachingWorkstations      int     `json:"schoolTeachingWorkstations" gorm:"type:int; not null; comment:校内实训实践工位数;"`
	ProfessionalTeachingResourceLib int     `json:"professionalTeachingResourceLib" gorm:"type:int; not null; comment:专业教学资源库数;"`
}

func (d TeachingData) TableComment() string {
	return "教学数据表"
}

func (d TeachingData) InitData() any {
	return nil
}

func (TeachingData) TableName() string {
	return "screen_new.t_teaching_data"
}

type InternshipData struct {
	InternshipBases       int     `json:"internshipBases" gorm:"type:int; not null; comment:校外实习基地数;"`
	InsideProvince        int     `json:"insideProvince" gorm:"type:int; not null; comment:省内实习生数;"`
	OutsideProvince       int     `json:"outsideProvince" gorm:"type:int; not null; comment:省外实习生数;"`
	EmploymentRate        float64 `json:"employmentRate" gorm:"type:decimal(10,2); not null; comment:就业率;"`
	MatchEmploymentRate   float64 `json:"matchEmploymentRate" gorm:"type:decimal(10,2); not null; comment:对口就业率;"`
	PurchaseInsuranceRate float64 `json:"purchaseInsuranceRate" gorm:"type:decimal(10,2); not null; comment:购买保险率;"`
	AverageSalary         float64 `json:"averageSalary" gorm:"type:decimal(10,2); not null; comment:平均薪酬;"`
}

func (d InternshipData) TableComment() string {
	return "实训数据表"
}

func (d InternshipData) InitData() any {
	return nil
}

func (InternshipData) TableName() string {
	return "screen_new.t_internship_data"
}

type TeacherData struct {
	Total             int `json:"total" gorm:"type:int; not null; comment:教师总数;"`
	DoubleAbility     int `json:"doubleAbility" gorm:"type:int; not null; comment:双师;"`
	FullTime          int `json:"fullTime" gorm:"type:int; not null; comment:专任教师;"`
	Professional      int `json:"professional" gorm:"type:int; not null; comment:专业教师;"`
	FullSeniorTitle   int `json:"fullSeniorTitle" gorm:"type:int; not null; comment:正高级职称;"`
	SeniorTitle       int `json:"seniorTitle" gorm:"type:int; not null; comment:高级职称;"`
	IntermediateTitle int `json:"intermediateTitle" gorm:"type:int; not null; comment:中级职称;"`
	PrimaryTitle      int `json:"primaryTitle" gorm:"type:int; not null; comment:初级职称;"`
	NoTitle           int `json:"noTitle" gorm:"type:int; not null; comment:无职称;"`
	MasterDegree      int `json:"masterDegree" gorm:"type:int; not null; comment:硕士学历;"`
	BachelorDegree    int `json:"bachelorDegree" gorm:"type:int; not null; comment:本科学历;"`
	CollegeDegree     int `json:"collegeDegree" gorm:"type:int; not null; comment:大专学历;"`
}

func (d TeacherData) TableComment() string {
	return "教师数据表"
}

func (d TeacherData) InitData() any {
	return nil
}

func (TeacherData) TableName() string {
	return "screen_new.t_teacher_data"
}

type StudentData struct {
	Recruit          int `json:"recruit" gorm:"type:int; not null; comment:招生人数;"`
	Graduate         int `json:"graduate" gorm:"type:int; not null; comment:毕业人数;"`
	EnterHigherGrade int `json:"enterHigherGrade" gorm:"type:int; not null; comment:升学人数;"`
	Employment       int `json:"employment" gorm:"type:int; not null; comment:就业人数;"`
}

func (d StudentData) TableComment() string {
	return "学生数据表"
}

func (d StudentData) InitData() any {
	return nil
}

func (StudentData) TableName() string {
	return "screen_new.t_student_data"
}

type StudentGradeData struct {
	Grade             int     `json:"grade" gorm:"type:int; not null; comment:年级;"`
	Recruit           int     `json:"recruit" gorm:"type:int; not null; comment:招生人数;"`
	Current           int     `json:"current" gorm:"type:int; not null; comment:当前人数;"`
	ConsolidationRate float64 `json:"consolidationRate" gorm:"type:decimal(10,2); not null; comment:巩固率;"`
}

func (d StudentGradeData) TableComment() string {
	return "学生年级表"
}

func (d StudentGradeData) InitData() any {
	return nil
}

func (StudentGradeData) TableName() string {
	return "screen_new.t_student_grade_data"
}

type CooperationData struct {
	Enterprise int `json:"enterprise" gorm:"type:int; not null; comment:校企合作企业数;"`
	Course     int `json:"course" gorm:"type:int; not null; comment:校企合作开发课程数;"`
	Textbook   int `json:"textbook" gorm:"type:int; not null; comment:校企合作开发教材数;"`
	People     int `json:"people" gorm:"type:int; not null; comment:校企合作订单培养人数;"`
}

func (d CooperationData) TableComment() string {
	return "校企合作表"
}

func (d CooperationData) InitData() any {
	return nil
}

func (CooperationData) TableName() string {
	return "screen_new.t_cooperation_data"
}

type AchievementData struct {
	Patent int `json:"patent" gorm:"type:int; not null; comment:专利;"`
	Paper  int `json:"paper" gorm:"type:int; not null; comment:论文;"`
}

func (d AchievementData) TableComment() string {
	return "成就数据表"
}

func (d AchievementData) InitData() any {
	return nil
}

func (AchievementData) TableName() string {
	return "screen_new.t_achievement_data"
}

type AwardData struct {
	Class     string    `json:"class" gorm:"type:varchar(100); not null; comment:获奖类别;"`
	Name      string    `json:"name" gorm:"type:varchar(100); not null; comment:获奖名称;"`
	Awards    string    `json:"awards" gorm:"type:varchar(100); not null; comment:奖项;"`
	Awardee   string    `json:"awardee" gorm:"type:varchar(100); not null; comment:获奖人;"`
	AwardTime time.Time `json:"awardTime" gorm:"type:timestamp(0); default:now(); comment:获奖时间;"`
}

func (d AwardData) TableComment() string {
	return "获奖数据表"
}

func (d AwardData) InitData() any {
	return nil
}

func (AwardData) TableName() string {
	return "screen_new.t_award_data"
}

type HonorData struct {
	Name       string    `json:"name" gorm:"type:varchar(100); not null; comment:获奖名称;"`
	Index      int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
}

func (d HonorData) TableComment() string {
	return "荣誉数据表"
}

func (d HonorData) InitData() any {
	return nil
}

func (HonorData) TableName() string {
	return "screen_new.t_honor_data"
}

// 基本办学指标
type BasicSchoolValue struct {
	Name         string  `json:"name" gorm:"type:varchar(100); not null; comment:指标名称;"`
	Unit         string  `json:"unit" gorm:"type:varchar(20); comment:指标单位;"`
	SchoolValue  float64 `json:"schoolValue" gorm:"type:decimal(10,2); comment:学校值;"`
	WarningValue float64 `json:"warningValue" gorm:"type:decimal(10,2); comment:预警值;"`
	Index        int     `json:"index" gorm:"type:smallint; comment:排序下标;"`
}

func (v BasicSchoolValue) InitData() any {
	var data = []*BasicSchoolValue{
		{Name: "专任教师数", Unit: "人", SchoolValue: 596.00, WarningValue: 60.00, Index: 1},
		{Name: "校园占地面积", Unit: "㎡", SchoolValue: 446973.25, WarningValue: 40000.00, Index: 2},
		{Name: "生均用地面积", Unit: "㎡", SchoolValue: 34.49, WarningValue: 33.00, Index: 3},
		{Name: "校舍建筑面积", Unit: "㎡", SchoolValue: 378441.89, WarningValue: 24000.00, Index: 4},
		{Name: "生均校舍建筑面积", Unit: "㎡", SchoolValue: 29.21, WarningValue: 20.00, Index: 5},
		{Name: "师生比", Unit: ":", SchoolValue: 19.92, WarningValue: 20.00, Index: 6},
		{Name: "仪器设备值", Unit: "万元", SchoolValue: 14139.96, WarningValue: 300.00, Index: 7},
		{Name: "生均仪器设备值", Unit: "元", SchoolValue: 10912.15, WarningValue: 2500.00, Index: 8},
		{Name: "生均图书", Unit: "册", SchoolValue: 32.89, WarningValue: 30.00, Index: 9},
		{Name: "中职学历教育在校生数", Unit: "人", SchoolValue: 12934.00, WarningValue: 3600.00, Index: 10},
	}
	return data
}

func (BasicSchoolValue) TableName() string {
	return "screen_new.t_basic_school_value"
}

// 表备注
func (BasicSchoolValue) TableComment() string {
	return "基本办学指标表"
}
