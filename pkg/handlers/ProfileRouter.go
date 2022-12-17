package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) ProfileRouter(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
