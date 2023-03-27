package models

import "github.com/jinzhu/gorm"

type Catagory struct {
	gorm.Model
	Name string
}

func GetCataId(name string) (uint, error) {
	cata := &Catagory{}
	err := DB.Table("catagories").Where("name = ?", name).Find(cata).Error
	if err != nil {
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
	return nil
}
