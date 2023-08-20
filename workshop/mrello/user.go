package mrello

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IsPasswordMatch(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}

func (u *User) NewCard(title string, description string) *Card {
	return &Card{
		ID:              uuid.New(),
		Title:           title,
		Description:     description,
		Column:          "todo",
		CreatedByUserID: u.ID,
		UpdatedByUserID: u.ID,
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, email string, passwordHash string) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
