package main

import (
	"errors"
	"fmt"
	"github.com/go-xuan/quanx/os/execx"
	"github.com/go-xuan/quanx/os/filex"
	"github.com/go-xuan/quanx/os/taskx"
	"github.com/go-xuan/quanx/utils/marshalx"
	"os"
	"path/filepath"
	"regexp"
	"restic_backup/common"
	"time"

	"restic_backup/model"
)

type Config struct {
	Restic     *model.Restic     `json:"restic"     yaml:"restic"     comment:"restic配置"`
	Datasource *model.Datasource `json:"datasource" yaml:"datasource" comment:"备份数据源"`
	Retry      *Retry            `json:"retry"      yaml:"retry"      comment:"重试机制"`
	Cron       string            `json:"cron"       yaml:"cron"       comment:"定时任务"`
}

// 重试机制
type Retry struct {
	Times    int `json:"times"    yaml:"times"    comment:"重试次数"`
	Interval int `json:"interval" yaml:"interval" comment:"重试间隔(秒)"`
}

func main() {
	var config = &Config{}
	if err := marshalx.UnmarshalFromFile("config.yaml", config); err != nil {
		panic(err)
	}
	if corn := config.Cron; corn != "" {
		taskx.Corn().Add("DoRestic", corn, func() {
			DoRestic(config)
		})
		taskx.Corn().Start()
	} else {
		if out, err := DoRestic(config); err != nil {
			panic(err)
		} else {
			fmt.Println(out)
		}
	}
}

func DoRestic(config *Config) (out string, err error) {
	var cmd = config.Restic.Command()
	if retry := config.Retry; retry != nil {
		if retry.Times > 0 && retry.Interval > 0 {
			if err = execx.Retry(retry.Times, 0, time.Duration(retry.Interval)*time.Second, func() error {
				if out, err = doRestic(config, cmd); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return
			}
		}
	} else {
		if out, err = doRestic(config, cmd); err != nil {
			return
		}
	}
	return
}

func doRestic(config *Config, cmd string) (out string, err error) {
	if config.Restic.Do == common.Backup {
		var dumpCmd, dumpPath string
		if dumpCmd, dumpPath = config.Datasource.Dump(); dumpCmd != "" {
			dir, _ := filepath.Split(dumpPath)
			filex.CreateDir(dir)
			if out, err = execx.ExecCommandAndLog(dumpCmd, "根据数据库配置执行dump"); err != nil {
				return
			}
		}
		// 将dump文件移动至备份路径
		backupPath := config.Restic.Backup.Path
		if filex.IsDir(backupPath) {
			backupPath = backupPath + string(os.PathSeparator)
		}
		mvCmd := fmt.Sprintf("mv %s %s", dumpPath, backupPath)
		if out, err = execx.ExecCommandAndLog(mvCmd, "移动dump文件至备份路径"); err != nil {
			return
		}
		// 检查备份路径是否存在或者为空
		if checkPath := config.Restic.Backup.Path; !filex.Exists(checkPath) {
			out = fmt.Sprintf("%s：此备份文件或者文件夹不存在", checkPath)
			err = errors.New(out)
			return
		} else if filex.IsDir(checkPath) && filex.IsEmptyDir(checkPath) {
			out = fmt.Sprintf("%s：此备份文件夹为空", checkPath)
			err = errors.New(out)
			return
		}
	}

	// 恢复任务准备工作
	if config.Restic.Do == common.Restore {
		// 检查恢复路径是否存在
		if checkPath := config.Restic.Restore.Target; filex.Exists(checkPath) {
			// 检查恢复路径是否是文件夹
			if !filex.IsDir(checkPath) {
				out = fmt.Sprintf("%s：此恢复路径不是文件夹", checkPath)
				err = errors.New(out)
				return
			}
		} else {
			filex.CreateDir(checkPath)
		}
	}

	// 执行restic命令
	if out, err = execx.ExecCommandAndLog(cmd, "执行restic命令"); err != nil {
		return
	} else {
		ExtractSnapshotId(out)
	}

	// restic restore命令生成的恢复文件需要移动至所需恢复路径
	if config.Restic.Do == common.Restore {
		if backupPath := config.Restic.Restore.Path; backupPath != "" {
			restorePath := config.Restic.Restore.Target
			backupPath = filepath.Join(restorePath, filex.Pwd(backupPath))
			mvCmd := fmt.Sprintf("mv %s %s/", backupPath, restorePath)
			// 操作不影响任务执行结果，仅记录报错日志
			_, _ = execx.ExecCommandAndLog(mvCmd, "移动恢复文件")
		}
	}
	return
}

// ExtractSnapshotId 提取快照ID
func ExtractSnapshotId(text string) string {
	re := regexp.MustCompile(`snapshot\s+(\w+)\s+saved`)
	match := re.FindStringSubmatch(text)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
