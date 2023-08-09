package models

import (
	"github.com/AlapanGhosh10/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publicaiton string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}