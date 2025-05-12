package controller

import (
	"forum-service/internal/entity"
	"forum-service/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postUC usecase.PostUseCase
}

func NewPostController(postUC usecase.PostUseCase) *PostController {
	return &PostController{postUC}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	var req entity.PostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorID := ctx.GetUint("userID")
	res, err := c.postUC.CreatePost(&req, authorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// Остальные методы контроллера...
