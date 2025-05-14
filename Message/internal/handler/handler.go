package handler

import (
	"github.com/gin-gonic/gin"
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
	mess := router.Group("/message")
	{
		mess.GET("/:chat_id", h.GetMessagesFromChat)
		mess.POST("/mark-read", h.ChangeMessageStatus)
	}

	return router
}
