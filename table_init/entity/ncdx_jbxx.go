package entity

type EduAppGrdaRsjbxx struct {
	Uid    string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Xm     string `json:"xm" gorm:"type:varchar(100); not null; comment:姓名;"`
	Zgh    string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Xb     string `json:"xb" gorm:"type:varchar(10); comment:性别;"`
	Mz     string `json:"mz" gorm:"type:varchar(20); comment:民族;"`
	Sfzh   string `json:"sfzh" gorm:"type:varchar(20); comment:身份证号;"`
	Csrq   string `json:"csrq" gorm:"type:date; comment:出生日期;"`
	Jg     string `json:"jg" gorm:"type:varchar(50); comment:籍贯;"`
	Zzmm   string `json:"zzmm" gorm:"type:varchar(20); comment:政治面貌;"`
	Jrsj   string `json:"jrsj" gorm:"type:timestamp(0); comment:加入时间;"`
	Xl     string `json:"xl" gorm:"type:varchar(20); comment:学历;"`
	Xw     string `json:"xw" gorm:"type:varchar(100); comment:学位;"`
	ZcZr   string `json:"zc_zr" gorm:"type:varchar(100); comment:职称_专任;"`
	Bkyx   string `json:"bkyx" gorm:"type:varchar(100); comment:本科院校;"`
	Bkzy   string `json:"bkzy" gorm:"type:varchar(100); comment:本科专业;"`
	Ssyx   string `json:"ssyx" gorm:"type:varchar(100); comment:硕士院校;"`
	Sszy   string `json:"sszy" gorm:"type:varchar(100); comment:硕士专业;"`
	Bsyx   string `json:"bsyx" gorm:"type:varchar(100); comment:博士院校;"`
	Bszy   string `json:"bszy" gorm:"type:varchar(100); comment:博士专业;"`
	Xy     string `json:"xy" gorm:"type:varchar(100); comment:学缘;"`
	Rzrq   string `json:"rzrq" gorm:"type:date; comment:入职日期;"`
	Lzrq   string `json:"lzrq" gorm:"type:date; comment:离职日期;"`
	Jzglx  string `json:"jzglx" gorm:"type:varchar(100); comment:教职工类型;"`
	Gwlx   string `json:"gwlx" gorm:"type:varchar(100); comment:岗位类型;"`
	Gwjb   string `json:"gwjb" gorm:"type:varchar(100); comment:岗位级别;"`
	Xzzw   string `json:"xzzw" gorm:"type:varchar(100); comment:行政职务;"`
	Bmmc   string `json:"bmmc" gorm:"type:varchar(100); comment:部门名称;"`
	Ksmc   string `json:"ksmc" gorm:"type:varchar(100); comment:科室名称;"`
	Dqzt   string `json:"dqzt" gorm:"type:varchar(100); comment:当前状态;"`
	Txdz   string `json:"txdz" gorm:"type:varchar(200); comment:通讯地址;"`
	Sfssx  string `json:"sfssx" gorm:"type:varchar(100); comment:是否双师型;"`
	Zzxx   string `json:"zzxx" gorm:"type:varchar(100); comment:在职信息;"`
	Sfzx   string `json:"sfzx" gorm:"type:varchar(100); comment:是否在校;"`
	Zyjszw string `json:"zyjszw" gorm:"type:varchar(100); comment:专业技术职务;"`
	Bzxx   string `json:"bzxx" gorm:"type:varchar(100); comment:编制信息;"`
	Prgw   string `json:"prgw" gorm:"type:varchar(100); comment:聘任岗位;"`
}

func (e EduAppGrdaRsjbxx) TableName() string {
	return "portal.edu_app_grda_rsjbxx"
}

func (e EduAppGrdaRsjbxx) TableComment() string {
	return "个人档案-人事基本信息"
}

func (e EduAppGrdaRsjbxx) InitData() any {
	return nil
}

type EduAppGrdaShjz struct {
	Uid      string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh      string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	ShjzMc   string `json:"shjz_mc" gorm:"type:varchar(100); comment:社会兼职名称;"`
	ShjzQsrq string `json:"shjz_qsrq" gorm:"type:date; comment:社会兼职起始日期;"`
	ShjzJsrq string `json:"shjz_jsrq" gorm:"type:date; comment:社会兼职结束日期;"`
	ShjzCzyy string `json:"shjz_czyy" gorm:"type:varchar(1000); comment:社会兼职辞职原因;"`
	ShjzZwmc string `json:"shjz_zwmc" gorm:"type:varchar(100); comment:兼职职务名称;"`
}

func (e EduAppGrdaShjz) TableName() string {
	return "portal.edu_app_grda_shjz"
}

func (e EduAppGrdaShjz) TableComment() string {
	return "个人档案-社会兼职"
}

