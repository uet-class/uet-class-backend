package models

import (
	"time"
)

type User struct {
	ID          int `gorm:"primaryKey"`
	Email       string
	Password    string
	FullName    string
	DateOfBirth time.Time
	AvatarUrl   string
	PhoneNumber string
	Class       string
}
