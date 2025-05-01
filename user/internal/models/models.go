package models

import "time"

type GetMeRequest struct {
}
type GetMeResponse struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateMeRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type UpdateMeResponse struct {
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	Users []*UserForGetUsers
}

type UserForGetUsers struct {
	ID        string    `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
