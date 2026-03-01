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

func TestOrganizationService_UpdateOrganization(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgID := "org-123"
		ownerID := "user-123"
		foundOrg := models.Organization{ID: orgID, OwnerID: ownerID}

		orgRepo.On("GetByID", orgID, mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)
		orgRepo.On("UpdateName", orgID, "Novo Nome").Return(nil)

		err := service.UpdateOrganization(orgID, ownerID, "Novo Nome")

		assert.NoError(t, err)
		orgRepo.AssertExpectations(t)
	})

	t.Run("org_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "missing", mock.AnythingOfType("*models.Organization")).
			Return(nil, goe.ErrNotFound)

		err := service.UpdateOrganization("missing", "user-123", "Nome")

		assert.Error(t, err)
		assert.Equal(t, goe.ErrNotFound, err)
		orgRepo.AssertNotCalled(t, "UpdateName")
	})

	t.Run("not_owner", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		foundOrg := models.Organization{ID: "org-123", OwnerID: "real-owner"}
		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)

		err := service.UpdateOrganization("org-123", "other-user", "Nome")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "only the owner")
		orgRepo.AssertNotCalled(t, "UpdateName")
	})

	t.Run("update_db_error", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		foundOrg := models.Organization{ID: "org-123", OwnerID: "user-123"}
		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)
		orgRepo.On("UpdateName", "org-123", "Nome").Return(errors.New("db error"))

		err := service.UpdateOrganization("org-123", "user-123", "Nome")

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
	})
}

func TestOrganizationService_DeleteOrganization(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgID := "org-123"
		ownerID := "user-123"
		foundOrg := models.Organization{ID: orgID, OwnerID: ownerID}

		orgRepo.On("GetByID", orgID, mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)
		orgRepo.On("SoftDelete", orgID).Return(nil)

		err := service.DeleteOrganization(orgID, ownerID)

		assert.NoError(t, err)
		orgRepo.AssertExpectations(t)
	})

	t.Run("org_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "missing", mock.AnythingOfType("*models.Organization")).
			Return(nil, goe.ErrNotFound)

		err := service.DeleteOrganization("missing", "user-123")

		assert.Error(t, err)
		assert.Equal(t, goe.ErrNotFound, err)
		orgRepo.AssertNotCalled(t, "SoftDelete")
	})

	t.Run("not_owner", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		foundOrg := models.Organization{ID: "org-123", OwnerID: "real-owner"}
		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&foundOrg, nil)

		err := service.DeleteOrganization("org-123", "intruder")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "only the owner")
		orgRepo.AssertNotCalled(t, "SoftDelete")
	})
}
