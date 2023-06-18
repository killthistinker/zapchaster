package repository

import (
	"database-api/app/core/domain/entities"
)

type ICarPhotoRepository interface {
	Add(d []entities.PartPhoto) error
	GetPhotoSById(partId int64) ([]entities.PartPhoto, error)
}
