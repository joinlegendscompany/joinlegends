package organization_test

import (
	"bytes"
	"encoding/json"
	"go-backend-stream/internal/domains/organization"
	"go-backend-stream/internal/models"
	"net/http/httptest"
	"testing"

	"github.com/go-goe/goe"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ─── MockOrganizationService ──────────────────────────────────────────────────

type MockOrganizationService struct {
	mock.Mock
}

func (m *MockOrganizationService) CreateOrganization(props organization.CreateOrganizationServiceProps) (*models.Organization, error) {
	args := m.Called(props)
	if o, ok := args.Get(0).(*models.Organization); ok {
		return o, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationService) GetOrganization(id string) (*models.Organization, error) {
	args := m.Called(id)
	if o, ok := args.Get(0).(*models.Organization); ok {
		return o, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationService) GetUserOrganizations(ownerID string) ([]models.Organization, error) {
	args := m.Called(ownerID)
	if o, ok := args.Get(0).([]models.Organization); ok {
		return o, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationService) UpdateOrganization(id string, ownerID string, name string) error {
	args := m.Called(id, ownerID, name)
	return args.Error(0)
}

func (m *MockOrganizationService) DeleteOrganization(id string, ownerID string) error {
	args := m.Called(id, ownerID)
	return args.Error(0)
}

func (m *MockOrganizationService) AddMember(orgID string, ownerID string, userID string, role string) (*models.Member, error) {
	args := m.Called(orgID, ownerID, userID, role)
	if mb, ok := args.Get(0).(*models.Member); ok {
		return mb, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationService) RemoveMember(orgID string, ownerID string, memberID string) error {
	args := m.Called(orgID, ownerID, memberID)
	return args.Error(0)
}

var _ organization.OrganizationService = (*MockOrganizationService)(nil)

// ─── helpers ──────────────────────────────────────────────────────────────────

const testUserID = "user-e2e-123"

func newOrgE2EApp(svc organization.OrganizationService) *fiber.App {
	app := fiber.New()
	ctrl := organization.NewOrganizationController(svc)

	// middleware de teste que injeta userId sem precisar de JWT real
	auth := func(c *fiber.Ctx) error {
		c.Locals("userId", testUserID)
		return c.Next()
	}

	v1 := app.Group("v1/organization")
	v1.Post("/", auth, ctrl.CreateOrganizationController)
	v1.Get("/", auth, ctrl.GetUserOrganizationsController)
	v1.Get("/:id", auth, ctrl.GetOrganizationController)
	v1.Patch("/:id", auth, ctrl.UpdateOrganizationController)
	v1.Delete("/:id", auth, ctrl.DeleteOrganizationController)
	v1.Post("/:id/members", auth, ctrl.AddMemberController)
	v1.Delete("/:id/members/:memberId", auth, ctrl.RemoveMemberController)

	return app
}

func toBody(t *testing.T, v any) *bytes.Buffer {
	t.Helper()
	b, err := json.Marshal(v)
	assert.NoError(t, err)
	return bytes.NewBuffer(b)
}

// ─── Create ───────────────────────────────────────────────────────────────────

func TestOrgE2E_CreateOrganization(t *testing.T) {
	t.Run("201_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		props := organization.CreateOrganizationServiceProps{Name: "Org Alpha", OwnerID: testUserID}
		org := &models.Organization{ID: "org-1", Name: "Org Alpha", OwnerID: testUserID}
		svc.On("CreateOrganization", props).Return(org, nil)

		req := httptest.NewRequest("POST", "/v1/organization/", toBody(t, map[string]string{"name": "Org Alpha"}))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		svc.AssertExpectations(t)
	})

	t.Run("400_invalid_body", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		req := httptest.NewRequest("POST", "/v1/organization/", bytes.NewBufferString("not-json"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		svc.AssertNotCalled(t, "CreateOrganization")
	})

	t.Run("500_service_error", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		props := organization.CreateOrganizationServiceProps{Name: "Org", OwnerID: testUserID}
		svc.On("CreateOrganization", props).Return(nil, fiber.NewError(0, "db error"))

		req := httptest.NewRequest("POST", "/v1/organization/", toBody(t, map[string]string{"name": "Org"}))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

// ─── Get By ID ────────────────────────────────────────────────────────────────

func TestOrgE2E_GetOrganization(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		org := &models.Organization{
			ID:      "org-1",
			Name:    "Org Alpha",
			Members: []models.Member{{ID: "m1", Role: "OWNER"}},
		}
		svc.On("GetOrganization", "org-1").Return(org, nil)

		req := httptest.NewRequest("GET", "/v1/organization/org-1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("404_not_found", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("GetOrganization", "missing").Return(nil, goe.ErrNotFound)

		req := httptest.NewRequest("GET", "/v1/organization/missing", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

// ─── List User Organizations ──────────────────────────────────────────────────

func TestOrgE2E_GetUserOrganizations(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		orgs := []models.Organization{
			{ID: "org-1", Name: "Alpha"},
			{ID: "org-2", Name: "Beta"},
		}
		svc.On("GetUserOrganizations", testUserID).Return(orgs, nil)

		req := httptest.NewRequest("GET", "/v1/organization/", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("200_empty_list", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("GetUserOrganizations", testUserID).Return([]models.Organization{}, nil)

		req := httptest.NewRequest("GET", "/v1/organization/", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

// ─── Update ───────────────────────────────────────────────────────────────────

func TestOrgE2E_UpdateOrganization(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("UpdateOrganization", "org-1", testUserID, "Novo Nome").Return(nil)

		req := httptest.NewRequest("PATCH", "/v1/organization/org-1", toBody(t, map[string]string{"name": "Novo Nome"}))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("403_not_owner", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("UpdateOrganization", "org-1", testUserID, "Nome").
			Return(fiber.NewError(0, "only the owner can update the organization"))

		req := httptest.NewRequest("PATCH", "/v1/organization/org-1", toBody(t, map[string]string{"name": "Nome"}))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)
	})

	t.Run("404_not_found", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("UpdateOrganization", "missing", testUserID, "Nome").Return(goe.ErrNotFound)

		req := httptest.NewRequest("PATCH", "/v1/organization/missing", toBody(t, map[string]string{"name": "Nome"}))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

// ─── Delete ───────────────────────────────────────────────────────────────────

func TestOrgE2E_DeleteOrganization(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("DeleteOrganization", "org-1", testUserID).Return(nil)

		req := httptest.NewRequest("DELETE", "/v1/organization/org-1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("403_not_owner", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("DeleteOrganization", "org-1", testUserID).
			Return(fiber.NewError(0, "only the owner can delete the organization"))

		req := httptest.NewRequest("DELETE", "/v1/organization/org-1", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)
	})
}

// ─── Add Member ───────────────────────────────────────────────────────────────

func TestOrgE2E_AddMember(t *testing.T) {
	t.Run("201_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		member := &models.Member{ID: "m-new", UserID: "user-new", OrganizationID: "org-1", Role: "MEMBER"}
		svc.On("AddMember", "org-1", testUserID, "user-new", "MEMBER").Return(member, nil)

		body := map[string]string{"user_id": "user-new", "role": "MEMBER"}
		req := httptest.NewRequest("POST", "/v1/organization/org-1/members", toBody(t, body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	t.Run("403_not_owner", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("AddMember", "org-1", testUserID, "user-new", "MEMBER").
			Return(nil, fiber.NewError(0, "only the owner can add members"))

		body := map[string]string{"user_id": "user-new", "role": "MEMBER"}
		req := httptest.NewRequest("POST", "/v1/organization/org-1/members", toBody(t, body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)
	})

	t.Run("400_invalid_body", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		req := httptest.NewRequest("POST", "/v1/organization/org-1/members", bytes.NewBufferString("bad"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		svc.AssertNotCalled(t, "AddMember", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	})
}

// ─── Remove Member ────────────────────────────────────────────────────────────

func TestOrgE2E_RemoveMember(t *testing.T) {
	t.Run("200_success", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("RemoveMember", "org-1", testUserID, "member-abc").Return(nil)

		req := httptest.NewRequest("DELETE", "/v1/organization/org-1/members/member-abc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("403_cannot_remove_owner", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("RemoveMember", "org-1", testUserID, "owner-member").
			Return(fiber.NewError(0, "cannot remove the owner from the organization"))

		req := httptest.NewRequest("DELETE", "/v1/organization/org-1/members/owner-member", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)
	})

	t.Run("404_member_not_found", func(t *testing.T) {
		svc := new(MockOrganizationService)
		app := newOrgE2EApp(svc)

		svc.On("RemoveMember", "org-1", testUserID, "ghost").
			Return(fiber.NewError(0, "member with id ghost not found"))

		req := httptest.NewRequest("DELETE", "/v1/organization/org-1/members/ghost", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}
