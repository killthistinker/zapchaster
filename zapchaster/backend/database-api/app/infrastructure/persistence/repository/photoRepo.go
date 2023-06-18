package persistence

import (
	"database-api/app/core/domain/entities"
	"gorm.io/gorm"
)

type PhotoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) *PhotoRepo {
	return &PhotoRepo{db}
}

func (p PhotoRepo) Add(ph []entities.PartPhoto) error {
	return p.db.Create(ph).Error
}

func (p PhotoRepo) GetPhotoSById(partId int64) ([]entities.PartPhoto, error) {
	var parPhotos []entities.PartPhoto
	res := p.db.Where("partid in (?)", partId).Find(&parPhotos)
	if res.Error != nil {
		return parPhotos, res.Error
	}
	return parPhotos, nil
}
