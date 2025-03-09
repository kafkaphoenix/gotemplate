package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
)

type UserService interface {
	Create(ctx context.Context, u *domain.User) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, uid uuid.UUID) error
	List(ctx context.Context, country string, limit, offset int) ([]*domain.User, error)
}

type userService struct {
	repo domain.UserRepo
}

func NewUserService(r domain.UserRepo) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	return s.repo.Create(ctx, u)
}

func (s *userService) Update(ctx context.Context, u *domain.User) error {
	return s.repo.Update(ctx, u)
}

func (s *userService) Delete(ctx context.Context, uid uuid.UUID) error {
	return s.repo.Delete(ctx, uid)
}

func (s *userService) List(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	return s.repo.List(ctx, country, limit, offset)
}
