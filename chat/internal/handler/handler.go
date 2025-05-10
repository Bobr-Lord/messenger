package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/bobr-lord-messenger/chat/docs"
	"gitlab.com/bobr-lord-messenger/chat/internal/middleware"
	"gitlab.com/bobr-lord-messenger/chat/internal/service"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.LoggerMiddleware())
	chats := router.Group("/chat")
	{
		chats.POST("/private", h.CreatePrivateChat)
		chats.POST("/public", h.CreatePublicChat)
		chats.GET("/", h.GetChats)
		chats.GET("/:id", h.GetChatHistory)
		chats.POST("/add", h.AddChat)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
