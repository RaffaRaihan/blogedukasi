package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReportArticle(c *gin.Context) {
	var report models.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah artikel ada
	var article models.Article
	if err := config.GetDB().First(&article, report.ArticleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	// Simpan report ke database
	if err := config.GetDB().Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat laporan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Laporan berhasil dikirim"})
}

func GetAllReports(c *gin.Context) {
	var reports []models.Report
	if err := config.GetDB().Find(&reports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil laporan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reports, "messages": "success"})
}
