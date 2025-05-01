package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/bobr-lord-messenger/user/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(cfg *config.Config, h *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           cfg.ServiceHost + ":" + cfg.ServicePort,
		Handler:        h,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
