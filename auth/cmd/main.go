package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/auth/internal/config"
	hand "gitlab.com/bobr-lord-messenger/auth/internal/handler"
	"gitlab.com/bobr-lord-messenger/auth/internal/jwtutil"
	"gitlab.com/bobr-lord-messenger/auth/internal/repository"
	"gitlab.com/bobr-lord-messenger/auth/internal/server"
	"gitlab.com/bobr-lord-messenger/auth/internal/service"
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
	logrus.Info(fmt.Sprintf("%+v", cfg))

	if err := jwtutil.LoadKeys(cfg.PrivateKeyPath, cfg.PublicKeyPath); err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewPostgres(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	handler := hand.NewHandler(srvc)
	srvr := server.NewServer()

	go func() {
		if err := srvr.Run(cfg, handler); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srvr.Shutdown(ctx); err != nil {
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
