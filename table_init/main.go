package main

import (
	"github.com/go-xuan/quanx/app"
	"github.com/go-xuan/quanx/utils/randx"

	"table-init/entity"
)

func main() {
	var engine = app.NewEngine(
		app.NonGin,        // 不启用gin
		app.MultiDatabase, // 多数据源
	)
	randx.IdCard()

	engine.AddSourceTable("nkdx",
		&entity.ScreenIndicesStats{},
		&entity.EnrollStudentsStats{},
		&entity.SubjectConstruction{},
		&entity.IndicesValue{},
		&entity.EvaluationDepartmentRate{},
		&entity.EvaluationCourseScore{},
		&entity.EvaluationTeacherRatio{},
	)
	//
	//engine.AddSourceTable("hbkj",
	//	&entity.HbkjAppUser{},
	//	&entity.HbkjAppDept{},
	//	&entity.HbkjAppEmployUser{},
	//	&entity.HbkjAppEmploySequence{},
	//)

	//engine.AddSourceTable("ywtb",
	//	&entity.Plate{},
	//	&entity.Module{},
	//	&entity.Content{},
	//)

	//engine.AddSourceTable("ncu_portal",
	//	&entity.EduAppGrdaRsjbxx{},
	//	&entity.EduAppGrdaShjz{},
	//	&entity.EduAppGrdaGjrcJbxx{},
	//	&entity.EduAppGrdaGnwyxqk{},
	//	&entity.EduAppGrdaKhxx{},
	//	&entity.EduAppGrdaKqxx{},
	//	&entity.EduAppGrdaDkxx{},
	//	&entity.EduAppGrdaWdxs{},
	//	&entity.EduAppGrdaKyxm{},
	//	&entity.EduAppGrdaZscq{},
	//	&entity.EduAppGrdaCbzz{},
	//	&entity.EduAppGrdaFblw{},
	//	&entity.EduAppGrdaKyjx{},
	//	&entity.EduAppGrdaYjbg{},
	//	&entity.EduAppGrdaKyjf{},
	//	&entity.EduAppGrdaTsjy{},
	//	&entity.EduAppGrdaWdzc{},
	//)

	//engine.AddSourceTable("nmkj",
	//&entity.Lookup{},
	//&entity.Report{},
	//&entity.ReportLog{},
	//&entity.ReportAttachment{},
	//&entity.ReportApprover{},
	//&entity.ReportRepairman{},
	//&entity.ReportEvaluation{},
	//&entity.Repairman{},
	//&entity.RepairmanGroup{},
	//&entity.ReportNotice{},
	//)

	//engine.AddSourceTable("scoring",
	//	&entity.ScoringData1{},
	//	&entity.ScoringData2{},
	//	&entity.ScoringData3{},
	//)

	//engine.AddSourceTable("default",
	//	&entity.EduAppRlzyJbxxLx{},
	//	&entity.EduAppRlzyJbxxSjb{},
	//	&entity.EduAppRlzyJbxxSjbXzjl{},
	//	&entity.EduAppRlzyJbxxXzmb{},
	//	&entity.EduAppRlzyJbxxXzzd{},
	//	&entity.ConfirmTask{},
	//	&entity.ConfirmTaskPerson{},
	//	&entity.ConfirmTaskWorkflow{},
	//	&entity.ConfirmTaskRecord{},
	//	&entity.Scoring{},
	//	&entity.ScoringPerson{},
	//	&entity.ScoringIndicator{},
	//)

	//engine.AddSourceTable("lizhi",
	//	&entity.SchoolData{},
	//	&entity.SchoolAssetData{},
	//	&entity.TeachingData{},
	//	&entity.InternshipData{},
	//	&entity.TeacherData{},
	//	&entity.StudentData{},
	//	&entity.StudentGradeData{},
	//	&entity.CooperationData{},
	//	&entity.AchievementData{},
	//	&entity.AwardData{},
	//	&entity.HonorData{},
	//	&entity.BasicSchoolValue{},
	//)

	//// 安徽馆-物业管理
	//engine.AddSourceTable("property",
	//	&table.Staff{},
	//	&table.Department{},
	//	&table.Post{},
	//	&table.Equipment{},
	//	&table.EquipmentAlarm{},
	//	&table.EquipmentRepair{},
	//	&table.EquipmentInspect{},
	//	&table.EquipmentStatus{},
	//)

	// 安徽馆
	//engine.AddSourceTable("default",
	//	&table.Benefit{},
	//	&table.BenefitObject{},
	//	&table.Level{},
	//	&table.Recommendation{},
	//	&table.User{},
	//	&table.LoginLog{},
	//	&table.ApiLog{},
	//)

	// 教育部大屏
	//engine.AddSourceTable("default",
	//	&table.AppSswcySyfb{},       // 十四五产业-省域分布
	//	&table.AppSswcySysltj{},     // 十四五产业-省域产业数量统计
	//	&table.AppSswcySyslfb{},     // 十四五产业-省域产业数量分布
	//	&table.AppSswcyZcqkGszcd{},  // 十四五产业-支撑情况-各省支撑度
	//	&table.AppSswcyZcqkZhzcd{},  // 十四五产业-支撑情况-综合支撑度
	//	&table.AppSswcyZcqkZcqktj{}, // 十四五产业-支撑情况-统计
	//	&table.AppSswcyHxzcxkzy{},   // 十四五产业-核心支撑学科专业
	//)

	//教育部大屏
	//engine.AddSourceTable("default",
	//	&table.AppSswcyGl{},          // 十四五产业
	//	&table.AppSswcyXkzy{},        // 十四五产业-学科专业
	//	&table.AppSswcySzsf{},        // 十四五产业-所在省份
	//	&table.AppSswcyYjsjyZcqk{},   // 十四五产业-研究生教育-支撑情况
	//	&table.AppSswcyXkpg{},        // 十四五产业-学科评估
	//	&table.AppSswcyQxgxkZyjsqk{}, // 十四五产业-强相关学科-专业建设情况
	//	&table.AppSswcyQxgxkXsslqk{}, // 十四五产业-强相关学科-学生数量情况
	//	&table.AppSswcyQxgxkJyqk{},   // 十四五产业-强相关学科-就业情况
	//)
	// 服务启动
	engine.RUN()
}
