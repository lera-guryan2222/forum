package router

import (
	"forum-service/internal/controller"
	"forum-service/pkg/auth"
	"forum-service/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter настраивает маршруты приложения
func SetupRouter(
	postController *controller.PostController,
	tokenManager auth.TokenManager,
	log logger.Logger,
) *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			// Public routes
			posts.GET("", postController.GetAllPosts)
			posts.GET("/:id", postController.GetPostByID)

			// Protected routes
			authorized := posts.Group("")
			authorized.Use(func(c *gin.Context) {
				token := c.GetHeader("Authorization")
				if token == "" {
					c.JSON(401, gin.H{"error": "unauthorized"})
					c.Abort()
					return
				}

				// Validate token and get user ID
				claims, err := tokenManager.ValidateToken(token)
				if err != nil {
					c.JSON(401, gin.H{"error": "invalid token"})
					c.Abort()
					return
				}

				c.Set("userID", claims.UserID)
				c.Next()
			})
			{
				authorized.POST("", postController.CreatePost)
				authorized.PUT("/:id", postController.UpdatePost)
				authorized.DELETE("/:id", postController.DeletePost)
			}
		}
	}

	return router
}
