package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
)

type CreateUserParams struct {
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
}

// UserService defines the interface for user-related business logic.
type UserService interface {
	CreateUser(ctx context.Context, params CreateUserParams) (*domain.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
}

// userService implements the UserService interface.
type userService struct {
	repo domain.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo domain.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser creates a new user in the system
func (s *userService) CreateUser(user *domain.User) error {
    // Generate a new UUID for the user
    user.ID = uuid.New().String()

    // Call the repository layer to save the user
    return s.userRepo.CreateUser(user)
}

// GetUserByID retrieves a user by ID.
func (s *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}
