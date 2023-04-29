package server

import (
	"database-api/app/core/services.abstractions/services"
	carPartService2 "database-api/app/core/services/carPartService"
	"database-api/app/core/services/parPhotoService"
	persistence "database-api/app/infrastructure/persistence/repository"
	"database-api/app/infrastructure/presentation/db"
)

func Configure() services.IServiceManager {
	db := db.InitDb()

	carPartRepo := persistence.NewCarPartRepo(db)
	photoRepo := persistence.NewPhotoRepo(db)

	carPartService := carPartService2.NewCarPartService(*carPartRepo)
	photoService := parPhotoService.New(*photoRepo)

	serviceManager := services.NewServiceManager(carPartService, photoService)
	return serviceManager
}
