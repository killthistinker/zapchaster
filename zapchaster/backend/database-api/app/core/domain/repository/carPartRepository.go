package repository

import (
	"database-api/app/core/domain/entities"
)

type ICarPartRepository interface {
	GetAll() ([]entities.CarPart, error)
	Add(d *entities.CarPart) error
	AddParts(d *[]entities.CarPart) (int, error)
	GetCarPartFromId(partId int64) (entities.CarPart, error)
	AddCount(counter *entities.Counter) (int, error)
}
