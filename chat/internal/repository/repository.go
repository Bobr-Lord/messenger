package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/chat/internal/models"
)

type Chat interface {
	CreatePrivateChat(userID string, req *models.CreatePrivateChatRequest) (string, error)
	CreatePublicChat(userID string, req *models.CreatePublicChatRequest) (string, error)
	GetChats(id string) ([]string, error)
	GetUsersChat(chatID string) ([]string, error)
}
type Repository struct {
	Chat Chat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Chat: NewChatRepository(db),
	}
}
