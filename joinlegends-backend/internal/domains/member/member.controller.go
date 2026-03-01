package member

import (
	"fmt"
	"go-backend-stream/internal/utilities/logger"
	"net/http"

	"github.com/go-goe/goe"
	"github.com/gofiber/fiber/v2"
)

type MemberController struct {
	service MemberService
}

func NewMemberController(service MemberService) *MemberController {
	return &MemberController{service: service}
}

func (c *MemberController) GetMemberController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	member, err := c.service.GetMember(id)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("membro com id %s não encontrado", id),
			})
		}
		logger.Error.Printf("error getting member: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"member": member,
	})
}

func (c *MemberController) GetUserMembersController(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(string)

	members, err := c.service.GetUserMembers(userID)
	if err != nil {
		logger.Error.Printf("error getting user members: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"members": members,
	})
}

func (c *MemberController) AddGameToMemberController(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(string)
	memberID := ctx.Params("id")

	var body AddGameToMemberDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Adding game %s to member %s", body.GameID, memberID)
	memberGame, err := c.service.AddGame(AddGameToMemberServiceProps{
		MemberID: memberID,
		GameID:   body.GameID,
		UserID:   userID,
	})
	if err != nil {
		if err.Error() == "você só pode adicionar jogos ao seu próprio perfil de membro" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err.Error() == "jogo já associado a este membro" {
			return ctx.Status(http.StatusConflict).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error adding game to member: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message":     "jogo adicionado ao membro com sucesso",
		"member_game": memberGame,
	})
}

func (c *MemberController) RemoveGameFromMemberController(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userId").(string)
	memberID := ctx.Params("id")
	gameID := ctx.Params("gameId")

	logger.Info.Printf("Removing game %s from member %s", gameID, memberID)
	err := c.service.RemoveGame(memberID, userID, gameID)
	if err != nil {
		if err.Error() == "você só pode remover jogos do seu próprio perfil de membro" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err.Error() == "jogo não associado a este membro" {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error removing game from member: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "jogo removido do membro com sucesso",
	})
}
