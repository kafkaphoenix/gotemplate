package http_server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/kafkaphoenix/gotemplate/internal/repository/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type HTTPServer struct {
	logger *slog.Logger
	server *echo.Echo
}

func New(l *slog.Logger) *HTTPServer {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Register swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.GET("/health", healthCheck)

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

type StatusResponse struct {
	Status string `json:"status" example:"ok"`
}

// healthCheck returns the health status of the service.
//
// @Summary      Health Check
// @Description  Returns the health status of the service.
// @ID           health-check
// @Tags         health
// @Produce      json
// @Success      200  {object} StatusResponse "ok"
// @Router       /health [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, StatusResponse{Status: "ok"})
}
