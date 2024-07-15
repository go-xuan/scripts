package model

import (
	"restic_backup/common"
	"strings"

	"github.com/go-xuan/quanx/os/filex"
)

type Restic struct {
	Do         string      `json:"do"         yaml:"do"         comment:"执行命令"`
	Backup     *Backup     `json:"backup"     yaml:"backup"     comment:"备份"`
	Restore    *Restore    `json:"restore"    yaml:"restore"    comment:"恢复"`
	Forget     *Forget     `json:"forget"     yaml:"forget"     comment:"删除"`
	Repository *Repository `json:"repository" yaml:"repository" comment:"存储库信息"`
}

// 生成restic命令
func (r *Restic) Command() string {
	var cmd = strings.Builder{}
	cmd.WriteString(r.Repository.SetEnv())
	switch strings.ToLower(r.Do) {
	case common.Backup:
		cmd.WriteString(r.Backup.Command())
	case common.Restore:
		cmd.WriteString(r.Restore.Command())
	case common.Forget:
		cmd.WriteString(r.Forget.Command())
	case common.Init:
		cmd.WriteString(`restic init --verbose`)
	default:
		cmd.WriteString(`restic snapshots --json`)
	}
	return cmd.String()
}

// 备份命令
type Backup struct {
	Host string `json:"host" yaml:"host" comment:"主机名称"`
	Tag  string `json:"tag"  yaml:"tag"  comment:"标签"`
	Path string `json:"path" yaml:"path" comment:"备份路径"`
}

func (b *Backup) Command() string {
	var cmd = strings.Builder{}
	cmd.WriteString(`restic backup `)
	cmd.WriteString(filex.Pwd(b.Path))
	cmd.WriteString(` --verbose `)
	if b.Host != "" {
		cmd.WriteString(` --host `)
		cmd.WriteString(b.Host)
	}
	if b.Tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(b.Tag)
	}
	return cmd.String()
}

// 恢复命令
type Restore struct {
	Host     string `json:"host"     yaml:"host"     comment:"主机名称"`
	Tag      string `json:"tag"      yaml:"tag"      comment:"标签"`
	Path     string `json:"path"     yaml:"path"     comment:"备份路径"`
	Target   string `json:"target"   yaml:"target"   comment:"恢复路径"`
	Snapshot string `json:"snapshot" yaml:"snapshot" comment:"恢复快照"`
}

func (r *Restore) Command() string {
	var cmd = strings.Builder{}
	cmd.WriteString(`restic restore `)
	if r.Snapshot != "" {
		cmd.WriteString(r.Snapshot)
	} else {
		cmd.WriteString(` latest `)
	}
	cmd.WriteString(` --target `)
	cmd.WriteString(filex.Pwd(r.Target))
	cmd.WriteString(` --verbose `)
	if r.Host != "" {
		cmd.WriteString(` --host `)
		cmd.WriteString(r.Host)
	}
	if r.Tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(r.Tag)
	}
	return cmd.String()
}

// 删除命令
type Forget struct {
	Host     string    `json:"host"     yaml:"host"     comment:"主机名称"`
	Tag      string    `json:"tag"      yaml:"tag"      comment:"标签"`
	Snapshot string    `json:"snapshot" yaml:"snapshot" comment:"删除快照"`
	Strategy *Strategy `json:"strategy" yaml:"strategy" comment:"删除策略"`
}

func (r *Forget) Command() string {
	var cmd = strings.Builder{}
	cmd.WriteString(`restic forget `)
	if r.Snapshot != "" {
		cmd.WriteString(r.Snapshot)
	} else if r.Strategy != nil {
		cmd.WriteString(r.Strategy.String())
	} else {
		cmd.WriteString(` --keep-last 3 `)
	}
	cmd.WriteString(` --verbose --prune `)
	if r.Host != "" {
		cmd.WriteString(` --host `)
		cmd.WriteString(r.Host)
	}
	if r.Tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(r.Tag)
	}
	return cmd.String()
}

