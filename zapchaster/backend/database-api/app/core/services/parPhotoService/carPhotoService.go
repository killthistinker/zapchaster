package parPhotoService

import (
	"database-api/app/core/contracts/dtos/partPhotoDtos"
	"database-api/app/core/contracts/mapper"
	"database-api/app/core/domain/entities"
	"database-api/app/core/domain/repository"
	persistence "database-api/app/infrastructure/persistence/repository"
	"strconv"
)

type ParPhotoService struct {
	repository.ICarPhotoRepository
}

func New(repo persistence.PhotoRepo) *ParPhotoService {
	r := ParPhotoService{repo}
	return &r
}

func (p ParPhotoService) AddPhoto(photo []partPhotoDtos.PartPhotoDto) int {
	entity := mapper.Map(photo, func(c partPhotoDtos.PartPhotoDto) entities.PartPhoto {
		return entities.PartPhoto{
			PartId:    c.PartId,
			PartPhoto: c.PartPhoto,
		}
	})

	err := p.ICarPhotoRepository.Add(entity)

	if err != nil {
		return -1
	}

	return 0
}

func (p ParPhotoService) GetPartPhotosFromId(partId string) ([]partPhotoDtos.PartPhotoDto, error) {
	id, err := strconv.ParseInt(partId, 10, 64)
	if err != nil {
		return nil, err
	}

	partPhotos, err := p.ICarPhotoRepository.GetPhotoSById(id)
	if err != nil {
		return nil, err
	}
	res := mapper.Map(partPhotos, func(ep entities.PartPhoto) partPhotoDtos.PartPhotoDto {
		return partPhotoDtos.PartPhotoDto{
			PartId:    ep.PartId,
			PartPhoto: ep.PartPhoto,
		}
	})
	return res, nil
}
