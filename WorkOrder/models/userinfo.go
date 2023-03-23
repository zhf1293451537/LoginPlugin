package models

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model
	// Grade    string
	Username string `gorm:"primary_key"`
	Password string
	NickName string
	Avatar   string
	Profile  string
	// Passwordverify string
}
