package repository

import (
	"database-api/app/core/domain/entities"
)

type ICarPartRepository interface {
	GetAll() ([]entities.CarPart, error)
	Add(d *entities.CarPart) error
}
