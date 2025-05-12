package main

import (
	"log"
	"os"

	"github.com/lera-guryan2222/forum-service/internal/controller"
	"github.com/lera-guryan2222/forum-service/internal/repository"
	"github.com/lera-guryan2222/forum-service/internal/router"
	"github.com/lera-guryan2222/forum-service/internal/usecase"
)

func main() {
	// Инициализация логгера
	logger := log.New(os.Stdout, "FORUM-SERVICE: ", log.LstdFlags|log.Lshortfile)

	// Инициализация приложения
	application, err := app.NewApp()
	if err != nil {
		logger.Fatalf("Failed to initialize app: %v", err)
	}

	// Инициализация зависимостей
	postRepo := repository.NewPostRepository(application.DB)
	postUC := usecase.NewPostUseCase(postRepo)
	postCtrl := controller.NewPostController(postUC, logger)

	// Настройка маршрутов
	r := router.SetupRouter(postCtrl, application.AuthMiddleware)

	// Запуск сервера
	logger.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Fatalf("Failed to run server: %v", err)
	}
}
