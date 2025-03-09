package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Nickname  string    `json:"nickname" validate:"required"`
	Password  string    `json:"password" validate:"required,min=8"`
	Email     string    `json:"email" validate:"required,email"`
	Country   string    `json:"country" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepo interface {
	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, uid uuid.UUID) error
	List(ctx context.Context, country string, limit, offset int) ([]*User, error)
}
