package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Username string `json:"username"`
}