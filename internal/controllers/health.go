package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck handles the health check request
// @Summary Health Check
// @Description Checks if the service is healthy
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Service is healthy",
	})
}
