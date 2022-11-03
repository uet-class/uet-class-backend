package database

import (
	"fmt"
	"log"

	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitPostgres() {
	config := config.GetConfig()

	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.GetString("POSTGRES_HOST"),
		config.GetString("POSTGRES_PORT"),
		config.GetString("POSTGRES_USER"),
		config.GetString("POSTGRES_PASSWORD"),
		config.GetString("POSTGRES_DATABASE"))

	if db, err = gorm.Open(postgres.Open(datasource), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(
		&models.User{},
		&models.Report{},
		&models.ReportType{},
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
