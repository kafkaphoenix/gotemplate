package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	"gorm.io/gorm"
)

// UserRepository is the GORM implementation of the UserRepository interface.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	// GORM's Create method is used to insert the user into the database
	return r.db.WithContext(ctx).Create(user).Error
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	// GORM's Save method is used for both insert (if the ID doesn't exist) and update (if it does)
	return r.db.WithContext(ctx).Save(user).Error
}

// DeleteUser deletes a user from the database.
func (r *UserRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	// GORM's Delete method is used to remove the user by their ID
	return r.db.WithContext(ctx).Where("id = ?", userID).Delete(&domain.User{}).Error
}

// GetUserByID retrieves a user by ID.
func (r *UserRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUsers retrieves a list of users, filtered by country and paginated.
func (r *UserRepository) GetUsers(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	var users []*domain.User
	// GORM's Find method retrieves users with pagination
	err := r.db.WithContext(ctx).
		Where("country = ?", country).
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
