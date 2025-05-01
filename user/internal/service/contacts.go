package service

import (
	"gitlab.com/bobr-lord-messenger/user/internal/repository"
)

type ContactsService struct {
	repo *repository.Repository
}

func NewContactsService(repo *repository.Repository) *ContactsService {
	return &ContactsService{repo: repo}
}
