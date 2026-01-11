package database

import (
	"go-backend-stream/internal/models"

	"github.com/go-goe/goe"
)

type StreamDB struct {
	Users         *models.User
	Recoveries    *models.Recovery
	Sessions      *models.Session
	Organizations *models.Organization
	Members       *models.Member
	*goe.DB
}
