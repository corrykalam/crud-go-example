package db

import (
	"fmt"
	"sesi12-final/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "final-project-hacktiv8"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success connect to DB using GORM")
	db.Debug().AutoMigrate(&models.User{}, &models.Social{}, &models.Photo{}, &models.Comment{})
	fmt.Println("Success migrate table")
}

func GetDB() *gorm.DB {
	return db
}
