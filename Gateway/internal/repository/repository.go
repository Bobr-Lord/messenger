package repository

import (
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	"gitlab.com/bobr-lord-messenger/gateway/internal/models"
)

type Auth interface {
	Register(req *models.RegisterRequest) (*models.RegisterResponse, error)
	Login(req *models.LoginRequest) (*models.LoginResponse, error)
}

type User interface {
	GetMe(id string) (*models.GetMeResponse, error)
	UpdateMe(id string, req *models.UpdateMeRequest) error
	GetUsers(id string) (*models.GetUsersResponse, error)
	GetUserById(req *models.GetUserByIdRequest) (*models.GetUserByIdResponse, error)
	GetUserByUsername(req *models.GetUserByUsernameRequest) (*models.GetUserByUsernameResponse, error)
}

type Chat interface {
	CreatePrivateChat(userID string, in *models.CreatePrivateChatRequest) (*models.CreatePrivateChatResponse, error)
	CreatePublicChat(userID string, in *models.CreatePublicChatRequest) (*models.CreatePublicChatResponse, error)
	GetMeChats(userID string) (*models.GetMeChatsResponse, error)
	GetChatUsers(UserID string, chatID string) (*models.GetChatUsersResponse, error)
}
type Repository struct {
	Auth Auth
	User User
	Chat Chat
}

func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Auth: NewAuthRepository(cfg),
		User: NewUserRepository(cfg),
		Chat: NewChatRepository(cfg),
	}
}
