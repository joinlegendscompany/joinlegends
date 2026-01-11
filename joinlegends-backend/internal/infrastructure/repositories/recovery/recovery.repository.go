package recovery_repository

import (
	"context"
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"

	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/update"
	"github.com/go-goe/goe/query/where"
)

type recoveryRepository struct {
	db *database.StreamDB
}

func NewRecoveryRepository(db *database.StreamDB) RecoveryRepository {
	return &recoveryRepository{db}
}

func (r *recoveryRepository) GetByUserID(userID string, recovery *models.Recovery) error {
	recoveries, err := goe.List(r.db.Recoveries).Where(
		where.Equals(&r.db.Recoveries.UserID, userID),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(recoveries) == 0 {
		return goe.ErrNotFound
	}
	*recovery = recoveries[0]
	return nil
}

func (r *recoveryRepository) GetByEmail(email string, recovery *models.Recovery) error {
	recoveries, err := goe.List(r.db.Recoveries).Where(
		where.Equals(&r.db.Recoveries.Email, email),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(recoveries) == 0 {
		return goe.ErrNotFound
	}
	*recovery = recoveries[0]
	return nil
}

func (r *recoveryRepository) GetByCode(code string, recovery *models.Recovery) error {
	recoveries, err := goe.List(r.db.Recoveries).Where(
		where.Equals(&r.db.Recoveries.Code, code),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(recoveries) == 0 {
		return goe.ErrNotFound
	}
	*recovery = recoveries[0]
	return nil
}

func (r *recoveryRepository) Create(recovery *models.Recovery) (models.Recovery, error) {
	if recovery.CreatedAt.IsZero() {
		recovery.CreatedAt = time.Now()
	}
	recovery.UpdatedAt = time.Now()

	err := goe.Insert(r.db.Recoveries).One(recovery)
	if err != nil {
		return models.Recovery{}, err
	}

	return *recovery, nil
}

func (r *recoveryRepository) IncrementAttempts(id int) error {
	return r.db.RawExecContext(context.Background(), `
		UPDATE recoveries
		SET attempts = attempts + 1, updated_at = $2
		WHERE id = $1
	`, id, time.Now())
}

func (r *recoveryRepository) ClearByID(id int) error {
	return r.db.RawExecContext(context.Background(), `
		UPDATE recoveries
		SET code = NULL,
		attempts = 0,
		expires_at = NULL,
		updated_at = NOW()
		WHERE id = $1
	`, id)
}

func (r *recoveryRepository) ClearByUserID(userID string) error {
	return r.db.RawExecContext(context.Background(), `
		UPDATE recoveries
		SET code = NULL,
		attempts = 0,
		expires_at = NULL,
		updated_at = NOW()
		WHERE user_id = $1
	`, userID)
}

func (r *recoveryRepository) ClearByEmail(email string) error {
	return r.db.RawExecContext(context.Background(), `
		UPDATE recoveries
		SET code = NULL,
		attempts = 0,
		expires_at = NULL,
		updated_at = NOW()
		WHERE email = $1
	`, email)
}

func (r *recoveryRepository) UpdateRecovery(id int, code string, attempts int, expiresAt time.Time, expired bool) error {
	err := goe.Update(r.db.Recoveries).
		Sets(
			update.Set(&r.db.Recoveries.Code, code),
			update.Set(&r.db.Recoveries.Attempts, attempts),
			update.Set(&r.db.Recoveries.ExpiresAt, expiresAt),
			update.Set(&r.db.Recoveries.Expired, expired),
			update.Set(&r.db.Recoveries.UpdatedAt, time.Now()),
		).
		Where(where.Equals(&r.db.Recoveries.ID, id))

	return err
}

func (r *recoveryRepository) MarkAsExpired(id int) error {
	err := goe.Update(r.db.Recoveries).
		Sets(
			update.Set(&r.db.Recoveries.Expired, true),
			update.Set(&r.db.Recoveries.UpdatedAt, time.Now()),
		).
		Where(where.Equals(&r.db.Recoveries.ID, id))

	return err
}
