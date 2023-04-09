package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        uint   `json:"cost"`
}
