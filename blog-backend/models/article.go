package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Label		string		`json:"label"`
	Title   	string 		`json:"title"`
	Content 	string 		`json:"content"`
	CategoryID  uint        `json:"category_id"` // Foreign Key
	Category    Category    `gorm:"foreignKey:CategoryID" json:"category"`
	Author  	string 		`json:"author"`
	FilePath    string    	`json:"file_path"` // Menyimpan path file yang diupload
    FileName    string    	`json:"file_name"` // Menyimpan nama file yang diupload
}
