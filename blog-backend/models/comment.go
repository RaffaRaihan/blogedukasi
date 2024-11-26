package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `json:"content"`
	ArticleID uint   `json:"article_id"`
	UserID    uint   `json:"user_id"`
}
