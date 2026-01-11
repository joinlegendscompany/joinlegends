package session_repository

import "go-backend-stream/internal/models"

type SessionRepository interface {
	Create(session *models.Session) (*models.Session, error)
	GetAllByUserID(userID string, sessions *[]models.Session) error
	GetByID(id string, session *models.Session) error
	DeleteByID(id string) error
	DeleteByUserId(userID string) error
}
