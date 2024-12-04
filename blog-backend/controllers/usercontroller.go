package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default role jika bukan admin
	if user.Role != "admin" {
		user.Role = "user"
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Save user
	if err := config.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}


// Get Users (Admin Only)
func GetUsers(c *gin.Context) {
	var users []models.User
	config.GetDB().Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	var user models.User
	// Cari user berdasarkan ID
	if err := config.GetDB().First(&user, id).Error; err != nil {
		// Jika tidak ditemukan, kembalikan pesan error
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Jika ditemukan, kembalikan data user
	c.JSON(http.StatusOK, user)
}
