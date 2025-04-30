package repository

import (
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
)

type Auth interface {
	Register(req *models.RegisterRequest) (*models.RegisterResponse, error)
}

type Repository struct {
	Auth Auth
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Auth: NewAuthRepository(cfg),
	}
}
