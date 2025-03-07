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
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return startHTTPServer(log.Logger, *userHandler)
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// initDB connects to the PostgreSQL database and runs migrations
func initDB() (*gorm.DB, error) {
    dbEndpoint := viper.GetString(config.DBEndpointKey)
    dbName := viper.GetString(config.DBNameKey)
    dsn := "user=" + dbEndpoint + " dbname=" + dbName + " sslmode=disable"

    // Connect to the database using GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Run auto migration to create the `users` table if it doesn't exist
    err = db.AutoMigrate(&domain.User{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

func startHTTPServer(logger zerolog.Logger, userHandler handler.UserHandler) error {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	url := viper.GetString(config.AppURLKey)
	logger.Info().Msgf("Starting server on %s", url)

	s := &http.Server{
		Addr:              url,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return s.ListenAndServe()
}
