package repository

import "github.com/jmoiron/sqlx"

type Chat interface {
}
type Repository struct {
	Chat Chat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Chat: NewChatRepository(db),
	}
}
