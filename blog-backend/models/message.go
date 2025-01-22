package models

import "gorm.io/gorm"

type Message struct {
    gorm.Model
    UserID  	uint   	`json:"user_id"`
    User        User    `json:"user" gorm:"foreignKey:UserID"`
    Nama        string  `json:"name"`
    Content   	string 	`json:"content" binding:"required"`
    Reply       string  `json:"reply"`
}
