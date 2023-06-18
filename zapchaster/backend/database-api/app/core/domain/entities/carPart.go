package entities

import "github.com/jinzhu/gorm"

type CarPart struct {
	gorm.Model
	Id           int64  `gorm:"column:id"`
	WinCode      string `gorm:"column:winCode"`
	Price        int64  `gorm:"column:price"`
	Title        string `gorm:"column:title"`
	Description  string `gorm:"column:description"`
	Manufacturer string `gorm:"column:manufacturer"`
	CarBand      string `gorm:"column:carBrand"`
	MainPhoto    string `gorm:"column:mainPhoto"`
}
