package session_repository

import (
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"
	"go-backend-stream/internal/utilities/logger"
	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/where"
)

type sessionRepository struct {
	db *database.StreamDB
}

func NewSessionRepository(db *database.StreamDB) SessionRepository {
	return &sessionRepository{db}
}

func (r *sessionRepository) GetAllByUserID(userID string, sessions *[]models.Session) error {
	result, err := goe.List(r.db.Sessions).Where(
		where.Equals(&r.db.Sessions.UserID, userID),
	).AsSlice()
	if err != nil {
		return err
	}
	*sessions = result
	return nil
}

func (r *sessionRepository) GetByID(
	id string,
	session *models.Session,
) error {
	sessions, err := goe.List(r.db.Sessions).Where(
		where.Equals(&r.db.Sessions.ID, id),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}

	if len(sessions) == 0 {
		return goe.ErrNotFound
	}

	*session = sessions[0]
	return nil
}

func (r *sessionRepository) Create(session *models.Session) (*models.Session, error) {
	// Print the ssession
	logger.Error.Println("Session data: ", session)

	if session.CreatedAt.IsZero() {
		session.CreatedAt = time.Now()
	}

	err := goe.Insert(r.db.Sessions).One(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *sessionRepository) DeleteByID(id string) error {
	err := goe.Delete(r.db.Sessions).Where(
		where.Equals(&r.db.Sessions.ID, id),
	)
	return err
}

func (r *sessionRepository) DeleteByUserId(userId string) error {
	err := goe.Delete(r.db.Sessions).Where(
		where.Equals(&r.db.Sessions.UserID, userId),
	)
	return err
}
