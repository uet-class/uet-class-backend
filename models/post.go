package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	CreatorID    int
	AttachmentID int
	ClassID      int
	Content      string
	Title        string
}
