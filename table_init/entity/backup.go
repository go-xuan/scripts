package entity

import (
	"time"
)

type BackupDatabase struct {
	Id           int       `json:"id"            gorm:"type:int4; comment:ID;"`
	UpdateAt     time.Time `json:"updateAt"      gorm:"type:timestamp; not null; comment:更新时间"`
	Type         int       `json:"type"          gorm:"type:int2; not null; comment:数据库类型（1-mysql；2-pgsql；3-mongo）"`
	Host         string    `json:"host"          gorm:"type:varchar(100); not null; comment:主机host"`
	Port         int       `json:"port"          gorm:"type:int4; not null; comment:端口"`
	Username     string    `json:"username"      gorm:"type:varchar(100); not null; comment:用户名"`
	Password     string    `json:"password"      gorm:"type:varchar(100); not null; comment:密码"`
	Name         string    `json:"name"          gorm:"type:varchar(100); not null; comment:数据库名"`
	Schema       string    `json:"schema"        gorm:"type:varchar(100); comment:模式名"`
	Table        string    `json:"table"         gorm:"type:varchar(100); comment:表名"`
	Status       int       `json:"status"        gorm:"type:int2; not null; comment:状态（1-正常；2-异常；3-禁用）"`
	LatestDumpId int       `json:"latestDumpId"  gorm:"type:int4; comment:最新dump"`
}

func (b BackupDatabase) TableName() string {
	return "backup_database"
}

func (b BackupDatabase) TableComment() string {
	return "数据库管理"
}

func (b BackupDatabase) InitData() any {
	return nil
}

type BackupDatabaseDump struct {
	Id           int       `json:"id"            gorm:"type:int4; comment:ID;"`
	UpdateAt     time.Time `json:"updateAt"      gorm:"type:timestamp; not null; comment:更新时间"`
	ClientId     int       `json:"clientId"      gorm:"type:int4; not null; comment:客户端ID"`
	DatabaseId   int       `json:"databaseId"    gorm:"type:int4; not null; comment:数据库ID"`
	DumpCmd      string    `json:"dumpCmd"       gorm:"type:varchar(2000); not null; comment:dump命令"`
	DumpPath     string    `json:"dumpPath"      gorm:"type:varchar(1000); not null; comment:dump路径"`
	Remark       string    `json:"remark"        gorm:"type:varchar(100); comment:执行备注"`
	StartTime    time.Time `json:"startTime"     gorm:"type:timestamp; comment:开始时间"`
	EndTime      time.Time `json:"endTime"       gorm:"type:timestamp; comment:结束时间"`
	Status       int       `json:"status"        gorm:"type:int2; not null; comment:状态（1-进行中；2-成功；3-失败）"`
	Used         bool      `json:"used"          gorm:"type:bool; not null; default:false; comment:是否已被使用"`
	StepRecordId int       `json:"stepRecordId"  gorm:"type:int4; comment:步骤记录ID"`
}

func (b BackupDatabaseDump) TableName() string {
	return "backup_database_dump"
}

func (b BackupDatabaseDump) TableComment() string {
	return "数据库dump管理"
}

func (b BackupDatabaseDump) InitData() any {
	return nil
}

type BackupClient struct {
	Id          int       `json:"id"          gorm:"type:int4; comment:ID;"`
	UpdateAt    time.Time `json:"updateAt"    gorm:"type:timestamp; not null; comment:更新时间"`
	Hostname    string    `json:"hostname"    gorm:"type:varchar(100); not null; comment:主机host"`
	Ip          string    `json:"ip"          gorm:"type:varchar(100); not null; comment:主机ip"`
	Port        int       `json:"port"        gorm:"type:int4; not null; comment:端口"`
	Status      int       `json:"status"      gorm:"type:int2; not null; comment:状态（1-正常；2-异常；3-禁用）"`
	ConnectTime time.Time `json:"connectTime" gorm:"type:timestamp; comment:上一次连接时间"`
}

func (b BackupClient) TableName() string {
	return "backup_client"
}

func (b BackupClient) TableComment() string {
	return "客户端管理"
}

func (b BackupClient) InitData() any {
	return nil
}

type BackupRepository struct {
	Id             int       `json:"id"             gorm:"type:int4; comment:ID;"`
	UpdateAt       time.Time `json:"updateAt"       gorm:"type:timestamp; not null;comment:更新时间"`
	Name           string    `json:"name"           gorm:"type:varchar(100); not null; comment:存储库名称"`
	Type           int       `json:"type"           gorm:"type:int2; not null; comment:存储库类型（1-SFTP；2-对象存储S3；3-本地存储）"`
	Uri            string    `json:"uri"            gorm:"type:varchar(100); not null; comment:存储库uri"`
	Password       string    `json:"password"       gorm:"type:varchar(100); not null; comment:存储库密码"`
	AccessUser     string    `json:"accessUser"     gorm:"type:varchar(100); comment:存储库访问用户"`
	AccessPassword string    `json:"accessPassword" gorm:"type:varchar(100); comment:存储库访问用户密码"`
	Status         int       `json:"status"         gorm:"type:int2; not null; comment:状态（1-正常；2-异常；3-禁用；4-未初始化）"`
}

func (b BackupRepository) TableName() string {
	return "backup_repository"
}

func (b BackupRepository) TableComment() string {
	return "存储库端管理"
}

func (b BackupRepository) InitData() any {
	return nil
}

