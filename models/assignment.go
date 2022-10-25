package models

import (
	"gorm.io/gorm"
	"time"
)

type Assignment struct {
	gorm.Model
	ClassID    int
	CreatorID  int
	Title      string
	Content    string
	DeadLine   time.Time
	Attachment []Attachment `gorm:"foreignKey:AssignmentID"`
	Submission []Submission `gorm:"foreignKey:AssignmentID"`
}
