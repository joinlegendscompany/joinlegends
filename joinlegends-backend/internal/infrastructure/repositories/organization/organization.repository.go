package organization_repository

import (
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/update"
	"github.com/go-goe/goe/query/where"
)

type organizationRepository struct {
	db *database.StreamDB
}

func NewOrganizationRepository(db *database.StreamDB) OrganizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) GetByID(id string, org *models.Organization) error {
	orgs, err := goe.List(r.db.Organizations).Where(
		where.Equals(&r.db.Organizations.ID, id),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(orgs) == 0 {
		return goe.ErrNotFound
	}
	*org = orgs[0]
	return nil
}

func (r *organizationRepository) GetByOwnerID(ownerID string, orgs *[]models.Organization) error {
	result, err := goe.List(r.db.Organizations).Where(
		where.Equals(&r.db.Organizations.OwnerID, ownerID),
	).AsSlice()

	if err != nil {
		return err
	}
	*orgs = result
	return nil
}

func (r *organizationRepository) Create(org *models.Organization) (models.Organization, error) {
	if org.CreatedAt.IsZero() {
		org.CreatedAt = time.Now()
	}

	err := goe.Insert(r.db.Organizations).One(org)
	if err != nil {
		return models.Organization{}, err
	}

	return *org, nil
}

func (r *organizationRepository) UpdateName(id string, name string) error {
	return goe.Update(r.db.Organizations).
		Sets(
			update.Set(&r.db.Organizations.Name, name),
			update.Set(&r.db.Organizations.UpdatedAt, time.Now()),
		).
		Where(where.Equals(&r.db.Organizations.ID, id))
}

func (r *organizationRepository) SoftDelete(id string) error {
	now := time.Now()
	return goe.Update(r.db.Organizations).
		Sets(
			update.Set(&r.db.Organizations.DeletedAt, &now),
			update.Set(&r.db.Organizations.UpdatedAt, now),
		).
		Where(where.Equals(&r.db.Organizations.ID, id))
}
