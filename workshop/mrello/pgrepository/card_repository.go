package pgrepository

import (
	"context"
	"database/sql"
	"mrello"

	"github.com/google/uuid"
)

type cardRepository struct {
	db *sql.DB
}

func NewCardRepository(db *sql.DB) *cardRepository {
	return &cardRepository{
		db: db,
	}
}

// CreateCard creates a card
func (r *cardRepository) CreateCard(ctx context.Context, card *mrello.Card) (*mrello.Card, error) {
	row := r.db.QueryRowContext(ctx, `INSERT INTO cards (id, title, description, "column", created_at, created_by_user_id, updated_at, updated_by_user_id) VALUES ($1, $2, $3, $4, NOW(), $5, NOW(), $6) RETURNING id, created_at, updated_at`, card.ID, card.Title, card.Description, card.Column, card.CreatedByUserID, card.UpdatedByUserID)

	err := row.Scan(&card.ID, &card.CreatedAt, &card.UpdatedAt)
	if err != nil {
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error creating card")
	}

	return card, nil
}

// FindCardByID returns a card with the given id.
func (r *cardRepository) FindCardByID(ctx context.Context, id uuid.UUID) (*mrello.Card, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, title, description, "column", created_at, created_by_user_id, updated_at, updated_by_user_id FROM cards WHERE id = $1`, id)

	var card mrello.Card

	err := row.Scan(&card.ID, &card.Title, &card.Description, &card.Column, &card.CreatedAt, &card.CreatedByUserID, &card.UpdatedAt, &card.UpdatedByUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, mrello.WrapErr(err, mrello.ErrCodeNotFound, "card not found")
		}
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error finding card")
	}

	return &card, nil
}

// SaveCard saves a card
func (r *cardRepository) SaveCard(ctx context.Context, card *mrello.Card) (*mrello.Card, error) {
	row := r.db.QueryRowContext(ctx, `UPDATE cards SET title = $2, description = $3, "column" = $4, updated_at = NOW(), updated_by_user_id = $5 WHERE id = $1 RETURNING updated_at`, card.ID, card.Title, card.Description, card.Column, card.UpdatedByUserID)

	err := row.Scan(&card.UpdatedAt)
	if err != nil {
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error saving card")
	}

	return card, nil
}

// GetAllCards returns all cards
func (r *cardRepository) GetAllCards(ctx context.Context) ([]*mrello.Card, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, title, description, "column", created_at, created_by_user_id, updated_at, updated_by_user_id FROM cards`)
	if err != nil {
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error getting cards")
	}
	defer rows.Close()

	var cards []*mrello.Card
	for rows.Next() {
		var card mrello.Card
		err := rows.Scan(&card.ID, &card.Title, &card.Description, &card.Column, &card.CreatedAt, &card.CreatedByUserID, &card.UpdatedAt, &card.UpdatedByUserID)
		if err != nil {
			return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error getting cards")
		}
		cards = append(cards, &card)
	}

	if err := rows.Err(); err != nil {
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error getting cards")
	}

	return cards, nil
}
