package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"sens.ai-backend/internal/logger"
)

// RequestLogger is a custom Fiber middleware for structured request/response logging
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Generate request ID
		requestID := uuid.New().String()
		c.Locals("requestID", requestID)

		// Start timer
		start := time.Now()

		// Log incoming request
		logger.Logger.Debug().
			Str("request_id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Str("user_agent", c.Get("User-Agent")).
			Msg("Incoming request")

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Determine log level based on status code
		statusCode := c.Response().StatusCode()
		logEvent := logger.Logger.Info()

		if statusCode >= 500 {
			logEvent = logger.Logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Logger.Warn()
		}

		// Log request completion
		logEvent.
			Str("request_id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", statusCode).
			Dur("duration_ms", duration).
			Str("ip", c.IP()).
			Msg("Request completed")

		return err
	}
}
