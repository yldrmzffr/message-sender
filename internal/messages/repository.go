package messages

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, message *CreateMessageRequest) (*Message, error) {
	query := `INSERT INTO messages (recipient, content) VALUES ($1, $2) RETURNING id, recipient, content, status`

	newMessage := &Message{}
	err := r.db.QueryRow(ctx, query, message.Recipient, message.Content).Scan(&newMessage.ID, &newMessage.Recipient, &newMessage.Content, &newMessage.Status)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}
