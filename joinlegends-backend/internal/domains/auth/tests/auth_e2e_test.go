package auth_test

import (
	"bytes"
	"encoding/json"
	"go-backend-stream/internal/domains/auth"
	"go-backend-stream/internal/models"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-goe/goe"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ─── MockAuthService ──────────────────────────────────────────────────────────

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) SignUp(dto auth.SignUpDto) (*models.User, string, error) {
	args := m.Called(dto)
	if u, ok := args.Get(0).(*models.User); ok {
		return u, args.String(1), args.Error(2)
	}
	return nil, args.String(1), args.Error(2)
}

func (m *MockAuthService) SignIn(dto auth.SignInDto) (*models.User, string, error) {
	args := m.Called(dto)
	if u, ok := args.Get(0).(*models.User); ok {
		return u, args.String(1), args.Error(2)
	}
	return nil, args.String(1), args.Error(2)
}

func (m *MockAuthService) SignUpWithSession(dto auth.SignUpDto, userAgent, ip string) (*models.User, string, *models.Session, error) {
	args := m.Called(dto, userAgent, ip)
	u, _ := args.Get(0).(*models.User)
	s, _ := args.Get(2).(*models.Session)
	return u, args.String(1), s, args.Error(3)
}

func (m *MockAuthService) SignInWithSession(dto auth.SignInDto, userAgent, ip string) (*models.User, string, *models.Session, error) {
	args := m.Called(dto, userAgent, ip)
	u, _ := args.Get(0).(*models.User)
	s, _ := args.Get(2).(*models.Session)
	return u, args.String(1), s, args.Error(3)
}

func (m *MockAuthService) RequestRecovery(email string) (string, error) {
	args := m.Called(email)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) ChangePassword(dto auth.ChangePasswordRequestRecoveryDto) error {
	args := m.Called(dto)
	return args.Error(0)
}

func (m *MockAuthService) GetAllSessions(userId string) ([]models.Session, error) {
	args := m.Called(userId)
	if s, ok := args.Get(0).([]models.Session); ok {
		return s, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuthService) DeleteSession(userId, sessionId string) error {
	args := m.Called(userId, sessionId)
	return args.Error(0)
}

var _ auth.AuthService = (*MockAuthService)(nil)

// ─── helpers ──────────────────────────────────────────────────────────────────

func newAuthE2EApp(svc auth.AuthService) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}})

	ctrl := auth.NewAuthController(svc)

	// test middleware that injects userId for protected routes
	injectUser := func(userID string) fiber.Handler {
		return func(c *fiber.Ctx) error {
			c.Locals("userId", userID)
			return c.Next()
		}
	}

	v1 := app.Group("v1/auth")
	session := v1.Group("/session")
	session.Post("/sign-up", ctrl.SignUpController)
	session.Post("/sign-in", ctrl.SignInController)
	session.Get("/all", injectUser("user-123"), ctrl.GetAllSessionsController)
	session.Delete("/:sessionId", injectUser("user-123"), ctrl.DeleteSessionController)

	recovery := v1.Group("/recovery")
	recovery.Post("/request/:email", ctrl.RequestRecoveryController)
	recovery.Post("/change-password", ctrl.ChangePasswordController)

	return app
}

func jsonBody(t *testing.T, v any) *bytes.Buffer {
	t.Helper()
	b, err := json.Marshal(v)
	assert.NoError(t, err)
	return bytes.NewBuffer(b)
}

// ─── Sign Up ──────────────────────────────────────────────────────────────────

