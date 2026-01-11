package user_repository

import (
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/update"
	"github.com/go-goe/goe/query/where"
)

type userRepository struct {
	db *database.StreamDB
}

func NewUserRepository(db *database.StreamDB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetById(id string, user *models.User) error {
	users, err := goe.List(r.db.Users).Where(
		where.Equals(&r.db.Users.ID, id),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(users) == 0 {
		return goe.ErrNotFound
	}
	*user = users[0]
	return nil
}

func (r *userRepository) GetByEmail(email string, user *models.User) error {
	users, err := goe.List(r.db.Users).Where(
		where.Equals(&r.db.Users.Email, email),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(users) == 0 {
		return goe.ErrNotFound
	}
	*user = users[0]
	return nil
}

func (r *userRepository) Create(user *models.User) (models.User, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	err := goe.Insert(r.db.Users).One(user)
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (r *userRepository) UpdatePassword(userID string, newPassword string) error {
	err := goe.Update(r.db.Users).
		Sets(
			update.Set(&r.db.Users.Password, newPassword),
			update.Set(&r.db.Users.UpdatedAt, time.Now()),
		).
		Where(where.Equals(&r.db.Users.ID, userID))

	return err
}
