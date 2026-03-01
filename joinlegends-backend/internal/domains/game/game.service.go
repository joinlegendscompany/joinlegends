package game

import (
	"fmt"
	game_repository "go-backend-stream/internal/infrastructure/repositories/game"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(props CreateGameServiceProps) (*models.Game, error)
	GetAllGames() ([]models.Game, error)
	GetGameByID(id string) (*models.Game, error)
	GetGamesByCategory(category string) ([]models.Game, error)
	UpdateGame(props UpdateGameServiceProps) (*models.Game, error)
	DeleteGame(id string) error
}

type gameService struct {
	gameRepo game_repository.GameRepository
}

func NewGameService(gameRepo game_repository.GameRepository) GameService {
	return &gameService{gameRepo: gameRepo}
}

func (s *gameService) CreateGame(props CreateGameServiceProps) (*models.Game, error) {
	game := models.Game{
		ID:          uuid.New().String(),
		Name:        props.Name,
		BannerID:    props.BannerID,
		Category:    props.Category,
		Description: props.Description,
		Developer:   props.Developer,
		Publisher:   props.Publisher,
		ReleaseYear: props.ReleaseYear,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := s.gameRepo.Create(&game)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *gameService) GetAllGames() ([]models.Game, error) {
	var games []models.Game
	if err := s.gameRepo.GetAll(&games); err != nil {
		return nil, err
	}
	return games, nil
}

func (s *gameService) GetGameByID(id string) (*models.Game, error) {
	var game models.Game
	if err := s.gameRepo.GetByID(id, &game); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("game with id %s not found", id)
		}
		return nil, err
	}
	return &game, nil
}

func (s *gameService) GetGamesByCategory(category string) ([]models.Game, error) {
	var games []models.Game
	if err := s.gameRepo.GetByCategory(category, &games); err != nil {
		return nil, err
	}
	return games, nil
}

func (s *gameService) UpdateGame(props UpdateGameServiceProps) (*models.Game, error) {
	var game models.Game
	if err := s.gameRepo.GetByID(props.ID, &game); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("game with id %s not found", props.ID)
		}
		return nil, err
	}

	if props.Name != nil {
		game.Name = *props.Name
	}
	if props.BannerID != nil {
		game.BannerID = props.BannerID
	}
	if props.Category != nil {
		game.Category = *props.Category
	}
	if props.Description != nil {
		game.Description = *props.Description
	}
	if props.Developer != nil {
		game.Developer = *props.Developer
	}
	if props.Publisher != nil {
		game.Publisher = props.Publisher
	}
	if props.ReleaseYear != nil {
		game.ReleaseYear = *props.ReleaseYear
	}
	if props.IsActive != nil {
		game.IsActive = *props.IsActive
	}

	if err := s.gameRepo.Update(&game); err != nil {
		return nil, err
	}

	return &game, nil
}

func (s *gameService) DeleteGame(id string) error {
	var game models.Game
	if err := s.gameRepo.GetByID(id, &game); err != nil {
		if err == goe.ErrNotFound {
			return fmt.Errorf("game with id %s not found", id)
		}
		return err
	}

	return s.gameRepo.SoftDelete(id)
}
