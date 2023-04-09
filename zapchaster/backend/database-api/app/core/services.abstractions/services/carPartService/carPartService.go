package carPartService

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
)

type ICarPartService interface {
	AddPart(d *carPartDtos.CarPartDto) int
	GetAll() []carPartDtos.CarPartDto
}
