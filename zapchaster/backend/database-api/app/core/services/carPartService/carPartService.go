package carPartService

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
	"database-api/app/core/contracts/mapper"
	"database-api/app/core/domain/entities"
	"database-api/app/core/domain/repository"
	persistence "database-api/app/infrastructure/persistence/repository"
)

type CarPartService struct {
	repository.ICarPartRepository
}

func NewCarPartService(repo persistence.CarPartRepo) *CarPartService {
	p := CarPartService{repo}
	return &p
}

func (c CarPartService) AddPart(d *carPartDtos.CarPartDto) int {
	entity := entities.CarPart{
		CarBand:      d.CarBand,
		Description:  d.Description,
		MainPhoto:    d.MainPhoto,
		Manufacturer: d.Manufacturer,
		Title:        d.Title,
		Type:         d.Type,
		WinCode:      d.WinCode,
	}
	r, err := c.ICarPartRepository.Add(&entity)
	if err != nil {
		return -2
	}
	return r
}

func (c CarPartService) GetAll() []carPartDtos.CarPartDto {
	parts, err := c.ICarPartRepository.GetAll()
	if err != nil {
		return nil
	}
	partDto := mapper.Map(parts, func(p entities.CarPart) carPartDtos.CarPartDto {
		return carPartDtos.CarPartDto{
			Id:           p.Id,
			Type:         p.Type,
			WinCode:      p.WinCode,
			Title:        p.Title,
			MainPhoto:    p.MainPhoto,
			Manufacturer: p.Manufacturer,
			Description:  p.Description,
			CarBand:      p.CarBand,
		}
	})
	return partDto
}
