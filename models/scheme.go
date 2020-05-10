package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Version  int64  `json:"version"`
}

func (p *Product) TableName() string {
	return "product"
}
