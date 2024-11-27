package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID		 uint	`gorm:"primayKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"` // "admin" atau "user"
}