package handlers

import (
	"github.com/gofiber/fiber/v2"
	"omniarr/internal/api/response"
)

func HealthHandler(c *fiber.Ctx) error {
	return response.OK(c, fiber.Map{
		"status":  "ok",
		"version": "1.0.0",
	})
}
