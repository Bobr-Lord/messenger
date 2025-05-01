package models

import "time"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetMeRequest struct {
}
type GetMeResponse struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UpdateMeRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type UpdateMeResponse struct{}

type WSInput struct {
}
