package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/bobr-lord-messenger/user/internal/errors"
	"gitlab.com/bobr-lord-messenger/user/internal/models"
	"net/http"
	"strings"
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

func (h *UserRepository) UpdateMe(id string, req *models.UpdateMeRequest) error {
	setParts := []string{}
	args := []interface{}{}
	argIdx := 1

	if req.Username != "" {
		setParts = append(setParts, fmt.Sprintf("username = $%d", argIdx))
		args = append(args, req.Username)
		argIdx++
	}
	if req.Email != "" {
		setParts = append(setParts, fmt.Sprintf("email = $%d", argIdx))
		args = append(args, req.Email)
		argIdx++
	}
	if req.Password != "" {
		setParts = append(setParts, fmt.Sprintf("password_hash = $%d", argIdx))
		args = append(args, req.Password)
		argIdx++
	}

	if len(setParts) == 0 {
		return nil
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(setParts, ", "), argIdx)
	args = append(args, id)

	_, err := h.db.Exec(query, args...)
	if err != nil {
		return errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %s", err))
	}

	return nil
}
