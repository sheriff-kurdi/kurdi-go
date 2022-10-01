package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AuthenticationMiddleware(c *fiber.Ctx) error {
	if false {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}
	return c.Next()

}
