package router

import (
	"forum-service/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine, postCtrl *controller.PostController, authMiddleware gin.HandlerFunc) {
	posts := r.Group("/posts")
	posts.Use(authMiddleware)

	posts.POST("", postCtrl.CreatePost)
	posts.GET("", postCtrl.GetAllPosts)
	posts.GET("/:id", postCtrl.GetPostByID)
	posts.PUT("/:id", postCtrl.UpdatePost)
	posts.DELETE("/:id", postCtrl.DeletePost)
}
