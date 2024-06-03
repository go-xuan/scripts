package common

import (
	"bufio"
	"bytes"
	"embed"
	"github.com/go-xuan/quanx/os/filex"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-xuan/quanx/types/stringx"
)

//go:embed resource/template/*
var TemplateFs embed.FS

type Tmpls []*Tmpl

type Tmpl struct {
	Frame    string           // 框架名
	Path     string           // 模板名称
	DataType string           // 模板用表（0不用表/1生成文件需要一张表/2生成文件需要所有表）
	FuncMap  template.FuncMap // 自定义模板参数方法
}

func (t *Tmpl) Template() *template.Template {
	if t.FuncMap != nil {
		return template.Must(template.New(t.Path).Funcs(t.FuncMap).ParseFS(TemplateFs, t.Path))
	} else {
		return template.Must(template.New(t.Path).ParseFS(TemplateFs, t.Path))
	}
}

func (t *Tmpl) OutputPath(table ...string) string {
	path := strings.TrimSuffix(t.Path, ".tmpl")
	split := strings.Split(path, string(os.PathSeparator))
	if split[0] == t.Frame {
		split = split[1:]
	}
	if len(table) > 0 {
		i := len(split) - 1
		split[i] = strings.Replace(split[i], "{{table}}", table[0], -1)
	}
	return filepath.Join(split...)
}

func (t *Tmpl) WriteCodeToFile(root string, data any, table ...string) (err error) {
	path := t.OutputPath(table...)
	path = filepath.Join(root, path)
	if filex.Exists(path) && ReadFirstLine(path) != OverwriteTag {
		return
	}
	var buf = &bytes.Buffer{}
	if err = t.Template().Execute(buf, data); err != nil {
		return
	}
	if err = filex.WriteFile(path, buf.String()); err != nil {
		return
	}
	return
}

func GoQuanxTemplates() []*Tmpl {
	var frame = GoQuanx
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"uc": stringx.ToUpperCamel,
		"lc": stringx.ToLowerCamel,
	}
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "main.go.tmpl"), DataType: AdapterData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "go.mod.tmpl"), DataType: AdapterData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "common", "consts.go.tmpl"), DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "conf", "config.yaml.tmpl"), DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "router", "router.go.tmpl"), DataType: AdapterData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "controller", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "logic", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "dao", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "{{table}}.go.tmpl"), DataType: TableData,
		FuncMap: template.FuncMap{"uc": stringx.ToUpperCamel, "lc": stringx.ToLowerCamel, "go_type": DB2GoType}})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "entity", "{{table}}.go.tmpl"), DataType: TableData,
		FuncMap: template.FuncMap{"uc": stringx.ToUpperCamel, "lc": stringx.ToLowerCamel, "go_type": DB2GoType, "gorm_type": DB2GormType}})
	return tmpls
}

func GoFrameTemplates() []*Tmpl {
	var frame = GoFrame
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "main.go.tmpl"), DataType: AdapterData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "go.mod.tmpl"), DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "consts", "consts.go.tmpl"), DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "manifest", "config", "global_conf", "nacos.yaml.tmpl"), DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "router", "router.go.tmpl"), DataType: AdapterData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "controller", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "logic", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "dao", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "do", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: filepath.Join(frame, "internal", "model", "entity", "{{table}}.go.tmpl"), DataType: TableData, FuncMap: funcMap})
	return tmpls
}

func SpringBootTemplates() []*Tmpl {
	var frame = SpringBoot
	var tmpls []*Tmpl
	var funcMap = template.FuncMap{
		"UC":  stringx.ToUpperCamel,
		"LC":  stringx.ToLowerCamel,
		"DTG": DB2GoType,
	}
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/main.go.tmpl", DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/common/consts.go.tmpl", DataType: NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/controller/{{table}}.go.tmpl", DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/logic/{{table}}.go.tmpl", DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/dao/{{table}}.go.tmpl", DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/model/{{table}}.go.tmpl", DataType: TableData, FuncMap: funcMap})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/model/entity/{{table}}.go.tmpl", DataType: TableData, FuncMap: funcMap})
	return tmpls
}

// 按行读取
func ReadFirstLine(path string) (text string) {
	var err error
	var file *os.File
	if file, err = os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
		return
	}
	defer file.Close()
	// 按行处理txt
	reader := bufio.NewReader(file)
	var line []byte
	if line, _, err = reader.ReadLine(); err == nil {
		text = string(line)
	}
	return
}
