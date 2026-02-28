package organization

import (
	"fmt"
	"go-backend-stream/internal/utilities/logger"
	"net/http"

	"github.com/go-goe/goe"
	"github.com/gofiber/fiber/v2"
)

type OrganizationController struct {
	service OrganizationService
}

func NewOrganizationController(service OrganizationService) *OrganizationController {
	return &OrganizationController{service: service}
}

func (c *OrganizationController) CreateOrganizationController(ctx *fiber.Ctx) error {
	ownerID := ctx.Locals("userId").(string)

	var body CreateOrganizationDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Creating organization %s for user %s", body.Name, ownerID)
	org, err := c.service.CreateOrganization(CreateOrganizationServiceProps{
		Name:    body.Name,
		OwnerID: ownerID,
	})
	if err != nil {
		logger.Error.Printf("error creating organization: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error when creating organization",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message":      "organização criada com sucesso",
		"organization": org,
	})
}

func (c *OrganizationController) GetOrganizationController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	org, err := c.service.GetOrganization(id)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("organização com id %s não encontrada", id),
			})
		}
		logger.Error.Printf("error getting organization: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"organization": org,
	})
}

func (c *OrganizationController) GetUserOrganizationsController(ctx *fiber.Ctx) error {
	ownerID := ctx.Locals("userId").(string)

	orgs, err := c.service.GetUserOrganizations(ownerID)
	if err != nil {
		logger.Error.Printf("error getting user organizations: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"organizations": orgs,
	})
}

func (c *OrganizationController) UpdateOrganizationController(ctx *fiber.Ctx) error {
	ownerID := ctx.Locals("userId").(string)
	id := ctx.Params("id")

	var body UpdateOrganizationDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Updating organization %s", id)
	err := c.service.UpdateOrganization(id, ownerID, body.Name)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("organização com id %s não encontrada", id),
			})
		}
		if err.Error() == "only the owner can update the organization" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error updating organization: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "organização atualizada com sucesso",
	})
}

func (c *OrganizationController) DeleteOrganizationController(ctx *fiber.Ctx) error {
	ownerID := ctx.Locals("userId").(string)
	id := ctx.Params("id")

	logger.Info.Printf("Deleting organization %s", id)
	err := c.service.DeleteOrganization(id, ownerID)
	if err != nil {
		if err == goe.ErrNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("organização com id %s não encontrada", id),
			})
		}
		if err.Error() == "only the owner can delete the organization" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error deleting organization: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "organização deletada com sucesso",
	})
}

func (c *OrganizationController) AddMemberController(ctx *fiber.Ctx) error {
	ownerID := ctx.Locals("userId").(string)
	orgID := ctx.Params("id")

	var body AddMemberDto
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logger.Info.Printf("Adding member %s to organization %s", body.UserID, orgID)
	member, err := c.service.AddMember(orgID, ownerID, body.UserID, body.Role)
	if err != nil {
		if err.Error() == "only the owner can add members" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error adding member: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "membro adicionado com sucesso",
		"member":  member,
	})
}

func (c *OrganizationController) RemoveMemberController(ctx *fiber.Ctx) error {
	ownerID := ctx.Locals("userId").(string)
	orgID := ctx.Params("id")
	memberID := ctx.Params("memberId")

	logger.Info.Printf("Removing member %s from organization %s", memberID, orgID)
	err := c.service.RemoveMember(orgID, ownerID, memberID)
	if err != nil {
		if err.Error() == "only the owner can remove members" || err.Error() == "cannot remove the owner from the organization" {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err == goe.ErrNotFound || err.Error() == fmt.Sprintf("member with id %s not found", memberID) {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		logger.Error.Printf("error removing member: %s", err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "membro removido com sucesso",
	})
}
