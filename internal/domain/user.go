package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User represents the user entity in the domain.
type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// UserRepository defines the interface for interacting with the User data.
type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetUsers(ctx context.Context, country string, limit, offset int) ([]*User, error)
}
