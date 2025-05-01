package repository

import "github.com/jmoiron/sqlx"

type ContactRepository struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) *ContactRepository {
	return &ContactRepository{db: db}
}
