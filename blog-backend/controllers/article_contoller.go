package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/// Create Article
func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid input", "error": err.Error()})
		return
	}

	// Validasi kategori
	var category models.Category
	if err := config.GetDB().First(&category, article.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid category ID", "error": err.Error()})
		return
	}

	if err := config.GetDB().Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create article", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": article})
}

// Get All Articles
func GetAll(c *gin.Context) {
	var articles []models.Article
	config.GetDB().Find(&articles)
	c.JSON(http.StatusOK, articles)
}

// Update Article
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.GetDB().First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Article not found", "error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid input", "error": err.Error()})
		return
	}

	// Validasi kategori
	var category models.Category
	if err := config.GetDB().First(&category, article.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid category ID", "error": err.Error()})
		return
	}

	if err := config.GetDB().Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update article", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": article})
}

// Delete Article
func Delete(c *gin.Context) {
	id := c.Param("id")
	if err := config.GetDB().Delete(&models.Article{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

// Search Articles
func SearchArticles(c *gin.Context) {
	// Ambil parameter pencarian
	query := c.Query("query")
	categoryID := c.Query("category_id")

	var articles []models.Article

	// Query dasar
	db := config.GetDB().Preload("Category")
	if query != "" {
		// Cari berdasarkan judul atau konten
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if categoryID != "" {
		// Filter berdasarkan kategori
		db = db.Where("category_id = ?", categoryID)
	}

	// Eksekusi query
	if err := db.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to search articles", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": articles})
}
