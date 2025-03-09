package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
)

type userRepo struct {
	storage *Storage
}

func NewUserRepo(s *Storage) domain.UserRepo {
	return &userRepo{
		storage: s,
	}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	// check if nickname already exists
	_, err := r.storage.Queries.GetUserByNickname(ctx, user.Nickname)
	if err == nil {
		return nil, fmt.Errorf("nickname '%s' already taken", user.Nickname)
	}

	// check if email already exists
	_, err = r.storage.Queries.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return nil, fmt.Errorf("email '%s' already taken", user.Email)
	}

	params := CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Country:   user.Country,
	}
	row, err := r.storage.Queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	user.ID = row.ID
	user.CreatedAt = row.CreatedAt.Time
	user.UpdatedAt = row.UpdatedAt.Time
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	params := UpdateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	}
	return r.storage.Queries.UpdateUser(ctx, params)
}

func (r *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.storage.Queries.DeleteUser(ctx, id)
}

func (r *userRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := r.storage.Queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Nickname:  row.Nickname,
		Password:  row.Password,
		Email:     row.Email,
		Country:   row.Country,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row, err := r.storage.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Nickname:  row.Nickname,
		Password:  row.Password,
		Email:     row.Email,
		Country:   row.Country,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *userRepo) GetByNickname(ctx context.Context, nickname string) (*domain.User, error) {
	row, err := r.storage.Queries.GetUserByNickname(ctx, nickname)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Nickname:  row.Nickname,
		Password:  row.Password,
		Email:     row.Email,
		Country:   row.Country,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}, nil
}

func (r *userRepo) List(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	params := ListUsersParams{
		Country: country,
		Limit:   int32(limit),
		Offset:  int32(offset),
	}
	rows, err := r.storage.Queries.ListUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	users := []*domain.User{}
	for _, row := range rows {
		users = append(users, &domain.User{
			ID:        row.ID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
			Nickname:  row.Nickname,
			Password:  row.Password,
			Email:     row.Email,
			Country:   row.Country,
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		})
	}

	return users, nil
}
