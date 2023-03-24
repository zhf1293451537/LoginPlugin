package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title       string
	Content     string
	Author      string
	PublishDate time.Time
}

func (a *Article) Create() error {
	return nil
}
