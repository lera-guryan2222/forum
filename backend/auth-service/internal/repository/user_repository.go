package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/lera-guryan2222/auth-service/internal/entity"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByRefreshToken(refreshToken string) (*entity.User, error)
	Create(user *entity.User) error
	UpdateRefreshToken(userID string, refreshToken string) error
}

type InMemoryUserRepository struct {
	users map[string]*entity.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *InMemoryUserRepository) FindByUsername(username string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindByRefreshToken(refreshToken string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.RefreshToken == refreshToken {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) Create(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user == nil {
		return errors.New("user cannot be nil")
	}

	user.ID = time.Now().Format("20060102150405")
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) UpdateRefreshToken(userID string, refreshToken string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user, exists := r.users[userID]; exists {
		user.RefreshToken = refreshToken
		return nil
	}
	return errors.New("user not found")
}
