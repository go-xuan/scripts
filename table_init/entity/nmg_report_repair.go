package entity

import "time"

// 字典配置
type Lookup struct {
	Id        int64  `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	Key       string `json:"key" gorm:"type:varchar(100); not null; comment:字典Key;"`
	Value     string `json:"value" gorm:"type:varchar(100); not null; comment:字典值;"`
	Class     string `json:"class" gorm:"type:varchar(100); not null; comment:分类;"`
	ClassName string `json:"className" gorm:"type:varchar(100); not null; comment:分类名称;"`
	Pid       int64  `json:"pid" gorm:"type:bigint; not null; default:0; comment:父级主键ID;"`
	Index     int    `json:"index" gorm:"type:smallint; default:1; comment:排序下标;"`
}

func (l Lookup) TableName() string {
	return "t_lookup"
}

func (l Lookup) TableComment() string {
	return "字典配置"
}

func (l Lookup) InitData() any {
	var data []*Lookup
	data = append(data, &Lookup{Key: "repair_completed", Value: "维修完成", Class: "report_result", ClassName: "维修结果", Pid: 0, Index: 1})
	data = append(data, &Lookup{Key: "accept_processing", Value: "已受理，持续处理中", Class: "report_result", ClassName: "维修结果", Pid: 0, Index: 2})
	data = append(data, &Lookup{Key: "accept_reporting", Value: "已受理，逐级报批中", Class: "report_result", ClassName: "维修结果", Pid: 0, Index: 3})
	data = append(data, &Lookup{Key: "incomplete_info", Value: "信息不全", Class: "report_result", ClassName: "维修结果", Pid: 0, Index: 4})
	data = append(data, &Lookup{Key: "pending", Value: "待处理", Class: "report_status", ClassName: "报修状态", Pid: 0, Index: 1})
	data = append(data, &Lookup{Key: "processing", Value: "处理中", Class: "report_status", ClassName: "报修状态", Pid: 0, Index: 2})
	data = append(data, &Lookup{Key: "processed", Value: "已处理", Class: "report_status", ClassName: "报修状态", Pid: 0, Index: 3})
	data = append(data, &Lookup{Key: "system", Value: "系统报修", Class: "report_way", ClassName: "报修方式", Pid: 0, Index: 1})
	data = append(data, &Lookup{Key: "phone", Value: "电话报修", Class: "report_way", ClassName: "报修方式", Pid: 0, Index: 2})
	data = append(data, &Lookup{Key: "yes", Value: "是", Class: "yes_or_no", ClassName: "是否", Pid: 0, Index: 1})
	data = append(data, &Lookup{Key: "no", Value: "否", Class: "yes_or_no", ClassName: "是否", Pid: 0, Index: 2})
	return data
}

type ReportEvaluation struct {
	Id         int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId   string    `json:"recordId" gorm:"varchar(100); not null; comment:表单记录ID;"`
	Resolved   bool      `json:"resolved" gorm:"bool; not null; default:false; comment:是否解决;"`
	Efficiency int       `json:"efficiency" gorm:"type:smallint; default:0; comment:处理时效;"`
	Attitude   int       `json:"attitude" gorm:"type:smallint; default:0; comment:服务态度;"`
	Technical  int       `json:"technical" gorm:"type:smallint; default:0; comment:技术水平;"`
	Evaluation string    `json:"evaluation" gorm:"type:text; default:0; comment:服务评价;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (r *ReportEvaluation) TableName() string {
	return "t_report_evaluation"
}

func (r *ReportEvaluation) TableComment() string {
	return "报修评价"
}

func (r *ReportEvaluation) InitData() any {
	return nil
}

// 报修数据
type Report struct {
	Id             int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId       string    `json:"recordId" gorm:"type:varchar(100); not null; comment:表单记录ID;"`
	Submitter      string    `json:"submitter" gorm:"type:varchar(100); not null; comment:发起人;"`
	Department     string    `json:"department" gorm:"type:varchar(100); not null; comment:所在部门;"`
	Contact        string    `json:"contact" gorm:"type:varchar(20); not null; comment:联系方式;"`
	Type           string    `json:"type" gorm:"type:varchar(100); not null; comment:报修类型;"`
	Project        string    `json:"project" gorm:"type:varchar(100); not null; comment:报修项目;"`
	Area           string    `json:"area" gorm:"type:varchar(200); not null; comment:报修区域;"`
	Place          string    `json:"place" gorm:"type:varchar(200); not null; comment:报修地点;"`
	Description    string    `json:"description" gorm:"type:varchar(1000); not null; comment:故障描述;"`
	Way            string    `json:"way" gorm:"type:varchar(100); comment:报修方式;"`
	Status         string    `json:"status" gorm:"type:varchar(100); comment:报修状态;"`
	Result         string    `json:"result" gorm:"type:varchar(100); comment:维修结果;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateUser     string    `json:"createUser" gorm:"type:varchar(100); comment:创建人姓名;"`
	CreateUsername string    `json:"createUsername" gorm:"type:varchar(100); comment:创建人用户名;"`
}

func (r *Report) TableName() string {
	return "t_report"
}

func (r *Report) TableComment() string {
	return "报修数据"
}

func (r *Report) InitData() any {
	return nil
}

// 报修记录日志
type ReportLog struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId   string    `json:"recordId" gorm:"type:varchar(100); not null; comment:表单记录ID;"`
	Status     string    `json:"status" gorm:"type:varchar(100); comment:报修状态;"`
	Title      string    `json:"title" gorm:"type:varchar(100); comment:记录标题;"`
	Content    string    `json:"content" gorm:"type:varchar(100); comment:记录内容;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	CreateUser string    `json:"createUser" gorm:"type:varchar(100); comment:创建人姓名;"`
	Remark     string    `json:"remark" gorm:"type:varchar(1000); comment:备注（200字以内）;"`
}

