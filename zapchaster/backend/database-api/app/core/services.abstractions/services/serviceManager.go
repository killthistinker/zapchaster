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
	res := s.CarPartService.GetAll()
	return res
}

func (s *ServiceManager) AddPhoto(photos []partPhotoDtos.PartPhotoDto) int {
	res := s.PartPhotoService.AddPhoto(photos)
	return res
}

func (s *ServiceManager) AddParts(e *[]carPartDtos.CarPartDto) int {
	res := s.CarPartService.AddParts(e)
	return res
}
func (s *ServiceManager) GetPart(partId string) (carPartDtos.CarPartDto, error) {
	res, err := s.CarPartService.GetPart(partId)
	if err != nil {
		return carPartDtos.CarPartDto{}, err
	}
	return res, nil
}

func (s *ServiceManager) GetPartPhotosFromId(partId string) ([]partPhotoDtos.PartPhotoDto, error) {
	res, err := s.PartPhotoService.GetPartPhotosFromId(partId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ServiceManager) GetPartDetails(partId string) (*carPartDtos.PartDetailDto, error) {
	var partDetail *carPartDtos.PartDetailDto
	partRes, err := s.GetPart(partId)
	photoRes, err := s.GetPartPhotosFromId(partId)
	if err != nil {
		return nil, err
	}
	partDetail.WinCode = partRes.WinCode
	partDetail.PartPhotos = photoRes
	return partDetail, nil
}

func (s *ServiceManager) AddCount(dto carPartDtos.CounterDto) (int, error) {
	res, err := s.CarPartService.AddCount(dto)
	return res, err
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
