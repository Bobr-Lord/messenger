package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/auth/internal/errors"
	"gitlab.com/bobr-lord-messenger/auth/internal/models"
	"net/http"
)

type PostgresAuthRepository struct {
	db *sqlx.DB
}

func NewPostgresAuthRepository(db *sqlx.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{
		db: db,
	}
}

func (r *PostgresAuthRepository) Register(req *models.RegisterRequest) (string, error) {
	query := "INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.db.QueryRow(query, req.Username, req.Password, req.Email).Scan(&id)
	if err != nil {
		return "", errors.NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return id, err
}

func (r *PostgresAuthRepository) Login(req *models.LoginRequest) (string, error) {
	query := "SELECT id FROM users WHERE username = $1 AND password_hash = $2"
	var id string
	err := r.db.QueryRow(query, req.Username, req.Password).Scan(&id)
	if err != nil {
		return "", errors.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	return id, err
}
