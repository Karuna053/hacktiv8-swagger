package database

import (
	"swagger/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=postgres.railway.internal user=postgres password=gUieLHoETvEyGCwWfjrovAGrXXRvsDVm dbname=carsapi port=railway sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Car{}, &models.Item{}, &models.Order{}, &models.User{})
	// DB.Debug().AutoMigrate(&models.Car{}, &models.Item{}, &models.Order{})
}

func GetDB() *gorm.DB {
	return DB
}
