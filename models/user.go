package models

import (
	"gorm.io/gorm"
	"time"
)

// User has many Comments, CreatorID is the foreign key
// User has manny Reports, CreatorID is the foreign key
// User has many Posts, CreatorID is the foreign key
// User has many Documents, CreatorID is the foreign key
// User has many Assignments, CreatorID is the foreign key
// User has many Submissions, CreatorID is the foreign key
// User has and belongs to many Class, teacherID and studentID are the foreign keys, reference to id
type User struct {
	gorm.Model
	Email       string
	Password    string
	FullName    string
	DateOfBirth time.Time
	AvatarUrl   string
	PhoneNumber string
	ClassName   string
	Comment     []Comment    `gorm:"foreignKey:CreatorID"`
	Report      []Report     `gorm:"foreignKey:CreatorID"`
	Post        []Post       `gorm:"foreignKey:CreatorID"`
	Document    []Document   `gorm:"foreignKey:CreatorID"`
	Assignment  []Assignment `gorm:"foreignKey:CreatorID"`
	Submission  []Submission `gorm:"foreignKey:StudentID"`
	Class       []Class      `gorm:"many2many:class;ForeignKey:teacherID,studentID;References:id"`
}
