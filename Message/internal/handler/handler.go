package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/bobr-lord-messenger/message/internal/handler/middleware"
	"gitlab.com/bobr-lord-messenger/message/internal/service"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.RequestMiddleware())
	mess := router.Group("/message")
	{
		mess.GET("/user/:user_id", h.GetUserMessage)
		mess.GET("/chat/:chat_id", h.GetMessagesByChatID)
	}

	return router
}
