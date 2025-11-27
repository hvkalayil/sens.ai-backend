package health

import (
	"github.com/gofiber/fiber/v2"

	"sens.ai-backend/internal/controllers"
)

// SetupRoutes configures the health routes
func SetupRoutes(router fiber.Router) {
	router.Get("/", controllers.HealthCheck)
}
