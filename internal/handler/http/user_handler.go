package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
)

// UserHandler is responsible for handling HTTP requests related to users.
type UserHandler struct {
	UserService usecase.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService usecase.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// CreateUser handles the HTTP request to create a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req domain.User

	// Decode the incoming JSON body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the user using the usecase
	err = h.UserService.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created user id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": req.ID})
}

// GetUser handles the HTTP request to get a user by ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the usecase to get the user by ID
	user, err := h.UserService.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
