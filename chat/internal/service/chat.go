package service

import "gitlab.com/bobr-lord-messenger/chat/internal/repository"

type ChatService struct {
	repo *repository.Repository
}

func NewChatService(repo *repository.Repository) *ChatService {
	return &ChatService{repo: repo}
}
