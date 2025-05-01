package service

import (
	"gitlab.com/bobr-lord-messenger/user/internal/hash"
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

func (s *UserService) UpdateMe(id string, req *models.UpdateMeRequest) error {
	if req.Password != "" {
		hashPass, err := hash.HashPass(req.Password)
		if err != nil {
			return err
		}
		req.Password = hashPass
	}
	return s.repo.User.UpdateMe(id, req)
}
