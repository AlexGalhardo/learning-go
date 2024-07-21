package handler

import "github.com/gofiber/fiber/v2"

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": true, "message": "API is on, let's goo!!!"})
}
