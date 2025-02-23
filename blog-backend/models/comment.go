package models

import "gorm.io/gorm"

type Comment struct {
    gorm.Model
    Content   string    `json:"content" gorm:"not null"`
    ArticleID uint      `json:"article_id" gorm:"not null"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user" gorm:"foreignKey:UserID"`
    ParentID  *uint     `json:"parent_id" gorm:"index"`
    Replies   []Comment `json:"replies" gorm:"foreignKey:ParentID"`
}
