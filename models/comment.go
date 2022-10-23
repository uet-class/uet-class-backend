package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CreatorID int
	PostID    int
	Content   string
}
