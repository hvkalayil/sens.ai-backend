package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

// Init initializes the global logger with configuration from environment variables
func Init() {
	// Get log level from environment (default: info)
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevelStr == "" {
		logLevelStr = "info"
	}

	// Parse log level
	var logLevel zerolog.Level
	switch logLevelStr {
	case "trace":
		logLevel = zerolog.TraceLevel
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn", "warning":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "fatal":
		logLevel = zerolog.FatalLevel
	case "panic":
		logLevel = zerolog.PanicLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	// Get log format from environment (default: console)
	logFormat := strings.ToLower(os.Getenv("LOG_FORMAT"))
	if logFormat == "" {
		logFormat = "console"
	}

	// Configure output format
	var output io.Writer = os.Stdout
	var logger zerolog.Logger

	if logFormat == "console" {
		// Pretty console output for development - cleaner and more readable
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.Kitchen, // Simpler time format (3:04PM)
			NoColor:    false,        // Keep colors for better readability
			PartsOrder: []string{
				zerolog.TimestampFieldName,
				zerolog.LevelFieldName,
				zerolog.MessageFieldName,
			},
		}

		// Create logger WITHOUT caller for cleaner console output
		logger = zerolog.New(output).
			Level(logLevel).
			With().
			Timestamp().
			Logger()
	} else {
		// JSON format - include all details including caller
		logger = zerolog.New(output).
			Level(logLevel).
			With().
			Timestamp().
			Caller().
			Logger()
	}

	Logger = logger

	// Set global logger
	log.Logger = Logger

	Logger.Info().
		Str("level", logLevelStr).
		Str("format", logFormat).
		Msg("Logger initialized")
}

// GetLogger returns the global logger instance
func GetLogger() *zerolog.Logger {
	return &Logger
}
