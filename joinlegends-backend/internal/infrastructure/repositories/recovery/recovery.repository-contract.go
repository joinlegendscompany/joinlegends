package recovery_repository

import (
	"go-backend-stream/internal/models"
	"time"
)

type RecoveryRepository interface {
	GetByEmail(email string, recovery *models.Recovery) error
	Create(recovery *models.Recovery) (models.Recovery, error)
	UpdateRecovery(id int, code string, attempts int, expiresAt time.Time, expired bool) error
	MarkAsExpired(id int) error
	IncrementAttempts(id int) error
	GetByUserID(userID string, recovery *models.Recovery) error
	GetByCode(code string, recovery *models.Recovery) error
	ClearByID(id int) error
	ClearByUserID(userID string) error
	ClearByEmail(email string) error
}
