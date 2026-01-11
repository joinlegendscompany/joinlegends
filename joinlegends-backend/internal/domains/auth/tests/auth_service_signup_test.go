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

func TestAuthService_SignUp(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockSessionRepo := new(MockSessionRepository)
	mockRecoveryRepo := new(MockRecoveryRepository)
	mockMailService := new(MockMailService)

	service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

	t.Run("success", func(t *testing.T) {
		dto := auth.SignUpDto{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}

		// Mock GetByEmail - User not found (good for signup)
		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)

		// Mock Create
		createdUser := models.User{
			ID:    "user-123",
			Name:  dto.Name,
			Email: dto.Email,
			Role:  "USER",
		}
		mockUserRepo.On("Create", mock.AnythingOfType("*models.User")).Return(createdUser, nil)

		// Mock Mail
		mockMailService.On("SendCreateAccount", dto.Name, dto.Email).Return(nil)

		user, token, err := service.SignUp(dto)
		time.Sleep(10 * time.Millisecond) // Allow go routine to run

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, createdUser.ID, user.ID)
		assert.NotEmpty(t, token)

		mockUserRepo.AssertExpectations(t)
		mockMailService.AssertExpectations(t)
	})

	t.Run("user_already_exists", func(t *testing.T) {
		dto := auth.SignUpDto{
			Name:     "Test User",
			Email:    "existing@example.com",
			Password: "password123",
		}

		// Mock GetByEmail - User found
		existingUser := models.User{ID: "existing-123", Email: dto.Email}
		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&existingUser, nil)

		user, token, err := service.SignUp(dto)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.Contains(t, err.Error(), "already registered")

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("database_error_on_check", func(t *testing.T) {
		dto := auth.SignUpDto{
			Email: "error@example.com",
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, errors.New("db error"))

		user, token, err := service.SignUp(dto)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.Equal(t, "db error", err.Error())
	})
}
