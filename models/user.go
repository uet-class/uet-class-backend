package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID          int
	Email       string
	Password    string
	FullName    string
	DateOfBirth time.Time
	AvatarUrl   string
	PhoneNumber string
	Class       string
}
