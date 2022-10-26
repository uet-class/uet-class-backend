package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	TeacherID   []User       `gorm:"many2many:teacher_class;"`
	StudentID   []User       `gorm:"many2many:student_class;"`
	Post        []Post       `gorm:"foreignKey:ClassID"`
	Assignment  []Assignment `gorm:"foreignKey:ClassID"`
	Document    []Document   `gorm:"foreignKey:ClassID"`
	ClassName   string
	Description string
}
