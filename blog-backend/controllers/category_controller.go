package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory membuat kategori baru
func CreateCategory(c *gin.Context) {
	var category models.Category

	// Bind JSON ke model Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid input", "error": err.Error()})
		return
	}

	// Simpan kategori ke database
	if err := config.GetDB().Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create category", "error": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": category})
}

// Get All category
func GetAllCategory(c *gin.Context) {
    var categories []models.Category
    if err := config.GetDB().Preload("Articles").Find(&categories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}
