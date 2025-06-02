package service

import (
	"gitlab.com/bobr-lord-messenger/chat/internal/models"
	"gitlab.com/bobr-lord-messenger/chat/internal/repository"
)

type Chat interface {
	CreatePrivateChat(userID string, req *models.CreatePrivateChatRequest) (string, error)
	CreatePublicChat(userID string, req *models.CreatePublicChatRequest) (string, error)
	GetChats(id string) ([]string, error)
	GetUsersChat(chatID string) ([]string, error)
}
type Service struct {
	Chat Chat
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Chat: NewChatService(repo),
	}
}