func TestAuthE2E_SignUp(t *testing.T) {
	t.Run("201_success", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.SignUpDto{Name: "Alice", Email: "alice@test.com", Password: "pass123"}
		user := &models.User{ID: "u1", Name: dto.Name, Email: dto.Email}

		svc.On("SignUp", dto).Return(user, "jwt-token", nil)

		req := httptest.NewRequest("POST", "/v1/auth/session/sign-up", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		svc.AssertExpectations(t)
	})

	t.Run("401_already_registered", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.SignUpDto{Name: "Bob", Email: "bob@test.com", Password: "pass"}

		svc.On("SignUp", dto).Return(nil, "", fiber.NewError(0, "user with email bob@test.com already registered"))

		req := httptest.NewRequest("POST", "/v1/auth/session/sign-up", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}

// ─── Sign In ──────────────────────────────────────────────────────────────────

func TestAuthE2E_SignIn(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.SignInDto{Email: "alice@test.com", Password: "pass123"}
		user := &models.User{ID: "u1", Email: dto.Email}

		svc.On("SignIn", dto).Return(user, "jwt-token", nil)

		req := httptest.NewRequest("POST", "/v1/auth/session/sign-in", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("401_invalid_password", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.SignInDto{Email: "alice@test.com", Password: "wrong"}

		svc.On("SignIn", dto).Return(nil, "", fiber.NewError(0, "password is invalid"))

		req := httptest.NewRequest("POST", "/v1/auth/session/sign-in", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("500_user_not_found", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.SignInDto{Email: "ghost@test.com", Password: "pass"}

		svc.On("SignIn", dto).Return(nil, "", goe.ErrNotFound)

		req := httptest.NewRequest("POST", "/v1/auth/session/sign-in", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

// ─── Get All Sessions ─────────────────────────────────────────────────────────

func TestAuthE2E_GetAllSessions(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		sessions := []models.Session{
			{ID: "s1", UserID: "user-123"},
			{ID: "s2", UserID: "user-123"},
		}

		body := map[string]string{"user_id": "user-123"}
		svc.On("GetAllSessions", "user-123").Return(sessions, nil)

		req := httptest.NewRequest("GET", "/v1/auth/session/all", jsonBody(t, body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

// ─── Delete Session ───────────────────────────────────────────────────────────

func TestAuthE2E_DeleteSession(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		svc.On("DeleteSession", "user-123", "session-abc").Return(nil)

		req := httptest.NewRequest("DELETE", "/v1/auth/session/session-abc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("406_not_owner_of_session", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		svc.On("DeleteSession", "user-123", "session-xyz").
			Return(fiber.NewError(0, "the session with id session-xyz belong another user"))

		req := httptest.NewRequest("DELETE", "/v1/auth/session/session-xyz", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotAcceptable, resp.StatusCode)
	})
}

// ─── Request Recovery ─────────────────────────────────────────────────────────

func TestAuthE2E_RequestRecovery(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		svc.On("RequestRecovery", "alice@test.com").
			Return("codigo enviado no email alice@test.com com sucesso", nil)

		req := httptest.NewRequest("POST", "/v1/auth/recovery/request/alice@test.com", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("404_email_not_found", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		svc.On("RequestRecovery", "ghost@test.com").Return("", goe.ErrNotFound)

		req := httptest.NewRequest("POST", "/v1/auth/recovery/request/ghost@test.com", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

// ─── Change Password ──────────────────────────────────────────────────────────

func TestAuthE2E_ChangePassword(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.ChangePasswordRequestRecoveryDto{
			Email:       "alice@test.com",
			NewPassword: "newpass123",
			Code:        "ABC123",
		}
		svc.On("ChangePassword", dto).Return(nil)

		req := httptest.NewRequest("POST", "/v1/auth/recovery/change-password", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("406_expired_code", func(t *testing.T) {
		svc := new(MockAuthService)
		app := newAuthE2EApp(svc)

		dto := auth.ChangePasswordRequestRecoveryDto{
			Email: "alice@test.com",
			Code:  "OLD",
		}
		svc.On("ChangePassword", dto).Return(fiber.NewError(0, "recovery code already expires"))

		req := httptest.NewRequest("POST", "/v1/auth/recovery/change-password", jsonBody(t, dto))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotAcceptable, resp.StatusCode)
	})
}

// ─── unused import guard ──────────────────────────────────────────────────────
var _ = time.Now
