package template

import (
	"bufio"
	"bytes"
	"embed"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-xuan/quanx/os/filex"
	"github.com/go-xuan/quanx/types/stringx"

	"gen_code_tool/common"
)

//go:embed tmpl/*
var fs embed.FS

var funcs = template.FuncMap{"uc": stringx.ToUpperCamel, "lc": stringx.ToLowerCamel}

type Tmpl struct {
	Frame    string           // 框架名
	Path     string           // 代码路径
	DataType string           // 数据类型
	FuncMap  template.FuncMap // 自定义模板参数方法
}

func (t *Tmpl) Template() *template.Template {
	var content, _ = fs.ReadFile(path.Join("tmpl/", t.Frame, t.Path))
	return template.Must(template.New(t.Path).Funcs(t.FuncMap).Parse(string(content)))
}

func (t *Tmpl) WriteCodeToFile(root string, data any, table ...string) (err error) {
	var dir, file = filepath.Split(t.Path)
	file = strings.TrimSuffix(file, ".tmpl")
	if len(table) > 0 {
		file = strings.Replace(file, "{{table}}", table[0], -1)
	}
	if file = filepath.Join(root, dir, file); needWrite(file) {
		var buf = &bytes.Buffer{}
		if err = t.Template().Execute(buf, data); err != nil {
			return
		}
		if err = filex.WriteFile(file, buf.String()); err != nil {
			return
		}
	}
	return
}

func GoQuanxTemplates() []*Tmpl {
	var frame = common.GoQuanx
	var tmpls []*Tmpl
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "main.go.tmpl", DataType: common.AppData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "go.mod.tmpl", DataType: common.AppData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "common/consts.go.tmpl", DataType: common.NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "conf/config.yaml.tmpl", DataType: common.AppData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "conf/database.yaml.tmpl", DataType: common.AppData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/router/router.go.tmpl", DataType: common.AppData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/controller/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/logic/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/dao/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/model/{{table}}.go.tmpl", DataType: common.TableData,
		FuncMap: template.FuncMap{"uc": stringx.ToUpperCamel, "lc": stringx.ToLowerCamel, "go_type": common.DB2GoType}})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/model/entity/{{table}}.go.tmpl", DataType: common.TableData,
		FuncMap: template.FuncMap{"uc": stringx.ToUpperCamel, "lc": stringx.ToLowerCamel, "go_type": common.DB2GoType, "gorm_type": common.DB2GormType}})
	return tmpls
}

func GoFrameTemplates() []*Tmpl {
	var frame = common.GoFrame
	var tmpls []*Tmpl
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "main.go.tmpl", DataType: common.AppData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "go.mod.tmpl", DataType: common.NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "consts/consts.go.tmpl", DataType: common.NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "manifest/config/global_conf/nacos.yaml.tmpl", DataType: common.NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/router/router.go.tmpl", DataType: common.AppData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/controller/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/logic/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/dao/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/model/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/model/do/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "internal/model/entity/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	return tmpls
}

func SpringBootTemplates() []*Tmpl {
	var frame = common.SpringBoot
	var tmpls []*Tmpl
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/main.go.tmpl", DataType: common.NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/common/consts.go.tmpl", DataType: common.NoData})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/controller/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/logic/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/dao/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/model/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	tmpls = append(tmpls, &Tmpl{Frame: frame, Path: "spring-boot/model/entity/{{table}}.go.tmpl", DataType: common.TableData, FuncMap: funcs})
	return tmpls
}

// 是否跳过文件写入
func needWrite(path string) bool {
	if filex.Exists(path) {
		var err error
		var file *os.File
		if file, err = os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
			return true
		}
		defer file.Close()
		var line []byte
		if line, _, err = bufio.NewReader(file).ReadLine(); err == nil {
			return string(line) == common.OverwriteTag
		}
	}
	return true
}
