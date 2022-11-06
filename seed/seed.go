package main

import (
	"fmt"
	"log"

	"github.com/uet-class/uet-class-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_HOST string = "localhost"
var DB_PORT string = "15432"
var DB_USER string = "uc_root"
var DB_PASSWORD string = "uc_pwd"
var DB_NAME string = "uet_class_dev"

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}

func Init() *gorm.DB {
	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_PASSWORD,
		DB_NAME)

	db, err := gorm.Open(postgres.Open(datasource), &gorm.Config{})
	if err != nil {
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

	return db
}

func generateUsers(quantity int) []models.User {
	var prefix_email string = "user"
	var prefix_password string = "resu"

	var users []models.User

	for i := 0; i < quantity; i++ {
		email := fmt.Sprintf("%s%d", prefix_email, i)
		password := fmt.Sprintf("%s%d", prefix_password, i)
		users = append(users, models.User{Email: email, Password: hashPassword(password)})
	}

	return users
}

func main() {
	db := Init()

	newUsers := generateUsers(10)
	db.Create(&newUsers)
}
