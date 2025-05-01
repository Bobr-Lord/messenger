package repository

import "github.com/jmoiron/sqlx"

type User interface {
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
