package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/lera-guryan2222/forum/backend/forum-service/internal/controller"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/delivery"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/entity"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/repository"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/router"
	"github.com/lera-guryan2222/forum/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "forum"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("database instance error: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping error: %w", err)
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.ChatMessage{},
		&entity.Token{},
		&entity.EmailVerification{},
	)
}

func main() {
	logger := log.New(os.Stdout, "[FORUM] ", log.LstdFlags|log.Lshortfile)

	db, err := Connect()
	if err != nil {
		logger.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	if err := autoMigrate(db); err != nil {
		logger.Fatalf("Migration failed: %v", err)
	}

	// Инициализация репозиториев
	postRepo := repository.NewPostRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Инициализация gRPC сервера
	grpcServer := grpc.NewServer()
	forumService := &ForumServiceServer{repo: postRepo}
	proto.RegisterForumServiceServer(grpcServer, forumService)

	// Запуск gRPC сервера
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			logger.Fatalf("failed to listen: %v", err)
		}
		logger.Println("gRPC server started on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	// Создание gRPC клиента
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Инициализация контроллеров (ИСПРАВЛЕНО: убран postRepo)
	postCtrl := controller.NewPostController(conn)

	// Middleware
	authMiddleware := delivery.NewAuthMiddleware(logger, userRepo)

	// Роутер
	router := router.SetupRouter(postCtrl, authMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Printf("HTTP server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server failed: %v", err)
	}
}

// ForumServiceServer реализует серверную часть gRPC сервиса
type ForumServiceServer struct {
	proto.UnimplementedForumServiceServer
	repo repository.PostRepository
}

func (s *ForumServiceServer) CreatePost(
	ctx context.Context,
	req *proto.CreatePostRequest,
) (*proto.Post, error) {
	post := &entity.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: uint(req.AuthorId),
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	return &proto.Post{
		Id:       uint64(post.ID),
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: uint64(post.AuthorID),
	}, nil
}

func (s *ForumServiceServer) GetPost(
	ctx context.Context,
	req *proto.GetPostRequest,
) (*proto.Post, error) {
	post, err := s.repo.GetByID(uint(req.PostId))
	if err != nil {
		return nil, err
	}

	return &proto.Post{
		Id:       uint64(post.ID),
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: uint64(post.AuthorID),
	}, nil
}

func (s *ForumServiceServer) ListPosts(
	ctx context.Context,
	req *proto.ListPostsRequest,
) (*proto.ListPostsResponse, error) {
	posts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var protoPosts []*proto.Post
	for _, post := range posts {
		protoPosts = append(protoPosts, &proto.Post{
			Id:       uint64(post.ID),
			Title:    post.Title,
			Content:  post.Content,
			AuthorId: uint64(post.AuthorID),
		})
	}

	return &proto.ListPostsResponse{Posts: protoPosts}, nil
}
