package user_repository

import "go-backend-stream/internal/models"

type UserRepository interface {
	GetByEmail(email string, user *models.User) error
	Create(user *models.User) (models.User, error)
	GetById(id string, user *models.User) error
	UpdatePassword(userID string, newPassword string) error
}
