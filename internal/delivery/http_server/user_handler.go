package http_server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	logger  *slog.Logger
	service usecase.UserService
}

func NewUserHandler(l *slog.Logger, s usecase.UserService) *UserHandler {
	return &UserHandler{
		logger:  l,
		service: s,
	}
}

func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/users")
	g.POST("", h.Create)
	g.PATCH("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	g.GET("", h.List)
}

// CreateUser godoc
// @Summary      Create an User
// @Description  Given a User object, it creates a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      201  {object}  domain.User
// @Failure      500  {object}  string
// @Router       /users [post]
func (h *UserHandler) Create(c echo.Context) error {
	var req domain.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	newUser, err := h.service.Create(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// only to showcase the logger no need to log this in production
	h.logger.Debug("User created", slog.String("id", newUser.ID.String()))

	return c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) Update(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var req domain.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	req.ID = uid
	if err := h.service.Update(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := h.service.Delete(c.Request().Context(), uid); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) List(c echo.Context) error {
	country := c.QueryParam("country")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	users, err := h.service.List(c.Request().Context(), country, parseInt(limit, 10), parseInt(offset, 0))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func parseInt(value string, defaultVal int) int {
	n, err := strconv.Atoi(value)
	if err != nil || n < 0 {
		return defaultVal
	}
	return n
}
