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

func TestAuthService_SessionManagement(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockSessionRepo := new(MockSessionRepository)
	mockRecoveryRepo := new(MockRecoveryRepository)
	mockMailService := new(MockMailService)

	service := auth.NewAuthService(mockUserRepo, mockSessionRepo, mockRecoveryRepo, mockMailService)

	t.Run("GetAllSessions_success", func(t *testing.T) {
		userID := "user-123"
		mockUserRepo.On("GetById", userID, mock.AnythingOfType("*models.User")).Return(&models.User{ID: userID}, nil)

		sessionsList := []models.Session{
			{ID: "s1", UserID: userID},
			{ID: "s2", UserID: userID},
		}
		mockSessionRepo.On("GetAllByUserID", userID, mock.AnythingOfType("*[]models.Session")).Return(sessionsList, nil)

		sessions, err := service.GetAllSessions(userID)

		assert.NoError(t, err)
		assert.Len(t, sessions, 2)
	})

	t.Run("GetAllSessions_user_not_found", func(t *testing.T) {
		userID := "missing-123"
		mockUserRepo.On("GetById", userID, mock.AnythingOfType("*models.User")).Return(nil, goe.ErrNotFound)

		sessions, err := service.GetAllSessions(userID)

		assert.Error(t, err)
		assert.Nil(t, sessions)
		assert.Equal(t, goe.ErrNotFound, err)
	})

	t.Run("DeleteSession_success", func(t *testing.T) {
		userID := "user-123"
		sessionID := "session-123"

		mockUserRepo.On("GetById", userID, mock.AnythingOfType("*models.User")).Return(&models.User{ID: userID}, nil)
		mockSessionRepo.On("GetByID", sessionID, mock.AnythingOfType("*models.Session")).Return(&models.Session{ID: sessionID, UserID: userID}, nil)
		mockSessionRepo.On("DeleteByID", sessionID).Return(nil)

		err := service.DeleteSession(userID, sessionID)

		assert.NoError(t, err)
	})

	t.Run("DeleteSession_not_owner", func(t *testing.T) {
		userID := "user-123"
		sessionID := "session-456"

		mockUserRepo.On("GetById", userID, mock.AnythingOfType("*models.User")).Return(&models.User{ID: userID}, nil)
		mockSessionRepo.On("GetByID", sessionID, mock.AnythingOfType("*models.Session")).Return(&models.Session{ID: sessionID, UserID: "other-user"}, nil)

		err := service.DeleteSession(userID, sessionID)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "belong another user")
	})

	t.Run("DeleteSession_session_not_found", func(t *testing.T) {
		userID := "user-123"
		sessionID := "missing"

		mockUserRepo.On("GetById", userID, mock.AnythingOfType("*models.User")).Return(&models.User{ID: userID}, nil)
		mockSessionRepo.On("GetByID", sessionID, mock.AnythingOfType("*models.Session")).Return(nil, errors.New("not found"))

		err := service.DeleteSession(userID, sessionID)

		assert.Error(t, err)
	})
}
