package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"` // Unique
	Password  string    `json:"password"` // Hashed password
	Email     string    `json:"email"`    // Unique
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepo interface {
	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, uid uuid.UUID) error
	GetByID(ctx context.Context, uid uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByNickname(ctx context.Context, nickname string) (*User, error)
	List(ctx context.Context, country string, limit, offset int) ([]*User, error)
}
