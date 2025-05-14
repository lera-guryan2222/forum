package controller

import (
	"log"

	"github.com/lera-guryan2222/forum/backend/forum-service/internal/entity"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/usecase"
)

type postControllerImpl struct {
	postUseCase usecase.PostUseCase
	logger      *log.Logger
}

func NewPostController(postUseCase usecase.PostUseCase, logger *log.Logger) PostController {
	return &postControllerImpl{
		postUseCase: postUseCase,
		logger:      logger,
	}
}

func (c *postControllerImpl) CreatePost(req *entity.PostRequest, authorID uint) (*entity.PostResponse, error) {
	c.logger.Printf("Creating post. Title: %s, AuthorID: %d", req.Title, authorID)

	if req.Title == "" || req.Content == "" {
		return nil, entity.ErrInvalidPostData
	}

	post := &entity.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: authorID,
	}

	createdPost, err := c.postUseCase.CreatePost(post)
	if err != nil {
		c.logger.Printf("Error creating post: %v", err)
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

func (c *postControllerImpl) GetAllPosts() ([]entity.PostResponse, error) {
	c.logger.Println("Fetching all posts")
	posts, err := c.postUseCase.GetAllPosts()
	if err != nil {
		c.logger.Printf("Error fetching posts: %v", err)
		return nil, err
	}

	var responses []entity.PostResponse
	for _, post := range posts {
		responses = append(responses, entity.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			AuthorID:  post.AuthorID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	c.logger.Printf("Retrieved %d posts", len(responses))
	return responses, nil
}

func (c *postControllerImpl) GetPostByID(id uint) (*entity.PostResponse, error) {
	c.logger.Printf("Fetching post with ID: %d", id)
	post, err := c.postUseCase.GetPostByID(id)
	if err != nil {
		c.logger.Printf("Error fetching post %d: %v", id, err)
		return nil, err
	}

	c.logger.Printf("Successfully fetched post with ID: %d", id)
	return &entity.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (c *postControllerImpl) UpdatePost(id uint, update *entity.PostRequest) (*entity.PostResponse, error) {
	c.logger.Printf("Updating post with ID: %d", id)

	post := &entity.Post{
		ID:      id,
		Title:   update.Title,
		Content: update.Content,
	}

	updatedPost, err := c.postUseCase.UpdatePost(post)
	if err != nil {
		c.logger.Printf("Error updating post %d: %v", id, err)
		return nil, err
	}

	c.logger.Printf("Successfully updated post with ID: %d", id)
	return &entity.PostResponse{
		ID:        updatedPost.ID,
		Title:     updatedPost.Title,
		Content:   updatedPost.Content,
		AuthorID:  updatedPost.AuthorID,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: updatedPost.UpdatedAt,
	}, nil
}

func (c *postControllerImpl) DeletePost(id uint) error {
	c.logger.Printf("Deleting post with ID: %d", id)
	err := c.postUseCase.DeletePost(id)
	if err != nil {
		c.logger.Printf("Error deleting post %d: %v", id, err)
		return err
	}
	c.logger.Printf("Successfully deleted post with ID: %d", id)
	return nil
}
