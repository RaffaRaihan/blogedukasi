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

func UpdateUser(c *gin.Context) {
    // Ambil ID dari parameter URL
    id := c.Param("id")

    var user models.User
    // Cari user berdasarkan ID
    if err := config.GetDB().First(&user, id).Error; err != nil {
        // Jika tidak ditemukan, kembalikan pesan error
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Bind data JSON dari request body ke user
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan perubahan ke database
    if err := config.GetDB().Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    // Kembalikan data user yang sudah diperbarui
    c.JSON(http.StatusOK, gin.H{
        "message": "User updated successfully",
        "user":    user,
    })
}


func UploadProfilePhoto(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.GetDB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Cek apakah foto profil sudah ada
	if user.Foto != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile photo already exists. Use PUT to update."})
		return
	}

	// Lanjutkan unggah file
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	user.Foto = file.Filename
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Profile photo uploaded successfully",
		"profile_photo": file.Filename,
	})
}

func UpdateProfilePhoto(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.GetDB().First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Ambil file dari form-data
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	// Simpan file ke folder uploads
	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Perbarui kolom ProfilePhoto
	user.Foto = file.Filename
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Profile photo updated successfully",
		"profile_photo": file.Filename,
	})
}

func DeleteUser(c *gin.Context) {
    // Ambil ID dari parameter URL
    id := c.Param("id")

    var user models.User
    // Cari user berdasarkan ID
    if err := config.GetDB().First(&user, id).Error; err != nil {
        // Jika user tidak ditemukan, kembalikan pesan error
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Hapus user dari database
    if err := config.GetDB().Delete(&user).Error; err != nil {
        // Jika gagal menghapus, kembalikan pesan error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    // Kembalikan respons sukses
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}