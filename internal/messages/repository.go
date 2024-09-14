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
	query := `INSERT INTO messages (recipient, content) VALUES ($1, $2) RETURNING id, recipient, content, status, created_at`

	newMessage := &Message{}
	err := r.db.QueryRow(ctx, query, message.Recipient, message.Content).Scan(&newMessage.ID, &newMessage.Recipient, &newMessage.Content, &newMessage.Status, &newMessage.CreatedAt)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}

func (r *Repository) GetSentMessages(ctx context.Context) ([]*Message, error) {
	query := `SELECT id, recipient, content, status, created_at, completed_at FROM messages WHERE status = 'sent'`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]*Message, 0)
	for rows.Next() {
		message := &Message{}
		err := rows.Scan(&message.ID, &message.Recipient, &message.Content, &message.Status, &message.CreatedAt, &message.CompletedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil

}
