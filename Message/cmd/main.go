package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/message/internal/config"
	"gitlab.com/bobr-lord-messenger/message/internal/handler"
	"gitlab.com/bobr-lord-messenger/message/internal/repository"
	"gitlab.com/bobr-lord-messenger/message/internal/server"
	"gitlab.com/bobr-lord-messenger/message/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%+v", cfg)
	db, err := repository.NewPostgres(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Successfully connected to postgres")
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	hand := handler.NewHandler(svc)
	srv := server.NewServer()
	go func() {
		if err := srv.Run(hand.InitRouter(), cfg); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()
	logrus.Info("Successfully started server")
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait
	if err := db.Close(); err != nil {
		logrus.Errorf("failed to close DB: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("server shutdown failed: %v", err)
	}
	select {
	case <-ctx.Done():
		logrus.Info("Shutting down")
		os.Exit(0)
	case <-time.After(30 * time.Second):
		logrus.Info("Timed out waiting for server to shutdown")
		os.Exit(1)
	}
}
