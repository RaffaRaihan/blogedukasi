package models

import (
	"time"

	"gorm.io/gorm"
)

type ArticleView struct {
	gorm.Model
	ArticleID uint      `gorm:"index" json:"article_id"`
	ViewedAt  time.Time `gorm:"autoCreateTime" json:"viewed_at"`
}
