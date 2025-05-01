package service

import (
	"gitlab.com/bobr-lord-messenger/user/internal/models"
	"gitlab.com/bobr-lord-messenger/user/internal/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetMe(id string) (*models.GetMeResponse, error) {
	return s.repo.User.GetMe(id)
}
