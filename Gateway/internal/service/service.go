package service

import "gitlab.com/bobr-lord-messenger/gateway/internal/repository"

type Websocket interface {
}

type Service struct {
	Websocket Websocket
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Websocket: NewWebsocketService(repo),
	}
}
