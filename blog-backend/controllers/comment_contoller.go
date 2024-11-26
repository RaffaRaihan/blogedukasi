package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add Comment
func CreateComment(c *gin.Context) {
	articleID := c.Param("id")

	// Konversi articleID dari string ke uint
	articleIDUint, err := strconv.ParseUint(articleID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ArticleID = uint(articleIDUint)
	if err := config.GetDB().Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// Get Comments
func GetByArticle(c *gin.Context) {
	articleID := c.Param("id")
	var comments []models.Comment
	if err := config.GetDB().Where("article_id = ?", articleID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
