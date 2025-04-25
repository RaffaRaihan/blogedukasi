package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Label		string		`json:"label"`
	Title   	string 		`json:"title"`
	Content 	string 		`json:"content"`
	CategoryID  uint        `json:"category_id"` // Foreign Key
	Category    Category    `gorm:"foreignKey:CategoryID" json:"category"`
	AuthorID	uint		`json:"author_id"`
	Author   	User   		`gorm:"foreignKey:AuthorID" json:"author"`
	FilePath    string    	`json:"file_path"` // Menyimpan path file yang diupload
    FileName    string    	`json:"file_name"` // Menyimpan nama file yang diupload
	Status		string		`json:"status"`
	Note		string		`json:"note"`
	Likes     []ArticleLike `gorm:"foreignKey:ArticleID"`
}

type ArticleStatus struct {
    Status string `json:"status"`
    Count  int64  `json:"count"`
}

type ArticleLike struct {
	gorm.Model
	UserID    uint
	ArticleID uint
}
