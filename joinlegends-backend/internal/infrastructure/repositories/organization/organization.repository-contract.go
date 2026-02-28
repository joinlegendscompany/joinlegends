package organization_repository

import "go-backend-stream/internal/models"

type OrganizationRepository interface {
	GetByID(id string, org *models.Organization) error
	GetByOwnerID(ownerID string, orgs *[]models.Organization) error
	Create(org *models.Organization) (models.Organization, error)
	UpdateName(id string, name string) error
	SoftDelete(id string) error
}