// --keep-last n                    keep the last n snapshots (use 'unlimited' to keep all snapshots)
// --keep-hourly n                  keep the last n hourly snapshots (use 'unlimited' to keep all hourly snapshots)
// --keep-daily n                   keep the last n daily snapshots (use 'unlimited' to keep all daily snapshots)
// --keep-weekly n                  keep the last n weekly snapshots (use 'unlimited' to keep all weekly snapshots)
// --keep-monthly n                 keep the last n monthly snapshots (use 'unlimited' to keep all monthly snapshots)
// --keep-yearly n                  keep the last n yearly snapshots (use 'unlimited' to keep all yearly snapshots)
// --keep-within duration           keep snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot
// --keep-within-hourly duration    keep hourly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot
// --keep-within-daily duration     keep daily snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot
// --keep-within-weekly duration    keep weekly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot
// --keep-within-monthly duration   keep monthly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot
// --keep-within-yearly duration    keep yearly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot
// --keep-tag tag-list              keep snapshots with this tag0list (can be specified multiple times) (default [])
type Strategy struct {
	KeepLast          string `json:"keep-last"           comment:"保留最近n个快照"`
	KeepHourly        string `json:"keep-hourly"         comment:"保留近n小时内的快照（使用“unlimited”保留所有小时快照）"`
	KeepDaily         string `json:"keep-daily"          comment:"保留近n天内的快照（使用“unlimited”保留所有每日快照）"`
	KeepWeekly        string `json:"keep-weekly"         comment:"保留近n周内的快照（使用“unlimited”保留所有每周快照）"`
	KeepMonthly       string `json:"keep-monthly"        comment:"保留近n月内的快照（使用“unlimited”保留所有每周快照）"`
	KeepYearly        string `json:"keep-yearly"         comment:"保留近n年内的快照（使用“unlimited”保留所有每周快照）"`
	KeepTag           string `json:"keep-tag"            comment:"根据此标签列表保留快照（可以指定多个）（默认[]）"`
	KeepWithin        string `json:"keep-within"         comment:"相对于最新快照，保留时间范围内（例如1y5m7d2h）的快照"`
	KeepWithinHourly  string `json:"keep-within-hourly"  comment:"相对于最新快照，保留时间范围内（例如1y5m7d2h）的每小时快照"`
	KeepWithinDaily   string `json:"keep-within-daily"   comment:"相对于最新快照，保留时间范围内（例如1y5m7d2h）的每日快照"`
	KeepWithinWeekly  string `json:"keep-within-weekly"  comment:"相对于最新快照，保留时间范围内（例如1y5m7d2h）的每周快照"`
	KeepWithinMonthly string `json:"keep-within-monthly" comment:"相对于最新快照，保留时间范围内（例如1y5m7d2h）的每月快照"`
	KeepWithinYearly  string `json:"keep-within-yearly"  comment:"相对于最新快照，保留时间范围内（例如1y5m7d2h）的每年快照"`
}

func (fs *Strategy) String() string {
	var sb = strings.Builder{}
	if fs.KeepLast != "" {
		sb.WriteString(` --keep-last `)
		sb.WriteString(fs.KeepLast)
	}
	if fs.KeepHourly != "" {
		sb.WriteString(` --keep-hourly `)
		sb.WriteString(fs.KeepHourly)
	}
	if fs.KeepDaily != "" {
		sb.WriteString(` --keep-daily `)
		sb.WriteString(fs.KeepDaily)
	}
	if fs.KeepWeekly != "" {
		sb.WriteString(` --keep-weekly `)
		sb.WriteString(fs.KeepWeekly)
	}
	if fs.KeepMonthly != "" {
		sb.WriteString(` --keep-monthly `)
		sb.WriteString(fs.KeepMonthly)
	}
	if fs.KeepYearly != "" {
		sb.WriteString(` --keep-yearly `)
		sb.WriteString(fs.KeepYearly)
	}
	if fs.KeepTag != "" {
		sb.WriteString(` --keep-tag `)
		sb.WriteString(fs.KeepTag)
	}
	if fs.KeepWithin != "" {
		sb.WriteString(` --keep-within `)
		sb.WriteString(fs.KeepWithin)
	}
	if fs.KeepWithinHourly != "" {
		sb.WriteString(` --keep-within-hourly `)
		sb.WriteString(fs.KeepWithinHourly)
	}
	if fs.KeepWithinDaily != "" {
		sb.WriteString(` --keep-within-daily `)
		sb.WriteString(fs.KeepWithinDaily)
	}
	if fs.KeepWithinWeekly != "" {
		sb.WriteString(` --keep-within-weekly `)
		sb.WriteString(fs.KeepWithinWeekly)
	}
	if fs.KeepWithinMonthly != "" {
		sb.WriteString(` --keep-within-monthly `)
		sb.WriteString(fs.KeepWithinMonthly)
	}
	if fs.KeepWithinYearly != "" {
		sb.WriteString(` --keep-within-yearly `)
		sb.WriteString(fs.KeepWithinYearly)
	}
	return sb.String()
}
