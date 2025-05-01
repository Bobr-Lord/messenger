package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/bobr-lord-messenger/user/internal/middleware"
	"gitlab.com/bobr-lord-messenger/user/internal/service"
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
	r := gin.New()
	r.Use(middleware.LoggerMiddleware())
	r.GET("/me", h.GetMe)
	r.PUT("/me", h.UpdateMe)
	r.GET("/users", h.GetUsers)
	user := r.Group("/user")
	{
		user.GET("/id", h.GetUserById)
		user.GET("/name", h.GetUserByName)
	}

	return r
}
