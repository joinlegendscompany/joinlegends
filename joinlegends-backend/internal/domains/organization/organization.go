package organization

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	//IMplement here, only admin can create a new organization
	return c.Status(http.StatusNotImplemented).JSON(fiber.Map{
		"message": "no implemented yet",
	})
}
