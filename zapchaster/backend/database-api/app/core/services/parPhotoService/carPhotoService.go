package parPhotoService

import (
	"database-api/app/core/contracts/dtos/partPhotoDtos"
	"database-api/app/core/domain/entities"
	"database-api/app/core/domain/repository"
	persistence "database-api/app/infrastructure/persistence/repository"
)

type parPhotoService struct {
	repository.ICarPhotoRepository
}

func New(repo persistence.PhotoRepo) *parPhotoService {
	r := parPhotoService{repo}
	return &r
}

func (p parPhotoService) AddPhoto(photo *partPhotoDtos.PartPhotoDto) int {
	entity := entities.PartPhoto{PartId: photo.PartId, PartPhoto: photo.PartPhoto}

	r, err := p.ICarPhotoRepository.Add(entity)

	if err != nil {
		return -1
	}

	return r
}
