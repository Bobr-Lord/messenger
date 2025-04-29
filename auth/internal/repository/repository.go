package repository

import "github.com/jmoiron/sqlx"

type PostgresAuth interface {
}
type Repository struct {
	PostgresAuth PostgresAuth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PostgresAuth: NewPostgresAuthRepository(db),
	}
}
