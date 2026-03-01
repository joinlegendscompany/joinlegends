package member

import (
	"fmt"
	game_repository "go-backend-stream/internal/infrastructure/repositories/game"
	member_repository "go-backend-stream/internal/infrastructure/repositories/member"
	member_game_repository "go-backend-stream/internal/infrastructure/repositories/member_game"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/google/uuid"
)

type MemberService interface {
	GetMember(id string) (*models.Member, error)
	GetUserMembers(userID string) ([]models.Member, error)
	AddGame(props AddGameToMemberServiceProps) (*models.MemberGame, error)
	RemoveGame(memberID string, userID string, gameID string) error
}

type memberService struct {
	memberRepo     member_repository.MemberRepository
	memberGameRepo member_game_repository.MemberGameRepository
	gameRepo       game_repository.GameRepository
}

func NewMemberService(
	memberRepo member_repository.MemberRepository,
	memberGameRepo member_game_repository.MemberGameRepository,
	gameRepo game_repository.GameRepository,
) MemberService {
	return &memberService{
		memberRepo:     memberRepo,
		memberGameRepo: memberGameRepo,
		gameRepo:       gameRepo,
	}
}

func (s *memberService) GetMember(id string) (*models.Member, error) {
	var member models.Member
	if err := s.memberRepo.GetByID(id, &member); err != nil {
		return nil, err
	}

	var memberGames []models.MemberGame
	if err := s.memberGameRepo.GetByMemberID(id, &memberGames); err != nil {
		return nil, err
	}

	games := make([]models.Game, 0, len(memberGames))
	for _, mg := range memberGames {
		var game models.Game
		if err := s.gameRepo.GetByID(mg.GameID, &game); err != nil {
			continue
		}
		games = append(games, game)
	}
	member.Games = games

	return &member, nil
}

func (s *memberService) GetUserMembers(userID string) ([]models.Member, error) {
	var members []models.Member
	if err := s.memberRepo.GetByUserID(userID, &members); err != nil {
		return nil, err
	}

	for i := range members {
		var memberGames []models.MemberGame
		if err := s.memberGameRepo.GetByMemberID(members[i].ID, &memberGames); err != nil {
			continue
		}

		games := make([]models.Game, 0, len(memberGames))
		for _, mg := range memberGames {
			var game models.Game
			if err := s.gameRepo.GetByID(mg.GameID, &game); err != nil {
				continue
			}
			games = append(games, game)
		}
		members[i].Games = games
	}

	return members, nil
}

func (s *memberService) AddGame(props AddGameToMemberServiceProps) (*models.MemberGame, error) {
	var member models.Member
	if err := s.memberRepo.GetByID(props.MemberID, &member); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("membro com id %s não encontrado", props.MemberID)
		}
		return nil, err
	}

	if member.UserID != props.UserID {
		return nil, fmt.Errorf("você só pode adicionar jogos ao seu próprio perfil de membro")
	}

	var game models.Game
	if err := s.gameRepo.GetByID(props.GameID, &game); err != nil {
		if err == goe.ErrNotFound {
			return nil, fmt.Errorf("jogo com id %s não encontrado", props.GameID)
		}
		return nil, err
	}

	var existing models.MemberGame
	if err := s.memberGameRepo.GetByMemberIDAndGameID(props.MemberID, props.GameID, &existing); err == nil {
		return nil, fmt.Errorf("jogo já associado a este membro")
	}

	memberGame := models.MemberGame{
		ID:        uuid.New().String(),
		MemberID:  props.MemberID,
		GameID:    props.GameID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := s.memberGameRepo.Create(&memberGame)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (s *memberService) RemoveGame(memberID string, userID string, gameID string) error {
	var member models.Member
	if err := s.memberRepo.GetByID(memberID, &member); err != nil {
		if err == goe.ErrNotFound {
			return fmt.Errorf("membro com id %s não encontrado", memberID)
		}
		return err
	}

	if member.UserID != userID {
		return fmt.Errorf("você só pode remover jogos do seu próprio perfil de membro")
	}

	var memberGame models.MemberGame
	if err := s.memberGameRepo.GetByMemberIDAndGameID(memberID, gameID, &memberGame); err != nil {
		if err == goe.ErrNotFound {
			return fmt.Errorf("jogo não associado a este membro")
		}
		return err
	}

	return s.memberGameRepo.DeleteByID(memberGame.ID)
}
