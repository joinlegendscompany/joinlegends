package auth

import (
	"fmt"
	recovery_repository "go-backend-stream/internal/infrastructure/repositories/recovery"
	session_repository "go-backend-stream/internal/infrastructure/repositories/session"
	user_repository "go-backend-stream/internal/infrastructure/repositories/user"
	"go-backend-stream/internal/models"
	"go-backend-stream/internal/utilities/code"
	"go-backend-stream/internal/utilities/config"
	"go-backend-stream/internal/utilities/date"
	"go-backend-stream/internal/utilities/jwt"
	"go-backend-stream/internal/utilities/mail"
	"time"

	"github.com/go-goe/goe"
	"github.com/google/uuid"
)

type MailService interface {
	SendCreateAccount(name, email string) error
	SendRecoveryRequestEmail(name, email, recoveryUrl string) error
	SendConfirmationChangePassword(name, email, url string) error
}

// DefaultMailService implements MailService using the mail package
type DefaultMailService struct{}

func (s *DefaultMailService) SendCreateAccount(name, email string) error {
	return mail.SendCreateAccount(name, email)
}

func (s *DefaultMailService) SendRecoveryRequestEmail(name, email, recoveryUrl string) error {
	return mail.SendRecoveryRequestEmail(name, email, recoveryUrl)
}

func (s *DefaultMailService) SendConfirmationChangePassword(name, email, url string) error {
	return mail.SendConfirmationChangePassword(name, email, url)
}

// AuthService Interfaces and Implementation

type AuthService interface {
	SignUp(dto SignUpDto) (*models.User, string, error)
	SignIn(dto SignInDto) (*models.User, string, error)
	SignUpWithSession(dto SignUpDto, userAgent, ip string) (*models.User, string, *models.Session, error)
	SignInWithSession(dto SignInDto, userAgent, ip string) (*models.User, string, *models.Session, error)
	RequestRecovery(email string) (string, error)
	ChangePassword(dto ChangePasswordRequestRecoveryDto) error
	GetAllSessions(userId string) ([]models.Session, error)
	DeleteSession(userId, sessionId string) error
}

type authService struct {
	userRepo     user_repository.UserRepository
	sessionRepo  session_repository.SessionRepository
	recoveryRepo recovery_repository.RecoveryRepository
	mailService  MailService
}

func NewAuthService(
	userRepo user_repository.UserRepository,
	sessionRepo session_repository.SessionRepository,
	recoveryRepo recovery_repository.RecoveryRepository,
	mailService MailService,
) AuthService {
	return &authService{
		userRepo:     userRepo,
		sessionRepo:  sessionRepo,
		recoveryRepo: recoveryRepo,
		mailService:  mailService,
	}
}

func (s *authService) SignUp(dto SignUpDto) (*models.User, string, error) {
	var user models.User
	err := s.userRepo.GetByEmail(dto.Email, &user)
	if err == nil {
		return nil, "", fmt.Errorf("user with email %s already registered", dto.Email)
	}
	if err != goe.ErrNotFound {
		return nil, "", err
	}

	role := "USER"
	if isRoot(dto.Email) {
		role = "ROOT"
	}

	user = models.User{
		ID:        uuid.New().String(),
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      role,
	}

	userCreated, err := s.userRepo.Create(&user)
	if err != nil {
		return nil, "", err
	}

	token, err := jwt.GenerateJwt(userCreated.ID)
	if err != nil {
		return nil, "", err
	}

	go func() {
		_ = s.mailService.SendCreateAccount(user.Name, user.Email)
	}()

	return &userCreated, token, nil
}

