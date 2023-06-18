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

func (c CarPartRepo) AddParts(d *[]entities.CarPart) (int, error) {
	res := c.db.Create(d)
	if res.Error != nil {
		return -1, res.Error
	}
	return 0, nil
}

func (c CarPartRepo) GetCarPartFromId(partId int64) (entities.CarPart, error) {
	var carPart entities.CarPart
	res := c.db.First(&carPart, partId)
	if res.Error != nil {
		return carPart, res.Error
	}
	return carPart, nil
}
