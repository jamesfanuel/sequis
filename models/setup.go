package models

import (
	"log"
	
	"gorm.io/driver/postgres"
  
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=jamesfanuel password=katakunci dbname=sequis port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db connection error")
	}

	DB = db
}