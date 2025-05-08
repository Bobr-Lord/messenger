package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/chat/internal/errors"
	"gitlab.com/bobr-lord-messenger/chat/internal/models"
	"net/http"
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) CreatePrivateChat(userID string, req *models.CreatePrivateChatRequest) (string, error) {
	tx := r.db.MustBegin()

	queryAddChat := "INSERT INTO chats (is_private) VALUES ($1) RETURNING id"
	var chatID string
	err := tx.QueryRow(queryAddChat, true).Scan(&chatID)
	if err != nil {
		errTX := tx.Rollback()
		if errTX != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, errTX.Error())
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, "could not create chat")
	}
	queryAddParticipant := "INSERT INTO chat_participants (chat_id, user_id) VALUES ($1, $2)"
	_, err = tx.Exec(queryAddParticipant, chatID, userID)
	if err != nil {
		errTX := tx.Rollback()
		if errTX != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, errTX.Error())
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, "could not add creator to chat")
	}

	_, err = tx.Exec(queryAddParticipant, chatID, req.FriendID)
	if err != nil {
		errTX := tx.Rollback()
		if errTX != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, errTX.Error())
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, "could not add friend to chat")
	}

	err = tx.Commit()
	if err != nil {
		return "", errors.NewCustomError(http.StatusInternalServerError, "could not commit transaction")
	}

	return chatID, nil

}
