package main

import (
	"github.com/lera-guryan2222/auth-service/internal/controller"
	"github.com/lera-guryan2222/auth-service/internal/repository"
	"github.com/lera-guryan2222/auth-service/internal/router"
	"github.com/lera-guryan2222/auth-service/internal/usecase"
	"github.com/lera-guryan2222/auth-service/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 1. Initialize logger
	log := logger.NewLogger()

	// 2. Initialize repository
	userRepo := repository.NewInMemoryUserRepository()

	// 3. Initialize usecase and controller
	authUsecase := usecase.NewAuthUsecase(userRepo)
	authController := controller.NewAuthController(authUsecase)

	// 4. Setup router
	r := router.SetupRouter(log, authController)

	// 5. Start server
	log.Info("Auth Service is running on port :8080", zap.String("port", "8080"))
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
