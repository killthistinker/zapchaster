package repository

import (
	"database-api/app/core/domain/entities"
)

type ICarPhotoRepository interface {
	Add(d entities.PartPhoto) (int, error)
}
