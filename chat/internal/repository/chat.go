package repository

import (
	"fmt"
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
			return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add participant to chat: %v", errTX))
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add participant to chat: %v", err))
	}
	queryAddParticipant := "INSERT INTO chat_participants (chat_id, user_id) VALUES ($1, $2)"
	_, err = tx.Exec(queryAddParticipant, chatID, userID)
	if err != nil {
		errTX := tx.Rollback()
		if errTX != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add participant to chat: %v", errTX))
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add participant to chat: %v", err))
	}

	_, err = tx.Exec(queryAddParticipant, chatID, req.FriendID)
	if err != nil {
		errTX := tx.Rollback()
		if errTX != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add participant to chat: %v", errTX))
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add creator to chat: %v", err))
	}

	err = tx.Commit()
	if err != nil {
		return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not commit transaction: %v", err))
	}

	return chatID, nil

}

func (h *ChatRepository) CreatePublicChat(userID string, req *models.CreatePublicChatRequest) (string, error) {
	tx := h.db.MustBegin()

	queryAddChat := "INSERT INTO chats (is_private, name) VALUES ($1, $2) RETURNING id"
	var chatID string
	err := tx.QueryRow(queryAddChat, false, req.Name).Scan(&chatID)
	if err != nil {
		errTX := tx.Rollback()
		if errTX != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add chat: %v", errTX))
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not create chat: %v", err))
	}
	req.ParticipantIDs = append(req.ParticipantIDs, userID)
	queryAddUsers := "INSERT INTO chat_participants (chat_id, user_id) VALUES ($1, $2)"
	for _, participantID := range req.ParticipantIDs {
		_, err := tx.Exec(queryAddUsers, chatID, participantID)
		if err != nil {
			errTX := tx.Rollback()
			if errTX != nil {
				return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add participant to chat: %v", errTX))
			}
			return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not add user to chat: %v", err))
		}
	}
	err = tx.Commit()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not commit transaction: %v", err))
		}
		return "", errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not create chat: %v", err))
	}
	return chatID, nil
}
