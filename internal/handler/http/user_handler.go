package handler

import (
	"encoding/json"
	"net/http"
	"github.com/kafkaphoenix/gotemplate/internal/usecase"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
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
	var userRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Nickname  string `json:"nickname"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		Country   string `json:"country"`
	}

	// Decode the request body into the user request struct
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Call the usecase to create the user
	user, err := h.UserService.CreateUser(r.Context(), userRequest.FirstName, userRequest.LastName, userRequest.Nickname, userRequest.Password, userRequest.Email, userRequest.Country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created user data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
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
