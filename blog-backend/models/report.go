package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Nama      string  `json:"name"`
	Reason    string  `json:"reason"`
	ArticleID uint    `json:"article_id"`
	Reply     string  `json:"reply"`
}