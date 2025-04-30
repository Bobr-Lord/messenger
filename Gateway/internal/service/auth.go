package service

import (
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
	"gitlab.com/bobr-lord-messenger/gateway/internal/repository"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req *models.RegisterRequest) (*models.RegisterResponse, error) {
	return s.repo.Auth.Register(req)
}
