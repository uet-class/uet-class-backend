package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	AssignmentID int
	StudentID    int
	Content      string
	Grade        float32
	Attachment   []Attachment `gorm:"foreignKey:SubmissionID"`
}
