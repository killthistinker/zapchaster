package entities

import "github.com/jinzhu/gorm"

type PartPhoto struct {
	gorm.Model
	Id        int64  `json:"id"`
	PartId    int64  `json:"partId"`
	PartPhoto string `json:"partPhoto"`
}
