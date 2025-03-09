package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
)

type userRepo struct {
	queries *sqlc.Queries
	db      *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) domain.UserRepo {
	queries := sqlc.New(db)
	return &userRepo{
		queries: queries,
		db:      db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	row, err := r.queries.Create(ctx, user.FirstName, user.LastName, user.Nickname, user.Password, user.Email, user.Country)
	if err != nil {
		return nil, err
	}
	user.ID = row.ID
	user.CreatedAt = row.CreatedAt
	user.UpdatedAt = row.UpdatedAt
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	_, err := r.queries.Update(ctx, user.FirstName, user.LastName, user.Nickname, user.Password, user.Email, user.Country, user.ID)
	return err
}

func (r *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.queries.Delete(ctx, id)
	return err
}

func (r *userRepo) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := r.queries.Get(ctx, id)
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
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (r *userRepo) List(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	rows, err := r.queries.List(ctx, country, limit, offset)
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
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return users, nil
}
