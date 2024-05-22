package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-xuan/quanx/os/shellx"
)

type Config struct {
	Restic     *Restic     `yaml:"restic"`     // restic配置信息
	DataSource []*Database `yaml:"dataSource"` // 备份数据源
}

type Restic struct {
	OS           string `yaml:"os"`           // 操作系统
	Repository   string `yaml:"repository"`   // 备份存储库
	BackupPath   string `yaml:"backupPath"`   // 备份目录
	PasswordFile string `yaml:"passwordFile"` // 存储库密码文件
	Corn         string `yaml:"corn"`         // 定时执行cron表达式
}

func (r *Restic) BackupCmd(tag string, options ...string) string {
	var cmd = strings.Builder{}
	cmd.WriteString(`restic `)
	// 设置存储库
	if r.Repository != "" {
		cmd.WriteString(` -r `)
		cmd.WriteString(r.Repository)
	}
	cmd.WriteString(` -v backup `)
	if len(options) > 0 {
		for _, option := range options {
			cmd.WriteString(option)
		}
	} else {
		cmd.WriteString(r.BackupPath)
	}
	// 快照标签参数
	if tag != "" {
		cmd.WriteString(` --tag `)
		cmd.WriteString(tag)
	}
	if r.PasswordFile != "" {
		// 密码文件参数
		cmd.WriteString(` --password-file `)
		cmd.WriteString(r.PasswordFile)
	}
	return cmd.String()
}

type Database struct {
	Type     string `yaml:"type"`     // 数据库类型
	Host     string `yaml:"host"`     // 主机
	Port     int    `yaml:"port"`     // 端口
	Username string `yaml:"username"` // 用户名
	Password string `yaml:"password"` // 密码
	Database string `yaml:"database"` // 数据库
}

func (d *Database) Dump(outPath string) error {
	var dump string
	var out = filepath.Join(outPath, d.Type)
	switch d.Type {
	case "mongo":
		// mongodump -h 127.0.0.1:27017 -d cgywtb_ticketManage -u cgywtb_ticketManage -p n5qxo0c7tHSI6DtM -o backup/mongo --gzip
		dump = fmt.Sprintf(
			`mongodump -h %s:%d -u %s -p %s -d %s -o %s --gzip`,
			d.Host, d.Port, d.Username, d.Password, d.Database, out,
		)
	case "mysql":
		// mongodump -h 192.168.2.250 -P 3306 -u root -p redbird@123 -d ticket > xxx.sql
		dump = fmt.Sprintf(
			`mysqldump -h %s -P %d -u %s -p %s -d %s > %s.sql`,
			d.Host, d.Port, d.Username, d.Password, d.Database, out,
		)
	}
	if dump != "" {
		if _, err := shellx.ExecCommand(dump); err != nil {
			return err
		}
	}
	fmt.Println("================================")
	fmt.Println(d.Type, "Dump Success ")
	return nil
}
