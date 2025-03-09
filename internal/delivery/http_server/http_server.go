package http_server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kafkaphoenix/gotemplate/internal/delivery"
	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
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
func (s *HTTPServer) Start() error {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", s.handler.Create).Methods("POST")
	router.HandleFunc("/users", s.handler.List).Methods("GET")
	router.HandleFunc("/users/{id}", s.handler.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", s.handler.Delete).Methods("DELETE")

	appPort := viper.GetInt(config.AppPortKey)
	s.logger.Info().Msgf("Starting server on :%d", appPort)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", appPort),
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return httpServer.ListenAndServe()
}
