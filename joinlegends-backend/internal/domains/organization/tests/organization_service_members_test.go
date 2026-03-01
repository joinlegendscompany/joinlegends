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

func TestOrganizationService_AddMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgID := "org-123"
		ownerID := "user-owner"
		userID := "user-new"

		orgRepo.On("GetByID", orgID, mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: orgID, OwnerID: ownerID}, nil)
		userRepo.On("GetById", userID, mock.AnythingOfType("*models.User")).
			Return(&models.User{ID: userID}, nil)

		createdMember := models.Member{ID: "m-new", UserID: userID, OrganizationID: orgID, Role: "MEMBER"}
		memberRepo.On("Create", mock.AnythingOfType("*models.Member")).Return(createdMember, nil)

		member, err := service.AddMember(orgID, ownerID, userID, "MEMBER")

		assert.NoError(t, err)
		assert.NotNil(t, member)
		assert.Equal(t, "MEMBER", member.Role)
		assert.Equal(t, userID, member.UserID)

		orgRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		memberRepo.AssertExpectations(t)
	})

	t.Run("org_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "missing-org", mock.AnythingOfType("*models.Organization")).
			Return(nil, goe.ErrNotFound)

		member, err := service.AddMember("missing-org", "owner", "user", "MEMBER")

		assert.Error(t, err)
		assert.Nil(t, member)
		assert.Contains(t, err.Error(), "not found")

		userRepo.AssertNotCalled(t, "GetById")
		memberRepo.AssertNotCalled(t, "Create")
	})

	t.Run("not_owner", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "real-owner"}, nil)

		member, err := service.AddMember("org-123", "intruder", "some-user", "MEMBER")

		assert.Error(t, err)
		assert.Nil(t, member)
		assert.Contains(t, err.Error(), "only the owner")

		memberRepo.AssertNotCalled(t, "Create")
	})

	t.Run("user_to_add_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "owner"}, nil)
		userRepo.On("GetById", "ghost-user", mock.AnythingOfType("*models.User")).
			Return(nil, goe.ErrNotFound)

		member, err := service.AddMember("org-123", "owner", "ghost-user", "MEMBER")

		assert.Error(t, err)
		assert.Nil(t, member)
		assert.Contains(t, err.Error(), "not found")

		memberRepo.AssertNotCalled(t, "Create")
	})

	t.Run("member_create_db_error", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "owner"}, nil)
		userRepo.On("GetById", "user-new", mock.AnythingOfType("*models.User")).
			Return(&models.User{ID: "user-new"}, nil)
		memberRepo.On("Create", mock.AnythingOfType("*models.Member")).
			Return(models.Member{}, errors.New("db error"))

		member, err := service.AddMember("org-123", "owner", "user-new", "ADMIN")

		assert.Error(t, err)
		assert.Nil(t, member)
	})
}

func TestOrganizationService_RemoveMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgID := "org-123"
		ownerID := "owner"
		memberID := "member-456"

		orgRepo.On("GetByID", orgID, mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: orgID, OwnerID: ownerID}, nil)
		memberRepo.On("GetByID", memberID, mock.AnythingOfType("*models.Member")).
			Return(&models.Member{ID: memberID, OrganizationID: orgID, Role: "MEMBER"}, nil)
		memberRepo.On("DeleteByID", memberID).Return(nil)

		err := service.RemoveMember(orgID, ownerID, memberID)

		assert.NoError(t, err)
		memberRepo.AssertExpectations(t)
	})

	t.Run("not_owner", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "real-owner"}, nil)

		err := service.RemoveMember("org-123", "intruder", "member-456")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "only the owner")
		memberRepo.AssertNotCalled(t, "GetByID")
	})

	t.Run("member_not_found", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "owner"}, nil)
		memberRepo.On("GetByID", "ghost-member", mock.AnythingOfType("*models.Member")).
			Return(nil, goe.ErrNotFound)

		err := service.RemoveMember("org-123", "owner", "ghost-member")

		assert.Error(t, err)
		memberRepo.AssertNotCalled(t, "DeleteByID")
	})

	t.Run("cannot_remove_owner", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "owner"}, nil)
		memberRepo.On("GetByID", "owner-member-id", mock.AnythingOfType("*models.Member")).
			Return(&models.Member{ID: "owner-member-id", OrganizationID: "org-123", Role: "OWNER"}, nil)

		err := service.RemoveMember("org-123", "owner", "owner-member-id")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot remove the owner")
		memberRepo.AssertNotCalled(t, "DeleteByID")
	})

	t.Run("member_belongs_to_other_org", func(t *testing.T) {
		orgRepo := new(MockOrganizationRepository)
		memberRepo := new(MockMemberRepository)
		userRepo := new(MockUserRepository)
		service := organization.NewOrganizationService(orgRepo, memberRepo, userRepo)

		orgRepo.On("GetByID", "org-123", mock.AnythingOfType("*models.Organization")).
			Return(&models.Organization{ID: "org-123", OwnerID: "owner"}, nil)
		memberRepo.On("GetByID", "member-from-other-org", mock.AnythingOfType("*models.Member")).
			Return(&models.Member{ID: "member-from-other-org", OrganizationID: "other-org", Role: "MEMBER"}, nil)

		err := service.RemoveMember("org-123", "owner", "member-from-other-org")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "does not belong")
		memberRepo.AssertNotCalled(t, "DeleteByID")
	})
}
