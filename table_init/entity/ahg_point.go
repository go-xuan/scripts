package entity

import "time"

// 用户信息
type User struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:用户ID;"`
	Name       string    `json:"userName" gorm:"type:varchar(100); not null; comment:用户姓名;"`
	Account    string    `json:"Account" gorm:"type:varchar(50); not null; comment:用户账号;"`
	UserType   string    `json:"userType" gorm:"type:varchar(10); comment:用户类型（MANAGE系统用户）;"`
	Phone      string    `json:"phone" gorm:"type:varchar(20); comment:手机;"`
	IdCard     string    `json:"idCard" gorm:"type:varchar(20); comment:身份证号;"`
	Email      string    `json:"email" gorm:"type:varchar(100); comment:邮箱;"`
	Gender     int       `json:"gender" gorm:"type:smallint; comment:性别（0未知 1男 2女）;"`
	Avatar     string    `json:"avatar" gorm:"type:varchar(500); comment:头像地址;"`
	Birthday   string    `json:"birthday" gorm:"type:date; comment:生日;"`
	Level      int       `json:"level" gorm:"type:int; comment:等级;"`
	Password   string    `json:"password" gorm:"type:varchar(100); not null; comment:密码;"`
	Status     int       `json:"status" gorm:"type:smallint; comment:帐号状态（1正常 2停用）;"`
	LoginIp    string    `json:"loginIp" gorm:"type:varchar(255); comment:最后登录IP;"`
	LoginDate  time.Time `json:"loginDate" gorm:"type:timestamp; comment:最后登录时间;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp; default:now();comment:创建时间;"`
	CreateBy   string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy   string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
	Remark     string    `json:"remark" gorm:"type:varchar(255); comment:备注;"`
}

func (User) TableName() string {
	return "sys_user"
}

func (u User) Comment() string {
	return "用户信息表"
}

func (u User) InitData() any {
	return nil
}

// 登录日志表
type LoginLog struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:日志ID;"`
	Account    int64     `json:"account" gorm:"type:varchar(20); comment:用户账号（手机号）;"`
	Status     string    `json:"status" gorm:"type:varchar(20); comment:登录状态;"`
	Content    string    `json:"content" gorm:"type:text; comment:日志内容;"`
	Ip         string    `json:"ip" gorm:"type:varchar(100); comment:当前IP;"`
	Address    string    `json:"address" gorm:"type:varchar(100); comment:地点;"`
	Browser    string    `json:"browser" gorm:"type:varchar(100); comment:浏览器;"`
	Os         string    `json:"os" gorm:"type:varchar(100); comment:操作系统;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateBy   string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
}

func (LoginLog) TableName() string {
	return "sys_login_log"
}

func (u LoginLog) Comment() string {
	return "登录日志表"
}

func (u LoginLog) InitData() any {
	return nil
}

// api日志表
type ApiLog struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:日志ID;"`
	UserId       int64     `json:"userId" gorm:"type:bigint; not null; comment:用户ID;"`
	UserName     string    `json:"userName" gorm:"type:varchar(100); not null; comment:用户名称;"`
	Module       string    `json:"module" gorm:"type:varchar(100); comment:操作模块;"`
	BusinessType string    `json:"businessType" gorm:"type:varchar(100); comment:功能类型;"`
	UserType     string    `json:"userType" gorm:"type:varchar(100); comment:用户类型;"`
	Status       string    `json:"status" gorm:"type:varchar(20); comment:状态;"`
	Content      string    `json:"content" gorm:"type:text; comment:日志内容;"`
	Method       string    `json:"method" gorm:"type:varchar(10); comment:请求方式;"`
	Url          string    `json:"url" gorm:"type:varchar(1000); comment:请求URL;"`
	Param        string    `json:"param" gorm:"type:text; comment:请求参数;"`
	Result       string    `json:"result" gorm:"type:text; comment:请求结果;"`
	Ip           string    `json:"ip" gorm:"type:varchar(100); comment:当前IP;"`
	Address      string    `json:"address" gorm:"type:varchar(100); comment:地点;"`
	Browser      string    `json:"browser" gorm:"type:varchar(100); comment:浏览器;"`
	Os           string    `json:"os" gorm:"type:varchar(100); comment:操作系统;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
}

func (ApiLog) TableName() string {
	return "sys_api_log"
}

func (u ApiLog) Comment() string {
	return "api日志表"
}

func (u ApiLog) InitData() any {
	return nil
}

// 用户积分记录表
type PointRecord struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	UserId       int64     `json:"userId" gorm:"type:bigint; not null; comment:用户ID;"`
	UserName     string    `json:"userName" gorm:"type:varchar(100); not null; comment:用户名称;"`
	ActivityName string    `json:"activityName" gorm:"type:varchar(100); not null; comment:活动名称;"`
	PointRule    string    `json:"pointRule" gorm:"type:varchar(100); not null; comment:积分规则名称;"`
	Duration     string    `json:"duration" gorm:"type:varchar(100); comment:时长;"`
	ChangePoint  int       `json:"changePoint" gorm:"type:int; not null; comment:积分变动;"`
	BeforePoint  int       `json:"beforePoint" gorm:"type:int; not null; comment:变动前积分;"`
	AfterPoint   int       `json:"afterPoint" gorm:"type:int; not null; comment:变动后积分;"`
	StartTime    time.Time `json:"startTime" gorm:"type:timestamp(0); comment:打卡开始时间;"`
	EndTime      time.Time `json:"endTime" gorm:"type:timestamp(0); comment:打卡结束时间;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateBy     string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy     string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (PointRecord) TableName() string {
	return "ahg_point_record"
}

