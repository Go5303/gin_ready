package main

import (
	"gin_ready/bootstrap"
	"gin_ready/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				return
			}
		}
	}()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		global.App.Log.Error("err:" + err.Error())
		return
	}
}
