package repository

import "github.com/jmoiron/sqlx"

type Message interface {
}
type Repository struct {
	Message Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
