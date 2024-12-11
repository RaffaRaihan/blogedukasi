package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Article
func CreateArticle(c *gin.Context) {
	var article models.Article
	// Mengikat JSON dari request ke model Article
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

	// Buat artikel baru
	if err := config.GetDB().Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to create article", "error": err.Error()})
		return
	}

	// Ambil artikel yang baru dibuat dengan preload relasi Category
	if err := config.GetDB().Preload("Category").First(&article, article.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to preload category", "error": err.Error()})
		return
	}

	// Kirim response sukses dengan artikel yang sudah memiliki relasi Category
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": article})
}


// Get All Articles
func GetAll(c *gin.Context) {
    var articles []models.Article
    db := config.GetDB()

    // Ambil query parameters
    categoryID := c.Query("categoryId")
    search := c.Query("search")

    // Filter berdasarkan kategori jika categoryId tersedia
    if categoryID != "" {
        db = db.Where("category_id = ?", categoryID)
    }

    // Filter berdasarkan pencarian jika search tersedia
    if search != "" {
        db = db.Where("title LIKE ?", "%"+search+"%")
    }

    // Muat artikel beserta kategori
    if err := db.Preload("Category").Find(&articles).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	if err := db.Order("created_at desc").Find(&articles).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, articles)
}


// Get Article by ID
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	// Mencari artikel berdasarkan ID
	if err := config.GetDB().Preload("Category").First(&article, id).Error; err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Article not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to retrieve article", "error": err.Error()})
		}
		return
	}

	// Jika artikel ditemukan, kembalikan data artikel
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": article})
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
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	sort := c.DefaultQuery("sort", "created_at desc")

	// Konversi page dan limit ke integer
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}

	var articles []models.Article
	var total int64

	// Query dasar
	db := config.GetDB().Preload("Category")
	if query != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if categoryID != "" {
		db = db.Where("category_id = ?", categoryID)
	}

	// Hitung total data
	db.Model(&models.Article{}).Count(&total)

	// Eksekusi query dengan pagination dan sorting
	if err := db.Order(sort).Offset((pageInt - 1) * limitInt).Limit(limitInt).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to search articles", "error": err.Error()})
		return
	}

	// Kirim response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   articles,
		"meta": gin.H{
			"page":       pageInt,
			"limit":      limitInt,
			"total_data": total,
		},
	})
}


// Upload File
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No file is received", "error": err.Error()})
		return
	}

	// Validasi ukuran dan tipe file
	if file.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "File size exceeds 2MB"})
		return
	}

	if contentType := file.Header.Get("Content-Type"); contentType != "image/jpeg" && contentType != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid file type"})
		return
	}

	// Simpan file
	uploadPath := "./uploads/"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	filePath := filepath.Join(uploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Unable to save file", "error": err.Error()})
		return
	}

	// Update artikel dengan nama file
	var article models.Article
	if err := config.GetDB().Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Article not found", "error": err.Error()})
		return
	}

	article.FileName = file.Filename
	article.FilePath = filePath // path lengkap ke file
	if err := config.GetDB().Save(&article).Error; err != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update article with file name", "error": err.Error()})
	    return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "File uploaded successfully", "data": article})
}

// Update File
func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	// Cari artikel berdasarkan ID
	if err := config.GetDB().Where("id = ?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Article not found", "error": err.Error()})
		return
	}

	// Terima file baru
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "No file is received", "error": err.Error()})
		return
	}

	// Validasi ukuran dan tipe file
	if file.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "File size exceeds 2MB"})
		return
	}
	if contentType := file.Header.Get("Content-Type"); contentType != "image/jpeg" && contentType != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid file type"})
		return
	}

	// Direktori upload
	uploadPath := "./uploads/"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}

	// Hapus file lama jika ada
	if article.FileName != "" {
		oldFilePath := filepath.Join(uploadPath, article.FileName)
		if _, err := os.Stat(oldFilePath); err == nil {
			os.Remove(oldFilePath)
		}
	}

	// Simpan file baru
	newFilePath := filepath.Join(uploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, newFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Unable to save new file", "error": err.Error()})
		return
	}

	// Update artikel dengan nama file baru
	article.FileName = file.Filename
	if err := config.GetDB().Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update article with new file name", "error": err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File updated successfully",
		"data":    article,
	})
}
