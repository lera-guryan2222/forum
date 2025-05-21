package controller

import (
	"context"
	"fmt"

	"backend.com/forum/proto"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/entity"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type PostController interface {
	GetAllPosts() ([]*entity.Post, error)
	GetPostByID(id uint) (*entity.Post, error)
	CreatePost(req *entity.PostRequest, authorID uint) (*entity.Post, error)
	UpdatePost(id uint, req *entity.PostRequest) (*entity.Post, error)
	DeletePost(id uint) error
}

type postController struct {
	client proto.ForumServiceClient
}

func NewPostController(conn *grpc.ClientConn) PostController {
	client := proto.NewForumServiceClient(conn)
	return &postController{client: client}
}

func (c *postController) GetAllPosts() ([]*entity.Post, error) {
	response, err := c.client.ListPosts(context.Background(), &proto.ListPostsRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	posts := make([]*entity.Post, 0, len(response.Posts))
	for _, post := range response.Posts {
		posts = append(posts, &entity.Post{
			Model:    gorm.Model{ID: uint(post.Id)}, // Используем правильное поле ID
			Title:    post.Title,
			Content:  post.Content,
			AuthorID: uint(post.AuthorId), // Прямое преобразование uint64 -> uint
		})
	}

	return posts, nil
}

func (c *postController) GetPostByID(id uint) (*entity.Post, error) {
	response, err := c.client.GetPost(context.Background(), &proto.GetPostRequest{
		PostId: uint64(id), // Прямая передача числового ID
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	return &entity.Post{
		Model:    gorm.Model{ID: uint(response.Id)},
		Title:    response.Title,
		Content:  response.Content,
		AuthorID: uint(response.AuthorId),
	}, nil
}

func (c *postController) CreatePost(req *entity.PostRequest, authorID uint) (*entity.Post, error) {
	response, err := c.client.CreatePost(context.Background(), &proto.CreatePostRequest{
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: uint64(authorID), // Прямое преобразование uint -> uint64
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return &entity.Post{
		Model:    gorm.Model{ID: uint(response.Id)},
		Title:    response.Title,
		Content:  response.Content,
		AuthorID: uint(response.AuthorId),
	}, nil
}

func (c *postController) UpdatePost(id uint, req *entity.PostRequest) (*entity.Post, error) {
	response, err := c.client.UpdatePost(context.Background(), &proto.UpdatePostRequest{
		PostId:  uint64(id), // Прямая передача числового ID
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}

	return &entity.Post{
		Model:    gorm.Model{ID: uint(response.Id)},
		Title:    response.Title,
		Content:  response.Content,
		AuthorID: uint(response.AuthorId),
	}, nil
}

func (c *postController) DeletePost(id uint) error {
	_, err := c.client.DeletePost(context.Background(), &proto.DeletePostRequest{
		PostId: uint64(id), // Прямая передача числового ID
	})
	return err
}
