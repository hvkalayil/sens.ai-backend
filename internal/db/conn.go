package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"sens.ai-backend/internal/logger"
)

// NewPool creates a new pgx connection pool using the DATABASE_URL env var.
// Example: postgres://postgres:postgres@localhost:5432/sensai?sslmode=disable
func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logger.Logger.Warn().Msg("DATABASE_URL is not set; using default local Postgres URL")
		dsn = "postgres://postgres:postgres@localhost:5432/sensai?sslmode=disable"
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// Optional pool tuning
	cfg.MaxConns = 10
	cfg.MinConns = 1
	cfg.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	logger.Logger.Info().Msg("Connected to Postgres")
	return pool, nil
}

// Pool is the global database connection pool.
var Pool *pgxpool.Pool

// Init initializes the global database connection pool.
func Init(ctx context.Context) error {
	var err error
	Pool, err = NewPool(ctx)
	return err
}

// Close closes the global database connection pool.
func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
