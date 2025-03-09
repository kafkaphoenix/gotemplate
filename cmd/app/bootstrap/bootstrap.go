package bootstrap

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"

	"github.com/kafkaphoenix/gotemplate/internal/delivery/http_server"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/postgres"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	initLogger()

	dbPool, err := initDB(cfg)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	userRepo := postgres.NewUserRepo(dbPool)
	userService := usecase.NewUserService(userRepo)
	userHandler := http_server.NewUserHandler(userService)

	server := http_server.NewHTTPServer(log.Logger, userHandler)
	return server.Start(cfg)
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func initDB(cfg *config.AppConfig) (*pgxpool.Pool, error) {
	// Create dsn for database connection
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Pass,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.SSL,
	)

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return dbPool, nil
}
