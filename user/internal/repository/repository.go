package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/user/internal/models"
)

type User interface {
	GetMe(id string) (*models.GetMeResponse, error)
}

type Contacts interface {
}

type Repository struct {
	User     User
	Contacts Contacts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserRepository(db),
		Contacts: NewContactRepository(db),
	}
}
