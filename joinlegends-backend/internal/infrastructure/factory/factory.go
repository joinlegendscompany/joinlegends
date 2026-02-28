package factory

import (
	"go-backend-stream/internal/domains/auth"
	"go-backend-stream/internal/domains/organization"
	"go-backend-stream/internal/infrastructure/database"
	member_repository "go-backend-stream/internal/infrastructure/repositories/member"
	organization_repository "go-backend-stream/internal/infrastructure/repositories/organization"
	recovery_repository "go-backend-stream/internal/infrastructure/repositories/recovery"
	session_repository "go-backend-stream/internal/infrastructure/repositories/session"
	user_repository "go-backend-stream/internal/infrastructure/repositories/user"
)

type AppControllers struct {
	AuthController         *auth.AuthController
	OrganizationController *organization.OrganizationController
}

func NewAppControllers(db *database.StreamDB) *AppControllers {
	// Repositories
	userRepo := user_repository.NewUserRepository(db)
	sessionRepo := session_repository.NewSessionRepository(db)
	recoveryRepo := recovery_repository.NewRecoveryRepository(db)
	orgRepo := organization_repository.NewOrganizationRepository(db)
	memberRepo := member_repository.NewMemberRepository(db)

	// Services
	mailService := &auth.DefaultMailService{}
	authService := auth.NewAuthService(userRepo, sessionRepo, recoveryRepo, mailService)
	orgService := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

	// Controllers
	authController := auth.NewAuthController(authService)
	orgController := organization.NewOrganizationController(orgService)

	return &AppControllers{
		AuthController:         authController,
		OrganizationController: orgController,
	}
}
