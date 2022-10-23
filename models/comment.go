package models

import (
	"time"
)

type Comment struct {
	ID        int `gorm:"primaryKey"`
	CreatorID int
	PostID    int
	Content   string
	CreatedAt time.Time
}
