package service

import "gitlab.com/bobr-lord-messenger/auth/internal/repository"

type Auth interface {
}
type Service struct {
	Auth Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