func (r *ReportLog) TableName() string {
	return "t_report_log"
}

func (r *ReportLog) TableComment() string {
	return "报修记录日志"
}

func (r *ReportLog) InitData() any {
	return nil
}

// 报修关联附件
type ReportAttachment struct {
	Id        int64  `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId  string `json:"recordId" gorm:"type:varchar(100); not null; comment:表单记录ID;"`
	Status    string `json:"status" gorm:"type:varchar(20); comment:附件所属报修状态;"`
	Type      string `json:"type" gorm:"type:varchar(20); comment:附件类型（picture/video）;"`
	MinioPath string `json:"minioPath" gorm:"type:varchar(500); comment:minio文件路径;"`
	Index     int    `json:"index" gorm:"type:smallint; not null; default:1; comment:排序下标;"`
}

func (r *ReportAttachment) TableName() string {
	return "t_report_attachment"
}

func (r *ReportAttachment) TableComment() string {
	return "报修关联附件"
}

func (r *ReportAttachment) InitData() any {
	return nil
}

// 报修关联审批人
type ReportApprover struct {
	Id               int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId         string    `json:"recordId" gorm:"type:varchar(100); not null; comment:表单记录ID;"`
	ApproverName     string    `json:"approverName" gorm:"type:varchar(100); comment:审批人姓名;"`
	ApproverUsername string    `json:"approverUsername" gorm:"type:varchar(100); not null; comment:审批人;"`
	ApprovalNodeId   string    `json:"approvalNodeId" gorm:"type:varchar(100); comment:审批节点ID;"`
	CreateTime       time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
}

func (r *ReportApprover) TableName() string {
	return "t_report_approver"
}

func (r *ReportApprover) TableComment() string {
	return "报修关联审批人"
}

func (r *ReportApprover) InitData() any {
	return nil
}

// 报修关联维修人表
type ReportRepairman struct {
	Id          int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId    string    `json:"recordId" gorm:"type:varchar(100); not null; comment:表单记录ID;"`
	RepairmanId int64     `json:"repairmanId" gorm:"type:bigint; not null; comment:维修人ID;"`
	CreateTime  time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
}

func (r *ReportRepairman) TableName() string {
	return "t_report_repairman"
}

func (r *ReportRepairman) TableComment() string {
	return "报修关联维修人"
}

func (r *ReportRepairman) InitData() any {
	return nil
}

// 维修人管理表
type Repairman struct {
	Id         int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	GroupId    int64     `json:"groupId" gorm:"type:bigint; not null; comment:分组ID;"`
	Name       string    `json:"name" gorm:"type:varchar(200); not null; comment:姓名;"`
	Contact    string    `json:"contact" gorm:"type:varchar(20); not null; comment:联系方式;"`
	Gender     *string   `json:"gender" gorm:"type:varchar(10); comment:性别;"`
	IdCard     string    `json:"idCard" gorm:"type:varchar(20); comment:身份证号;"`
	IsDelete   int       `json:"isDelete" gorm:"type:smallint; default:0; comment:是否删除;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp; default:now(); comment:更新时间;"`
	Remark     string    `json:"remark" gorm:"type:varchar(1000); comment:备注（200字以内）;"`
}

func (r *Repairman) TableName() string {
	return "t_repairman"
}

func (r *Repairman) TableComment() string {
	return "维修人"
}

func (r *Repairman) InitData() any {
	return nil
}

// 维修人分组
type RepairmanGroup struct {
	Id    int64  `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	Name  string `json:"name" gorm:"type:varchar(200); not null; comment:姓名;"`
	Index int    `json:"index" gorm:"type:smallint; not null; default:1; comment:排序下标;"`
}

func (t *RepairmanGroup) TableName() string {
	return "t_repairman_group"
}

func (t *RepairmanGroup) TableComment() string {
	return "维修人分组"
}

func (t *RepairmanGroup) InitData() any {
	return nil
}

type ReportNotice struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RecordId     string    `json:"recordId" gorm:"type:varchar(100); not null; comment:表单记录ID;"`
	Title        string    `json:"title" gorm:"type:varchar(100); comment:记录标题;"`
	Content      string    `json:"content" gorm:"type:varchar(100); comment:记录内容;"`
	SendTime     time.Time `json:"sendTime" gorm:"type:timestamp; default:now(); comment:创建时间;"`
	Receiver     string    `json:"receiver" gorm:"type:varchar(100); comment:接收人;"`
	ReceiverRole string    `json:"receiverRole" gorm:"type:varchar(100); comment:接收人角色;"`
	Read         int       `json:"read" gorm:"type:varchar(100); comment:已读（1-已读；0-未读）;"`
}

func (t *ReportNotice) TableName() string {
	return "t_report_notice"
}

func (t *ReportNotice) TableComment() string {
	return "报修通知"
}

func (t *ReportNotice) InitData() any {
	return nil
}
