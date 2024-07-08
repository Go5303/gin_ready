package app

import (
	"gin_ready/app/common/request"
	"gin_ready/app/common/response"
	"gin_ready/app/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
