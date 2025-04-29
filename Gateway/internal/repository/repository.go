package repository

type Websocket interface {
}

type Repository struct {
	Websocket Websocket
}

func NewRepository() *Repository {
	return &Repository{
		Websocket: NewWebsocketRepo(),
	}
}
