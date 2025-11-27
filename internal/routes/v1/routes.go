package v1

import (
	"github.com/gofiber/fiber/v2"

	"sens.ai-backend/internal/routes/v1/health"
)

// SetupRoutes configures the v1 routes
func SetupRoutes(router fiber.Router) {
	healthGroup := router.Group("/health")
	health.SetupRoutes(healthGroup)
}
