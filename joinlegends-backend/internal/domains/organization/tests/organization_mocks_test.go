package organization_test

import (
	member_repository "go-backend-stream/internal/infrastructure/repositories/member"
	organization_repository "go-backend-stream/internal/infrastructure/repositories/organization"
	user_repository "go-backend-stream/internal/infrastructure/repositories/user"
	"go-backend-stream/internal/models"

	"github.com/stretchr/testify/mock"
)

// ─── OrganizationRepository mock ────────────────────────────────────────────

type MockOrganizationRepository struct {
	mock.Mock
}

func (m *MockOrganizationRepository) GetByID(id string, org *models.Organization) error {
	args := m.Called(id, org)
	if o, ok := args.Get(0).(*models.Organization); ok && o != nil {
		*org = *o
	}
	return args.Error(1)
}

func (m *MockOrganizationRepository) GetByOwnerID(ownerID string, orgs *[]models.Organization) error {
	args := m.Called(ownerID, orgs)
	if o, ok := args.Get(0).([]models.Organization); ok {
		*orgs = o
	}
	return args.Error(1)
}

func (m *MockOrganizationRepository) Create(org *models.Organization) (models.Organization, error) {
	args := m.Called(org)
	return args.Get(0).(models.Organization), args.Error(1)
}

func (m *MockOrganizationRepository) UpdateName(id string, name string) error {
	args := m.Called(id, name)
	return args.Error(0)
}

func (m *MockOrganizationRepository) SoftDelete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ organization_repository.OrganizationRepository = (*MockOrganizationRepository)(nil)

// ─── MemberRepository mock ────────────────────────────────────────────────────

type MockMemberRepository struct {
	mock.Mock
}

func (m *MockMemberRepository) Create(member *models.Member) (models.Member, error) {
	args := m.Called(member)
	return args.Get(0).(models.Member), args.Error(1)
}

func (m *MockMemberRepository) GetByOrganizationID(orgID string, members *[]models.Member) error {
	args := m.Called(orgID, members)
	if ms, ok := args.Get(0).([]models.Member); ok {
		*members = ms
	}
	return args.Error(1)
}

func (m *MockMemberRepository) GetByID(id string, member *models.Member) error {
	args := m.Called(id, member)
	if mb, ok := args.Get(0).(*models.Member); ok && mb != nil {
		*member = *mb
	}
	return args.Error(1)
}

func (m *MockMemberRepository) GetByUserID(userID string, members *[]models.Member) error {
	args := m.Called(userID, members)
	if ms, ok := args.Get(0).([]models.Member); ok {
		*members = ms
	}
	return args.Error(1)
}

func (m *MockMemberRepository) DeleteByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

var _ member_repository.MemberRepository = (*MockMemberRepository)(nil)

// ─── UserRepository mock ──────────────────────────────────────────────────────

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmail(email string, user *models.User) error {
	args := m.Called(email, user)
	if u, ok := args.Get(0).(*models.User); ok && u != nil {
		*user = *u
	}
	return args.Error(1)
}

func (m *MockUserRepository) Create(user *models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) GetById(id string, user *models.User) error {
	args := m.Called(id, user)
	if u, ok := args.Get(0).(*models.User); ok && u != nil {
		*user = *u
	}
	return args.Error(1)
}

func (m *MockUserRepository) UpdatePassword(userID string, newPassword string) error {
	args := m.Called(userID, newPassword)
	return args.Error(0)
}

var _ user_repository.UserRepository = (*MockUserRepository)(nil)
