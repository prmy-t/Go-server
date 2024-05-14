package models

import (
	"example/go-server/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Title string `gorm:""json:"title"`
	Author string `gorm:""json:"autor"`
	Quantity string `gorm:""json:"quantity"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}