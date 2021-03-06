package models

import (
	"github.com/ayannahindonesia/basemodel"
	"github.com/lib/pq"
)

type (
	Bank struct {
		basemodel.BaseModel
		Name     string        `json:"name" gorm:"column:name;type:varchar(255)"`
		Image    string        `json:"image" gorm:"column:image;type:text"`
		Type     int           `json:"type" gorm:"column:type;type:varchar(255)"`
		Address  string        `json:"address" gorm:"column:address;type:text"`
		Province string        `json:"province" gorm:"column:province;type:varchar(255)"`
		City     string        `json:"city" gorm:"column:city;type:varchar(255)"`
		PIC      string        `json:"pic" gorm:"column:pic;type:varchar(255)"`
		Phone    string        `json:"phone" gorm:"column:phone;type:varchar(255)"`
		Services pq.Int64Array `json:"services" gorm "column:services"`
		Products pq.Int64Array `json:"products" gorm "column:products"`
	}
)

func (model *Bank) Create() error {
	err := basemodel.Create(&model)
	return err
}

func (model *Bank) FirstOrCreate() (err error) {
	return basemodel.FirstOrCreate(&model)
}

func (model *Bank) Save() error {
	err := basemodel.Save(&model)
	return err
}

func (model *Bank) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

func (model *Bank) FindbyID(id uint64) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

func (model *Bank) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	banks := []Bank{}
	order := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&banks, page, rows, order, sorts, filter)

	return result, err
}

func (model *Bank) FilterSearchSingle(filter interface{}) (err error) {
	err = basemodel.SingleFindFilter(&model, filter)
	return err
}
