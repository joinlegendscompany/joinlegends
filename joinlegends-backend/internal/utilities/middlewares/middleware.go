package middlewares

import (
	"encoding/json"
	"go-backend-stream/internal/infrastructure/database"
	sessionRepository "go-backend-stream/internal/infrastructure/repositories/session"
	userRepository "go-backend-stream/internal/infrastructure/repositories/user"
	"go-backend-stream/internal/models"
	"go-backend-stream/internal/utilities/date"
	"go-backend-stream/internal/utilities/jwt"
	"go-backend-stream/internal/utilities/logger"
	"net/http"
	"strings"

	"github.com/go-goe/goe"
	"github.com/gofiber/fiber/v2"
)

type SessionContext struct {
	User      *models.User `json:"user"`
	SessionId string       `json:"session_id"`
}

func JwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing authorization header",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token format",
		})
	}

	token := parts[1]

	sub, err := jwt.ParseJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired token",
		})
	}

	c.Locals("userId", sub)

	return c.Next()
}

func NewJwtSessionMiddleware(db *database.StreamDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		token := parts[1]

		sub, err := jwt.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		logger.Info.Println("check if session exist")

		sessionRepository := sessionRepository.NewSessionRepository(db)

		var session models.Session
		if err = sessionRepository.GetByID(sub, &session); err != nil && err != goe.ErrNotFound {
			logger.Error.Printf("internal server error when get session by id %s. error: %s", sub, err.Error())
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "erro interno no servidor ao tentar buscar a sessão",
				"error":   err.Error(),
			})
		}
		if err == goe.ErrNotFound {
			logger.Warn.Printf("session with id %s no found", sub)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "sessão não encontrada",
			})
		}
		if !date.IsNotExpired(session.ExpiresAt) {
			logger.Warn.Printf("session already expired at %s", session.ExpiresAt)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "sessão inválida ou expirada",
			})
		}

		logger.Info.Printf("get user from session with id: %s", session.UserID)
		userRepository := userRepository.NewUserRepository(db)

		var user models.User
		if err = userRepository.GetById(session.UserID, &user); err != nil {
			if err == goe.ErrNotFound {
				logger.Warn.Printf("user with id %s no found", session.UserID)
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"message": "usuário não encontrado",
				})
			}
			logger.Error.Printf("internal server error when get user by id %s. error: %s", session.UserID, err.Error())
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "erro interno no servidor ao tentar buscar o usuário",
				"error":   err.Error(),
			})
		}

		sessionContext := &SessionContext{
			User:      &user,
			SessionId: session.ID,
		}
		jsonSession, err := json.Marshal(sessionContext)
		if err != nil {
			logger.Error.Printf("internal server error when marshal session context. error: %s", err.Error())
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "erro interno no servidor ao tentar marshal a sessão",
				"error":   err.Error(),
			})
		}
		c.Locals("session-context", jsonSession)
		c.Locals("userId", user.ID)

		return c.Next()
	}
}
