package service

import (
	customErr "gitlab.com/bobr-lord-messenger/auth/internal/errors"
	"gitlab.com/bobr-lord-messenger/auth/internal/jwt"
	"gitlab.com/bobr-lord-messenger/auth/internal/models"
	"gitlab.com/bobr-lord-messenger/auth/internal/repository"
	"net/http"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req *models.RegisterRequest) (string, error) {
	id, err := s.repo.PostgresAuth.Register(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *AuthService) Login(req *models.LoginRequest) (string, error) {
	id, err := s.repo.PostgresAuth.Login(req)
	if err != nil {
		return "", err
	}
	token, err := jwt.CreateJWT(id)
	if err != nil {
		return "", customErr.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	return token, nil
}
