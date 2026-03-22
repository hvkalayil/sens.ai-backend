package controllers

import (
	"github.com/gofiber/fiber/v2"

	"sens.ai-backend/internal/db"
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
	dbStatus := "down"
	if db.Pool != nil {
		if err := db.Pool.Ping(c.Context()); err == nil {
			dbStatus = "up"
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Service is healthy",
		"database": dbStatus,
	})
}
