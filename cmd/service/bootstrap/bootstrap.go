package bootstrap

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/spf13/viper"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	handler "github.com/kafkaphoenix/gotemplate/internal/handler/http"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/postgres"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
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

	defer func() {
		if err := closeDB(db); err != nil {
			log.Error().Err(err).Msg("Error closing the database connection")
		}
	}()

	userRepo := postgres.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return startHTTPServer(log.Logger, *userHandler)
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// initDB connects to the PostgreSQL database and runs migrations.
func initDB() (*gorm.DB, error) {
	dbEndpoint := viper.GetString(config.DBEndpointKey)

	// Connect to the database using GORM.
	db, err := gorm.Open(pg.Open(dbEndpoint), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run auto migration to create the `users` table if it doesn't exist.
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func closeDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Gracefully close the underlying sql.DB connection.
	return sqlDB.Close()
}

func startHTTPServer(logger zerolog.Logger, userHandler handler.UserHandler) error {
	router := mux.NewRouter()

	// Define user-related routes
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	url := viper.GetString(config.AppURLKey)
	logger.Info().Msgf("Starting server on %s", url)

	s := &http.Server{
		Addr:              url,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return s.ListenAndServe()
}
