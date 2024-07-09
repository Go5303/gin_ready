package main

import (
	"gin_ready/bootstrap"
	"gin_ready/global"
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

	//初始化redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 路由注册
	bootstrap.RunServer(global.App.Config.App.Port)
}
