package http_server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/entities"
	"github.com/kafkaphoenix/gotemplate/internal/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	logger  *slog.Logger
	service usecases.UserService
}

func NewUserHandler(l *slog.Logger, s usecases.UserService) *UserHandler {
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

// Create creates a new user based on the data provided.
//
// @Summary      Create a User
// @Description  Creates a new User based on the data provided.
// @ID           create-user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        FirstName  path  string  true  "First name of the user"
// @Param        LastName   path  string  true  "Last name of the user"
// @Param        Nickname   path  string  true  "Nickname of the user"
// @Param        Password   path  string  true  "Password of the user"
// @Param        Email      path  string  true  "Email of the user"
// @Param        Country    path  string  true  "Country of the user"
// @Success      201  {object} entities.User
// @Failure      400  "Invalid input"
// @Failure      500  "Internal server error"
// @Router       /users [post]
func (h *UserHandler) Create(c echo.Context) error {
	var req entities.User
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

// Update updates a user based on the given ID and the provided data.
//
// @Summary      Update a User
// @Description  Updates a User based on the given ID and the provided data.
// @ID           update-user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id         path  string  true  "ID of the user"
// @Param        FirstName  path  string  true  "First name of the user"
// @Param        LastName   path  string  true  "Last name of the user"
// @Param        Nickname   path  string  true  "Nickname of the user"
// @Param        Password   path  string  true  "Password of the user"
// @Param        Email      path  string  true  "Email of the user"
// @Param        Country    path  string  true  "Country of the user"
// @Success      204  "No content"
// @Failure      400  "Invalid user ID"
// @Failure      400  "Invalid input"
// @Failure      500  "Internal server error"
// @Router       /users/{id} [patch]
func (h *UserHandler) Update(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var req entities.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	req.ID = uid
	if err := h.service.Update(c.Request().Context(), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// Delete removes a User based on the given ID.
//
// @Summary      Delete a User
// @Description  Removes a User based on the given ID.
// @ID           delete-user
// @Produce      json
// @Tags         users
// @Param        id  path  string  true  "ID of the user"
// @Success      204  "No content"
// @Failure      400  "Invalid user ID"
// @Failure      500  "Internal server error"
// @Router       /users/{id} [delete]
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

// List returns a list of Users optionally filtered by country, limit and offset.
//
// @Summary      List Users
// @Description  Returns a list of Users optionally filtered by country, limit and offset.
// @ID           list-users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        country  query  string  false  "Country of the user"
// @Param        limit    query  int     false  "Limit of users to be listed"
// @Param        offset   query  int     false  "Offset of users to be listed"
// @Success      200  {array} entities.User
// @Failure      500  "Internal server error"
// @Router       /users [get]
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
