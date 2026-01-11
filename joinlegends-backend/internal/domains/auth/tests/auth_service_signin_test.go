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

func TestAuthService_SignIn(t *testing.T) {
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

		foundUser := models.User{
			ID:       "user-123",
			Email:    dto.Email,
			Password: dto.Password,
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&foundUser, nil)

		user, token, err := service.SignIn(dto)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, foundUser.ID, user.ID)
		assert.NotEmpty(t, token)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("user_not_found", func(t *testing.T) {
		dto := auth.SignInDto{
			Email:    "notfound@example.com",
			Password: "password123",
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)

		user, token, err := service.SignIn(dto)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.Equal(t, goe.ErrNotFound, err)
	})

	t.Run("invalid_password", func(t *testing.T) {
		dto := auth.SignInDto{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		foundUser := models.User{
			ID:       "user-123",
			Email:    dto.Email,
			Password: "correctpassword",
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&foundUser, nil)

		user, token, err := service.SignIn(dto)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.Contains(t, err.Error(), "password is invalid")
	})

	t.Run("database_error", func(t *testing.T) {
		dto := auth.SignInDto{
			Email: "error@example.com",
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(nil, errors.New("db error"))

		user, token, err := service.SignIn(dto)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.Equal(t, "db error", err.Error())
	})
}
