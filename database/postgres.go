package database

import (
	"fmt"
	"log"
	"os"

	"github.com/uet-class/uet-class-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitPostgres() {

	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"))

	if db, err = gorm.Open(postgres.Open(datasource), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(
		&models.User{},
		&models.Report{},
		&models.Class{},
		&models.Post{},
		&models.Document{},
		&models.Assignment{},
		&models.Submission{},
		&models.Assignment{},
		&models.Comment{},
		&models.Attachment{}); err != nil {
		log.Fatal(err)
	}
}

func GetDatabase() *gorm.DB {
	return db
}
