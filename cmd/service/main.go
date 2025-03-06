package main

import (
	"github.com/kafkaphoenix/gotemplate/internal/repository/config"

	"github.com/rs/zerolog"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	"github.com/kafkaphoenix/gotemplate/internal/repository"
	"github.com/kafkaphoenix/gotemplate/internal/handler"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	config.Init()

	// Setup the database connection using the loaded config
	db, err := postgres.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Setup the repository
	userRepo := postgres.NewUserRepository(db)

	// Setup usecase/service layer
	userService := usecase.NewUserService(userRepo)

	// Setup the HTTP router and handlers
	router := mux.NewRouter()

	// Initialize user handler
	userHandler := handler.NewUserHandler(userService)

	// Define routes
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	// Add other routes...

	// Start the HTTP server
	port := viper.GetString(config.APP_PORT_KEY)
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
