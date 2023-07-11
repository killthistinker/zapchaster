package persistence

import (
	"database-api/app/core/domain/entities"
	"errors"
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

func (c CarPartRepo) AddCount(co *entities.Counter) (int, error) {
	var counter entities.Counter

	res := c.db.Where("wincode = ?", co.WinCode).First(&counter)
	if res.Error != nil {
		cr := c.db.Create(co)
		if cr.Error != nil {
			return -1, errors.New("ошибка при создании объекта")
		}
		co.Count++
		if err := c.db.Model(&entities.Counter{}).Where("wincode = ?", co.WinCode).Update("count", co.Count).Error; err != nil {
			return -2, errors.New("ошибка при обновлении данных")
		}
		return 0, nil
	}
	counter.Count++
	if err := c.db.Model(&entities.Counter{}).Where("wincode = ?", counter.WinCode).Update("count", counter.Count).Error; err != nil {
		return -3, errors.New("ошибка при обновлении данных")
	}
	return 0, nil
}
