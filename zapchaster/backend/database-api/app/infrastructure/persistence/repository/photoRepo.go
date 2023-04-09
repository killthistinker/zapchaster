package persistence

import (
	"database-api/app/core/domain/entities"
)

type PhotoRepo struct {
}

func (p PhotoRepo) Add(entities.PartPhoto) (int, error) {

}
