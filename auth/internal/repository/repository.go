package repository

type PostgresAuth interface {
}
type Repository struct {
	PostgresAuth PostgresAuth
}

func NewRepository() *Repository {
	return &Repository{}
}
