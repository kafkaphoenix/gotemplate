package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// UserHandler defines the HTTP handler for managing users.
type UserHandler struct {
	userService usecase.UserService
	logger      zerolog.Logger
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService usecase.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger(),
	}
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	// Decode the incoming JSON request body into the User struct.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the required fields.
	if user.FirstName == "" || user.LastName == "" || user.Nickname == "" || user.Email == "" || user.Country == "" {
		h.logger.Warn().Msg("Missing required fields")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Call the service layer to create the user.
	if err := h.userService.CreateUser(r.Context(), &user); err != nil {
		h.logger.Error().Err(err).Msg("Failed to create user")
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return a success response with the created user data.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser handles retrieving a user by ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the URL.
	vars := mux.Vars(r)
	userID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.logger.Error().Err(err).Msg("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service layer to get the user by ID.
	user, err := h.userService.GetUserByID(r.Context(), userID)
	if err != nil {
		h.logger.Error().Err(err).Msg("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user data.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles retrieving users, possibly with filtering by country and pagination.
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters for pagination and filtering.
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "all" // Default to fetching all countries if not specified
	}
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	// Convert the pagination parameters to integers.
	// Use default values if parameters are not set or are invalid.
	limitInt := 10 // Default limit
	offsetInt := 0 // Default offset

	if limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			limitInt = l
		}
	}

	if offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			offsetInt = o
		}
	}

	// Call the service layer to get users with pagination and filtering.
	users, err := h.userService.GetUsers(r.Context(), country, limitInt, offsetInt)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to fetch users")
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Return the users list.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles updating an existing user.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.logger.Error().Err(err).Msg("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user domain.User
	// Decode the incoming JSON request body into the User struct.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error().Err(err).Msg("Failed to decode request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the ID of the user to update
	user.ID = userID.String()

	// Call the service layer to update the user.
	if err := h.userService.UpdateUser(r.Context(), &user); err != nil {
		h.logger.Error().Err(err).Msg("Failed to update user")
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Return the updated user data.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.logger.Error().Err(err).Msg("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the service layer to delete the user.
	if err := h.userService.DeleteUser(r.Context(), userID); err != nil {
		h.logger.Error().Err(err).Msg("Failed to delete user")
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	// Return a success message.
	w.WriteHeader(http.StatusNoContent)
}

// HealthCheck is used to confirm the API is working.
func (h *UserHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
