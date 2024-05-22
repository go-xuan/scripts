package entity

import "time"

type HbkjAppUser struct {
	Id       string `json:"id" gorm:"type:varchar(20); not null; primary_key; comment:用户ID;"`
	Name     string `json:"name" gorm:"type:varchar(100); not null; comment:用户姓名;"`
	DeptId   string `json:"deptId" gorm:"type:varchar(20); not null; comment:部门ID;"`
	DeptName string `json:"deptName" gorm:"type:varchar(100); comment:部门名称;"`
}

func (l HbkjAppUser) TableName() string {
	return "app_user"
}

func (l HbkjAppUser) TableComment() string {
	return "用户表"
}

func (l HbkjAppUser) InitData() any {
	return nil
}

type HbkjAppDept struct {
	Id       string `json:"id" gorm:"type:varchar(20); not null; primary_key; comment:部门ID;"`
	Name     string `json:"name" gorm:"type:varchar(100); not null; comment:部门名称;"`
	Prefix   string `json:"prefix" gorm:"type:varchar(10); not null; default:001; comment:部门前缀;"`
	First    int    `json:"first" gorm:"type:smallint; not null; default:0; comment:工号首位;"`
	Sequence int    `json:"sequence" gorm:"type:smallint; not null; default:1; comment:工号序列;"`
}

func (l HbkjAppDept) TableName() string {
	return "app_dept"
}

func (l HbkjAppDept) TableComment() string {
	return "部门工号序列维护表"
}

func (l HbkjAppDept) InitData() any {
	return nil
}

type HbkjAppEmployUser struct {
	Id         string `json:"id" gorm:"type:varchar(20); not null; primary_key; comment:用户ID;"`
	Name       string `json:"name" gorm:"type:varchar(100); not null; comment:用户姓名;"`
	DeptId     string `json:"deptId" gorm:"type:varchar(20); not null; comment:部门ID;"`
	DeptName   string `json:"deptName" gorm:"type:varchar(100); comment:部门名称;"`
	EmployType string `json:"employType" gorm:"type:varchar(10); not null; comment:聘用类型;"`
	EmployName string `json:"employName" gorm:"type:varchar(100); comment:聘用类型名称;"`
}

func (l HbkjAppEmployUser) TableName() string {
	return "app_employ_user"
}

func (l HbkjAppEmployUser) TableComment() string {
	return "雇佣用户表"
}

func (l HbkjAppEmployUser) InitData() any {
	return nil
}

type HbkjAppEmploySequence struct {
	Id       string `json:"id" gorm:"type:varchar(20); not null; primary_key; comment:雇佣类型ID;"`
	Name     string `json:"name" gorm:"type:varchar(100); not null; comment:雇佣类型名称;"`
	Year     int    `json:"year" gorm:"type:smallint; not null;  comment:年份;"`
	Sequence int    `json:"sequence" gorm:"type:smallint; not null; default:1; comment:工号序列;"`
}

func (l HbkjAppEmploySequence) TableName() string {
	return "app_employ_sequence"
}

func (l HbkjAppEmploySequence) TableComment() string {
	return "雇佣工号序列维护表"
}

func (l HbkjAppEmploySequence) InitData() any {
	var data []HbkjAppEmploySequence
	data = append(data, HbkjAppEmploySequence{"98", "院聘", time.Now().Year(), 1})
	data = append(data, HbkjAppEmploySequence{"99", "劳务派遣", time.Now().Year(), 1})
	return data
}
