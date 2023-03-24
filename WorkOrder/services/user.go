package services

import (
	"WorkOrder/models"
	"WorkOrder/serializer"
	"WorkOrder/utils"
	"log"
)

type UserLoginInfo struct {
	Username string `json:"username" binding:"required,min=1,max=10"`
	Password string `json:"password" binding:"required"`
}

func (u *UserLoginInfo) Userverification() (*serializer.Response, bool) {
	// name := "zhf"
	// hashPass, _ := utils.EncryptPassword("123456") //模拟密码为前段md5加密后
	// _ = models.RedisClient.Set(name, hashPass, 0).Err()
	user := &models.UserInfo{Username: u.Username}
	err := models.DB.Where("username = ?", user.Username).First(&user).Error
	// log.Println(user)
	if err != nil {
		log.Println("账户不存在")
		return &serializer.Response{
			Stauts: 20002,
			Msg:    "fail",
			Error:  "账户名或密码错误",
		}, false
	}
	IsTrue := utils.EqualsPassword(u.Password, user.Password)
	if IsTrue {
		return &serializer.Response{
			Stauts: 20001,
			Msg:    "true",
		}, true
	}
	return &serializer.Response{
		Stauts: 20002,
		Msg:    "fail",
		Error:  "账户名或密码错误",
	}, false
}
