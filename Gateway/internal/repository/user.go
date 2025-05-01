package repository

import (
	"encoding/json"
	"fmt"
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	"gitlab.com/bobr-lord-messenger/gateway/internal/errors"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
	"io"
	"net/http"
)

type UserRepository struct {
	cfg *config.Config
}

func NewUserRepository(cfg *config.Config) *UserRepository {
	return &UserRepository{cfg: cfg}
}

func (r *UserRepository) GetMe(id string) (*models.GetMeResponse, error) {
	req, err := http.NewRequest("GET", "http://"+r.cfg.UserServiceHost+":"+r.cfg.UserServicePort+"/me", nil)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not create request: %w", err))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("id", id)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not send request: %w", err))
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not read response: %w", err))
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.NewCustomError(http.StatusNotFound, "user not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewCustomError(resp.StatusCode, string(body))
	}
	var user models.GetMeResponse
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, errors.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("could not unmarshal response: %w", err))
	}
	return &user, nil
}
