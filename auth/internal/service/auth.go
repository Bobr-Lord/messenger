package service

import "gitlab.com/bobr-lord-messenger/auth/internal/repository"

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}
