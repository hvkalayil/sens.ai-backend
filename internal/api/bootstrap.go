package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	"sens.ai-backend/docs"
	"sens.ai-backend/internal/logger"
	"sens.ai-backend/internal/middleware"
	v1 "sens.ai-backend/internal/routes/v1"
)

func Start(host, port string) {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Middleware
	app.Use(middleware.CORS())
	app.Use(recover.New())
	app.Use(middleware.RequestLogger())

	// Setup Routes
	app.Static("/", "./public")

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)
	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1Group := api.Group("/v1")
	v1.SetupRoutes(v1Group)

	// Start server
	addr := fmt.Sprintf("%s:%s", host, port)

	logger.Logger.Info().
		Str("host", host).
		Str("port", port).
		Str("address", addr).
		Msg("Starting server")

	if err := app.Listen(addr); err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
