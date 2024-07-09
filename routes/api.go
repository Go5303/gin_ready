package routes

import (
	"gin_ready/app/controller/app"
	"gin_ready/app/middleware"
	"gin_ready/app/services"
	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.Engine) {
	routerGroup := router.Group("/api")
	{
		//不需要token签名
		routerGroup.POST("/auth/register", app.Register)
		routerGroup.POST("/auth/login", app.Login)

		// 需要token签名
		authRouter := routerGroup.Group("").Use(middleware.JwtAuth(services.AppGuardName))
		{
			authRouter.POST("/auth/info", app.Info)
			authRouter.POST("/auth/logout", app.Logout)
		}
	}

}
