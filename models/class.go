package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	TeacherID   int
	StudentID   int
	ClassName   string
	Description string
}
