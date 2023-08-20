package mrello

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Column          string    `json:"column"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedByUserID uuid.UUID `json:"created_by_user_id"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedByUserID uuid.UUID `json:"updated_by_user_id"`
}

func (c *Card) MoveToColumn(column string, moveBy uuid.UUID) {
	c.Column = column
	c.UpdatedByUserID = moveBy
}

func (c *Card) Update(title string, description string, updateBy uuid.UUID) {
	c.Title = title
	c.Description = description
	c.UpdatedByUserID = updateBy
}

type CardRepository interface {
	CreateCard(context.Context, *Card) (*Card, error)
	FindCardByID(ctx context.Context, id uuid.UUID) (*Card, error)
	SaveCard(ctx context.Context, card *Card) (*Card, error)
	GetAllCards(ctx context.Context) ([]*Card, error)
}
