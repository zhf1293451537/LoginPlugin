package models

import "github.com/jinzhu/gorm"

//用户文章点赞表
type UserLike struct {
	gorm.Model
	UserName  string
	ArticleID uint
}
