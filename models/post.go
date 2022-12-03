package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ClassID   int       `gorm:"not null"`
	CreatorID int       `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Comment   []Comment `gorm:"foreignKey:PostID"`
}
