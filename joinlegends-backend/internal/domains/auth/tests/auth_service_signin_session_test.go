package auth_test

import (
	"errors"
	"go-backend-stream/internal/domains/auth"
	"go-backend-stream/internal/models"
	"testing"

	"github.com/go-goe/goe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_SignInWithSession(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockSessionRepo := new(MockSessionRepository)
	mockRecoveryRepo := new(MockRecoveryRepository)
	mockMailService := new(MockMailService)

	service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

	t.Run("success", func(t *testing.T) {
		dto := auth.SignInDto{
			Email:    "test@example.com",
			Password: "password123",
		}
		userAgent := "Mozilla/5.0"
		ip := "127.0.0.1"

		foundUser := models.User{
			ID:       "user-123",
			Email:    dto.Email,
			Password: dto.Password,
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&foundUser, nil)

		createdSession := models.Session{ID: "session-123", UserID: foundUser.ID}
		mockSessionRepo.On("Create", mock.AnythingOfType("*models.Session")).Return(&createdSession, nil)

		user, token, session, err := service.SignInWithSession(dto, userAgent, ip)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.NotNil(t, session)
		assert.NotEmpty(t, token)

		mockUserRepo.AssertExpectations(t)
		mockSessionRepo.AssertExpectations(t)
	})

	t.Run("user_not_found", func(t *testing.T) {
		dto := auth.SignInDto{Email: "missing@example.com", Password: "pwd"}
		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)

		user, token, session, err := service.SignInWithSession(dto, "ua", "ip")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Nil(t, session)
		assert.Empty(t, token)
	})

	t.Run("session_creation_fails", func(t *testing.T) {
		dto := auth.SignInDto{Email: "test@example.com", Password: "password123"}
		foundUser := models.User{ID: "u1", Email: dto.Email, Password: dto.Password}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&foundUser, nil)
		mockSessionRepo.On("Create", mock.AnythingOfType("*models.Session")).Return(nil, errors.New("db fail"))

		user, token, session, err := service.SignInWithSession(dto, "ua", "ip")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Nil(t, session)
		assert.Empty(t, token)
	})
}
