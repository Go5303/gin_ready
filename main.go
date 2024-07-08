package main

import (
	"gin_ready/bootstrap"
	"gin_ready/config"
	"gin_ready/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var (
	Log  *zap.Logger
	Yaml config.Configuration
)

func main() {
	//初始化环境配置
	bootstrap.InitializeConfig()
	Yaml = global.App.Config

	//初始化日志
	global.App.Log = bootstrap.InitializeLog()
	Log = global.App.Log
	Log.Info("log init success!")

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	err := r.Run(":" + Yaml.App.Port)
	if err != nil {
		return
	}
}
