package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "sens.ai-backend/docs"
	"sens.ai-backend/internal/logger"
	"sens.ai-backend/internal/middleware"
	v1 "sens.ai-backend/internal/routes/v1"
)

// @title Sens.ai Backend API
// @description This is the API for the Sens.ai backend.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		// Use fmt.Println before logger is initialized
		fmt.Println("No .env file found")
	}

	// Initialize logger
	logger.Init()

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

	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1Group := api.Group("/v1")
	v1.SetupRoutes(v1Group)

	// Start server
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
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
