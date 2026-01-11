package auth

import (
	"fmt"
	"go-backend-stream/internal/utilities/logger"
	"net/http"

	"github.com/go-goe/goe"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) SignUpController(ctx *fiber.Ctx) error {
	var body SignUpDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Creating user %s", body.Email)
	user, token, err := c.service.SignUp(body)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with email %s already registered", body.Email) {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error when create user: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when create a new user",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "usuário cadastrado com sucesso",
		"user":    user,
		"access_token": fiber.Map{
			"jwt":     token,
			"expires": "devo pegar das envs, dps eu faço isso",
		},
	})
}

func (c *AuthController) SignInController(ctx *fiber.Ctx) error {
	var body SignInDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Signing in user %s", body.Email)
	user, token, err := c.service.SignIn(body)
	if err != nil {
		if err == goe.ErrNotFound {
			logger.Warn.Printf("user with email %s unregistered", body.Email)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": fmt.Sprintf("user with email %s unregistered", body.Email),
				"error":   err.Error(),
			})
		}
		if err.Error() == "password is invalid" {
			logger.Warn.Printf("password of user %s is invalid", body.Email)
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "password is inválid",
			})
		}

		logger.Error.Printf("error when sign in: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "usuário logado com sucesso",
		"user":    user,
		"access_token": fiber.Map{
			"jwt":        token,
			"expires_in": "depois eu faço isso...",
		},
	})
}

func (c *AuthController) SignUpWithSessionController(ctx *fiber.Ctx) error {
	var body SignUpDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	browser := ctx.Get("User-Agent")
	ip := ctx.IP()

	logger.Info.Printf("Creating user with session %s", body.Email)
	user, token, session, err := c.service.SignUpWithSession(body, browser, ip)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with email %s already registered", body.Email) {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error when create user with session: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when create a new user",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "usuário cadastrado com sucesso",
		"user":    user,
		"access_token": fiber.Map{
			"jwt":     token,
			"expires": "devo pegar das envs, dps eu faço isso",
		},
		"session": session,
	})
}

func (c *AuthController) SignInWithSession(ctx *fiber.Ctx) error {
	var body SignInDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	browser := ctx.Get("User-Agent")
	ip := ctx.IP()

	logger.Info.Printf("Signing in user with session %s", body.Email)
	user, token, session, err := c.service.SignInWithSession(body, browser, ip)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": fmt.Sprintf("user with email %s unregistered", body.Email),
				"error":   err.Error(),
			})
		}
		if err.Error() == "password is invalid" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "password is inválid",
			})
		}

		logger.Error.Printf("error when sign in with session: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "usuário logado com sucesso",
		"user":    user,
		"access_token": fiber.Map{
			"jwt":        token,
			"expires_in": "depois eu faço isso...",
		},
		"session": session,
	})
}

func (c *AuthController) RequestRecoveryController(ctx *fiber.Ctx) error {
	email := ctx.Params("email")

	logger.Info.Printf("Requesting recovery for %s", email)
	msg, err := c.service.RequestRecovery(email)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("usuário com email %s não foi cadastrado", email),
			})
		}
		logger.Error.Printf("error request recovery: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error interno no servidor",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": msg,
	})
}

func (c *AuthController) ChangePasswordController(ctx *fiber.Ctx) error {
	var body ChangePasswordRequestRecoveryDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Changing password for %s", body.Email)
	err := c.service.ChangePassword(body)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "usuário ou dados de recuperação não encontrados",
				"error":   err.Error(),
			})
		}
		// Check for specific error messages (fragile but effective for now)
		if err.Error() == "recovery code already expires" || err.Error() == "number of attempts exceed" || fmt.Sprintf("invalid code, remaining attempts: %d", 10) == err.Error() /* simplified check */ {
			return ctx.Status(http.StatusNotAcceptable).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		logger.Error.Printf("error changing password: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error interno no servidor",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "senha alterada com sucesso",
	})
}

func (c *AuthController) GetAllSessionsController(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(string)

	sessions, err := c.service.GetAllSessions(userId)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "usuário não cadastrado",
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error interno no servidor",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "sessões encontradas com successo",
		"sessions": sessions,
	})
}

func (c *AuthController) DeleteSessionController(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(string)
	sessionId := ctx.Params("sessionId")

	err := c.service.DeleteSession(userId, sessionId)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": fmt.Sprintf("sessão com id: %s não encontrada", sessionId),
				"error":   err.Error(),
			})
		}
		if err.Error() == fmt.Sprintf("the session with id %s belong another user", sessionId) {
			return ctx.Status(http.StatusNotAcceptable).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "error interno no servidor",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "sessão deletada com sucesso",
	})
}
