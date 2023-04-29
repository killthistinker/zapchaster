package persistence

import (
	"database-api/app/core/domain/entities"
	"gorm.io/gorm"
)

type CarPartRepo struct {
	db *gorm.DB
}

func NewCarPartRepo(db *gorm.DB) *CarPartRepo {
	return &CarPartRepo{db}
}

func (c CarPartRepo) Add(d *entities.CarPart) error {
	return c.db.Create(d).Error
}

func (c CarPartRepo) GetAll() ([]entities.CarPart, error) {
	var carParts []entities.CarPart

	err := c.db.Find(&carParts).Error
	if err != nil {
		return nil, err
	}
	return carParts, nil
}
