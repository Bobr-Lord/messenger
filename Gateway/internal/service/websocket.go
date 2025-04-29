package service

import "gitlab.com/bobr-lord-messenger/gateway/internal/repository"

type WebsocketService struct {
	repo *repository.Repository
}

func NewWebsocketService(repo *repository.Repository) *WebsocketService {
	return &WebsocketService{
		repo: repo,
	}
}
