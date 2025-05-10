package usecase

import (
	"auth-service/internal/entity"
	"auth-service/internal/repository"
	"auth-service/pkg/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(request LoginRequest) (*LoginResponse, error)
	Register(request RegisterRequest) (*RegisterResponse, error)
	Refresh(request RefreshRequest) (*RefreshResponse, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string       `json:"access_token"`
		RefreshToken string       `json:"refresh_token"`
		User         *entity.User `json:"user"`
	}

	RegisterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		User *entity.User `json:"user"`
	}

	RefreshRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (uc *authUsecase) Login(request LoginRequest) (*LoginResponse, error) {
	user, err := uc.userRepo.FindByEmail(request.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.UpdateRefreshToken(user.ID, refreshToken); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

func (uc *authUsecase) Register(request RegisterRequest) (*RegisterResponse, error) {
	if request.Username == "" || request.Email == "" || request.Password == "" {
		return nil, errors.New("username, email and password are required")
	}

	existingUser, _ := uc.userRepo.FindByUsername(request.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	existingEmailUser, _ := uc.userRepo.FindByEmail(request.Email)
	if existingEmailUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &RegisterResponse{User: user}, nil
}

func (uc *authUsecase) Refresh(request RefreshRequest) (*RefreshResponse, error) {
	user, err := uc.userRepo.FindByRefreshToken(request.RefreshToken)
	if err != nil || user == nil {
		return nil, errors.New("invalid refresh token")
	}

	accessToken, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.UpdateRefreshToken(user.ID, newRefreshToken); err != nil {
		return nil, err
	}

	return &RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
