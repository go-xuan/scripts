package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-xuan/quanx/os/filex"
)

type ResticCmdParam struct {
	Repository     *Repository `json:"repository"     description:"存储库信息"`
	Path           string      `json:"path"           description:"路径"`
	Tag            string      `json:"tag"            description:"标签"`
	Hostname       string      `json:"hostname"       description:"主机名称"`
	SnapshotId     string      `json:"snapshotId"     description:"快照ID（用于恢复/删除）"`
	ForgetStrategy string      `json:"forgetStrategy" description:"删除策略"`
}

// Backup 备份快照命令
// eg: restic backup -r local_repo --password-command 'echo 123' --verbose data_b
// restic backup -r local_repo --verbose --password-command 'echo 123' --stdin --stdin-filename data_a/mysql_nmkj_lightapps.sql
func (param *ResticCmdParam) Backup() string {
	var cmd = strings.Builder{}
	cmd.WriteString(param.Repository.SetEnv())
	cmd.WriteString(`restic backup `)
	cmd.WriteString(filex.Pwd(param.Path))
	cmd.WriteString(` --verbose `)
	if param.Hostname != "" {
		cmd.WriteString(` --host `)
		cmd.WriteString(param.Hostname)
	}
	if param.Tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(param.Tag)
	}
	return cmd.String()
}

// Forget 策略删除快照命令
// --keep-last n						# 保留最近的n个快照。
// --keep-hourly n						# 对于有一个或多个快照的最后n几个小时，每个小时只保留最近的一个。
// --keep-daily n	 					# 对于有一个或多个快照的最后n几天，每天只保留最近的一个。
// --keep-weekly n 						# 对于有一个或多个快照的最后n几周，每周只保留最近的一个。
// --keep-monthly n 					# 对于有一个或多个快照的最后n几个月，每个月只保留最近的一个。
// --keep-yearly n 						# 对于有一个或多个快照的最后n几年，每年只保留最近的一个。
// --keep-tag 							# 保留具有此标签的所有快照，可以指定多个标签。
// --keep-within [duration] 			# 将所有具有时间戳的快照保留在最新快照的指定持续时间内
// --keep-within-hourly [duration] 		# 保留在最新快照的指定持续时间内制作的所有每小时快照。
// --keep-within-daily [duration] 		# 保留在最新快照的指定持续时间内制作的所有每日快照。
// --keep-within-weekly [duration]  	# 保留在最新快照的指定持续时间内制作的所有每周快照。
// --keep-within-monthly [duration]  	# 保留在最新快照的指定持续时间内制作的所有月度快照。
// --keep-within-yearly [duration]		# 保留在最新快照的指定持续时间内制作的所有年度快照
// eg: restic forget  -r local_repo --password-command 'echo 123' --verbose --prune --keep-last 1
func (param *ResticCmdParam) Forget() string {
	var cmd = strings.Builder{}
	cmd.WriteString(param.Repository.SetEnv())
	cmd.WriteString(`restic forget `)
	if param.SnapshotId != "" {
		cmd.WriteString(param.SnapshotId)
	} else {
		cmd.WriteString(param.ForgetStrategy)
	}
	cmd.WriteString(` --verbose --prune `)
	if param.Hostname != "" {
		cmd.WriteString(` --host `)
		cmd.WriteString(param.Hostname)
	}
	if param.Tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(param.Tag)
	}
	return cmd.String()
}

// Restore 恢复快照命令
func (param *ResticCmdParam) Restore() string {
	var cmd = strings.Builder{}
	cmd.WriteString(param.Repository.SetEnv())
	cmd.WriteString(`restic restore `)
	if param.SnapshotId != "" {
		cmd.WriteString(param.SnapshotId)
	} else {
		cmd.WriteString(`latest`)
	}
	cmd.WriteString(` --target `)
	cmd.WriteString(filex.Pwd(param.Path))
	cmd.WriteString(` --verbose `)
	if param.Hostname != "" {
		cmd.WriteString(` --host `)
		cmd.WriteString(param.Hostname)
	}
	if param.Tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(param.Tag)
	}
	return cmd.String()
}

type Repository struct {
	Uri                string `json:"uri" comment:"存储库uri"`
	Password           string `json:"password" comment:"存储库密码"`
	AwsAccessKeyId     string `json:"awsAccessKeyId" comment:"AmazonS3秘钥ID（用于配置当前客户端环境变量）"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey" comment:"AmazonS3秘钥（用于配置当前客户端环境变量）"`
}

// 设置环境变量
func (r *Repository) SetEnv() string {
	switch runtime.GOOS {
	case "linux":
		return fmt.Sprintf(
			`export RESTIC_REPOSITORY="%s";export RESTIC_PASSWORD="%s";export AWS_ACCESS_KEY_ID="%s";export AWS_SECRET_ACCESS_KEY="%s";`,
			r.Uri, r.Password, r.AwsAccessKeyId, r.AwsSecretAccessKey)
	case "windows":
		return fmt.Sprintf(
			`set RESTIC_REPOSITORY=%s&&set RESTIC_PASSWORD=%s&&set AWS_ACCESS_KEY_ID=%s&&set AWS_SECRET_ACCESS_KEY=%s&&`,
			r.Uri, r.Password, r.AwsAccessKeyId, r.AwsSecretAccessKey)
	default:
		return ""
	}
}

// 初始化存储库
func (r *Repository) Init() string {
	var cmd = strings.Builder{}
	cmd.WriteString(r.SetEnv())
	cmd.WriteString(`restic init --verbose `)
	return cmd.String()
}

// 存储库快照查询
func (r *Repository) Snapshots() string {
	var cmd = strings.Builder{}
	cmd.WriteString(r.SetEnv())
	cmd.WriteString(`restic snapshots --json`)
	return cmd.String()
}
