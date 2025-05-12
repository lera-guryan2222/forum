package usecase

import (
	"forum-service/internal/entity"
	"forum-service/internal/repository"
)

type PostUseCase interface {
	CreatePost(post *entity.PostRequest, authorID uint) (*entity.PostResponse, error)
	GetAllPosts() ([]entity.PostResponse, error)
	GetPostByID(id uint) (*entity.PostResponse, error)
	UpdatePost(id uint, update *entity.PostRequest) (*entity.PostResponse, error)
	DeletePost(id uint) error
}

type postUseCase struct {
	postRepo repository.PostRepository
}

func NewPostUseCase(postRepo repository.PostRepository) PostUseCase {
	return &postUseCase{postRepo}
}

func (uc *postUseCase) CreatePost(req *entity.PostRequest, authorID uint) (*entity.PostResponse, error) {
	post := &entity.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: authorID,
	}

	createdPost, err := uc.postRepo.Create(post)
	if err != nil {
		return nil, err
	}

	return &entity.PostResponse{
		ID:        createdPost.ID,
		Title:     createdPost.Title,
		Content:   createdPost.Content,
		AuthorID:  createdPost.AuthorID,
		CreatedAt: createdPost.CreatedAt,
		UpdatedAt: createdPost.UpdatedAt,
	}, nil
}

// Реализации остальных методов аналогичны...
