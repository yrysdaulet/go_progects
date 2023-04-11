package database

import (
	"github.com/yrysdaulet/go_progects/assignment-3/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB       *gorm.DB
	dbError  error
	psqlInfo string = "postgres://postgres:postgres@db:5432/assignment-3"
)

func Connect() {
	DB, dbError = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot ot connect to db!")
	}
}
func Migrate() {
	dbError = DB.AutoMigrate(&models.Book{})
	if dbError != nil {
		panic(dbError.Error())
		return
	}
}
