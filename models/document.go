package models

import (
	"gorm.io/gorm"
	"time"
)

type Document struct {
	gorm.Model
	ID        int
	CreatorID int
	AttachmentID    int
	ClassID    int
	Title string
	CreatedAt time.Time
}
