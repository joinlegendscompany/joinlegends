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

func TestAuthService_RequestRecovery(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockSessionRepo := new(MockSessionRepository)
	mockRecoveryRepo := new(MockRecoveryRepository)
	mockMailService := new(MockMailService)

	service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

	t.Run("success_new_recovery", func(t *testing.T) {
		email := "test@example.com"
		user := models.User{ID: "user-123", Email: email, Name: "Test User"}

		mockUserRepo.On("GetByEmail", email, mock.AnythingOfType("*models.User")).Return(&user, nil)
		mockRecoveryRepo.On("GetByEmail", email, mock.AnythingOfType("*models.Recovery")).Return(nil, goe.ErrNotFound)

		mockRecoveryRepo.On("Create", mock.AnythingOfType("*models.Recovery")).Return(models.Recovery{}, nil)
		mockMailService.On("SendRecoveryRequestEmail", user.Name, user.Email, mock.Anything).Return(nil)

		msg, err := service.RequestRecovery(email)
		time.Sleep(10 * time.Millisecond)

		assert.NoError(t, err)
		assert.Contains(t, msg, "enviado no email")

		mockUserRepo.AssertExpectations(t)
		mockRecoveryRepo.AssertExpectations(t)
		mockMailService.AssertExpectations(t)
	})

	t.Run("success_update_existing_recovery", func(t *testing.T) {
		email := "test@example.com"
		user := models.User{ID: "user-123", Email: email, Name: "Test User"}
		existingRecovery := models.Recovery{ID: 1, UserID: user.ID, Email: email}

		mockUserRepo.On("GetByEmail", email, mock.AnythingOfType("*models.User")).Return(&user, nil)
		mockRecoveryRepo.On("GetByEmail", email, mock.AnythingOfType("*models.Recovery")).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Recovery)
			*arg = existingRecovery
		}).Return(nil, nil)

		mockRecoveryRepo.On("UpdateRecovery", existingRecovery.ID, mock.Anything, 0, mock.AnythingOfType("time.Time"), false).Return(nil)
		mockMailService.On("SendRecoveryRequestEmail", user.Name, user.Email, mock.Anything).Return(nil)

		msg, err := service.RequestRecovery(email)
		time.Sleep(10 * time.Millisecond)

		assert.NoError(t, err)
		assert.Contains(t, msg, "enviado no email")
	})

	t.Run("user_not_found", func(t *testing.T) {
		email := "missing@example.com"
		mockUserRepo.On("GetByEmail", email, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)

		msg, err := service.RequestRecovery(email)

		assert.Error(t, err)
		assert.Empty(t, msg)
		assert.Equal(t, goe.ErrNotFound, err)
	})

	t.Run("recovery_db_error", func(t *testing.T) {
		email := "error@example.com"
		user := models.User{ID: "u1", Email: email}

		mockUserRepo.On("GetByEmail", email, mock.AnythingOfType("*models.User")).Return(&user, nil)
		mockRecoveryRepo.On("GetByEmail", email, mock.AnythingOfType("*models.Recovery")).Return(nil, errors.New("db error"))

		msg, err := service.RequestRecovery(email)

		assert.Error(t, err)
		assert.Empty(t, msg)
	})
}

func TestAuthService_ChangePassword(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockSessionRepo := new(MockSessionRepository)
	mockRecoveryRepo := new(MockRecoveryRepository)
	mockMailService := new(MockMailService)

	service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

	t.Run("success", func(t *testing.T) {
		dto := auth.ChangePasswordRequestRecoveryDto{
			Email:       "test@example.com",
			NewPassword: "newpassword123",
			Code:        "123456",
		}

		user := models.User{ID: "user-123", Email: dto.Email, Name: "Test User"}
		recovery := models.Recovery{
			ID:        1,
			Email:     dto.Email,
			Code:      dto.Code,
			Attempts:  0,
			ExpiresAt: time.Now().Add(1 * time.Hour),
			Expired:   false,
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&user, nil)
		mockRecoveryRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.Recovery")).Return(&recovery, nil)
		mockUserRepo.On("UpdatePassword", user.ID, dto.NewPassword).Return(nil)
		mockSessionRepo.On("DeleteByUserId", user.ID).Return(nil)
		mockRecoveryRepo.On("MarkAsExpired", recovery.ID).Return(nil)
		mockMailService.On("SendConfirmationChangePassword", user.Name, user.Email, mock.Anything).Return(nil)

		// Allow IncrementAttempts to be called (even though it shouldn't) to debug flow
		// mockRecoveryRepo.On("IncrementAttempts", mock.Anything).Return(nil)
		
		err := service.ChangePassword(dto)
		time.Sleep(10 * time.Millisecond)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
		mockRecoveryRepo.AssertExpectations(t)
	})

	t.Run("expired_code", func(t *testing.T) {
		dto := auth.ChangePasswordRequestRecoveryDto{Email: "test@example.com", Code: "123"}
		// Set ExpiresAt to future to avoid that check failing (we want to test Expired flag)
		// Set Code to match dto to avoid IncrementAttempts if Expired check fails
		recovery := models.Recovery{
			Expired:   true, 
			ExpiresAt: time.Now().Add(1 * time.Hour),
			Code:      "123",
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&models.User{}, nil)
		// Use Run to ensure side effect happens if standard Return matching is failing
		mockRecoveryRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.Recovery")).Return(&recovery, nil)

		// Allow IncrementAttempts to be called (even though it shouldn't) to debug flow
		// mockRecoveryRepo.On("IncrementAttempts", mock.Anything).Return(nil)
		
		err := service.ChangePassword(dto)

		assert.Error(t, err)
		assert.Equal(t, "recovery code already expires", err.Error())
	})

	t.Run("invalid_code", func(t *testing.T) {
		// Fresh mocks for this run to avoid pollution
		mockUserRepo := new(MockUserRepository)
		mockRecoveryRepo := new(MockRecoveryRepository)
		mockSessionRepo := new(MockSessionRepository)
		mockMailService := new(MockMailService)
		service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

		dto := auth.ChangePasswordRequestRecoveryDto{Email: "test@example.com", Code: "WRONG"}
		recovery := models.Recovery{
			ID:        1,
			Code:      "RIGHT",
			ExpiresAt: time.Now().Add(1 * time.Hour),
			Attempts:  0,
		}

		mockUserRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.User")).Return(&models.User{}, nil)
		// Use Run to ensure side effect happens if standard Return matching is failing
		mockRecoveryRepo.On("GetByEmail", dto.Email, mock.AnythingOfType("*models.Recovery")).Return(&recovery, nil)
		mockRecoveryRepo.On("IncrementAttempts", recovery.ID).Return(nil)

		// Allow IncrementAttempts to be called (even though it shouldn't) to debug flow
		// mockRecoveryRepo.On("IncrementAttempts", mock.Anything).Return(nil)
		
		err := service.ChangePassword(dto)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid code")
	})
}
