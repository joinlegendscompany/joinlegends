package auth_test

import (
	"errors"
	"go-backend-stream/internal/domains/auth"
	"go-backend-stream/internal/models"
	"testing"
	"time"

	"github.com/go-goe/goe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_SignUpWithSession(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockSessionRepo := new(MockSessionRepository)
	mockRecoveryRepo := new(MockRecoveryRepository)
	mockMailService := new(MockMailService)

	service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

	t.Run("success", func(t *testing.T) {
		dto := auth.SignUpDto{
			Name:     "Test User",
			Email:    "session@example.com",
			Password: "password123",
		}
		userAgent := "Mozilla/5.0"
		ip := "127.0.0.1"

		// Mock User Check
		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)

		// Mock User Create
		createdUser := models.User{ID: "user-123", Email: dto.Email}
		mockUserRepo.On("Create", mock.AnythingOfType("*models.User")).Return(createdUser, nil)

		// Mock Session Create
		createdSession := models.Session{ID: "session-123", UserID: createdUser.ID}
		mockSessionRepo.On("Create", mock.AnythingOfType("*models.Session")).Return(&createdSession, nil)

		// Mock Email
		mockMailService.On("SendCreateAccount", dto.Name, dto.Email).Return(nil)

		user, token, session, err := service.SignUpWithSession(dto, userAgent, ip)
		time.Sleep(10 * time.Millisecond) // Allow go routine to run

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.NotNil(t, session)
		assert.NotEmpty(t, token)
		assert.Equal(t, createdUser.ID, user.ID)
		assert.Equal(t, createdSession.ID, session.ID)

		mockUserRepo.AssertExpectations(t)
		mockSessionRepo.AssertExpectations(t)
		mockMailService.AssertExpectations(t)
	})

	t.Run("session_create_error", func(t *testing.T) {
		dto := auth.SignUpDto{
			Name:     "Test User",
			Email:    "failsession@example.com",
			Password: "password123",
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)
		mockUserRepo.On("Create", mock.AnythingOfType("*models.User")).Return(models.User{ID: "user-123", Name: dto.Name, Email: dto.Email}, nil)
		mockSessionRepo.On("Create", mock.AnythingOfType("*models.Session")).Return(nil, errors.New("session db error"))

		user, token, session, err := service.SignUpWithSession(dto, "ua", "ip")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Nil(t, session)
		assert.Empty(t, token)
		assert.Equal(t, "session db error", err.Error())
	})
}
