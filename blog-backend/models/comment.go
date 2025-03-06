package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string  `json:"content" gorm:"type:text;not null"`
	ArticleID uint    `json:"article_id" gorm:"not null;index"`
	ParentID  *uint   `json:"parent_id" gorm:"index"`
	UserID    uint    `json:"user_id" gorm:"not null"`
	User      User    `json:"user" gorm:"foreignKey:UserID"`
	Replies   []Comment `json:"replies" gorm:"foreignKey:ParentID"`
}
