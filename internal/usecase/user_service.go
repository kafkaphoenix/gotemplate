package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
)

// UserService defines the interface for user-related business logic.
type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
}

// userService implements the UserService interface.
type userService struct {
	userRepo domain.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// CreateUser creates a new user in the system
func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	// Generate a new UUID for the user
	user.ID = uuid.New().String()

	// Call the repository layer to save the user
	return s.userRepo.CreateUser(ctx, user)
}

// GetUserByID retrieves a user by ID.
func (s *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}
