package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (Post) TableName() string {
	return "posts"
}
