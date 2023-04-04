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
	Cataid      uint
	Likes       int
	Views       int
	PublishDate time.Time
}

func (a *Article) Create() error {
	err := DB.Create(a).Error
	if err != nil {
		return err
	}
	return nil
}

func GetArticleByID(id string) (*Article, error) {
	// id = uint(id)
	art := &Article{}
	result := DB.Where("id = ?", id).Find(&art)
	if result.Error != nil {
		return nil, result.Error
	}
	err := DB.Table("articles").Where("id = ?", id).UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
	if err != nil {
		return nil, err
	}
	return art, nil
}

func (a *Article) UpdateById(id string) error {
	err := DB.Table("articles").Where("id = ?", id).Update(a).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id string) error {
	err := DB.Table("articles").Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetArticleList() ([]*Article, error) {
	result := []*Article{}
	err := DB.Table("articles").Select("id,title,content,author").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetArticleListByLike() ([]*Article, error) {
	result := []*Article{}
	err := DB.Table("articles").Select("id,title,content,author,likes").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func GetArticleListByCata(cataid string) ([]*Article, error) {
	result := []*Article{}
	err := DB.Table("articles").Where("cataid = ?", cataid).Select("id,title,content,author").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetArticleByArchive() ([]*Article, error) {
	result := []*Article{}
	err := DB.Table("articles").Select("id,title,publish_date").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

//获取文章阅读量表
func GetArticleByViews() ([]*Article, error) {
	result := []*Article{}
	err := DB.Table("articles").Select("id,title,author,views").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetArtTitleById(articleID string) (string, error) {
	art := &Article{}
	err := DB.Table("articles").Where("id = ?", articleID).Select("title").Find(&art).Error
	if err != nil {
		return "", err
	}
	return art.Title, nil
}

func GetTitleCataidById(articleID string) (string, uint, error) {
	art := &Article{}
	err := DB.Table("articles").Where("id = ?", articleID).Select("title,cataid").Find(&art).Error
	if err != nil {
		return "", 0, err
	}
	return art.Title, art.Cataid, nil
}
