package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID        int
	CreatorID int
	AttachmentID    int
	ClassID    int
	Content    string
	Title string
	CreatedAt time.Time
}
