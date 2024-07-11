package main

import (
	"fmt"
	"github.com/go-xuan/quanx/utils/marshalx"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func main() {
	var err error
	var config = &Config{}
	if err = marshalx.UnmarshalFromFile("restic_backup_tool.yaml", config); err != nil {
		panic(err)
	}
	var restic = config.Restic
	if restic.cron != "" {
		var myCron = cron.New(cron.WithSeconds(),
			cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)),
			cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))),
		)
		var entryId cron.EntryID
		if entryId, err = myCron.AddFunc(restic.cron, func() {
			if err = Backup(config); err != nil {
				panic(err)
			}
		}); err != nil {
			panic(err)
		}
		fmt.Println("定时备份启动成功!", entryId)
	} else {
		// 仅备份一次
		if err = Backup(config); err != nil {
			panic(err)
		}
	}
}
