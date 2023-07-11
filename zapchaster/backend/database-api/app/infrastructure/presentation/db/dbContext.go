package db

import (
	"database-api/app/core/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "host=postgres-container user=postgres password=postgre dbname=zapchaster port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.PartPhoto{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.Counter{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.CarPart{})

	if err != nil {
		panic(err)
	}

	return db
}
