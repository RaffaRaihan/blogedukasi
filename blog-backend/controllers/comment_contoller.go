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

func ReplyToComment(c *gin.Context) {
    var input struct {
        CommentID uint   `json:"comment_id" binding:"required"`
        UserID    uint   `json:"user_id" binding:"required"`
        Content   string `json:"content" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Periksa apakah komentar induk ada
    var parentComment models.Comment
    if err := config.GetDB().First(&parentComment, input.CommentID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Parent comment not found"})
        return
    }

    // Buat balasan
    reply := models.Comment{
        ArticleID: parentComment.ArticleID,
        UserID:    input.UserID,
        Content:   input.Content,
        ParentID:  &input.CommentID, // Menyimpan ID komentar induk
    }

    if err := config.GetDB().Create(&reply).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Reply added successfully", "reply": reply})
}

// Edit Comment or Reply
func EditComment(c *gin.Context) {
	commentID := c.Param("id")
	var comment models.Comment

	// Cari komentar berdasarkan ID
	if err := config.GetDB().First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.GetDB().Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// Delete Comment or Reply
func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")

	// Hapus komentar berdasarkan ID
	if err := config.GetDB().Delete(&models.Comment{}, commentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// Get Comments
func GetByArticle(c *gin.Context) {
    articleID := c.Param("id")
    var comments []models.Comment

    // Gunakan Preload untuk memuat data User
    if err := config.GetDB().Preload("User").Where("article_id = ?", articleID).Find(&comments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, comments)
}