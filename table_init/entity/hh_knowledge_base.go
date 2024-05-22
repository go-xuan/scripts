package entity

import "time"

type Question struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	Question       string    `json:"Question" gorm:"type:varchar(350); not null; comment:问题;"`
	Answer         string    `json:"answer" gorm:"type:text; not null; comment:答案;"`
	BigClassId     string    `json:"bigClassId" gorm:"type:varchar(100); not null; comment:大类ID;"`
	NeverExpire    bool      `json:"neverExpire" gorm:"type:bool; comment:是否永不过期;"`
	ExpireTime     time.Time `json:"expireTime" gorm:"type:timestamp(0); comment:过期时间;"`
	AdvancedEnable bool      `json:"advancedEnable" gorm:"type:bool; comment:启用高级设置;"`
	GreetType      string    `json:"greetType" gorm:"type:varchar(20); comment:寒暄语类型（greet/end/unknown）;"`
	AskInterval    int       `json:"askInterval" gorm:"type:int; comment:咨询间隔（分钟）;"`
	Status         int       `json:"status" gorm:"type:smallint; not null; comment:问题状态;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:创建人ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId   int       `json:"updateUserId" gorm:"type:bigint; comment:更新人ID;"`
	UpdateUserName string    `json:"updateUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	UpdateTime     time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (Question) TableName() string {
	return "t_question"
}

type QuestionClass struct {
	Id           int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	QuestionId   int       `json:"questionId" gorm:"type:bigint; not null; comment:问题ID;"`
	SmallClassId int       `json:"smallClassId" gorm:"type:bigint; not null; comment:问题分类ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (QuestionClass) TableName() string {
	return "t_question_class"
}

type QuestionRelation struct {
	Id                 int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	QuestionId         int       `json:"questionId" gorm:"type:bigint; not null; comment:问题ID;"`
	RelationQuestionId int       `json:"relationQuestionId" gorm:"type:bigint; not null; comment:关联问题ID;"`
	CreateTime         time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime         time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (QuestionRelation) TableName() string {
	return "t_question_relation"
}

type QuestionSimilar struct {
	Id         string    `json:"id" gorm:"type:varchar(100); not null; primary_key; comment:主键ID;"`
	QuestionId int       `json:"questionId" gorm:"type:bigint; not null; comment:标准问题ID;"`
	Question   string    `json:"Question" gorm:"type:varchar(350); not null; comment:相似问题题干;"`
	BigClassId string    `json:"bigClassId" gorm:"type:varchar(100); not null; comment:大类ID;"`
	IsStandard int       `json:"isStandard" gorm:"type:smallint; not null; comment:是否标准问题;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (QuestionSimilar) TableName() string {
	return "t_question_similar"
}

type Sensitive struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	Word           string    `json:"word" gorm:"type:varchar(350); not null; comment:敏感词汇;"`
	ShieldMethod   int       `json:"shieldMethod" gorm:"type:smallint; not null; comment:屏蔽方式;"`
	SmallClassId   int       `json:"smallClassId" gorm:"type:bigint; not null; comment:小类ID;"`
	Status         int       `json:"status" gorm:"type:smallint; not null; comment:问题状态;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:创建人ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId   int       `json:"updateUserId" gorm:"type:bigint; comment:更新人ID;"`
	UpdateUserName string    `json:"updateUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	UpdateTime     time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (Sensitive) TableName() string {
	return "t_sensitive"
}

type ChatRecord struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	UserQuestion   string    `json:"userQuestion" gorm:"type:varchar(500); not null; comment:用户问题;"`
	RealQuestion   string    `json:"realQuestion" gorm:"type:varchar(500); not null; comment:用户实际问题;"`
	AnswerType     int       `json:"answerType" gorm:"type:smallint; not null; comment:回答类型（0-未识别；1-知识库应答；2-寒暄库应答；3-欺诈检测；4-插入信息）;"`
	QuestionId     int       `json:"questionId" gorm:"type:bigint; comment:匹配问题ID;"`
	Question       string    `json:"question" gorm:"type:varchar(500); comment:匹配问题题干;"`
	BigClass       string    `json:"bigClass" gorm:"type:varchar(100); comment:匹配分类;"`
	Answer         string    `json:"answer" gorm:"type:text; comment:回答;"`
	RelationList   string    `json:"relationList" gorm:"type:text; comment:关联问题;"`
	InsertInfoList string    `json:"insertInfoList" gorm:"type:text; comment:插入信息（推荐回答）;"`
	AskTime        time.Time `json:"askTime" gorm:"type:timestamp(0); comment:提问时间;"`
	AnswerTime     time.Time `json:"answerTime" gorm:"type:timestamp(0); comment:回答时间;"`
	SolveType      int       `json:"solveType" gorm:"type:smallint; default:0; comment:解答评价（0-未评价；1-未解决；2-解决）;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:创建人ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime     time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (ChatRecord) TableName() string {
	return "t_chat_record"
}

type Class struct {
	Id         int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	SmallClass string    `json:"smallClass" gorm:"type:varchar(100); not null; comment:分类小类;"`
	BigClassId string    `json:"bigClassId" gorm:"type:varchar(100); not null; comment:分类大类key;"`
	BigClass   string    `json:"bigClass" gorm:"type:varchar(100); not null; comment:分类大类;"`
	Index      int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateTime time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (Class) TableName() string {
	return "t_class"
}

type LeaveMsg struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	Phone          string    `json:"phone" gorm:"type:varchar(20); not null; comment:电话号码;"`
	Content        string    `json:"content" gorm:"type:varchar(2000); not null; comment:述求内容;"`
	Attachment     string    `json:"attachment" gorm:"type:text; comment:附件;"`
	Anonymous      bool      `json:"anonymous" gorm:"type:bool; comment:是否匿名;"`
	HasRead        bool      `json:"hasRead" gorm:"type:bool; comment:是否已读;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:用户ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:用户姓名;"`
	CreateDeptName string    `json:"createDeptName" gorm:"type:varchar(100); comment:所在部门;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
}

