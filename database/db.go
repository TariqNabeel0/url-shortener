package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TariqNabeel0/url-shortener/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	psqlconn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect the db")
	}

err = db.AutoMigrate(&models.URL{})
if err != nil {
	log.Fatal("Failed to auto migrate", err)
}

	fmt.Println("DB connected successfully")
	DB = db
}