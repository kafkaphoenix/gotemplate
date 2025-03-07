package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetUsers(ctx context.Context, country string, limit, offset int) ([]*domain.User, error)
}

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.userRepo.DeleteUser(ctx, userID)
}

func (s *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *userService) GetUsers(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	return s.userRepo.GetUsers(ctx, country, limit, offset)
}
