package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/bobr-lord-messenger/chat/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(cfg *config.Config, hand *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           cfg.ServerHost + ":" + cfg.ServerPort,
		Handler:        hand,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
