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
	psqlInfo string = "host=localhost user=postgres password=postgres dbname=assignment-3 port=5432 sslmode=disable "
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
