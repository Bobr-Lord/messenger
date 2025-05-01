package service

import (
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
	"gitlab.com/bobr-lord-messenger/gateway/internal/repository"
)

type Auth interface {
	Register(req *models.RegisterRequest) (*models.RegisterResponse, error)
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
}

type Service struct {
	Auth Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
