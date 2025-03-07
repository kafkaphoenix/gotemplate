package bootstrap

import (
	"database/sql"
	"net/http"
	"time"

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
	if err := config.Init("config.yaml"); err != nil {
		return err
	}

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

	dbEndpoint := viper.GetString(config.DBEndpointKey)
	dbName := viper.GetString(config.DBNameKey)

	db, err := sql.Open(dbName, dbEndpoint)
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

	port := viper.GetString(config.AppURLKey)
	logger.Debug().Msgf("Starting server on port %s", port)

	s := &http.Server{
		Addr:              port,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return s.ListenAndServe()
}
