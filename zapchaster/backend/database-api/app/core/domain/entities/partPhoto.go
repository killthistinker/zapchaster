package entities

import "github.com/jinzhu/gorm"

type PartPhoto struct {
	gorm.Model
	Id        int64  `gorm:"column:id"`
	PartId    int64  `gorm:"column:partid"`
	PartPhoto string `gorm:"column:PartPhoto"`
}
