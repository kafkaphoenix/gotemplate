package usecase

import (
	"context"
	"time"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"github.com/google/uuid"
)

// UserService defines the interface for user-related business logic.
type UserService interface {
	CreateUser(ctx context.Context, firstName, lastName, nickname, password, email, country string) (*domain.User, error)
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
func (s *userService) CreateUser(ctx context.Context, firstName, lastName, nickname, password, email, country string) (*domain.User, error) {
	userID := uuid.New()
	now := time.Now()

	user := &domain.User{
		ID:        userID.String(),
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Password:  password,  // In a real-world app, you would hash the password
		Email:     email,
		Country:   country,
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
