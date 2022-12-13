package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	AssignmentID uint
	CreatorID    uint
	CreatorName  string
	FileName     string
}
