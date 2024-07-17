package crontab

import (
	"fmt"
	"gin_ready/global"
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

const (
	Task1Space = "*/1 * * * * *" // every second
)

func InitCron() {
	c := cron.New()
	err1 := c.AddFunc(Task1Space, task1)
	if err1 != nil {
		global.App.Log.Error("crontab", zap.Any("err", err1))
	}
	// 启动cron
	c.Start()
}

// task1
func task1() {
	fmt.Println("task1")
}
