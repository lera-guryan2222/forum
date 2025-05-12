package repository

import (
	"forum-service/internal/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) (*entity.Post, error)
	GetAll() ([]entity.Post, error)
	GetByID(id uint) (*entity.Post, error)
	Update(post *entity.Post) (*entity.Post, error)
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(post *entity.Post) (*entity.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) GetAll() ([]entity.Post, error) {
	var posts []entity.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetByID(id uint) (*entity.Post, error) {
	var post entity.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(post *entity.Post) (*entity.Post, error) {
	if err := r.db.Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Post{}, id).Error
}
