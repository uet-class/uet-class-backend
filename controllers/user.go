package controllers

import (
	"github.com/uet-class/uet-class-backend/models"
	"gorm.io/gorm"
	"time"
)

func parseTime(t string) time.Time {
	result, err := time.Parse("02-01-2006", t)
	if err != nil {
		panic(err)
	}
	return result
}

func BatchInsertData(db *gorm.DB) {
	var users = []models.User{
		{
			Email:       "minh@gmail.com",
			Password:    "minh",
			FullName:    "Phạm Vũ Minh",
			DateOfBirth: parseTime("05-11-2001"),
		},
		{
			Email:       "thai@gmail.com",
			Password:    "thai",
			FullName:    "Nguyễn Minh Thái",
			DateOfBirth: parseTime("02-03-2001"),
		},
		{
			Email:       "quan@gmail.com",
			Password:    "quan",
			FullName:    "Võ Minh Quân",
			DateOfBirth: parseTime("30-08-2001"),
		},
	}

	db.Create(&users)
}

func MigrateModel(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Report{}, &models.ReportType{}, &models.Class{},
		&models.Post{}, &models.Document{}, &models.Assignment{}, &models.Submission{}, &models.Assignment{},
		&models.Comment{}, &models.Attachment{})
}
