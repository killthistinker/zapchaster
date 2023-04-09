package persistence

import (
	"database-api/app/core/domain/entities"
)

type CarPartRepo struct {
}

func (c CarPartRepo) Add(d *entities.CarPart) (int, error) {

}

func (c CarPartRepo) GetAll() ([]entities.CarPart, error) {

}
