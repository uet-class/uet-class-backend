package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	ClassID      uint
	AssignmentID uint
	CreatorID    uint
	CreatorName  string
	BucketName   string
	FileName     string
}
