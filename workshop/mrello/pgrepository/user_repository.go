package pgrepository

import (
	"context"
	"database/sql"
	"mrello"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

// FindUserByEmail returns a user with the given email.
func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*mrello.User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT user_id, email, password_hash, created_at, updated_at FROM users WHERE email = $1`, email)

	var user mrello.User

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, mrello.WrapErr(err, mrello.ErrCodeNotFound, "user not found")
		}
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error finding user")
	}

	return &user, nil
}

// CreateUser creates a user with the given email and password hash.
func (r *userRepository) CreateUser(ctx context.Context, email string, passwordHash string) (*mrello.User, error) {
	row := r.db.QueryRowContext(ctx, `INSERT INTO users (user_id, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING user_id,  created_at, updated_at`, uuid.NewString(), email, passwordHash)

	var user mrello.User
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error creating user")
	}
	user.Email = email
	user.PasswordHash = passwordHash

	return &user, nil
}
