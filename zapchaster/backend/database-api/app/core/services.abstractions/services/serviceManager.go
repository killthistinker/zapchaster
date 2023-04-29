package services

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
	"database-api/app/core/contracts/dtos/partPhotoDtos"
	"database-api/app/core/services.abstractions/services/carPartService"
	"database-api/app/core/services.abstractions/services/partPhoto"
)

type ServiceManager struct {
	CarPartService   carPartService.ICarPartService
	PartPhotoService partPhoto.IPartPhotoService
}

func (s *ServiceManager) AddPart(d *carPartDtos.CarPartDto) int {
	res := s.CarPartService.AddPart(d)
	return res
}

func (s *ServiceManager) GetAll() []carPartDtos.CarPartDto {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceManager) AddPhoto(photo *partPhotoDtos.PartPhotoDto) int {
	//TODO implement me
	panic("implement me")
}

func NewServiceManager(c carPartService.ICarPartService, p partPhoto.IPartPhotoService) IServiceManager {
	return &ServiceManager{
		CarPartService:   c,
		PartPhotoService: p,
	}
}

type IServiceManager interface {
	carPartService.ICarPartService
	partPhoto.IPartPhotoService
}

// Интерфейс IServiceManager должен быть реализован структурой ServiceManager.
// Методы из carPartService.ICarPartService и partPhoto.IPartPhotoService не нужно реализовывать явно,
// так как они уже определены в этих интерфейсах.
