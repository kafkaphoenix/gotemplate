package bootstrap

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/spf13/viper"

	handler "github.com/kafkaphoenix/gotemplate/internal/handler/http"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/postgres"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run() error {
	config.Init()

	initLogger()

	db, err := initDB()
	if err != nil {
		return err
	}

	userRepo := postgres.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return startHTTPServer(log.Logger, *userHandler)
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func initDB() (*sql.DB, error) {
	var err error
	dbUser := viper.GetString(config.DbUserKey)
	dbPassword := viper.GetString(config.DbPasswordKey)
	dbPort := viper.GetString(config.DbPortKey)
	dbName := viper.GetString(config.DbNameKey)
	sslMode := viper.GetString(config.DbSSLModeKey)

	connStr := fmt.Sprintf("postgres://%s:%s@postgres:%s/%s?sslmode=%s", dbUser, dbPassword, dbPort, dbName, sslMode)

	db, err := sql.Open(dbName, connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db, nil
}

func startHTTPServer(logger zerolog.Logger, userHandler handler.UserHandler) error {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	port := viper.GetString(config.AppPortKey)
	logger.Debug().Msgf("Starting server on port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
