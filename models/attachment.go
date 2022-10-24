package models

import (
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	CreatorID  int
	Attachment string
	DocumentID int
	PostID int
	AssignmentID int
	SubmissionID int
}
