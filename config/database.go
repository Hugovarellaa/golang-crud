package config

import (
	"github.com/Hugovarellaa/golang-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=localhost user=postgres	password=postgres dbname=postgres port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Customers{})
	DB = db

}
