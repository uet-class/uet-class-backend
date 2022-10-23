package models

import (
	"gorm.io/gorm"
	"time"
)

type Attachment struct {
	gorm.Model
	ID        int
	CreatorID int
	Attachment string
	CreatedAt time.Time
}
