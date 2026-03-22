package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"sens.ai-backend/internal/api"
	"sens.ai-backend/internal/db"
	"sens.ai-backend/internal/logger"
	"sens.ai-backend/internal/worker"
)

// @title Sens.ai Backend API
// @description This is the API for the Sens.ai backend.
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:4000
// @BasePath /api

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// Initialize logger
	logger.Init()

	// Initialize DB
	ctx := context.Background()
	if err := db.Init(ctx); err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	// Start Worker
	go func() {
		worker.Start()
	}()

	// Start API Server (Blocking)
	api.Start(host, port)
}
