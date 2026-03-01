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

func TestOrganizationService_GetOrganization(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgID := "org-123"
		foundOrg := models.Organization{ID: orgID, Name: "Org Alpha", OwnerID: "user-123"}
		members := []models.Member{
			{ID: "m1", UserID: "user-123", OrganizationID: orgID, Role: "OWNER"},
			{ID: "m2", UserID: "user-456", OrganizationID: orgID, Role: "MEMBER"},
		}

		orgRepo.On("GetByID", orgID, mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)
		memberRepo.On("GetByOrganizationID", orgID, mock.AnythingOfType("*[]models.Member")).
			Return(members, nil)

		org, err := service.GetOrganization(orgID)

		assert.NoError(t, err)
		assert.NotNil(t, org)
		assert.Equal(t, orgID, org.ID)
		assert.Len(t, org.Members, 2)

		orgRepo.AssertExpectations(t)
		memberRepo.AssertExpectations(t)
	})

	t.Run("org_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "missing", mock.AnythingOfType("*models.Organization")).
			Return(nil, goe.ErrNotFound)

		org, err := service.GetOrganization("missing")

		assert.Error(t, err)
		assert.Nil(t, org)
		assert.Equal(t, goe.ErrNotFound, err)

		memberRepo.AssertNotCalled(t, "GetByOrganizationID")
	})

	t.Run("members_db_error", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		foundOrg := models.Organization{ID: "org-123"}
		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)
		memberRepo.On("GetByOrganizationID", "org-123", mock.AnythingOfType("*[]models.Member")).
			Return(nil, errors.New("db error"))

		org, err := service.GetOrganization("org-123")

		assert.Error(t, err)
		assert.Nil(t, org)
	})
}

func TestOrganizationService_GetUserOrganizations(t *testing.T) {
	t.Run("success_with_results", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		ownerID := "user-123"
		orgs := []models.Organization{
			{ID: "org-1", Name: "Alpha", OwnerID: ownerID},
			{ID: "org-2", Name: "Beta", OwnerID: ownerID},
		}

		orgRepo.On("GetByOwnerID", ownerID, mock.AnythingOfType("*[]models.Organization")).
			Return(orgs, nil)

		result, err := service.GetUserOrganizations(ownerID)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("success_empty", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByOwnerID", "user-no-orgs", mock.AnythingOfType("*[]models.Organization")).
			Return([]models.Organization{}, nil)

		result, err := service.GetUserOrganizations("user-no-orgs")

		assert.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("db_error", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByOwnerID", "user-123", mock.AnythingOfType("*[]models.Organization")).
			Return(nil, errors.New("db error"))

		result, err := service.GetUserOrganizations("user-123")

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
