package worker

import (
	"time"

	"sens.ai-backend/internal/logger"
)

func Start() {
	logger.Logger.Info().Msg("Worker started")

	// Simulate worker loop in a separate goroutine if needed,
	// but since this is called with 'go worker.Start()' from main,
	// we can just run the loop here.
	for {
		// Your worker logic here
		logger.Logger.Info().Msg("Worker processing job...")
		time.Sleep(10 * time.Second)
	}
}
