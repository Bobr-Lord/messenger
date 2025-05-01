package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/user/internal/config"
	"gitlab.com/bobr-lord-messenger/user/internal/handler"
	"gitlab.com/bobr-lord-messenger/user/internal/repository"
	"gitlab.com/bobr-lord-messenger/user/internal/server"
	"gitlab.com/bobr-lord-messenger/user/internal/service"
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
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)
	srv := server.NewServer()
	go func() {
		if err := srv.Run(cfg, h.InitRouter()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("Server shutdown error: %v", err)
	} else {
		logrus.Info("Server shutdown complete.")
	}

	if err := db.DB.Close(); err != nil {
		logrus.Errorf("Database close error: %v", err)
	} else {
		logrus.Info("Database connection closed.")
	}

}
