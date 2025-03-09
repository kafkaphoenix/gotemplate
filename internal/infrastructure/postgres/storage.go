package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
)

type Storage struct {
	Queries *Queries
	DB      *pgxpool.Pool
}

// NewStorage initializes the database connection and repositories.
func NewStorage(cfg *config.AppConfig) (*Storage, error) {
	// Create DSN
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Pass,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.SSL,
	)

	// Initialize DB Pool
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Create shared Storage instance
	return &Storage{
		Queries: New(dbPool),
		DB:      dbPool,
	}, nil
}
