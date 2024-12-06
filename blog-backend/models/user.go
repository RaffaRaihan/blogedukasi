package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID		 uint	`gorm:"primayKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Comments []Comment `json:"comments"`
	Foto	 string	`json:"foto"`
	Role     string `json:"role" gorm:"default:user"` // "admin" atau "user"
}