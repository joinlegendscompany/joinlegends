package game_repository

import "go-backend-stream/internal/models"

type GameRepository interface {
	GetByID(id string, game *models.Game) error
	GetAll(games *[]models.Game) error
	GetByCategory(category string, games *[]models.Game) error
	Create(game *models.Game) (models.Game, error)
	Update(game *models.Game) error
	SoftDelete(id string) error
}
