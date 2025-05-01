package repository

import (
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
)

type Auth interface {
	Register(req *models.RegisterRequest) (*models.RegisterResponse, error)
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
}

type User interface {
	GetMe(id string) (*models.GetMeResponse, error)
	UpdateMe(id string, req *models.UpdateMeRequest) error
}
type Repository struct {
	Auth Auth
	User User
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Auth: NewAuthRepository(cfg),
		User: NewUserRepository(cfg),
	}
}
