package main

import (
	"gin_ready/bootstrap"
	"gin_ready/config"
	"gin_ready/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

var (
	Log  *zap.Logger
	Yaml config.Configuration
	DB   *gorm.DB
)

func main() {
	//初始化yaml配置
	bootstrap.InitializeConfig()
	Yaml = global.App.Config

	//初始化log
	global.App.Log = bootstrap.InitializeLog()
	Log = global.App.Log
	Log.Info("log init success!")

	//数据库初始化
	DB = bootstrap.InitializeDB()
	Log.Info("db init success")

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
