package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TrackView mencatat tampilan artikel
func TrackView(c *gin.Context) {
	var input struct {
		ArticleID uint `json:"article_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	view := models.ArticleView{
		ArticleID: input.ArticleID,
		ViewedAt:  time.Now(),
	}

	if err := config.GetDB().Create(&view).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan view"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "View berhasil dicatat"})
}

// GetPopularArticles mengambil daftar artikel yang paling sering dilihat
func GetPopularArticles(c *gin.Context) {
	var results []struct {
		ArticleID uint   `json:"article_id"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
		Views     int64  `json:"views"`
	}

	// Query dengan JOIN ke tabel articles
	config.GetDB().Table("article_views").
		Select("articles.id as article_id, articles.title, articles.created_at, COUNT(article_views.article_id) as views").
		Joins("JOIN articles ON articles.id = article_views.article_id").
		Group("article_views.article_id, articles.id, articles.title, articles.created_at").
		Order("views DESC").
		Limit(10).
		Scan(&results)

	// Kirim response ke frontend
	c.JSON(http.StatusOK, results)
}
