package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yrysdaulet/go_progects/assignment-3/database"
	"github.com/yrysdaulet/go_progects/assignment-3/models"
	"net/http"
)

type Book models.Book

func GetBookByID(c *gin.Context) {
	var book Book
	id := c.Param("id")
	print(id)
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context) {
	var books []Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBookByID(c *gin.Context) {
	var book Book
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBookByID(c *gin.Context) {
	var book Book
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	database.DB.Delete(&book)
	c.Status(http.StatusNoContent)
}

func SearchBooksByTitle(c *gin.Context) {
	title := c.Query("title")
	var books []Book
	database.DB.Where("title LIKE ?", "%"+title+"%").Find(&books)
	c.JSON(http.StatusOK, books)
}

func SortBooksByCost(c *gin.Context) {

	order := c.DefaultQuery("order", "desc")
	var books []Book
	database.DB.Order("cost " + order).Find(&books)
	c.JSON(http.StatusOK, books)
}
