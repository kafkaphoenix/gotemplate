package http_server

import (
	"fmt"
	"log/slog"

	"github.com/kafkaphoenix/gotemplate/internal/infrastructure/config"
	"github.com/labstack/echo/v4"
)

type HTTPServer struct {
	logger *slog.Logger
	server *echo.Echo
}

func New(l *slog.Logger) *HTTPServer {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	return &HTTPServer{
		logger: l,
		server: e,
	}
}

func (s *HTTPServer) RegisterRoutes(handlers ...func(e *echo.Echo)) {
	for _, h := range handlers {
		h(s.server)
	}
}

// Start initiate the HTTP server and listens for incoming requests.
func (s *HTTPServer) Start(cfg *config.AppConfig) error {
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	s.logger.Info("Starting server", slog.Int("port", cfg.App.Port))
	return s.server.Start(addr)
}
