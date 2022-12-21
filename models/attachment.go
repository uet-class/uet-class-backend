// Currently we will not use this model due to the lack of time before the MVP debut.
package models

import (
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	CreatorID    int
	Attachment   string
	DocumentID   int
	PostID       int
	AssignmentID int
	SubmissionID int
}
