package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/entities"
)

type UserService interface {
	Create(ctx context.Context, u *entities.User) (*entities.User, error)
	Update(ctx context.Context, u *entities.User) error
	Delete(ctx context.Context, uid uuid.UUID) error
	List(ctx context.Context, country string, limit, offset int) ([]*entities.User, error)
}

type userService struct {
	repo entities.UserRepo
}

func NewUserService(r entities.UserRepo) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, u *entities.User) (*entities.User, error) {
	return s.repo.Create(ctx, u)
}

func (s *userService) Update(ctx context.Context, u *entities.User) error {
	return s.repo.Update(ctx, u)
}

func (s *userService) Delete(ctx context.Context, uid uuid.UUID) error {
	return s.repo.Delete(ctx, uid)
}

func (s *userService) List(ctx context.Context, country string, limit, offset int) ([]*entities.User, error) {
	return s.repo.List(ctx, country, limit, offset)
}
