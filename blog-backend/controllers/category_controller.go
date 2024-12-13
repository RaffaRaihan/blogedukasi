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

// GetCategoryByID mendapatkan kategori berdasarkan ID
func GetCategoryByID(c *gin.Context) {
	var category models.Category

	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Cari kategori berdasarkan ID
	if err := config.GetDB().Preload("Articles").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Category not found", "error": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// UpdateCategory memperbarui kategori berdasarkan ID
func UpdateCategory(c *gin.Context) {
	var category models.Category

	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Cari kategori berdasarkan ID
	if err := config.GetDB().First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Category not found", "error": err.Error()})
		return
	}

	// Bind JSON ke model Category untuk memperbarui data
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid input", "error": err.Error()})
		return
	}

	// Simpan perubahan ke database
	if err := config.GetDB().Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update category", "error": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": category})
}

// DeleteCategory menghapus kategori berdasarkan ID
func DeleteCategory(c *gin.Context) {
	var category models.Category

	// Ambil ID dari parameter URL
	id := c.Param("id")

	// Cari kategori berdasarkan ID
	if err := config.GetDB().First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Category not found", "error": err.Error()})
		return
	}

	// Hapus kategori dari database
	if err := config.GetDB().Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete category", "error": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Category deleted successfully"})
}
