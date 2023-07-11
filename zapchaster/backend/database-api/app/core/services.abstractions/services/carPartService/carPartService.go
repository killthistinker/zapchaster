package carPartService

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
)

type ICarPartService interface {
	AddPart(d *carPartDtos.CarPartDto) int
	GetAll() []carPartDtos.CarPartDto
	AddParts(c *[]carPartDtos.CarPartDto) int
	GetPart(partId string) (carPartDtos.CarPartDto, error)
	AddCount(dto carPartDtos.CounterDto) (int, error)
}
