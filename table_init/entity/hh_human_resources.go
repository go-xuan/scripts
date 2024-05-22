package entity

import (
	"time"

	"github.com/go-xuan/quanx/utils/randx"
)

var hhSchema = "portal."

// 业绩信息认定任务
type ConfirmTask struct {
	Id         string    `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	Name       string    `json:"name" gorm:"type:varchar(100); not null; comment:任务名称;"`
	StartTime  time.Time `json:"startTime" gorm:"type:timestamp(0); not null; default:now(); comment:任务开始时间;"`
	EndTime    time.Time `json:"endTime" gorm:"type:timestamp(0); not null; default:now(); comment:任务结束时间;"`
	DataSource string    `json:"dataSource" gorm:"type:varchar(100); not null; comment:数据来源;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (c ConfirmTask) TableName() string {
	return hhSchema + "hr_confirm_task"
}

func (c ConfirmTask) TableComment() string {
	return "业绩信息认定-任务"
}

func (c ConfirmTask) InitData() any {
	return nil
}

// 业绩信息认定参评人员
type ConfirmTaskPerson struct {
	Id       string `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	TaskId   string `json:"taskId" gorm:"type:varchar(50); not null; comment:任务ID;"`
	Username string `json:"username" gorm:"type:varchar(100); not null; comment:用户名;"`
	Name     string `json:"name" gorm:"type:varchar(100); not null; comment:用户姓名;"`
}

func (c ConfirmTaskPerson) TableName() string {
	return hhSchema + "hr_confirm_task_person"
}

func (c ConfirmTaskPerson) TableComment() string {
	return "业绩信息认定-参评人员"
}

func (c ConfirmTaskPerson) InitData() any {
	return nil
}

// 业绩信息认定关联表单
type ConfirmTaskWorkflow struct {
	Id       string `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	TaskId   string `json:"taskId" gorm:"type:varchar(50); not null; comment:任务ID;"`
	FormYid  string `json:"formYid" gorm:"type:varchar(100); not null; comment:表单源ID;"`
	FormId   string `json:"formId" gorm:"type:varchar(100); not null; comment:表单ID;"`
	FormName string `json:"formName" gorm:"type:varchar(100); not null; comment:表单名称;"`
}

func (c ConfirmTaskWorkflow) TableName() string {
	return hhSchema + "hr_confirm_task_workflow"
}

func (c ConfirmTaskWorkflow) TableComment() string {
	return "业绩信息认定-关联表单"
}

func (c ConfirmTaskWorkflow) InitData() any {
	return nil
}

