package main

import (
	"blog-backend/config"
	"blog-backend/models"
	"blog-backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// load env
	config.LoadEnv()
	// connect database
	config.Connect()

    config.GetDB().AutoMigrate(&models.Article{}, &models.Comment{}, &models.Category{}, &models.User{}, &models.Message{},&models.ArticleView{}, &models.Report{}, &models.ArticleLike{})

    r := gin.Default()
    routes.SetupRoutes(r)

    port := os.Getenv("PORT")
	r.Run(":" + port)
}
