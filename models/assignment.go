package models

import (
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	ClassID     int    `gorm:"not null"`
	CreatorID   int    `gorm:"not null"`
	CreatorName string `gorm:"not null"`
	Title       string
	Content     string `gorm:"not null"`
	Duedate     string
	Attachment  []Attachment `gorm:"foreignKey:AssignmentID"`
	Submission  []Submission `gorm:"foreignKey:AssignmentID"`
}
