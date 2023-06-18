package carPartService

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
	"database-api/app/core/contracts/mapper"
	"database-api/app/core/domain/entities"
	"database-api/app/core/domain/repository"
	persistence "database-api/app/infrastructure/persistence/repository"
	"fmt"
	"strconv"
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
		Price:        d.Price,
		Description:  d.Description,
		MainPhoto:    d.MainPhoto,
		Manufacturer: d.Manufacturer,
		Title:        d.Title,
		WinCode:      d.WinCode,
	}
	err := c.ICarPartRepository.Add(&entity)
	if err != nil {
		return -2
	}
	return 0
}

func (c CarPartService) GetPart(partId string) (carPartDtos.CarPartDto, error) {
	id, err := strconv.ParseInt(partId, 10, 64)
	if err != nil {
		return carPartDtos.CarPartDto{}, err
	}
	resEntity, err := c.GetCarPartFromId(id)
	if err != nil {
		return carPartDtos.CarPartDto{}, err
	}

	res := carPartDtos.CarPartDto{
		Id:           resEntity.Id,
		Price:        resEntity.Price,
		Manufacturer: resEntity.Manufacturer,
		Title:        resEntity.Title,
		WinCode:      resEntity.WinCode,
		MainPhoto:    resEntity.MainPhoto,
		CarBand:      resEntity.CarBand,
		Description:  resEntity.Description,
	}
	return res, nil
}

func (c CarPartService) GetAll() []carPartDtos.CarPartDto {
	parts, err := c.ICarPartRepository.GetAll()
	if err != nil {
		return nil
	}
	partDto := mapper.Map(parts, func(p entities.CarPart) carPartDtos.CarPartDto {
		return carPartDtos.CarPartDto{
			Id:           p.Id,
			Price:        p.Price,
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

func (c CarPartService) AddParts(d *[]carPartDtos.CarPartDto) int {
	parts := mapper.Map(*d, func(car carPartDtos.CarPartDto) entities.CarPart {
		return entities.CarPart{
			Id:           car.Id,
			WinCode:      car.WinCode,
			Price:        car.Price,
			Title:        car.Title,
			Description:  car.Description,
			Manufacturer: car.Manufacturer,
			CarBand:      car.CarBand,
			MainPhoto:    car.MainPhoto,
		}
	})

	res, err := c.ICarPartRepository.AddParts(&parts)
	if err != nil {
		fmt.Println(err)
		return res
	}

	return res
}
