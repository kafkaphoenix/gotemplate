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

// CreateUser handles the creation of a new user.
func (s *userService) CreateUser(ctx context.Context, params CreateUserParams) (*domain.User, error) {
	userID := uuid.New()
	now := time.Now()

	user := &domain.User{
		ID:        userID.String(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Nickname:  params.Nickname,
		Password:  params.Password, // In a real-world app, you would hash the password
		Email:     params.Email,
		Country:   params.Country,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID.
func (s *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}
