package controller

import "github.com/lera-guryan2222/forum/backend/forum-service/internal/entity"

type PostController interface {
	CreatePost(post *entity.PostRequest, authorID uint) (*entity.PostResponse, error)
	GetAllPosts() ([]entity.PostResponse, error)
	GetPostByID(id uint) (*entity.PostResponse, error)
	UpdatePost(id uint, update *entity.PostRequest) (*entity.PostResponse, error)
	DeletePost(id uint) error
}
