package entities

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	FirstName string    `json:"first_name" example:"John" validate:"required"`
	LastName  string    `json:"last_name" example:"Doe" validate:"required"`
	Nickname  string    `json:"nickname" example:"johndoe" validate:"required"`
	Password  string    `json:"password" example:"password" validate:"required" min:"8"`
	Email     string    `json:"email" example:"johndoe@aexample.com" validate:"required,email"`
	Country   string    `json:"country" example:"UK" validate:"required"`
	CreatedAt time.Time `json:"created_at" example:"2021-07-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-07-01T00:00:00Z"`
}

type UserRepo interface {
	Create(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, uid uuid.UUID) error
	List(ctx context.Context, country string, limit, offset int) ([]*User, error)
}

// LogValue implements the slog.Valuer interface
// to provide a loggable value for the user entity preventing
// user's fields that should not be logged from being exposed.
func (u User) LogValue() slog.Value {
	return slog.StringValue(u.ID.String())
}
