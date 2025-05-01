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

    config.GetDB().AutoMigrate(&models.Article{}, &models.Comment{}, &models.Category{}, &models.User{}, &models.Message{},&models.ArticleView{}, &models.Report{}, &models.ArticleLike{})

    r := gin.Default()
    routes.SetupRoutes(r)

    port := os.Getenv("PORT")
	if port == "" {
	    port = "8080"
	}
	r.Run(":" + port)
}
