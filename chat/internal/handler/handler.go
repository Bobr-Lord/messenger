package handler

import (
	"github.com/gin-gonic/gin"
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
	return router
}
