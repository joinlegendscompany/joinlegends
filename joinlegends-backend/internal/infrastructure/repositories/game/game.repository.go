package game_repository

import (
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/update"
	"github.com/go-goe/goe/query/where"
)

type gameRepository struct {
	db *database.StreamDB
}

func NewGameRepository(db *database.StreamDB) GameRepository {
	return &gameRepository{db}
}

func (r *gameRepository) GetByID(id string, game *models.Game) error {
	games, err := goe.List(r.db.Games).Where(
		where.Equals(&r.db.Games.ID, id),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(games) == 0 {
		return goe.ErrNotFound
	}
	*game = games[0]
	return nil
}

func (r *gameRepository) GetAll(games *[]models.Game) error {
	result, err := goe.List(r.db.Games).Where(
		where.Equals(&r.db.Games.DeletedAt, (*time.Time)(nil)),
	).AsSlice()

	if err != nil {
		return err
	}
	*games = result
	return nil
}

func (r *gameRepository) GetByCategory(category string, games *[]models.Game) error {
	result, err := goe.List(r.db.Games).Where(
		where.And(
			where.Equals(&r.db.Games.Category, category),
			where.Equals(&r.db.Games.DeletedAt, (*time.Time)(nil)),
		),
	).AsSlice()

	if err != nil {
		return err
	}
	*games = result
	return nil
}

func (r *gameRepository) Create(game *models.Game) (models.Game, error) {
	if game.CreatedAt.IsZero() {
		game.CreatedAt = time.Now()
	}

	err := goe.Insert(r.db.Games).One(game)
	if err != nil {
		return models.Game{}, err
	}

	return *game, nil
}

func (r *gameRepository) Update(game *models.Game) error {
	return goe.Update(r.db.Games).
		Sets(
			update.Set(&r.db.Games.Name, game.Name),
			update.Set(&r.db.Games.BannerID, game.BannerID),
			update.Set(&r.db.Games.Category, game.Category),
			update.Set(&r.db.Games.Description, game.Description),
			update.Set(&r.db.Games.Developer, game.Developer),
			update.Set(&r.db.Games.Publisher, game.Publisher),
			update.Set(&r.db.Games.ReleaseYear, game.ReleaseYear),
			update.Set(&r.db.Games.IsActive, game.IsActive),
			update.Set(&r.db.Games.UpdatedAt, time.Now()),
		).
		Where(where.Equals(&r.db.Games.ID, game.ID))
}

func (r *gameRepository) SoftDelete(id string) error {
	now := time.Now()
	return goe.Update(r.db.Games).
		Sets(
			update.Set(&r.db.Games.DeletedAt, &now),
			update.Set(&r.db.Games.UpdatedAt, now),
		).
		Where(where.Equals(&r.db.Games.ID, id))
}
