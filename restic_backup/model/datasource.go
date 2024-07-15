package model

import (
	"fmt"
	"github.com/go-xuan/quanx/server/gormx"
	"path/filepath"
	"strconv"
	"strings"
)

type Datasource struct {
	Type     string `json:"type"     yaml:"type"     description:"数据库类型（mysql/postgre/mongo）"`
	Host     string `json:"host"     yaml:"host"     description:"主机host"`
	Port     int    `json:"port"     yaml:"port"     description:"端口"`
	Username string `json:"username" yaml:"username" description:"用户名"`
	Password string `json:"password" yaml:"password" description:"密码"`
	Name     string `json:"name"     yaml:"name"     description:"数据库名"`
	Schema   string `json:"schema"   yaml:"schema"   description:"模式名"`
	Table    string `json:"table"    yaml:"table"    description:"表名"`
}

func (d *Datasource) Dump(dir ...string) (dump string, path string) {
	path = filepath.Join("", d.Type)
	if len(dir) > 0 {
		path = filepath.Join(path, dir[0])
	}
	switch d.Type {
	case gormx.Mysql:
		dump, path = d.mysqlDump(path)
	case gormx.Postgres:
		dump, path = d.pgsqlDump(path)
	case "mongo":
		dump, path = d.mongoDump(path)
	}
	return
}

func (d *Datasource) mysqlDump(path ...string) (string, string) {
	var filePath = d.Name
	var dump = strings.Builder{}
	dump.WriteString(`mysqldump`)
	dump.WriteString(` -h `)
	dump.WriteString(d.Host)
	dump.WriteString(` -P `)
	dump.WriteString(strconv.Itoa(d.Port))
	dump.WriteString(` -u `)
	dump.WriteString(d.Username)
	dump.WriteString(` --password=`)
	dump.WriteString(d.Password)
	dump.WriteString(` `)
	dump.WriteString(d.Name)
	if d.Table != "" {
		dump.WriteString(` `)
		dump.WriteString(d.Table)
		filePath = filepath.Join(filePath, d.Table)
	}
	if len(path) > 0 {
		filePath = filepath.Join(path[0], filePath)
	}
	filePath = filePath + ".sql"
	dump.WriteString(` > `)
	dump.WriteString(filePath)
	return dump.String(), filePath
}

func (d *Datasource) pgsqlDump(path ...string) (string, string) {
	var filePath = d.Name
	var dump = strings.Builder{}
	dump.WriteString(`pg_dump`)
	dump.WriteString(fmt.Sprintf(` "host=%s port=%d user=%s password=%s dbname=%s" `, d.Host, d.Port, d.Username, d.Password, d.Name))
	if d.Schema != "" {
		dump.WriteString(` -n `)
		dump.WriteString(d.Schema)
		filePath = filepath.Join(filePath, d.Schema)
		if d.Table != "" {
			dump.WriteString(` -t `)
			dump.WriteString(d.Table)
			filePath = filepath.Join(filePath, d.Table)
		}
	}
	if len(path) > 0 {
		filePath = filepath.Join(path[0], filePath)
	}
	filePath = filePath + ".sql"
	dump.WriteString(` -f `)
	dump.WriteString(filePath)
	return dump.String(), filePath
}

func (d *Datasource) mongoDump(path ...string) (string, string) {
	var filePath = d.Name
	var dump = strings.Builder{}
	dump.WriteString(`mongodump`)
	dump.WriteString(` -h `)
	dump.WriteString(d.Host)
	dump.WriteString(` --port `)
	dump.WriteString(strconv.Itoa(d.Port))
	dump.WriteString(` -u `)
	dump.WriteString(d.Username)
	dump.WriteString(` -p `)
	dump.WriteString(d.Password)
	dump.WriteString(` -d `)
	dump.WriteString(d.Name)
	if d.Table != "" {
		dump.WriteString(` -c `)
		dump.WriteString(d.Table)
		filePath = filepath.Join(filePath, d.Table)
	}
	if len(path) > 0 {
		filePath = filepath.Join(path[0], filePath)
	}
	dump.WriteString(` -o `)
	dump.WriteString(filePath)
	dump.WriteString(` --gzip`)
	return dump.String(), filePath
}
