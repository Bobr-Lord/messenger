package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("register", func(ctx *gin.Context) {})
		auth.POST("login", func(ctx *gin.Context) {})
	}
	return r
}
