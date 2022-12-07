package models

import (
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	ClassID    int
	CreatorID  int
	Title      string
	Content    string
	Duedate    string
	Attachment []Attachment `gorm:"foreignKey:AssignmentID"`
	Submission []Submission `gorm:"foreignKey:AssignmentID"`
}
