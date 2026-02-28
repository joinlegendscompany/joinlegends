package member_repository

import "go-backend-stream/internal/models"

type MemberRepository interface {
	Create(member *models.Member) (models.Member, error)
	GetByOrganizationID(orgID string, members *[]models.Member) error
	GetByID(id string, member *models.Member) error
	DeleteByID(id string) error
}
