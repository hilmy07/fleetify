package middleware

import (
	"fleetify/internal/config"
	"fleetify/internal/models"

	"github.com/gofiber/fiber/v2"
)

func RBACMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userID := c.Get("X-User-ID")

		if userID == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "missing user id",
			})
		}

		var user models.User

		err := config.DB.First(&user, userID).Error
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"message": "user not found",
			})
		}

		if user.Role != role {
			return c.Status(403).JSON(fiber.Map{
				"message": "forbidden",
			})
		}

		c.Locals("user_id", user.ID)

		return c.Next()
	}
}