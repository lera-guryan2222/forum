package controller

import (
	"net/http"

	"github.com/lera-guryan2222/forum/backend/auth-service/internal/usecase"
	"github.com/lera-guryan2222/forum/backend/auth-service/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{authUsecase: authUsecase}
}

func (ac *AuthController) Login(c *gin.Context) {
	var request usecase.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ac.authUsecase.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ac *AuthController) Register(c *gin.Context) {
	var request usecase.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ac.authUsecase.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate tokens after registration
	accessToken, err := auth.GenerateAccessToken(response.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	refreshToken, err := auth.GenerateRefreshToken(response.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"access_token":  accessToken,
		"user":          response.User,
		"refresh_token": refreshToken,
	})
}

func (ac *AuthController) Refresh(c *gin.Context) {
	var request usecase.RefreshRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ac.authUsecase.Refresh(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
