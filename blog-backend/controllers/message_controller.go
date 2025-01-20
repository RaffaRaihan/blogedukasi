package controllers

import (
    "blog-backend/config"
    "blog-backend/models"
    "net/http"

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