type BackupTask struct {
	Id            int       `json:"id"            gorm:"type:int4; comment:ID;"`
	UpdateAt      time.Time `json:"updateAt"      gorm:"type:timestamp; not null;comment:更新时间"`
	RepositoryId  int       `json:"repositoryId"  gorm:"type:int4; not null; comment:存储库ID"`
	Name          string    `json:"name"          gorm:"type:varchar(100); not null; comment:任务名称"`
	Corn          string    `json:"corn"          gorm:"type:varchar(100); not null; comment:定时任务配置"`
	KeepDay       int       `json:"keepDay"       gorm:"type:int4; comment:保留天数"`
	KeepNum       int       `json:"keepNum"       gorm:"type:int4; comment:保留快照数"`
	Retry         int       `json:"retry"         gorm:"type:int4; comment:重试次数"`
	RetryInterval int       `json:"retryInterval" gorm:"type:int4; comment:重试间隔(秒)"`
	TaskTimeout   int       `json:"taskTimeout"   gorm:"type:int4; comment:任务过期时间(秒)"`
	LastTime      time.Time `json:"lastTime"      gorm:"type:timestamp; comment:上一次执行时间"`
	Disable       bool      `json:"disable"       gorm:"type:bool; not null; default:false; comment:是否禁用"`
}

func (b BackupTask) TableName() string {
	return "backup_task"
}

func (b BackupTask) TableComment() string {
	return "任务管理"
}

func (b BackupTask) InitData() any {
	return nil
}

type BackupTaskRecord struct {
	Id        int       `json:"id"        gorm:"type:int4; comment:ID;"`
	UpdateAt  time.Time `json:"updateAt"  gorm:"type:timestamp; not null;comment:更新时间"`
	TaskId    int       `json:"taskId"    gorm:"type:int4; not null; comment:任务ID"`
	Remark    string    `json:"remark"    gorm:"type:varchar(100); not null; comment:执行备注"`
	StartTime time.Time `json:"startTime" gorm:"type:timestamp; comment:任务开始执行时间"`
	EndTime   time.Time `json:"endTime"   gorm:"type:timestamp; comment:任务结束执行时间"`
	Status    int       `json:"status"    gorm:"type:int2; not null; comment:状态（1-进行中；2-成功；3-失败）"`
}

func (b BackupTaskRecord) TableName() string {
	return "backup_task_record"
}

func (b BackupTaskRecord) TableComment() string {
	return "任务记录管理"
}

func (b BackupTaskRecord) InitData() any {
	return nil
}

type BackupTaskStep struct {
	Id          int       `json:"id"          gorm:"type:int4; comment:ID;"`
	UpdateAt    time.Time `json:"updateAt"    gorm:"type:timestamp; not null;comment:更新时间"`
	TaskId      int       `json:"taskId"      gorm:"type:int4; not null; comment:任务ID"`
	ClientId    int       `json:"clientId"    gorm:"type:int4; not null; comment:客户端ID"`
	Category    int       `json:"category"    gorm:"type:int2; not null; comment:类型（1-备份；2-恢复；3-删除）"`
	DatabaseId  int       `json:"databaseId"  gorm:"type:int4; comment:数据库ID"`
	BackupCmd   string    `json:"backupCmd"   gorm:"type:varchar(2000); comment:备份命令"`
	BackupPath  string    `json:"backupPath"  gorm:"type:varchar(1000); comment:备份路径"`
	RestorePath string    `json:"restorePath" gorm:"type:varchar(1000); comment:恢复路径"`
	SnapshotId  string    `json:"snapshotId"  gorm:"type:varchar(100); comment:快照ID（用于恢复/删除）"`
	Disable     bool      `json:"disable"     gorm:"type:bool; not null; default:false; comment:是否禁用"`
}

func (b BackupTaskStep) TableName() string {
	return "backup_task_step"
}

func (b BackupTaskStep) TableComment() string {
	return "任务步骤管理"
}

func (b BackupTaskStep) InitData() any {
	return nil
}

type BackupTaskStepRecord struct {
	Id           int       `json:"id"            gorm:"type:int4; comment:ID;"`
	UpdateAt     time.Time `json:"updateAt"      gorm:"type:timestamp; not null;comment:更新时间"`
	StepId       int       `json:"stepId"        gorm:"type:int4; not null; comment:步骤ID"`
	TaskId       int       `json:"taskId"        gorm:"type:int4; not null; comment:任务ID"`
	TaskRecordId int       `json:"taskRecordId"  gorm:"type:int4; not null; comment:任务记录ID"`
	StartTime    time.Time `json:"startTime"     gorm:"type:timestamp; comment:开始执行时间"`
	EndTime      time.Time `json:"endTime"       gorm:"type:timestamp; comment:结束执行时间"`
	Status       int       `json:"status"        gorm:"type:int2; not null; comment:状态（1-进行中；2-成功；3-失败）"`
	ClientIp     string    `json:"clientIp"      gorm:"type:varchar(100); comment:客户端IP"`
	SnapshotId   string    `json:"snapshotId"    gorm:"type:varchar(100); comment:快照ID"`
	Cmd          string    `json:"cmd"           gorm:"type:text; comment:执行命令"`
	Result       string    `json:"result"        gorm:"type:text; comment:输出结果"`
}

func (b BackupTaskStepRecord) TableName() string {
	return "backup_task_step_record"
}

func (b BackupTaskStepRecord) TableComment() string {
	return "任务步骤记录管理"
}

func (b BackupTaskStepRecord) InitData() any {
	return nil
}
