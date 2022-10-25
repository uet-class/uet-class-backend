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
	Email        string `gorm:"unique"`
	Password     string
	FullName     string
	DateOfBirth  time.Time
	AvatarUrl    string
	PhoneNumber  string
	ClassName    string
	Comment      []Comment    `gorm:"foreignKey:CreatorID"`
	Report       []Report     `gorm:"foreignKey:ReporterID"`
	Post         []Post       `gorm:"foreignKey:CreatorID"`
	Document     []Document   `gorm:"foreignKey:CreatorID"`
	Assignment   []Assignment `gorm:"foreignKey:CreatorID"`
	Submission   []Submission `gorm:"foreignKey:StudentID"`
	ClassStudent []Class      `gorm:"many2many:teacher_class;"`
	ClassTeacher []Class      `gorm:"many2many:teacher_class;"`
}
