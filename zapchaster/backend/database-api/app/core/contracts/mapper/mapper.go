package mapper

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
	"database-api/app/core/contracts/dtos/partPhotoDtos"
	"database-api/app/core/domain/entities"
)

func Map[T, U any](items []T, f func(T) U) []U {
	mapped := make([]U, len(items))
	for i, item := range items {
		mapped[i] = f(item)
	}
	return mapped
}

func MapPartDtoToCarPart(car carPartDtos.CarPartDto) entities.CarPart {
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
}

func MapToPartDetail(c carPartDtos.CarPartDto, ph []partPhotoDtos.PartPhotoDto) carPartDtos.PartDetailDto {
	return carPartDtos.PartDetailDto{
		WinCode:      c.WinCode,
		Price:        c.Price,
		Title:        c.Title,
		Description:  c.Description,
		Manufacturer: c.Manufacturer,
		CarBand:      c.CarBand,
		PartPhotos:   ph,
	}
}
