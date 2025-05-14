package usecase

import (
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/entity"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/repository"
)

type PostUseCase interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
	GetAllPosts() ([]entity.Post, error)
	GetPostByID(id uint) (*entity.Post, error)
	UpdatePost(post *entity.Post) (*entity.Post, error)
	DeletePost(id uint) error
}

type postUseCase struct {
	postRepo repository.PostRepository
}

func NewPostUseCase(postRepo repository.PostRepository) PostUseCase {
	return &postUseCase{postRepo: postRepo}
}

func (uc *postUseCase) CreatePost(post *entity.Post) (*entity.Post, error) {
	return uc.postRepo.Create(post)
}

func (uc *postUseCase) GetAllPosts() ([]entity.Post, error) {
	return uc.postRepo.GetAll()
}

func (uc *postUseCase) GetPostByID(id uint) (*entity.Post, error) {
	return uc.postRepo.GetByID(id)
}

func (uc *postUseCase) UpdatePost(post *entity.Post) (*entity.Post, error) {
	return uc.postRepo.Update(post)
}

func (uc *postUseCase) DeletePost(id uint) error {
	return uc.postRepo.Delete(id)
}
