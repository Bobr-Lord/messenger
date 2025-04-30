package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	fmt.Println(322332)
	query := "INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.db.QueryRow(query, req.Username, req.Password, req.Email).Scan(&id)
	if err != nil {
		// Преобразуем к pq.Error
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" && pgErr.Constraint == "users_username_key" {
				return "", errors.NewHttpError(http.StatusConflict, "username already in use")
			}
		}
		// Фолбэк
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
