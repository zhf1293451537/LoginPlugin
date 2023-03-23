package services

import (
	"WorkOrder/models"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type RegisInfo struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=15"`
	Password string `json:"password" form:"password" binding:"required"`
	Nickname string `json:"nickname" form:"nickname" binding:"required,min=1,max=10"`
	Profile  string `json:"profile" form:"profile" binding:"required,min=1,max=100"`
}

func (i *RegisInfo) CheckUsernameExists() bool {
	userinfo := &models.UserInfo{}
	result := models.DB.Where("username = ?", i.Username).First(&userinfo)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (i *RegisInfo) RegisUser(pwd string) error {
	userinfo := &models.UserInfo{
		Username: i.Username,
		Password: pwd,
		NickName: i.Nickname,
		Profile:  i.Profile,
	}
	//创建用户信息
	result := models.DB.Create(&userinfo)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