func (e EduAppGrdaShjz) InitData() any {
	return nil
}

type EduAppGrdaGjrcJbxx struct {
	Uid      string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh      string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Rcch     string `json:"rcch" gorm:"type:varchar(100); comment:人才称号;"`
	ZjchJbmc string `json:"zjch_jbmc" gorm:"type:varchar(100); comment:专家称号级别名称;"`
	ZjchHdsj string `json:"zjch_hdsj" gorm:"type:timestamp(0); comment:专家称号获得时间;"`
	Bxprny   string `json:"bxprny" gorm:"type:varchar(20); comment:本校聘任年月;"`
	Zjlxmc   string `json:"zjlxmc" gorm:"type:varchar(100); comment:专家类型名称;"`
}

func (e EduAppGrdaGjrcJbxx) TableName() string {
	return "portal.edu_app_grda_gjrc_jbxx"
}

func (e EduAppGrdaGjrcJbxx) TableComment() string {
	return "个人档案-高级人才-基本信息"
}

func (e EduAppGrdaGjrcJbxx) InitData() any {
	return nil
}

type EduAppGrdaGnwyxqk struct {
	Uid      string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh      string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Rq       string `json:"rq" gorm:"type:date; comment:日期;"`
	Yxfx     string `json:"yxfx" gorm:"type:varchar(100); comment:研修方向;"`
	SqdwZwmc string `json:"sqdw_zwmc" gorm:"type:varchar(100); comment:所去单位中文名称;"`
	Pcdw     string `json:"pcdw" gorm:"type:varchar(100); comment:派出单位;"`
}

func (e EduAppGrdaGnwyxqk) TableName() string {
	return "portal.edu_app_grda_gnwyxqk"
}

func (e EduAppGrdaGnwyxqk) TableComment() string {
	return "个人档案-国内外研修情况"
}

func (e EduAppGrdaGnwyxqk) InitData() any {

	return nil
}

type EduAppGrdaKhxx struct {
	Uid  string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Khnf string `json:"khnf" gorm:"type:varchar(10); comment:考核年份;"`
	Gwlx string `json:"gwlx" gorm:"type:varchar(100); comment:岗位类型;"`
	Khjg string `json:"khjg" gorm:"type:varchar(200); comment:考核结果;"`
}

func (e EduAppGrdaKhxx) TableName() string {
	return "portal.edu_app_grda_khxx"
}

func (e EduAppGrdaKhxx) TableComment() string {
	return "个人档案-考核信息"
}

func (e EduAppGrdaKhxx) InitData() any {
	return nil
}

type EduAppGrdaKqxx struct {
	Uid  string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Rq   string `json:"rq" gorm:"type:date; comment:日期;"`
	Kqjg string `json:"kqjg" gorm:"type:varchar(100); comment:考勤结果;"`
	Dkdd string `json:"dkdd" gorm:"type:varchar(200); comment:打卡地点;"`
	Dksj string `json:"dksj" gorm:"type:timestamp(0); comment:打卡时间;"`
}

func (e EduAppGrdaKqxx) TableName() string {
	return "portal.edu_app_grda_kqxx"
}

func (e EduAppGrdaKqxx) TableComment() string {
	return "个人档案-考勤信息"
}

func (e EduAppGrdaKqxx) InitData() any {
	return nil
}

type EduAppGrdaDkxx struct {
	Uid  string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Nd   string `json:"nd" gorm:"type:varchar(20); comment:年度;"`
	Xq   string `json:"xq" gorm:"type:varchar(20); comment:学期;"`
	Kcmc string `json:"kcmc" gorm:"type:varchar(100); comment:课程名称;"`
	Ks   int    `json:"ks" gorm:"type:int; comment:课时;"`
}

func (e EduAppGrdaDkxx) TableName() string {
	return "portal.edu_app_grda_dkxx"
}

func (e EduAppGrdaDkxx) TableComment() string {
	return "个人档案-带课信息"
}

func (e EduAppGrdaDkxx) InitData() any {
	return nil
}

type EduAppGrdaWdxs struct {
	Uid  string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Xm   string `json:"xm" gorm:"type:varchar(100); not null; comment:学生姓名;"`
	Sfzh string `json:"sfzh" gorm:"type:varchar(100); comment:身份证号;"`
	Xb   string `json:"xb" gorm:"type:varchar(100); comment:性别;"`
	Mz   string `json:"mz" gorm:"type:varchar(100); comment:民族;"`
	Csrq string `json:"csrq" gorm:"type:date; comment:出生日期;"`
	Jg   string `json:"jg" gorm:"type:varchar(20); comment:籍贯;"`
	Yx   string `json:"yx" gorm:"type:varchar(100); comment:院系;"`
	Zy   string `json:"zy" gorm:"type:varchar(100); comment:专业;"`
	Rxny string `json:"rxny" gorm:"type:varchar(20); comment:入学年月;"`
	Nj   string `json:"nj" gorm:"type:varchar(100); comment:年级;"`
}

