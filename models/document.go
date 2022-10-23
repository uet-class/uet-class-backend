package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	CreatorID    int
	AttachmentID int
	ClassID      int
	Title        string
}
