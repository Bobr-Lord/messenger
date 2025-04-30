package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/bobr-lord-messenger/auth/internal/service"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
	return r
}
