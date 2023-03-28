package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Catagory struct {
	gorm.Model
	Name string
}

func GetCataId(name string) (uint, error) {
	cata := &Catagory{}
	err := DB.Table("catagories").Where("name = ?", name).Find(cata).Error
	if err == gorm.ErrRecordNotFound {
		return 0, errors.New("没有这个分类,请先添加分类")
	} else if err != nil {
		return 0, err
	}
	return cata.ID, nil
}

func GetCataList() ([]*Catagory, error) {
	result := []*Catagory{}
	err := DB.Table("catagories").Select("id,name").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Catagory) Create() error {
	err := DB.Create(c).Error
	if err != nil {
		return err
	}
	return nil
}
func GetCataByID(id string) (string, error) {
	cata := &Catagory{}
	err := DB.Table("catagories").Where("id = ?", id).Select("name").Find(cata).Error
	if err != nil {
		return "", err
	}
	return cata.Name, nil
}
