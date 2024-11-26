package routes

import (
	"blog-backend/middlewares"
	"blog-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// Routes Public
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.GET("/articles", controllers.GetAll)
	r.GET("/articles/:id/comments", controllers.GetByArticle)

	// Group Admin (dengan middleware Auth)
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middlewares.AuthMiddleware())
	{
		// Admin User Management
		adminRoutes.GET("/users", controllers.GetUsers)

		// Admin Article Management
		adminRoutes.GET("/articles", controllers.GetAll)
		adminRoutes.POST("/articles", controllers.CreateArticle)
		adminRoutes.PUT("/articles/:id", controllers.Update)
		adminRoutes.DELETE("/articles/:id", controllers.Delete)
	}

	// Group User (dengan middleware Auth)
	userRoutes := r.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		// User Comment Management
		userRoutes.POST("/articles/:id/comments", controllers.CreateComment)
	}
}
