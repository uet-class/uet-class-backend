package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	CreatorID    int
	AttachmentID []Attachment	`gorm:"foreignKey:PostID"`
	ClassID      int
	Content      string
	Title        string
}
