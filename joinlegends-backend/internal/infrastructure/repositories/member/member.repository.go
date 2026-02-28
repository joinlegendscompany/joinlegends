package member_repository

import (
	"go-backend-stream/internal/infrastructure/database"
	"go-backend-stream/internal/models"
	"time"

	"github.com/go-goe/goe"
	"github.com/go-goe/goe/query/where"
)

type memberRepository struct {
	db *database.StreamDB
}

func NewMemberRepository(db *database.StreamDB) MemberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) Create(member *models.Member) (models.Member, error) {
	if member.CreatedAt.IsZero() {
		member.CreatedAt = time.Now()
	}

	err := goe.Insert(r.db.Members).One(member)
	if err != nil {
		return models.Member{}, err
	}

	return *member, nil
}

func (r *memberRepository) GetByOrganizationID(orgID string, members *[]models.Member) error {
	result, err := goe.List(r.db.Members).Where(
		where.Equals(&r.db.Members.OrganizationID, orgID),
	).AsSlice()

	if err != nil {
		return err
	}
	*members = result
	return nil
}

func (r *memberRepository) GetByID(id string, member *models.Member) error {
	members, err := goe.List(r.db.Members).Where(
		where.Equals(&r.db.Members.ID, id),
	).Take(1).AsSlice()

	if err != nil {
		return err
	}
	if len(members) == 0 {
		return goe.ErrNotFound
	}
	*member = members[0]
	return nil
}

func (r *memberRepository) DeleteByID(id string) error {
	return goe.Delete(r.db.Members).Where(
		where.Equals(&r.db.Members.ID, id),
	)
}
