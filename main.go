package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// ... 添加你的路由配置 ...

	// 创建一个errgroup来管理goroutine
	var eg errgroup.Group

	// 启动HTTP服务器
	eg.Go(func() error {
		return g.Run(":8080") // 你的监听端口
	})

	// 监听退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// 当收到信号时，等待所有请求处理完成
	<-quit
	log.Println("Shutting down server...")

	// 设置超时，确保在一定时间内关闭
	timeout := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 优雅关闭HTTP服务器
	if err := g.Shutdown(ctx); err != nil {
		log.Printf("Could not gracefully shutdown the server: %v\n", err)
	}

	// 等待所有goroutine完成
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
