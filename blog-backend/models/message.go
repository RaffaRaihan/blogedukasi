package models

import "gorm.io/gorm"

type Message struct {
    gorm.Model
    SenderID  	uint   	`json:"sender_id"`
    Nama        string  `json:"name"`
    Content   	string 	`json:"content" binding:"required"`
}