func (a PointRecord) Comment() string {
	return "用户积分记录表"
}

func (a PointRecord) InitData() any {
	return nil
}

// 用户积分规则表
type PointRule struct {
	Id          int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:规则ID;"`
	Name        string    `json:"name" gorm:"type:varchar(100); not null; comment:规则名;"`
	Each        string    `json:"each" gorm:"type:varchar(10); not null; comment:每（次/小时）;"`
	Point       int       `json:"point" gorm:"type:int; not null; comment:得分;"`
	Description string    `json:"description" gorm:"type:varchar(1000); not null; comment:规则说明;"`
	CreateTime  time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	UpdateBy    string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
	UpdateTime  time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	CreateBy    string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
}

func (PointRule) TableName() string {
	return "ahg_point_rule"
}

func (a PointRule) Comment() string {
	return "用户积分规则表"
}

func (a PointRule) InitData() any {
	return nil
}

// 用户积分等级表
type Level struct {
	Id         int       `json:"id" gorm:"type:int; not null; primary_key; comment:等级ID;"`
	Name       string    `json:"name" gorm:"type:varchar(100); not null; comment:等级名;"`
	PointMin   int       `json:"pointMin" gorm:"type:int; not null; comment:积分下限;"`
	PointMax   int       `json:"pointMax" gorm:"type:int; not null; comment:积分上限;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateBy   string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy   string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (Level) TableName() string {
	return "ahg_point_level"
}

func (a Level) Comment() string {
	return "用户积分等级表"
}

func (a Level) InitData() any {
	return nil
}

// 用户推荐表
type Recommendation struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:推荐ID;"`
	GoodsId      string    `json:"goodsId" gorm:"type:varchar(100); not null; comment:商品ID;"`
	GoodsName    string    `json:"goodsName" gorm:"type:varchar(100); not null; comment:商品名称;"`
	GoodsPrice   int       `json:"goodsPrice" gorm:"type:int; not null; comment:商品价格;"`
	GoodsPicture int       `json:"goodsPicture" gorm:"type:varchar(2000); comment:商品图片;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateBy     string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy     string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (Recommendation) TableName() string {
	return "ahg_recommendation"
}

func (a Recommendation) Comment() string {
	return "用户推荐表"
}

func (a Recommendation) InitData() any {
	return nil
}

// 用户权益表
type Benefit struct {
	Id               int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:权益ID;"`
	BenefitType      string    `json:"benefitType" gorm:"type:varchar(100); not null; comment:权益内容类型;"`
	BenefitContent   string    `json:"benefitContent" gorm:"type:varchar(100); not null; comment:权益内容（积分/商品）;"`
	BenefitPoint     int       `json:"benefitPoint" gorm:"type:int; comment:权益积分;"`
	BenefitGoodsId   int       `json:"benefitGoodsId" gorm:"type:bigint; comment:权益商品ID;"`
	BenefitGoodsName string    `json:"benefitGoodsName" gorm:"type:varchar(100); comment:权益商品名称;"`
	Status           bool      `json:"status" gorm:"type:bool; not null; comment:状态;"`
	Description      string    `json:"description" gorm:"type:varchar(1000); not null; comment:说明;"`
	CreateTime       time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateBy         string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime       time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy         string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (Benefit) TableName() string {
	return "ahg_benefit"
}

func (a Benefit) Comment() string {
	return "用户权益表"
}

func (a Benefit) InitData() any {
	return nil
}

// 用户权益对象表
type BenefitObject struct {
	Id        int64  `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	BenefitId int64  `json:"benefitId" gorm:"type:bigint; not null; comment:权益ID;"`
	ObjId     int64  `json:"memberId" gorm:"type:bigint; not null; comment:对象ID;"`
	ObjName   string `json:"memberName" gorm:"type:varchar(100); not null; comment:对象名称;"`
}

func (BenefitObject) TableName() string {
	return "ahg_benefit_object"
}

func (a BenefitObject) Comment() string {
	return "用户权益对象表"
}

func (a BenefitObject) InitData() any {
	return nil
}
