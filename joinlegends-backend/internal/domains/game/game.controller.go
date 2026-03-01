package game

import (
	"fmt"
	"net/http"

	"go-backend-stream/internal/utilities/logger"

	"github.com/gofiber/fiber/v2"
)

type GameController struct {
	service GameService
}

func NewGameController(service GameService) *GameController {
	return &GameController{service: service}
}

func (c *GameController) CreateGameController(ctx *fiber.Ctx) error {
	var body CreateGameDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("creating game %s", body.Name)

	game, err := c.service.CreateGame(CreateGameServiceProps{
		Name:        body.Name,
		BannerID:    body.BannerID,
		Category:    body.Category,
		Description: body.Description,
		Developer:   body.Developer,
		Publisher:   body.Publisher,
		ReleaseYear: body.ReleaseYear,
	})
	if err != nil {
		logger.Error.Printf("error when create game: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when create a new game",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "jogo cadastrado com sucesso",
		"game":    game,
	})
}

func (c *GameController) GetAllGamesController(ctx *fiber.Ctx) error {
	category := ctx.Query("category")

	if category != "" {
		games, err := c.service.GetGamesByCategory(category)
		if err != nil {
			logger.Error.Printf("error when get games by category: %s", err.Error())
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "internal server error when get games by category",
				"error":   err.Error(),
			})
		}
		return ctx.JSON(fiber.Map{
			"games": games,
		})
	}

	games, err := c.service.GetAllGames()
	if err != nil {
		logger.Error.Printf("error when get all games: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when get all games",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"games": games,
	})
}

func (c *GameController) GetGameController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	game, err := c.service.GetGameByID(id)
	if err != nil {
		if err.Error() == fmt.Sprintf("game with id %s not found", id) {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error when get game by id: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when get game",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"game": game,
	})
}

func (c *GameController) UpdateGameController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var body UpdateGameDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	game, err := c.service.UpdateGame(UpdateGameServiceProps{
		ID:          id,
		Name:        body.Name,
		BannerID:    body.BannerID,
		Category:    body.Category,
		Description: body.Description,
		Developer:   body.Developer,
		Publisher:   body.Publisher,
		ReleaseYear: body.ReleaseYear,
		IsActive:    body.IsActive,
	})
	if err != nil {
		if err.Error() == fmt.Sprintf("game with id %s not found", id) {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error when update game %s: %s", id, err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when update game",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "jogo atualizado com sucesso",
		"game":    game,
	})
}

func (c *GameController) DeleteGameController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.service.DeleteGame(id); err != nil {
		if err.Error() == fmt.Sprintf("game with id %s not found", id) {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error when delete game %s: %s", id, err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when delete game",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "jogo removido com sucesso",
	})
}
