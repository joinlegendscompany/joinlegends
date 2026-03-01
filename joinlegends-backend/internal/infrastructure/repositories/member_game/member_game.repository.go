package member_game_repository

import (
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/where"
)

type memberGameRepository struct {
	db *database.StreamDB
}

func NewMemberGameRepository(db *database.StreamDB) MemberGameRepository {
	return &memberGameRepository{db}
}

func (r *memberGameRepository) Create(memberGame *models.MemberGame) (models.MemberGame, error) {
	if memberGame.CreatedAt.IsZero() {
		memberGame.CreatedAt = time.Now()
	}
	if memberGame.UpdatedAt.IsZero() {
		memberGame.UpdatedAt = time.Now()
	}

	err := goe.Insert(r.db.MemberGames).One(memberGame)
	if err != nil {
		return models.MemberGame{}, err
	}

	return *memberGame, nil
}

func (r *memberGameRepository) GetByMemberID(memberID string, memberGames *[]models.MemberGame) error {
	result, err := goe.List(r.db.MemberGames).Where(
		where.Equals(&r.db.MemberGames.MemberID, memberID),
	).AsSlice()

	if err != nil {
		return err
	}
	*memberGames = result
	return nil
}

func (r *memberGameRepository) GetByMemberIDAndGameID(memberID string, gameID string, memberGame *models.MemberGame) error {
	result, err := goe.List(r.db.MemberGames).Where(
		where.And(
			where.Equals(&r.db.MemberGames.MemberID, memberID),
			where.Equals(&r.db.MemberGames.GameID, gameID),
		),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(result) == 0 {
		return goe.ErrNotFound
	}
	*memberGame = result[0]
	return nil
}

func (r *memberGameRepository) DeleteByID(id string) error {
	return goe.Delete(r.db.MemberGames).Where(
		where.Equals(&r.db.MemberGames.ID, id),
	)
}
