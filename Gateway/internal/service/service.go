package service

import "gitlab.com/bobr-lord-messenger/gateway/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
