package routes

import (
	"gin_ready/app/common/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
