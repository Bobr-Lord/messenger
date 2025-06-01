package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/message/internal/errors"
	"gitlab.com/bobr-lord-messenger/message/internal/models"
	"net/http"
)

type MessageRepo struct {
	db *sqlx.DB
}

func NewMessageRepo(db *sqlx.DB) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Save(msg *models.Message) (string, error) {
	query := "INSERT INTO messages (chat_id, sender_id, content) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.db.QueryRow(query, msg.ChatID, msg.SenderID, msg.Content).Scan(&id)
	if err != nil {
		return id, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to save message: %v", err))
	}
	return id, nil
}

func (r *MessageRepo) GetUserMessages(userID string) ([]*models.Message, error) {
	var messages []*models.Message

	query := `
		SELECT m.* 
		FROM messages m 
		JOIN chat_participants cp ON m.chat_id = cp.chat_id 
		WHERE cp.user_id = $1 
		ORDER BY m.created_at DESC
	`

	err := r.db.Select(&messages, query, userID)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to get user messages: %v", err))
	}
	return messages, nil
}

func (r *MessageRepo) GetMessagesByChatID(chatID string) ([]*models.Message, error) {
	var messages []*models.Message
	query := "SELECT * FROM messages WHERE chat_id = $1"
	err := r.db.Select(&messages, query, chatID)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to get user messages: %v", err))
	}
	return messages, nil
}

func (r *MessageRepo) UsersSendMess(chatID string, senderID string) (*[]string, error) {
	var users []string
	query := "SELECT user_id FROM chat_participants WHERE chat_id = $1 AND user_id != $2"
	err := r.db.Select(&users, query, chatID, senderID)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to get user messages: %v", err))
	}
	return &users, nil
}
