package bootstrap

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/spf13/viper"

	"github.com/kafkaphoenix/gotemplate/internal/delivery/http_server"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/postgres"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run() error {
	config.Init()

	initLogger()

	dbPool, err := initDB()
	if err != nil {
		return err
	}
	defer dbPool.Close()

	userRepo := postgres.NewUserRepo(dbPool)
	userService := usecase.NewUserService(userRepo)
	userHandler := http_server.NewUserHandler(userService)

	server := http_server.NewHTTPServer(log.Logger, userHandler)
	return server.Start()
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func initDB() (*pgxpool.Pool, error) {
	// Create dsn for database connection
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString(config.DBHostKey),
		viper.GetInt(config.DBPortKey),
		viper.GetString(config.DBUserKey),
		viper.GetString(config.DBPassKey),
		viper.GetString(config.DBNameKey),
		viper.GetString(config.DBSSLKey),
	)

	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return dbPool, nil
}
