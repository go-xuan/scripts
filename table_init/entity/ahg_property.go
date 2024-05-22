package entity

import "time"

type Staff struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:人员ID;"`
	Name         string    `json:"name" gorm:"type:varchar(100); not null; comment:人员姓名;"`
	Contact      string    `json:"contact" gorm:"type:varchar(20); comment:联系方式;"`
	Sex          int       `json:"sex" gorm:"type:smallint;not null; comment:性别（0未知 1男 2女）;"`
	IdCard       string    `json:"idCard" gorm:"type:varchar(20);not null; comment:身份证号;"`
	Birthday     string    `json:"birthday" gorm:"type:date; comment:生日;"`
	Nation       string    `json:"nation" gorm:"type:varchar(500); comment:民族;"`
	DepartmentId int       `json:"departmentId" gorm:"type:bigint; not null; comment:部门;"`
	PostId       int       `json:"postId" gorm:"type:bigint; comment:岗位;"`
	EntryTime    time.Time `json:"entryTime" gorm:"type:timestamp; default:now();comment:入职时间;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp; default:now();comment:创建时间;"`
	CreateBy     string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy     string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (s Staff) TableName() string {
	return "property_staff"
}

func (s Staff) Comment() string {
	return "物业管理-员工信息表"
}

func (s Staff) InitData() any {

	return nil
}

type Department struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:部门ID;"`
	Name       string    `json:"name" gorm:"type:varchar(100); not null; comment:部门名称;"`
	Pid        int64     `json:"contact" gorm:"type:bigint; default:0; comment:父部门ID;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp; default:now();comment:创建时间;"`
	CreateBy   string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy   string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (d Department) TableName() string {
	return "property_department"
}

func (d Department) Comment() string {
	return "物业管理-部门信息表"
}

func (d Department) InitData() any {
	return nil
}

type Post struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:岗位ID;"`
	Code         string    `json:"code" gorm:"type:varchar(100); not null; comment:岗位编号;"`
	Name         string    `json:"name" gorm:"type:varchar(100); not null; comment:岗位名称;"`
	DepartmentId int       `json:"departmentId" gorm:"type:bigint; not null; comment:部门;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注说明;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp; default:now();comment:创建时间;"`
	CreateBy     string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy     string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (p Post) TableName() string {
	return "property_post"
}

func (p Post) Comment() string {
	return "物业管理-岗位信息表"
}

func (p Post) InitData() any {
	return nil
}

type Equipment struct {
	Id              int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:设备ID;"`
	Type            int       `json:"type" gorm:"type:smallint; not null; default:1; comment:设备类型（1：消防设备，2：保洁设备）;"`
	Code            string    `json:"code" gorm:"type:varchar(100); not null; comment:设备编号;"`
	Name            string    `json:"name" gorm:"type:varchar(100); not null; comment:设备名称;"`
	PurchasePrice   float64   `json:"purchasePrice" gorm:"type:decimal(10,2); default:0; comment:采购价格;"`
	PurchaseTime    time.Time `json:"purchaseTime" gorm:"type:timestamp; default:now(); comment:采购时间;"`
	Manager         string    `json:"manager" gorm:"type:varchar(100); comment:设备负责人;"`
	Supplier        string    `json:"supplier" gorm:"type:varchar(100); comment:供应商;"`
	Installer       string    `json:"Installer" gorm:"type:varchar(100); comment:安装人员;"`
	InstallPosition string    `json:"installPosition" gorm:"type:varchar(1000); comment:设备安装位置;"`
	InstallTime     time.Time `json:"installTime" gorm:"type:timestamp; default:now(); comment:设备安装时间;"`
	Status          string    `json:"status" gorm:"type:varchar(100); comment:状态编号;"`
	StatusSource    string    `json:"statusSource" gorm:"type:varchar(100); comment:状态来源;"`
	Remark          string    `json:"remark" gorm:"type:varchar(1000); comment:备注说明;"`
	CreateTime      time.Time `json:"createTime" gorm:"type:timestamp; default:now();comment:创建时间;"`
	CreateBy        string    `json:"createBy" gorm:"type:varchar(255); comment:创建者;"`
	UpdateTime      time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	UpdateBy        string    `json:"updateBy" gorm:"type:varchar(255); comment:更新者;"`
}

func (e Equipment) TableName() string {
	return "property_equipment"
}

func (e Equipment) Comment() string {
	return "物业管理-设备信息表"
}

func (e Equipment) InitData() any {
	return nil
}

type EquipmentStatus struct {
	Id     int64  `json:"id" gorm:"type:bigint; not null; primary_key; comment:状态ID;"`
	Code   string `json:"code" gorm:"type:varchar(100); not null; comment:状态编号;"`
	Name   string `json:"name" gorm:"type:varchar(100); not null; comment:状态名称;"`
	Remark string `json:"remark" gorm:"type:varchar(1000); comment:备注说明;"`
}

func (e EquipmentStatus) TableName() string {
	return "property_equipment_status"
}

func (e EquipmentStatus) Comment() string {
	return "物业管理-设备状态配置表"
}

func (e EquipmentStatus) InitData() any {
	return nil
}

type EquipmentAlarm struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:记录ID;"`
	EquipmentId  int64     `json:"equipmentId" gorm:"type:bigint; not null; comment:设备ID;"`
	AlarmMessage string    `json:"alarmMessage" gorm:"type:varchar(20); comment:报警信息;"`
	AlarmTime    time.Time `json:"alarmTime" gorm:"type:timestamp; default:now(); comment:报警时间;"`
}

func (e EquipmentAlarm) TableName() string {
	return "property_equipment_alarm"
}

func (e EquipmentAlarm) Comment() string {
	return "物业管理-设备报警记录表"
}

func (e EquipmentAlarm) InitData() any {
	return nil
}

type EquipmentRepair struct {
	Id            int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:记录ID;"`
	EquipmentId   int64     `json:"equipmentId" gorm:"type:bigint; not null; comment:设备ID;"`
	RepairType    string    `json:"repairType" gorm:"type:varchar(20); not null; comment:维修类型;"`
	ReportTime    time.Time `json:"reportTime" gorm:"type:timestamp; default:now(); comment:报修时间;"`
	CompletedTime time.Time `json:"completedTime" gorm:"type:timestamp; default:now(); comment:维修完成时间;"`
	Repairman     string    `json:"repairman" gorm:"type:varchar(100); comment:维修人;"`
}

func (e EquipmentRepair) TableName() string {
	return "property_equipment_repair"
}

func (e EquipmentRepair) Comment() string {
	return "物业管理-设备报修记录表"
}

func (e EquipmentRepair) InitData() any {
	return nil
}

type EquipmentInspect struct {
	Id          int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:记录ID;"`
	EquipmentId int64     `json:"equipmentId" gorm:"type:bigint; not null; comment:设备ID;"`
	InspectPlan string    `json:"inspectPlan" gorm:"type:varchar(20); not null; comment:巡检计划;"`
	InspectTime time.Time `json:"inspectTime" gorm:"type:timestamp; default:now(); comment:巡检时间;"`
	Inspector   string    `json:"inspector" gorm:"type:varchar(100); comment:巡检人;"`
}

func (e EquipmentInspect) TableName() string {
	return "property_equipment_inspect"
}

func (e EquipmentInspect) Comment() string {
	return "物业管理-设备巡检记录表"
}

func (e EquipmentInspect) InitData() any {
	return nil
}
