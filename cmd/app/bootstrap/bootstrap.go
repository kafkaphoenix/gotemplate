package bootstrap

import (
	"fmt"

	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/logger"

	"github.com/kafkaphoenix/gotemplate/internal/delivery/http_server"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/postgres"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logger := logger.Init(cfg.App.LogLevel)

	storage, err := postgres.NewStorage(cfg)
	if err != nil {
		return err
	}
	defer storage.DB.Close()

	// create repo, service per entity in the domain
	userRepo := postgres.NewUserRepo(storage)
	userService := usecase.NewUserService(userRepo)

	switch cfg.App.ServerType {
	case config.ServerTypeGRPC:
		return nil
	case config.ServerTypeHTTP:
		// create handler per entity in the domain
		userHandler := http_server.NewUserHandler(logger, userService)

		server := http_server.New(logger)
		server.RegisterRoutes(userHandler.RegisterRoutes)

		return server.Start(cfg)
	default:
		return fmt.Errorf("unsupported server type: %s", cfg.App.ServerType)
	}
}