func (LeaveMsg) TableName() string {
	return "t_leave_msg"
}

type LeaveMsgComment struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	LeaveMsgId     int       `json:"leaveMsgId" gorm:"type:bigint; comment:留言ID;"`
	Anonymous      bool      `json:"anonymous" gorm:"type:bool; comment:是否匿名;"`
	Comment        string    `json:"comment" gorm:"type:varchar(2000); comment:评论;"`
	CreateUserType int       `json:"createUserType" gorm:"type:smallint; comment:评论人类型（1-智能客服；2-普通用户）;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:用户ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:用户姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
}

func (LeaveMsgComment) TableName() string {
	return "t_leave_msg_comment"
}

type LeaveMsgLike struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	LikeObject     int       `json:"likeObject" gorm:"type:smallint; comment:点赞对象（1-留言；2-留言评论）;"`
	LikeObjectId   int       `json:"likeObjectId" gorm:"type:bigint; comment:点赞对象ID;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:创建人ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
}

func (LeaveMsgLike) TableName() string {
	return "t_leave_msg_like"
}

type SysConfig struct {
	Id             int       `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	ConfigClass    string    `json:"configClass" gorm:"type:varchar(100); comment:配置分类;"`
	ConfigKey      string    `json:"configKey" gorm:"type:varchar(100); comment:配置项KEY;"`
	ConfigValue    string    `json:"configValue" gorm:"type:text; comment:配置项值;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:用户ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:用户姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId   int       `json:"updateUserId" gorm:"type:bigint; comment:更新人ID;"`
	UpdateUserName string    `json:"updateUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	UpdateTime     time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (SysConfig) TableName() string {
	return "t_sys_config"
}

type QuickEntry struct {
	Id             string    `json:"id" gorm:"type:varchar(100); not null; primary_key; comment:主键ID;"`
	Class          string    `json:"class" gorm:"type:varchar(100); comment:分类;"`
	Name           string    `json:"name" gorm:"type:varchar(500); comment:名称;"`
	Icon           string    `json:"icon" gorm:"type:varchar(200); comment:图标;"`
	Url            string    `json:"url" gorm:"type:varchar(1000); comment:URL;"`
	Server         string    `json:"server" gorm:"type:varchar(100); comment:事项关联服务名（事项专属字段）;"`
	ServerId       string    `json:"serverId" gorm:"type:varchar(100); comment:事项关联服务ID（事项专属字段）;"`
	App            string    `json:"app" gorm:"type:varchar(100); comment:事项关联表单所属应用名（事项专属字段）;"`
	AppType        string    `json:"appType" gorm:"type:varchar(100); comment:应用类型（轻应用专属字段，other-第三方；customize-自定义）;"`
	Index          int       `json:"index" gorm:"type:smallint; not null; comment:排序下标;"`
	CreateUserId   int       `json:"createUserId" gorm:"type:bigint; comment:用户ID;"`
	CreateUserName string    `json:"createUserName" gorm:"type:varchar(30); comment:用户姓名;"`
	CreateTime     time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId   int       `json:"updateUserId" gorm:"type:bigint; comment:更新人ID;"`
	UpdateUserName string    `json:"updateUserName" gorm:"type:varchar(30); comment:创建人姓名;"`
	UpdateTime     time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (e QuickEntry) TableComment() string {
	return "系统设置-快捷入口表"
}

func (e QuickEntry) InitData() any {
	return nil
}

func (QuickEntry) TableName() string {
	return "t_quick_entry"
}

type InsertInfo struct {
	Id         string    `json:"id" gorm:"type:varchar(100); not null; primary_key; comment:主键ID;"`
	Class      string    `json:"class" gorm:"type:varchar(100); comment:分类;"`
	Name       string    `json:"option" gorm:"type:varchar(500); not null; comment:选项;"`
	Url        string    `json:"insertUrl" gorm:"type:varchar(100); not null; comment:链接;"`
	CreateTime time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
}

func (InsertInfo) TableName() string {
	return "t_insert_info"
}
