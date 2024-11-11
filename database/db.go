package database

import (
	"fmt"
	"os"
	"swagger/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%s 
	user=%s `+`
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Car{}, &models.Item{}, &models.Order{}, &models.User{})
	// DB.Debug().AutoMigrate(&models.Car{}, &models.Item{}, &models.Order{})
}

func GetDB() *gorm.DB {
	return DB
}