func (s *authService) SignIn(dto SignInDto) (*models.User, string, error) {
	var user models.User
	if err := s.userRepo.GetByEmail(dto.Email, &user); err != nil {
		return nil, "", err
	}

	if user.Password != dto.Password {
		return nil, "", fmt.Errorf("password is invalid")
	}

	token, err := jwt.GenerateJwt(user.ID)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (s *authService) SignUpWithSession(dto SignUpDto, userAgent, ip string) (*models.User, string, *models.Session, error) {
	var user models.User
	err := s.userRepo.GetByEmail(dto.Email, &user)
	if err == nil {
		return nil, "", nil, fmt.Errorf("user with email %s already registered", dto.Email)
	}
	if err != goe.ErrNotFound {
		return nil, "", nil, err
	}

	role := "USER"
	if isRoot(dto.Email) {
		role = "ROOT"
	}

	user = models.User{
		ID:        uuid.New().String(),
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      role,
	}

	userCreated, err := s.userRepo.Create(&user)
	if err != nil {
		return nil, "", nil, err
	}

	sessionExpiresAt, err := date.GenerateFutureDate(1, "days")
	if err != nil {
		return nil, "", nil, err
	}

	session := models.Session{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Browser:   &userAgent,
		IP:        &ip,
		CreatedAt: time.Now(),
		ExpiresAt: sessionExpiresAt,
	}

	sessionCreated, err := s.sessionRepo.Create(&session)
	if err != nil {
		return nil, "", nil, err
	}

	token, err := jwt.GenerateJwt(session.ID)
	if err != nil {
		return nil, "", nil, err
	}

	go func() {
		_ = s.mailService.SendCreateAccount(user.Name, user.Email)
	}()

	return &userCreated, token, sessionCreated, nil
}

func (s *authService) SignInWithSession(dto SignInDto, userAgent, ip string) (*models.User, string, *models.Session, error) {
	var user models.User
	if err := s.userRepo.GetByEmail(dto.Email, &user); err != nil {
		return nil, "", nil, err
	}

	if user.Password != dto.Password {
		return nil, "", nil, fmt.Errorf("password is invalid")
	}

	sessionExpiresAt, err := date.GenerateFutureDate(1, "days")
	if err != nil {
		return nil, "", nil, err
	}

	session := models.Session{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Browser:   &userAgent,
		IP:        &ip,
		CreatedAt: time.Now(),
		ExpiresAt: sessionExpiresAt,
	}

	sessionCreated, err := s.sessionRepo.Create(&session)
	if err != nil {
		return nil, "", nil, err
	}

	token, err := jwt.GenerateJwt(session.ID)
	if err != nil {
		return nil, "", nil, err
	}

	return &user, token, sessionCreated, nil
}

func (s *authService) RequestRecovery(email string) (string, error) {
	var user models.User
	if err := s.userRepo.GetByEmail(email, &user); err != nil {
		return "", err
	}

	recoveryCode, err := code.GenerateRecoveryCode(120)
	if err != nil {
		return "", err
	}

	expiresAt, err := date.GenerateFutureDate(5, "minutes")
	if err != nil {
		return "", err
	}

	var recovery models.Recovery
	err = s.recoveryRepo.GetByEmail(email, &recovery)
	if err != nil && err != goe.ErrNotFound {
		return "", err
	}

	if err == goe.ErrNotFound {
		_, err := s.recoveryRepo.Create(&models.Recovery{
			ID:        0, // DB handles ID
			UserID:    user.ID,
			Email:     user.Email,
			Code:      recoveryCode,
			Attempts:  0,
			ExpiresAt: expiresAt,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Expired:   false,
		})
		if err != nil {
			return "", err
		}
	} else {
		if err = s.recoveryRepo.UpdateRecovery(recovery.ID, recoveryCode, 0, expiresAt, false); err != nil {
			return "", err
		}
	}

	recoveryUrl := fmt.Sprintf("https://dominio-do-front/recuperar/tocar-senha?code=%s", recoveryCode)

	go func() {
		_ = s.mailService.SendRecoveryRequestEmail(user.Name, user.Email, recoveryUrl)
	}()

	return fmt.Sprintf("o codigo de recuperacao foi enviado no email %s com sucesso", email), nil
}

func (s *authService) ChangePassword(dto ChangePasswordRequestRecoveryDto) error {
	var user models.User
	if err := s.userRepo.GetByEmail(dto.Email, &user); err != nil {
		return err
	}

	var recovery models.Recovery
	if err := s.recoveryRepo.GetByEmail(dto.Email, &recovery); err != nil {
		return err
	}

	if recovery.Expired {
		return fmt.Errorf("recovery code already expires")
	}

	if recovery.Attempts > 10 {
		_ = s.recoveryRepo.MarkAsExpired(recovery.ID)
		return fmt.Errorf("number of attempts exceed")
	}

	if !date.IsNotExpired(recovery.ExpiresAt) {
		_ = s.recoveryRepo.MarkAsExpired(recovery.ID)
		return fmt.Errorf("recovery code already expires")
	}

	if dto.Code != recovery.Code {
		_ = s.recoveryRepo.IncrementAttempts(recovery.ID)
		return fmt.Errorf("invalid code, remaining attempts: %d", 10-recovery.Attempts)
	}

	if err := s.userRepo.UpdatePassword(user.ID, dto.NewPassword); err != nil {
		return err
	}

	if err := s.sessionRepo.DeleteByUserId(user.ID); err != nil {
		return err
	}

	if err := s.recoveryRepo.MarkAsExpired(recovery.ID); err != nil {
		return err
	}

	go func() {
		_ = s.mailService.SendConfirmationChangePassword(user.Name, user.Email, "http://dominio-do-front-que-tem-que-terlogi")
	}()

	return nil
}

func (s *authService) GetAllSessions(userId string) ([]models.Session, error) {
	var user models.User
	if err := s.userRepo.GetById(userId, &user); err != nil {
		return nil, err
	}

	var sessions []models.Session
	if err := s.sessionRepo.GetAllByUserID(userId, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *authService) DeleteSession(userId, sessionId string) error {
	var user models.User
	if err := s.userRepo.GetById(userId, &user); err != nil {
		return err
	}

	var session models.Session
	if err := s.sessionRepo.GetByID(sessionId, &session); err != nil {
		return err
	}

	if session.UserID != userId {
		return fmt.Errorf("the session with id %s belong another user", sessionId)
	}

	if err := s.sessionRepo.DeleteByID(sessionId); err != nil {
		return err
	}

	return nil
}

func isRoot(email string) bool {
	return email == config.ROOT_EMAIL
}
