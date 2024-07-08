package services

import (
	"errors"
	"gin_ready/app/common/request"
	"gin_ready/app/models"
	"gin_ready/global"
	"gin_ready/utils"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	userInfo := models.GetUserInfo(params.Mobile)

	if userInfo.ID.ID > 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}
