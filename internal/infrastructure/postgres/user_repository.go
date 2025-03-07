package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/kafkaphoenix/gotemplate/internal/domain"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var ErrNotFound = errors.New("user not found")

// userRepository implements the UserRepository interface using PostgreSQL.
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of userRepository.
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// CreateUser inserts a new user into the PostgreSQL database.
func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, first_name, last_name, nickname, password, email, country, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Nickname, user.Password,
		user.Email, user.Country, user.CreatedAt, user.UpdatedAt)

	return err
}

// UpdateUser updates an existing user in the database.
func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, nickname = $3, password = $4, email = $5,
	 country = $6, updated_at = $7 WHERE id = $8`
	_, err := r.db.ExecContext(ctx, query, user.FirstName, user.LastName, user.Nickname, user.Password, user.Email,
		user.Country, user.UpdatedAt, user.ID)

	return err
}

// DeleteUser removes a user by ID from the database.
func (r *userRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)

	return err
}

// GetUserByID retrieves a user by their ID.
func (r *userRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	query := `SELECT id, first_name, last_name, nickname, password, email, country, created_at,
	 updated_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, userID)

	var user domain.User

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Nickname, &user.Password, &user.Email,
		&user.Country, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}

// GetUsers retrieves users filtered by country, with pagination.
func (r *userRepository) GetUsers(ctx context.Context, country string, limit, offset int) ([]*domain.User, error) {
	query := `SELECT id, first_name, last_name, nickname, password, email, country, created_at, updated_at 
			  FROM users WHERE country = $1 LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, country, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Nickname, &user.Password, &user.Email,
			&user.Country, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
