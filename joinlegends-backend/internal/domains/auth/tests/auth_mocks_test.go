package auth_test

import (
	recovery_repository "go-backend-stream/internal/infrastructure/repositories/recovery"
	session_repository "go-backend-stream/internal/infrastructure/repositories/session"
	user_repository "go-backend-stream/internal/infrastructure/repositories/user"
	"go-backend-stream/internal/models"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmail(email string, user *models.User) error {
	args := m.Called(email, user)
	if u, ok := args.Get(0).(*models.User); ok && u != nil {
		*user = *u
	}
	return args.Error(1)
}

func (m *MockUserRepository) Create(user *models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) GetById(id string, user *models.User) error {
	args := m.Called(id, user)
	if u, ok := args.Get(0).(*models.User); ok && u != nil {
		*user = *u
	}
	return args.Error(1)
}

func (m *MockUserRepository) UpdatePassword(userID string, newPassword string) error {
	args := m.Called(userID, newPassword)
	return args.Error(0)
}

var _ user_repository.UserRepository = (*MockUserRepository)(nil)

// MockSessionRepository
type MockSessionRepository struct {
	mock.Mock
}

func (m *MockSessionRepository) Create(session *models.Session) (*models.Session, error) {
	args := m.Called(session)
	if s, ok := args.Get(0).(*models.Session); ok {
		return s, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSessionRepository) GetAllByUserID(userID string, sessions *[]models.Session) error {
	args := m.Called(userID, sessions)
	if s, ok := args.Get(0).([]models.Session); ok {
		*sessions = s
	}
	return args.Error(1)
}

func (m *MockSessionRepository) GetByID(id string, session *models.Session) error {
	args := m.Called(id, session)
	if s, ok := args.Get(0).(*models.Session); ok && s != nil {
		*session = *s
	}
	return args.Error(1)
}

func (m *MockSessionRepository) DeleteByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSessionRepository) DeleteByUserId(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

var _ session_repository.SessionRepository = (*MockSessionRepository)(nil)

// MockRecoveryRepository
type MockRecoveryRepository struct {
	mock.Mock
}

func (m *MockRecoveryRepository) GetByEmail(email string, recovery *models.Recovery) error {
	args := m.Called(email, recovery)
	if r, ok := args.Get(0).(*models.Recovery); ok && r != nil {
		*recovery = *r
	}
	return args.Error(1)
}

func (m *MockRecoveryRepository) Create(recovery *models.Recovery) (models.Recovery, error) {
	args := m.Called(recovery)
	return args.Get(0).(models.Recovery), args.Error(1)
}

func (m *MockRecoveryRepository) UpdateRecovery(id int, code string, attempts int, expiresAt time.Time, expired bool) error {
	args := m.Called(id, code, attempts, expiresAt, expired)
	return args.Error(0)
}

func (m *MockRecoveryRepository) MarkAsExpired(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRecoveryRepository) IncrementAttempts(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRecoveryRepository) GetByUserID(userID string, recovery *models.Recovery) error {
	args := m.Called(userID, recovery)
	if r, ok := args.Get(0).(*models.Recovery); ok && r != nil {
		*recovery = *r
	}
	return args.Error(1)
}

func (m *MockRecoveryRepository) GetByCode(code string, recovery *models.Recovery) error {
	args := m.Called(code, recovery)
	if r, ok := args.Get(0).(*models.Recovery); ok && r != nil {
		*recovery = *r
	}
	return args.Error(1)
}

func (m *MockRecoveryRepository) ClearByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRecoveryRepository) ClearByUserID(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockRecoveryRepository) ClearByEmail(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

var _ recovery_repository.RecoveryRepository = (*MockRecoveryRepository)(nil)

// MockMailService
type MockMailService struct {
	mock.Mock
}

func (m *MockMailService) SendCreateAccount(name, email string) error {
	args := m.Called(name, email)
	return args.Error(0)
}

func (m *MockMailService) SendRecoveryRequestEmail(name, email, recoveryUrl string) error {
	args := m.Called(name, email, recoveryUrl)
	return args.Error(0)
}

func (m *MockMailService) SendConfirmationChangePassword(name, email, url string) error {
	args := m.Called(name, email, url)
	return args.Error(0)
}
