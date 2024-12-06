package routes

import (
	"blog-backend/controllers"
	"blog-backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// setup cors dengan konfigurasi yang tepat
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Hanya izinkan asal dari frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Metode yang diizinkan
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Izinkan header Authorization
		AllowCredentials: true, // Izinkan cookies untuk dikirimkan
	}))

	// Routes Public
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.GET("/articles", controllers.GetAll)
	r.GET("/category", controllers.GetAllCategory)
	r.GET("/articles/:id/comments", controllers.GetByArticle)

	// Uploads
	r.Static("/uploads", "./uploads")

	// Group Admin (dengan middleware Auth)
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		// Admin User Management
		adminRoutes.GET("/users", controllers.GetUsers)

		// Admin Article Management
		adminRoutes.GET("/articles", controllers.GetAll)
		adminRoutes.POST("/articles", controllers.CreateArticle)
		adminRoutes.PUT("/articles/:id", controllers.UpdateArticle)
		adminRoutes.DELETE("/articles/:id", controllers.Delete)

		//upload file
		adminRoutes.POST("/articles/:id/uploads", controllers.UploadFile)
	}

	// Group User (dengan middleware Auth)
	userRoutes := r.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		// User Comment Management
		userRoutes.POST("/articles/:id/comments", controllers.CreateComment)
		userRoutes.GET("/articles/search", controllers.SearchArticles)
		userRoutes.GET("/articles/:id", controllers.GetArticleByID)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.POST("/:id/foto", controllers.UploadProfilePhoto)
		userRoutes.PUT("/:id/foto", controllers.UpdateProfilePhoto)
	}
}
