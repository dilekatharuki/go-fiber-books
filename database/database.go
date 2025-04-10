package database

import (
	"log"

	"github.com/dilekatharuki/go-fiber-books/models"
	"github.com/glebarez/sqlite" // Pure Go SQLite driver that doesn't require cgo
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	
	database, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }
    database.AutoMigrate(&models.Book{})
    DB = database
}
