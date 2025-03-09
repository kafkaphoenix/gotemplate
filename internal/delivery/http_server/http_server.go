package http_server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kafkaphoenix/gotemplate/internal/delivery"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/rs/zerolog"
)

type HTTPServer struct {
	logger  zerolog.Logger
	handler *UserHandler
}

func NewHTTPServer(logger zerolog.Logger, handler *UserHandler) delivery.Server {
	return &HTTPServer{
		logger:  logger,
		handler: handler,
	}
}

// Start initiate the HTTP server and listens for incoming requests.
func (s *HTTPServer) Start(cfg *config.AppConfig) error {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", s.handler.Create).Methods("POST")
	router.HandleFunc("/users", s.handler.List).Methods("GET")
	router.HandleFunc("/users/{id}", s.handler.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", s.handler.Delete).Methods("DELETE")

	s.logger.Info().Msgf("Starting server on :%d", cfg.App.Port)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.App.Port),
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return httpServer.ListenAndServe()
}
