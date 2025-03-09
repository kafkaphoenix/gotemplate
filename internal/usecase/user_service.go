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
	GetByID(ctx context.Context, uid uuid.UUID) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByNickname(ctx context.Context, nickname string) (*domain.User, error)
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

func (s *userService) GetByID(ctx context.Context, uid uuid.UUID) (*domain.User, error) {
	return s.repo.GetByID(ctx, uid)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *userService) GetByNickname(ctx context.Context, nickname string) (*domain.User, error) {
	return s.repo.GetByNickname(ctx, nickname)
}

func (s *userService) List(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	return s.repo.List(ctx, country, limit, offset)
}
