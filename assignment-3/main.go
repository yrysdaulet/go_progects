package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/yrysdaulet/go_progects/assignment-3/controllers"
	"github.com/yrysdaulet/go_progects/assignment-3/database"
)

var (
	db  *sql.DB
	err error
)

type User struct {
	Name     string
	Password string
}

type Item struct {
	ID       int
	Name     string
	Price    float32
	Raiting  float32
	Quantity int
}

func main() {
	r := gin.Default()

	database.Connect()
	database.Migrate()
	r.GET("/books/:id", controllers.GetBookByID)
	r.GET("/books", controllers.GetAllBooks)
	r.POST("/books", controllers.AddBook)
	r.PUT("/books/:id", controllers.UpdateBookByID)
	r.DELETE("/books/:id", controllers.DeleteBookByID)
	r.GET("/books/search", controllers.SearchBooksByTitle)
	r.GET("/books/sort", controllers.SortBooksByCost)

	r.Run(":8080")
}
