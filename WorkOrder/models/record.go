package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type History struct {
	gorm.Model
	UserID     string
	ArticleID  string
	Title      string
	RecordTime time.Time
}

//创建浏览记录
func (h *History) Create() error {
	err := DB.Create(h).Error
	if err != nil {
		return err
	}
	return nil
}

//根据用户id获取用户个人浏览记录
func GetRecordList(userID string) ([]*History, error) {
	history := []*History{}
	err := DB.Table("histories").Where("user_id = ?", userID).Select("id,title,article_id,record_time").Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}
