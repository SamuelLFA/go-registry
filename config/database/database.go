package database

import (
	"log"

	"github.com/SamuelLFA/goregister/person"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectWithDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=register port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error in connection with database")
	}

	DB.AutoMigrate(&person.PersonModel{})
}
