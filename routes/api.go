package routes

import (
	"gin_ready/app/controller/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)
}
