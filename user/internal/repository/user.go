package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/user/internal/errors"
	"gitlab.com/bobr-lord-messenger/user/internal/models"
	"net/http"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (h *UserRepository) GetMe(id string) (*models.GetMeResponse, error) {
	query := "SELECT id, username, password_hash, email, created_at, updated_at FROM users WHERE id = $1"
	var user models.GetMeResponse
	err := h.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreateAt,
		&user.UpdateAt,
	)
	fmt.Println(user)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusNotFound, fmt.Sprintf("user not found: %s", err))
	}
	return &user, nil
}
