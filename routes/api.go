package routes

import (
	"gin_ready/app/controller/app"
	"gin_ready/app/middleware"
	"gin_ready/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JwtAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
	}
}
