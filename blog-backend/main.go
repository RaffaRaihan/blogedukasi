package main

import (
	"blog-backend/config"
	"blog-backend/models"
	"blog-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// load env
	config.LoadEnv()
	// connect database
	config.Connect()

    config.GetDB().AutoMigrate(&models.Article{}, &models.User{}, &models.Comment{})

    r := gin.Default()
    routes.SetupRoutes(r)

    r.Run(":8080")
}
