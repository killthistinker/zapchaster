package services

import (
	"database-api/app/core/services.abstractions/services/carPartService"
	"database-api/app/core/services.abstractions/services/partPhoto"
)

type IServiceManager interface {
	carPartService.ICarPartService
	partPhoto.IPartPhotoService
}