// 业绩信息认定填报记录
type ConfirmTaskRecord struct {
	Id             string `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	TaskId         string `json:"taskId" gorm:"type:varchar(50); not null; comment:任务ID;"`
	Username       string `json:"username" gorm:"type:varchar(100); not null; comment:用户名;"`
	Name           string `json:"name" gorm:"type:varchar(100); not null; comment:姓名;"`
	DeptName       string `json:"deptName" gorm:"type:varchar(100); not null; comment:部门名称;"`
	FormYid        string `json:"formYid" gorm:"type:varchar(100); not null; comment:表单源ID;"`
	FormId         string `json:"formId" gorm:"type:varchar(100); not null; comment:表单ID;"`
	FormName       string `json:"formName" gorm:"type:varchar(200); not null; comment:表单名称;"`
	CompleteStatus string `json:"complete_status" gorm:"type:varchar(200); comment:填报状态（0-未填报；1-已填报）;"`
	ApproveStatus  string `json:"approveStatus" gorm:"type:varchar(200);  comment:审批状态（0-未审批；1-已审批）;"`
	ResultStatus   string `json:"resultStatus" gorm:"type:varchar(200);  comment:结果状态（0-进行中；1-已完结）;"`
}

func (c ConfirmTaskRecord) TableName() string {
	return hhSchema + "hr_confirm_task_record"
}

func (c ConfirmTaskRecord) TableComment() string {
	return "业绩信息认定-填报记录"
}

func (c ConfirmTaskRecord) InitData() any {
	return nil
}

// 业绩量化评分
type Scoring struct {
	Id          string    `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	Name        string    `json:"name" gorm:"type:varchar(100); not null; comment:指标体系名称;"`
	Description string    `json:"description" gorm:"type:varchar(1000); not null; comment:指标体系描述;"`
	CreateTime  time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime  time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (s Scoring) TableName() string {
	return hhSchema + "hr_scoring"
}

func (s Scoring) TableComment() string {
	return "业绩量化评分"
}

func (s Scoring) InitData() any {
	return nil
}

// 参评人员
type ScoringPerson struct {
	Id        string `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	ScoringId string `json:"scoringId" gorm:"type:varchar(50); not null; comment:评分ID;"`
	Username  string `json:"username" gorm:"type:varchar(100); not null; comment:用户名;"`
	Name      string `json:"name" gorm:"type:varchar(100); not null; comment:用户姓名;"`
}

func (s ScoringPerson) TableName() string {
	return hhSchema + "hr_scoring_person"
}

func (s ScoringPerson) TableComment() string {
	return "业绩量化评分-参评人员"
}

func (s ScoringPerson) InitData() any {
	return nil
}

// 评分指标
type ScoringIndicator struct {
	Id         string  `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	ScoringId  string  `json:"scoringId" gorm:"type:varchar(50); not null; comment:评分ID;"`
	Name       string  `json:"name" gorm:"type:varchar(100); not null; comment:指标名称;"`
	Table      string  `json:"table" gorm:"type:varchar(100); not null; comment:指标所在表名;"`
	Column     string  `json:"column" gorm:"type:varchar(100); not null; comment:指标字段名;"`
	UserColumn string  `json:"userColumn" gorm:"type:varchar(100); not null; comment:用户字段名;"`
	Weight     float64 `json:"weight" gorm:"type:float; default:0; comment:指标权重;"`
}

func (s ScoringIndicator) TableName() string {
	return hhSchema + "hr_scoring_indicator"
}

func (s ScoringIndicator) TableComment() string {
	return "业绩量化评分-指标配置"
}

func (s ScoringIndicator) InitData() any {
	return nil
}

// 评分指标
type ScoringData1 struct {
	Username string  `json:"username" gorm:"type:varchar(100); not null; comment:用户名;"`
	Score11  float64 `json:"scoreAa" gorm:"type:numeric; not null; comment:评分11;"`
	Score12  float64 `json:"scoreAb" gorm:"type:numeric; not null; comment:评分12;"`
	Score13  float64 `json:"scoreAc" gorm:"type:numeric; not null; comment:评分13;"`
}

var users = []string{"zhou",
	"quanchao",
	"zjh",
	"test002",
	"wcl",
	"test",
	"ces1"}

func (s ScoringData1) TableName() string {
	return "scoring_data_1"
}

func (s ScoringData1) TableComment() string {
	return "评分测试数据表1"
}

func (s ScoringData1) InitData() any {
	var data []*ScoringData1
	for _, user := range users {
		data = append(data, &ScoringData1{
			Username: user,
			Score11:  randx.Float64Range(1, 100, 2),
			Score12:  randx.Float64Range(1, 100, 2),
			Score13:  randx.Float64Range(1, 100, 2),
		})
	}
	return data
}

type ScoringData2 struct {
	Username string  `json:"username" gorm:"type:varchar(100); not null; comment:用户名;"`
	Score21  float64 `json:"scoreBa" gorm:"type:numeric; not null; comment:评分21;"`
	Score22  float64 `json:"scoreBb" gorm:"type:numeric; not null; comment:评分22;"`
	Score23  float64 `json:"scoreBc" gorm:"type:numeric; not null; comment:评分23;"`
}

func (s ScoringData2) TableName() string {
	return "scoring_data_2"
}

func (s ScoringData2) TableComment() string {
	return "评分测试数据表2"
}

func (s ScoringData2) InitData() any {
	var data []*ScoringData2
	for _, user := range users {
		data = append(data, &ScoringData2{
			Username: user,
			Score21:  randx.Float64Range(1, 100, 2),
			Score22:  randx.Float64Range(1, 100, 2),
			Score23:  randx.Float64Range(1, 100, 2),
		})
	}
	return data
}

type ScoringData3 struct {
	Username string  `json:"username" gorm:"type:varchar(100); not null; comment:用户名;"`
	Score31  float64 `json:"scoreCa" gorm:"type:numeric; not null; comment:评分31;"`
	Score32  float64 `json:"scoreCb" gorm:"type:numeric; not null; comment:评分32;"`
	Score33  float64 `json:"scoreCc" gorm:"type:numeric; not null; comment:评分33;"`
}

func (s ScoringData3) TableName() string {
	return "scoring_data_3"
}

func (s ScoringData3) TableComment() string {
	return "评分测试数据表3"
}

func (s ScoringData3) InitData() any {
	var data []*ScoringData3
	for _, user := range users {
		data = append(data, &ScoringData3{
			Username: user,
			Score31:  randx.Float64Range(1, 100, 2),
			Score32:  randx.Float64Range(1, 100, 2),
			Score33:  randx.Float64Range(1, 100, 2),
		})
	}
	return data
}

type EduAppRlzyJbxxLx struct {
	SjbId     string `json:"sjbId" gorm:"type:varchar(50); not null; comment:数据包id;"`
	Tx        string `json:"tx" gorm:"type:varchar(555); comment:头像;"`
	Xm        string `json:"xm" gorm:"type:varchar(255); comment:姓名;"`
	Sfzh      string `json:"sfzh" gorm:"type:varchar(255); comment:身份证号;"`
	Jx        string `json:"jx" gorm:"type:varchar(255); comment:军衔;"`
	Xb        string `json:"xb" gorm:"type:varchar(255); comment:性别;"`
	Bzb       string `json:"bzb" gorm:"type:varchar(255); comment:部职别;"`
	Jrzh      string `json:"jrzh" gorm:"type:varchar(255); comment:军人证号;"`
	Csrq      string `json:"csrq" gorm:"type:varchar(255); comment:出生日期;"`
	Rwny      string `json:"rwny" gorm:"type:varchar(255); comment:入伍年月;"`
	Jg        string `json:"jg" gorm:"type:varchar(255); comment:籍贯;"`
	Rdrq      string `json:"rdrq" gorm:"type:varchar(255); comment:入党日期;"`
	Zj        string `json:"zj" gorm:"type:varchar(255); comment:职级;"`
	Zjny      string `json:"zjny" gorm:"type:varchar(255); comment:职级年月;"`
	Zgxl      string `json:"zgxl" gorm:"type:varchar(255); comment:最高学历;"`
	Xzzw      string `json:"xzzw" gorm:"type:varchar(255); comment:行政职务;"`
	Xy        string `json:"xy" gorm:"type:varchar(255); comment:学缘;"`
	Dwmc      string `json:"dwmc" gorm:"type:varchar(255); comment:单位名称;"`
	Hyzt      string `json:"hyzt" gorm:"type:varchar(255); comment:婚姻状态;"`
	Dbdgzny   string `json:"dbdgzny" gorm:"type:varchar(255); comment:到部队工作年月;"`
	Bddh      string `json:"bddh" gorm:"type:varchar(255); comment:部队代号;"`
	Bzkzmbzh  string `json:"bzkzmbzh" gorm:"type:varchar(255); comment:保障卡正面保障号;"`
	Bzkbmyhkh string `json:"bzkbmyhkh" gorm:"type:varchar(255); comment:保障卡背面银行卡号;"`
	Mz        string `json:"mz" gorm:"type:varchar(255); comment:民族;"`
	Dyjb      string `json:"dyjb" gorm:"type:varchar(255); comment:待遇级别;"`
	Dyjbny    string `json:"dyjbny" gorm:"type:varchar(255); comment:待遇级别年月;"`
	Zyjsdjny  string `json:"zyjsdjny" gorm:"type:varchar(255); comment:专业技术等级年月;"`
	Rgny      string `json:"rgny" gorm:"type:varchar(255); comment:任干年月;"`
	Csd       string `json:"csd" gorm:"type:varchar(255); comment:出生地;"`
	Rwd       string `json:"rwd" gorm:"type:varchar(255); comment:入伍地;"`
	Zgxw      string `json:"zgxw" gorm:"type:varchar(255); comment:最高学位;"`
	Zyjszw    string `json:"zyjszw" gorm:"type:varchar(255); comment:专业技术职务;"`
	Jxny      string `json:"jxny" gorm:"type:varchar(255); comment:军衔年月;"`
	Zzmm      string `json:"zzmm" gorm:"type:varchar(255); comment:政治面貌;"`
	Zyjndj    string `json:"zyjndj" gorm:"type:varchar(255); comment:专业技能等级;"`
	Nl        string `json:"nl" gorm:"type:varchar(255); comment:年龄;"`
	Rtrq      string `json:"rtrq" gorm:"type:varchar(255); comment:入团日期;"`
	Hkszd     string `json:"hkszd" gorm:"type:varchar(255); comment:户口所在地;"`
	Yktkh     string `json:"yktkh" gorm:"type:varchar(255); comment:一卡通卡号;"`
	Yktlx     string `json:"yktlx" gorm:"type:varchar(255); comment:一卡通类型;"`
	Bgsh      string `json:"bgsh" gorm:"type:varchar(255); comment:办公室号;"`
	Rylx      string `json:"rylx" gorm:"type:varchar(255); comment:人员类型;"`
	Zy        string `json:"zy" gorm:"type:varchar(255); comment:专业;"`
	Hrjzgny   string `json:"hrjzgny" gorm:"type:varchar(255); comment:获任教资格年月;"`
	Rysmdj    string `json:"rysmdj" gorm:"type:varchar(255); comment:人员涉密等级;"`
	Cjgzny    string `json:"cjgzny" gorm:"type:varchar(255); comment:参加工作年月;"`
	Jtdz      string `json:"jtdz" gorm:"type:varchar(255); comment:家庭地址;"`
	Lysj      string `json:"lysj" gorm:"type:varchar(255); comment:录用时间;"`
	Zc        string `json:"zc" gorm:"type:varchar(255); comment:职称;"`
	Wzzjh     string `json:"wzzjh" gorm:"type:varchar(255); comment:文职证件号;"`
	Gwlb      string `json:"gwlb" gorm:"type:varchar(255); comment:岗位类别;"`
	Wydj      string `json:"wydj" gorm:"type:varchar(255); comment:文员等级;"`
	Wydjny    string `json:"wydjny" gorm:"type:varchar(255); comment:文员等级年月;"`
}

func (s EduAppRlzyJbxxLx) TableName() string {
	return hhSchema + "edu_app_rlzy_jbxx_lx"
}

func (s EduAppRlzyJbxxLx) TableComment() string {
	return "人员基本信息-离线数据"
}

func (s EduAppRlzyJbxxLx) InitData() any {
	return nil
}

type EduAppRlzyJbxxSjb struct {
	Id    string    `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	Sjbmc string    `json:"sjbmc" gorm:"type:varchar(100); not null; comment:数据包名称;"`
	Czr   string    `json:"czr" gorm:"type:varchar(100); comment:操作人;"`
	Czrxm string    `json:"czrxm" gorm:"type:varchar(100); comment:操作人姓名;"`
	Czsj  time.Time `json:"czsj" gorm:"type:timestamp(0); default:now(); comment:操作时间;"`
}

func (s EduAppRlzyJbxxSjb) TableName() string {
	return hhSchema + "edu_app_rlzy_jbxx_sjb"
}

func (s EduAppRlzyJbxxSjb) TableComment() string {
	return "人员基本信息-离线数据包"
}

func (s EduAppRlzyJbxxSjb) InitData() any {
	return nil
}

type EduAppRlzyJbxxSjbXzjl struct {
	Id    string    `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	SjbId string    `json:"sjbId" gorm:"type:varchar(50); not null; comment:数据包ID;"`
	Czr   string    `json:"czr" gorm:"type:varchar(100); comment:操作人;"`
	Czrxm string    `json:"czrxm" gorm:"type:varchar(100); comment:操作人姓名;"`
	Czsj  time.Time `json:"czsj" gorm:"type:timestamp(0); default:now(); comment:操作时间;"`
}

func (s EduAppRlzyJbxxSjbXzjl) TableName() string {
	return hhSchema + "edu_app_rlzy_jbxx_sjb_xzjl"
}

func (s EduAppRlzyJbxxSjbXzjl) TableComment() string {
	return "人员基本信息-离线数据包-下载记录"
}

func (s EduAppRlzyJbxxSjbXzjl) InitData() any {
	return nil
}

type EduAppRlzyJbxxXzmb struct {
	Id    string    `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	Mbmc  string    `json:"mbmc" gorm:"type:varchar(100); not null; comment:模板名称;"`
	Czr   string    `json:"czr" gorm:"type:varchar(100); comment:操作人;"`
	Czrxm string    `json:"czrxm" gorm:"type:varchar(100); comment:操作人姓名;"`
	Czsj  time.Time `json:"czsj" gorm:"type:timestamp(0); default:now(); comment:操作时间;"`
}

func (s EduAppRlzyJbxxXzmb) TableName() string {
	return hhSchema + "edu_app_rlzy_jbxx_xzmb"
}

func (s EduAppRlzyJbxxXzmb) TableComment() string {
	return "人员基本信息-下载模板"
}

func (s EduAppRlzyJbxxXzmb) InitData() any {
	return nil
}

type EduAppRlzyJbxxXzzd struct {
	Id    string    `json:"id" gorm:"type:varchar(50); comment:主键ID;"`
	MbId  string    `json:"mbId" gorm:"type:varchar(50); not null; comment:模板ID（默认为0）;"`
	Zd    string    `json:"zd" gorm:"type:varchar(100); not null; comment:字段;"`
	Zdmc  string    `json:"zdmc" gorm:"type:varchar(100); not null; comment:字段名称;"`
	Px    int       `json:"px" gorm:"type:smallint; default:1; comment:排序;"`
	Czr   string    `json:"czr" gorm:"type:varchar(100); comment:操作人;"`
	Czrxm string    `json:"czrxm" gorm:"type:varchar(100); comment:操作人姓名;"`
	Czsj  time.Time `json:"czsj" gorm:"type:timestamp(0); default:now(); comment:操作时间;"`
}

func (s EduAppRlzyJbxxXzzd) TableName() string {
	return hhSchema + "edu_app_rlzy_jbxx_xzzd"
}

func (s EduAppRlzyJbxxXzzd) TableComment() string {
	return "人员基本信息-下载字段"
}

func (s EduAppRlzyJbxxXzzd) InitData() any {
	return nil
}
