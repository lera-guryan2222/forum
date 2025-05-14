package main

import (
	"github.com/lera-guryan2222/forum/backend/auth-service/internal/controller"
	"github.com/lera-guryan2222/forum/backend/auth-service/internal/repository"
	"github.com/lera-guryan2222/forum/backend/auth-service/internal/router"
	"github.com/lera-guryan2222/forum/backend/auth-service/internal/usecase"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()

	// Инициализация репозитория
	userRepo := repository.NewInMemoryUserRepository()

	// Инициализация usecase и controller
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authController := controller.NewAuthController(authUsecase) // Убрали log

	// Настройка роутера
	r := router.SetupRouter(authController) // Убрали log

	// Запуск сервера
	log.Info("Auth Service is running",
		zap.String("port", "8080"),
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server",
			zap.Error(err),
		)
	}
}
