package service

import (
	"gitlab.com/bobr-lord-messenger/chat/internal/models"
	"gitlab.com/bobr-lord-messenger/chat/internal/repository"
)

type Chat interface {
	CreatePrivateChat(userID string, req *models.CreatePrivateChatRequest) (string, error)
}
type Service struct {
	Chat Chat
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Chat: NewChatService(repo),
	}
}
