package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-xuan/quanx/os/shellx"
	"github.com/go-xuan/quanx/types/timex"
	"github.com/olivere/elastic/v7"
)

const (
	Prefix  = "filebeat-7.10.2-"
	DateFmt = "2006.01.02"
	Backup  = ""
)

var client *elastic.Client

func main() {
	ctx := context.Background()
	var host = flag.String("host", "localhost", "输入ES服务主机IP，例如：-host=localhost")
	var port = flag.Int("port", 9200, "输入ES服务连接端口，例如：-port=9200")
	var prefix = flag.String("prefix", Prefix, "输入ES索引名前缀，例如：-prefix=xxx_xxx_")
	var maxDays = flag.Int("maxDays", 1, "输入索引最大保留月份数，例如：-maxDays=1")
	var backup = flag.String("backup", Backup, "输入ES备份数据文件，例如：-backup=data, 当-backup= 为空时不进行备份")
	var confirm = flag.Bool("confirm", false, "是否确认删除索引，例如：-confirm=true，如需删除索引，必须加上此参数")

	flag.Parse() // 解析输入的参数
	fmt.Println("ES服务主机IP:", *host)
	fmt.Println("ES服务连接端口:", *port)
	fmt.Println("ES索引名前缀:", *prefix)
	fmt.Println("索引最大保留天数:", *maxDays)
	fmt.Println("ES备份数据文件:", *backup)
	fmt.Println("确认删除索引:", *confirm)

	var err error
	url := fmt.Sprintf("http://%s:%d", *host, *port)
	if client, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	); err != nil {
		panic(err)
	}
	var result *elastic.PingResult
	var code int
	if result, code, err = client.Ping(url).Do(ctx); err != nil && code != 200 {
		panic(err)
	}
	fmt.Println("当前ES版本:", result.Version.Number)
	fmt.Println()
	// 获取所有索引
	var rows []elastic.CatIndicesResponseRow
	if rows, err = client.CatIndices().Do(ctx); err != nil {
		panic(err)
	}
	fmt.Println("========================================")
	fmt.Printf("%-28s|%v\n", "index_name", "if_delete")
	fmt.Println("----------------------------------------")
	// 筛选需要删除的索引
	pl, now := len(*prefix), time.Now()
	var deleteIndices []string
	for _, row := range rows {
		index := row.Index
		if strings.HasPrefix(index, *prefix) {
			var doDelete bool
			date := timex.TimeParse(index[pl:], DateFmt)
			if timex.DayInterval(date, now) >= *maxDays {
				deleteIndices = append(deleteIndices, index)
				doDelete = true
			}
			fmt.Printf("%-28s|%v\n", index, doDelete)
		}
	}
	fmt.Println("========================================")
	// 备份
	if *backup != "" {
		backupFile := *backup + "_" + now.Format("20060102")
		if !Exists(backupFile) {
			if _, err = shellx.ExecCommand(fmt.Sprintf("cp -r %s %s", *backup, backupFile)); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("备份已存在 ：", backupFile)
		}
	}
	var total = len(deleteIndices)
	fmt.Println("删除索引占比：", total, "/", len(rows))
	if *confirm && total > 0 {
		// 未避免http请求超过4096限制，每次最多删除50个索引
		var start, limit = 0, 50
		for start < total {
			if start+limit > total {
				limit = total - start
			}
			var end = start + limit
			var resp *elastic.IndicesDeleteResponse
			if resp, err = client.DeleteIndex(deleteIndices[start:end]...).Do(ctx); err != nil {
				panic(err)
			}
			fmt.Printf("删除索引[%d-%d]结果：%v\n", start, end, resp.Acknowledged)
			start = end
		}
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
