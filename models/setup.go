package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("playlist.db"), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database!")
	}

	err = database.AutoMigrate(&Playlist{})

	if err != nil {
		return
	}

	DB = database
}