func (e EduAppGrdaWdxs) TableName() string {
	return "portal.edu_app_grda_wdxs"
}

func (e EduAppGrdaWdxs) TableComment() string {
	return "个人档案-我的学生"
}

func (e EduAppGrdaWdxs) InitData() any {
	return nil
}

type EduAppGrdaKyxm struct {
	Uid   string  `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh   string  `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Xmmc  string  `json:"xmmc" gorm:"type:varchar(100); comment:项目名称;"`
	Xmbh  string  `json:"xmbh" gorm:"type:varchar(100); comment:项目编号;"`
	Fzrxm string  `json:"fzrxm" gorm:"type:varchar(100); comment:负责人姓名;"`
	Htjf  float64 `json:"htjf" gorm:"type:numeric; comment:合同经费（元）;"`
	Ksrq  string  `json:"ksrq" gorm:"type:date; comment:开始日期;"`
	Lxrq  string  `json:"lxrq" gorm:"type:date; comment:立项日期;"`
}

func (e EduAppGrdaKyxm) TableName() string {
	return "portal.edu_app_grda_kyxm"
}

func (e EduAppGrdaKyxm) TableComment() string {
	return "个人档案-科研项目"
}

func (e EduAppGrdaKyxm) InitData() any {
	return nil
}

type EduAppGrdaZscq struct {
	Uid    string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh    string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Zlmc   string `json:"zlmc" gorm:"type:varchar(100); comment:专利名称;"`
	Zllx   string `json:"zllx" gorm:"type:varchar(100); comment:专利类型;"`
	Zlzt   string `json:"zlzt" gorm:"type:varchar(100); comment:专利状态;"`
	Dyfmr  string `json:"dyfmr" gorm:"type:varchar(100); comment:第一发明人;"`
	Dyfmdw string `json:"dyfmdw" gorm:"type:varchar(100); comment:第一发明单位;"`
	Sqsj   string `json:"sqsj" gorm:"type:timestamp(0); comment:申请时间;"`
}

func (e EduAppGrdaZscq) TableName() string {
	return "portal.edu_app_grda_zscq"
}

func (e EduAppGrdaZscq) TableComment() string {
	return "个人档案-知识产权"
}

func (e EduAppGrdaZscq) InitData() any {
	return nil
}

type EduAppGrdaCbzz struct {
	Uid   string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh   string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Dyzz  string `json:"dyzz" gorm:"type:varchar(100); comment:第一作者;"`
	Zzmc  string `json:"zzmc" gorm:"type:varchar(100); comment:著作名称;"`
	Zzlx  string `json:"zzlx" gorm:"type:varchar(100); comment:著作类型;"`
	Cbdw  string `json:"cbdw" gorm:"type:varchar(100); comment:出版单位;"`
	Cbsjb string `json:"cbsjb" gorm:"type:varchar(100); comment:出版社级别;"`
	Cbh   string `json:"cbh" gorm:"type:varchar(100); comment:出版号;"`
}

func (e EduAppGrdaCbzz) TableName() string {
	return "portal.edu_app_grda_cbzz"
}

func (e EduAppGrdaCbzz) TableComment() string {
	return "个人档案-出版著作"
}

func (e EduAppGrdaCbzz) InitData() any {
	return nil
}

type EduAppGrdaFblw struct {
	Uid    string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh    string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Lwbt   string `json:"lwbt" gorm:"type:varchar(100); comment:论文标题;"`
	Cbn    string `json:"cbn" gorm:"type:varchar(100); comment:出版年;"`
	Ly     string `json:"ly" gorm:"type:varchar(100); comment:来源;"`
	Zzcsjg string `json:"zzcsjg" gorm:"type:varchar(100); comment:作者从属机构;"`
	Sfdydw int    `json:"sfdydw" gorm:"type:smallint; comment:是否第一单位;"`
	Syh    string `json:"syh" gorm:"type:varchar(100); comment:索引号;"`
}

func (e EduAppGrdaFblw) TableName() string {
	return "portal.edu_app_grda_fblw"
}

func (e EduAppGrdaFblw) TableComment() string {
	return "个人档案-发表论文"
}

func (e EduAppGrdaFblw) InitData() any {
	return nil
}

type EduAppGrdaKyjx struct {
	Uid     string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh     string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Hjmc    string `json:"hjmc" gorm:"type:varchar(100); comment:获奖名称;"`
	Dyhjrxm string `json:"dyhjrxm" gorm:"type:varchar(100); comment:第一获奖人姓名;"`
	Ssdw    string `json:"ssdw" gorm:"type:varchar(100); comment:所属单位;"`
	Zsbh    string `json:"zsbh" gorm:"type:varchar(100); comment:获奖等级;"`
	Hjrq    string `json:"hjrq" gorm:"type:varchar(100); comment:证书编号;"`
}

