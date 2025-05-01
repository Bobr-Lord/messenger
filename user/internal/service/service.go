package service

import (
	"gitlab.com/bobr-lord-messenger/user/internal/models"
	"gitlab.com/bobr-lord-messenger/user/internal/repository"
)

type User interface {
	GetMe(id string) (*models.GetMeResponse, error)
}

type Contacts interface {
}
type Service struct {
	User     User
	Contacts Contacts
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:     NewUserService(repo),
		Contacts: NewContactsService(repo),
	}
}
