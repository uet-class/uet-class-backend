package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ID        int
	CreatorID int
	PostID    int
	Content   string
	CreatedAt time.Time
}
