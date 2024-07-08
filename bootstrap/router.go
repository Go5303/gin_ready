package bootstrap

import (
	"context"
	"gin_ready/global"
	"gin_ready/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	//静态资源路由
	router.StaticFile("/", "./static/dist/index.html")
	//api相关路由
	routes.SetApiGroupRoutes(router)

	return router
}

func RunServer(port string) {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.App.Log.Error(err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.App.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.App.Log.Error("Shutdown err ...")
	}
	global.App.Log.Info("Server Exiting ...")
}
