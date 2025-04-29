package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "gitlab.com/bobr-lord-messenger/gateway/docs" // тут будет документация
	"gitlab.com/bobr-lord-messenger/gateway/internal/middleware"
	"gitlab.com/bobr-lord-messenger/gateway/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		service: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestMiddleware())
	r.GET("/ws", h.Websocket)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
	user := r.Group("/user")
	{
		user.GET("/me", h.GetMe)
		user.PUT("/me", h.UpdateMe)
		user.GET("/users", h.GetUsers)
		user.POST("/contacts/add", h.AddContacts)
		user.GET("/contacts", h.GetContacts)
		user.GET("/id", h.GetUserByID)
		user.GET("/name", h.GetUserByUsername)
	}

	chat := r.Group("/chat")
	{
		chat.POST("/", h.CreateChat)
		chat.GET("/", h.GetChats)
		chat.GET("/:id", h.GetMessagesFromChat)
	}

	message := r.Group("/message")
	{
		message.PUT("/upd", h.UpdateMessageStatus)
		message.GET("/:id", h.GetUnsentMessages)
	}

	return r
}
