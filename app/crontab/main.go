package crontab

import (
	"gin_ready/app/services"
	"gin_ready/global"
	"github.com/Go5303/uuid"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"time"
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
	randStr := uuid.Uuid()
	redisService := services.RedisService{RedisClient: global.App.Redis}
	err := redisService.HSetRedisString("uuid", randStr, "uuid:"+randStr, 3600*10*time.Second)
	if err != nil {
		global.App.Log.Error("task1", zap.Any("err", err))
	}
}
