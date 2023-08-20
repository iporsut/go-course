package mrello

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IsPasswordMatch(passwordHash string) bool {
	return u.PasswordHash == passwordHash
}

type UserRepository interface {
	CreateUser(ctx context.Context, email string, passwordHash string) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
