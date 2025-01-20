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
	r.GET("/articles/:id", controllers.GetArticleByID)
	r.GET("/category/:id", controllers.GetCategoryByID)
	r.GET("/articles/:id/comments", controllers.GetByArticle)

	// Uploads
	r.Static("/uploads", "./uploads")

	// Group Admin (dengan middleware Auth)
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		// Admin User Management
		adminRoutes.GET("/users", controllers.GetUsers)
		adminRoutes.GET("/users/:id", controllers.GetUserByID)
		adminRoutes.PUT("/users/:id", controllers.UpdateUser)
		adminRoutes.DELETE("/users/:id", controllers.DeleteUser)
		adminRoutes.POST("/users/:id/foto", controllers.UploadProfilePhoto)
		adminRoutes.PUT("/users/:id/foto", controllers.UpdateProfilePhoto)

		// Admin Article Management
		adminRoutes.GET("/articles", controllers.GetAll)
		adminRoutes.GET("/articles/:id", controllers.GetArticleByID)
		adminRoutes.POST("/articles", controllers.CreateArticle)
		adminRoutes.PUT("/articles/:id", controllers.UpdateArticle)
		adminRoutes.DELETE("/articles/:id", controllers.Delete)

		// Admin File Management
		adminRoutes.POST("/articles/:id/uploads", controllers.UploadFile)
		adminRoutes.PUT("/articles/:id/uploads", controllers.UpdateFile)

		// Admin Categories
		adminRoutes.POST("/category", controllers.CreateCategory)
		adminRoutes.PUT("/category/:id", controllers.UpdateCategory)
		adminRoutes.DELETE("/category/:id", controllers.DeleteCategory)

		// Admin Messages
		adminRoutes.GET("/messages", controllers.GetAllMessages)
		adminRoutes.POST("/send-message", controllers.SendMessage)

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
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.POST("/:id/foto", controllers.UploadProfilePhoto)
		userRoutes.PUT("/:id/foto", controllers.UpdateProfilePhoto)
		userRoutes.POST("/send-message", controllers.SendMessage)
	}
}
