package models

import (
	"gin_ready/global"
	"gorm.io/gorm"
)

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}

func GetUserInfo(mobile string) *User {
	var user User
	err := global.App.DB.Table("users").
		Where("mobile = ?", mobile).
		First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		global.App.Log.Error(err.Error())
	}
	return &user
}
