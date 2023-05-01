package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Database *gorm.DB
}

func NewConnection(url string) *DbConnection {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	}
	return &DbConnection{db}
}
