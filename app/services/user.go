package services

import (
	"errors"
	"gin_ready/app/common/request"
	"gin_ready/app/models"
	"gin_ready/global"
	"gin_ready/utils"
	"strconv"
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

// Login 登陆
func (userService *userService) Login(params request.Login) (err error, user models.User) {
	err = global.App.DB.Where("mobile = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
