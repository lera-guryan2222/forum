package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/controller"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/delivery"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/repository"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/router"
	"github.com/lera-guryan2222/forum/backend/forum-service/internal/usecase"
)

func main() {
	logger := log.New(os.Stdout, "FORUM-SERVICE: ", log.LstdFlags|log.Lshortfile)

	// Подключение к PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=forum port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("Failed to get DB instance: %v", err)
	}
	defer sqlDB.Close()

	// Инициализация зависимостей
	authMiddleware := delivery.NewAuthMiddleware(logger)
	postRepo := repository.NewPostRepository(db)
	postUC := usecase.NewPostUseCase(postRepo)
	postCtrl := controller.NewPostController(postUC, logger)

	// Настройка Gin
	app := gin.Default()
	app.Use(delivery.LoggerMiddleware(logger))

	// Настройка маршрутов
	router.SetupRouter(postCtrl, authMiddleware)

	// Запуск сервера
	logger.Println("Starting server on :8081")
	if err := app.Run(":8081"); err != nil {
		logger.Fatalf("Failed to run server: %v", err)
	}
}
