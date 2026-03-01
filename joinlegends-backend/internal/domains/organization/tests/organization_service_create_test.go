package organization_test

import (
	"errors"
	"go-backend-stream/internal/domains/organization"
	"go-backend-stream/internal/models"
	"testing"

	"github.com/go-goe/goe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrganizationService_CreateOrganization(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		ownerID := "user-123"
		props := organization.CreateOrganizationServiceProps{Name: "Minha Org", OwnerID: ownerID}

		userRepo.On("GetById", ownerID, mock.AnythingOfType("*models.User")).
			Return(&models.User{ID: ownerID}, nil)

		createdOrg := models.Organization{ID: "org-123", Name: props.Name, OwnerID: ownerID}
		orgRepo.On("Create", mock.AnythingOfType("*models.Organization")).
			Return(createdOrg, nil)

		ownerMember := models.Member{ID: "member-123", UserID: ownerID, OrganizationID: "org-123", Role: "OWNER"}
		memberRepo.On("Create", mock.AnythingOfType("*models.Member")).
			Return(ownerMember, nil)

		org, err := service.CreateOrganization(props)

		assert.NoError(t, err)
		assert.NotNil(t, org)
		assert.Equal(t, "org-123", org.ID)
		assert.Equal(t, "Minha Org", org.Name)
		assert.Len(t, org.Members, 1)
		assert.Equal(t, "OWNER", org.Members[0].Role)

		orgRepo.AssertExpectations(t)
		memberRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})

	t.Run("owner_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		props := organization.CreateOrganizationServiceProps{Name: "Org", OwnerID: "missing-user"}

		userRepo.On("GetById", "missing-user", mock.AnythingOfType("*models.User")).
			Return(nil, goe.ErrNotFound)

		org, err := service.CreateOrganization(props)

		assert.Error(t, err)
		assert.Nil(t, org)
		assert.Contains(t, err.Error(), "not found")

		orgRepo.AssertNotCalled(t, "Create")
	})

	t.Run("org_create_db_error", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		props := organization.CreateOrganizationServiceProps{Name: "Org", OwnerID: "user-123"}

		userRepo.On("GetById", "user-123", mock.AnythingOfType("*models.User")).
			Return(&models.User{ID: "user-123"}, nil)
		orgRepo.On("Create", mock.AnythingOfType("*models.Organization")).
			Return(models.Organization{}, errors.New("db error"))

		org, err := service.CreateOrganization(props)

		assert.Error(t, err)
		assert.Nil(t, org)
		assert.Equal(t, "db error", err.Error())

		memberRepo.AssertNotCalled(t, "Create")
	})

	t.Run("member_create_db_error", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		props := organization.CreateOrganizationServiceProps{Name: "Org", OwnerID: "user-123"}

		userRepo.On("GetById", "user-123", mock.AnythingOfType("*models.User")).
			Return(&models.User{ID: "user-123"}, nil)
		orgRepo.On("Create", mock.AnythingOfType("*models.Organization")).
			Return(models.Organization{ID: "org-123", Name: "Org", OwnerID: "user-123"}, nil)
		memberRepo.On("Create", mock.AnythingOfType("*models.Member")).
			Return(models.Member{}, errors.New("member db error"))

		org, err := service.CreateOrganization(props)

		assert.Error(t, err)
		assert.Nil(t, org)
		assert.Equal(t, "member db error", err.Error())
	})
}
