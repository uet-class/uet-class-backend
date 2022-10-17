package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID       int
	Email    string
	Password string
}