func (e EduAppGrdaKyjx) TableName() string {
	return "portal.edu_app_grda_kyjx"
}

func (e EduAppGrdaKyjx) TableComment() string {
	return "个人档案-科研奖项"
}

func (e EduAppGrdaKyjx) InitData() any {
	return nil
}

type EduAppGrdaYjbg struct {
	Uid     string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh     string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Dyzzxm  string `json:"dyzzxm" gorm:"type:varchar(100); comment:第一作者姓名;"`
	Dyzzzgh string `json:"dyzzzgh" gorm:"type:varchar(100); comment:第一作者职工号;"`
	Dyzzzc  string `json:"dyzzzc" gorm:"type:varchar(100); comment:第一作者职称;"`
	Dyzzxb  string `json:"dyzzxb" gorm:"type:varchar(100); comment:第一作者性别;"`
	Dyzzxl  string `json:"dyzzxl" gorm:"type:varchar(100); comment:第一作者学历;"`
	Dyzzxw  string `json:"dyzzxw" gorm:"type:varchar(100); comment:第一作者学位;"`
	Ssjys   string `json:"ssjys" gorm:"type:varchar(100); comment:所属教研室;"`
	Zzs     int    `json:"zzs" gorm:"type:int; comment:作者数;"`
	Sfldps  int    `json:"sfldps" gorm:"type:smallint; comment:是否领导批示;"`
	Cndw    string `json:"cndw" gorm:"type:varchar(100); comment:采纳单位;"`
}

func (e EduAppGrdaYjbg) TableName() string {
	return "portal.edu_app_grda_yjbg"
}

func (e EduAppGrdaYjbg) TableComment() string {
	return "个人档案-研究报告"
}

func (e EduAppGrdaYjbg) InitData() any {
	return nil
}

type EduAppGrdaKyjf struct {
	Uid  string  `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string  `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Xmmc string  `json:"xmmc" gorm:"type:varchar(100); comment:项目名称;"`
	Xmbh string  `json:"xmbh" gorm:"type:varchar(100); comment:项目编号;"`
	Dzje float64 `json:"dzje" gorm:"type:numeric; comment:到账金额（元）;"`
	Dzsj string  `json:"dzsj" gorm:"type:timestamp(0); comment:到账时间;"`
	Wbje float64 `json:"wbje" gorm:"type:numeric; comment:外拨金额（元）;"`
	Bkdw string  `json:"bkdw" gorm:"type:varchar(100); comment:拨款单位;"`
}

func (e EduAppGrdaKyjf) TableName() string {
	return "portal.edu_app_grda_kyjf"
}

func (e EduAppGrdaKyjf) TableComment() string {
	return "个人档案-科研经费"
}

func (e EduAppGrdaKyjf) InitData() any {
	return nil
}

type EduAppGrdaTsjy struct {
	Uid  string `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Sjmc string `json:"sjmc" gorm:"type:varchar(200); comment:书籍名称;"`
	Jysj string `json:"jysj" gorm:"type:timestamp(0); comment:借阅时间;"`
	Hssj string `json:"hssj" gorm:"type:timestamp(0); comment:还书时间;"`
}

func (e EduAppGrdaTsjy) TableName() string {
	return "portal.edu_app_grda_tsjy"
}

func (e EduAppGrdaTsjy) TableComment() string {
	return "个人档案-图书借阅"
}

func (e EduAppGrdaTsjy) InitData() any {
	return nil
}

type EduAppGrdaWdzc struct {
	Uid  string  `json:"uid" gorm:"type:varchar(50); not null; comment:UID;"`
	Zgh  string  `json:"zgh" gorm:"type:varchar(50); not null; comment:职工号;"`
	Zcmc string  `json:"zcmc" gorm:"type:varchar(100); comment:资产名称;"`
	Zclb string  `json:"zclb" gorm:"type:varchar(100); comment:资产类别;"`
	Cfdd string  `json:"cfdd" gorm:"type:varchar(200); comment:存放地点;"`
	Sl   int     `json:"sl" gorm:"type:int; comment:数量;"`
	Jg   float64 `json:"jg" gorm:"type:numeric; comment:价格（元）;"`
}

func (e EduAppGrdaWdzc) TableName() string {
	return "portal.edu_app_grda_wdzc"
}

func (e EduAppGrdaWdzc) TableComment() string {
	return "个人档案-我的资产"
}

func (e EduAppGrdaWdzc) InitData() any {
	return nil
}
