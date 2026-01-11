package factory

import (
	"go-backend-stream/internal/domains/auth"
	"go-backend-stream/internal/infrastructure/database"
	recovery_repository "go-backend-stream/internal/infrastructure/repositories/recovery"
	session_repository "go-backend-stream/internal/infrastructure/repositories/session"
	user_repository "go-backend-stream/internal/infrastructure/repositories/user"
)

type AppControllers struct {
	AuthController *auth.AuthController
}

func NewAppControllers(db *database.StreamDB) *AppControllers {
	// Repositories
	userRepo := user_repository.NewUserRepository(db)
	sessionRepo := session_repository.NewSessionRepository(db)
	recoveryRepo := recovery_repository.NewRecoveryRepository(db)

	// Services
	mailService := &auth.DefaultMailService{}
	authService := auth.NewAuthService(userRepo, sessionRepo, recoveryRepo, mailService)

	// Controllers
	authController := auth.NewAuthController(authService)

	return &AppControllers{
		AuthController: authController,
	}
}
