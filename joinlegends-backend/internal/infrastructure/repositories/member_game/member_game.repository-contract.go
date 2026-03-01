package member_game_repository

import "go-backend-stream/internal/models"

type MemberGameRepository interface {
	Create(memberGame *models.MemberGame) (models.MemberGame, error)
	GetByMemberID(memberID string, memberGames *[]models.MemberGame) error
	GetByMemberIDAndGameID(memberID string, gameID string, memberGame *models.MemberGame) error
	DeleteByID(id string) error
}
