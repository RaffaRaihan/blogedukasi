package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
    var message models.Message

    // Bind data JSON dari request body ke model Message
    if err := c.ShouldBindJSON(&message); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan pesan ke database
    if err := config.GetDB().Create(&message).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Message sent successfully",
        "data":    message,
    })
}

func GetAllMessages(c *gin.Context) {
    var messages []models.Message

    // Ambil semua pesan dari database
    if err := config.GetDB().Find(&messages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Messages fetched successfully",
        "data":    messages,
    })
}

func ReplyToMessage(c *gin.Context) {
    var replyInput struct {
        Reply string `json:"reply" binding:"required"`
    }

    // Ambil ID pesan dari parameter URL
    id := c.Param("id")

    // Bind data JSON dari request body ke input balasan
    if err := c.ShouldBindJSON(&replyInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var message models.Message

    // Cari pesan berdasarkan ID
    if err := config.GetDB().First(&message, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
        return
    }

    // Perbarui kolom reply
    message.Reply = replyInput.Reply

    // Simpan perubahan ke database
    if err := config.GetDB().Save(&message).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reply to message"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Reply added successfully",
        "data":    message,
    })
}

func GetReply(c *gin.Context) {
    // Ambil ID pesan dari parameter URL
    id := c.Param("id")

    var message models.Message

    // Cari pesan berdasarkan ID
    if err := config.GetDB().First(&message, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
        return
    }

    // Pastikan ada balasan untuk pesan
    if message.Reply == "" {
        c.JSON(http.StatusOK, gin.H{
            "message": "No reply for this message",
            "data":    nil,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Reply fetched successfully",
        "data":    message.Reply,
    })
}

func GetMessagesByUserID(c *gin.Context) {
    // Ambil user_id dari parameter URL
    userIDParam := c.Param("id")

    // Konversi user_id ke uint
    userID, err := strconv.ParseUint(userIDParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var messages []models.Message

    // Ambil semua pesan yang dimiliki user dari database
    if err := config.GetDB().Where("user_id = ?", uint(userID)).Find(&messages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Messages fetched successfully",
        "data":    messages,
    })
}
