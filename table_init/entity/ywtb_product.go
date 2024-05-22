package entity

// 板块
type Plate struct {
	Id             string `json:"id" gorm:"type:varchar(20); primary_key; comment:主键ID;"`
	Name           string `json:"name" gorm:"type:varchar(100); not null; comment:板块名称;"`
	Source         int    `json:"source" gorm:"type:smallint; default:1; comment:数据源（1-自定义/2-默认）;"`
	SourceId       string `json:"sourceId" gorm:"type:varchar(20); comment:源ID;"`
	SourceName     string `json:"sourceName" gorm:"type:varchar(100); comment:源名称;"`
	IconType       int    `json:"iconType" gorm:"type:smallint; comment:图标类型（1-自定义/2-默认）;"`
	Icon           string `json:"icon" gorm:"type:varchar(1000); comment:图标URL;"`
	IconChooseType int    `json:"iconChooseType" gorm:"type:smallint; comment:选中图标类型（1-自定义/2-默认））;"`
	IconChoose     string `json:"iconChoose" gorm:"type:varchar(1000); comment:选中图标URL;"`
	Index          int    `json:"index" gorm:"type:smallint; default:1; comment:排序下标;"`
	Status         int    `json:"status" gorm:"type:smallint; default:1; comment:状态（1-开启/2-关闭/3-默认）;"`
	TenantId       string `json:"tenantId" gorm:"type:varchar(100); default:default; comment:租户ID;"`
}

func (s Plate) TableName() string {
	return "test.t_plate"
}

func (s Plate) TableComment() string {
	return "板块设置"
}

func (s Plate) InitData() any {
	return nil
}

// 模块
type Module struct {
	Id         string `json:"id" gorm:"type:varchar(20); primary_key; comment:主键ID;"`
	PlateId    string `json:"plateId" gorm:"type:varchar(20); comment:板块ID;"`
	Name       string `json:"name" gorm:"type:varchar(100); not null; comment:模块名称;"`
	Source     int    `json:"source" gorm:"type:smallint; default:1; comment:数据源（1-自定义/2-默认）;"`
	SourceId   string `json:"sourceId" gorm:"type:varchar(20); comment:源ID;"`
	SourceName string `json:"sourceName" gorm:"type:varchar(100); comment:源名称;"`
	Layout     string `json:"layout" gorm:"type:varchar(20); not null; comment:布局;"`
	LayoutName string `json:"layoutName" gorm:"type:varchar(20); comment:布局名称;"`
	DisplayNum int    `json:"displayNum" gorm:"type:smallint; comment:显示数量;"`
	Frequency  int    `json:"frequency" gorm:"type:smallint; default:500; comment:轮播频率（单位ms）;"`
	Index      int    `json:"index" gorm:"type:smallint; default:1; comment:排序下标;"`
	Status     int    `json:"status" gorm:"type:smallint; default:1; comment:状态（1-开启/2-关闭/3-默认）;"`
	TenantId   string `json:"tenantId" gorm:"type:varchar(100); default:default; comment:租户ID;"`
}

func (s Module) TableName() string {
	return "test.t_module"
}

func (s Module) TableComment() string {
	return "模块设置"
}

func (s Module) InitData() any {
	return nil
}

// 模块内容
type Content struct {
	Id          string `json:"id" gorm:"type:varchar(20); primary_key; comment:主键ID;"`
	Name        string `json:"name" gorm:"type:varchar(1000); not null; comment:名称;"`
	ModuleId    string `json:"moduleId" gorm:"type:varchar(20); comment:模块ID;"`
	Description string `json:"description" gorm:"type:varchar(2000); comment:简介描述;"`
	Icon        string `json:"icon" gorm:"type:varchar(1000); comment:图标URL;"`
	UrlType     int    `json:"urlType" gorm:"type:smallint; not null; comment:链接类型（1-H5连接/2-内部链接/3-外部小程序）;"`
	Url         string `json:"url" gorm:"type:varchar(1000); comment:链接地址;"`
	AppId       string `json:"appId" gorm:"type:varchar(100); comment:外部小程序APPID;"`
	Index       int    `json:"index" gorm:"type:smallint; default:500; comment:排序下标;"`
	Status      int    `json:"status" gorm:"type:smallint; default:1; comment:状态（1-开启/2-关闭）;"`
	TenantId    string `json:"tenantId" gorm:"type:varchar(100); default:default; comment:租户ID;"`
}

func (s Content) TableName() string {
	return "test.t_module_content"
}

func (s Content) TableComment() string {
	return "模块内容设置"
}

func (s Content) InitData() any {
	return nil
}
