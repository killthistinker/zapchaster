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

func (p PhotoRepo) Add(ph entities.PartPhoto) error {
	return p.db.Create(ph).Error
}
