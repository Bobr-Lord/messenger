package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/auth/internal/config"
	hand "gitlab.com/bobr-lord-messenger/auth/internal/handler"
	"gitlab.com/bobr-lord-messenger/auth/internal/jwtutil"
	"gitlab.com/bobr-lord-messenger/auth/internal/repository"
	"gitlab.com/bobr-lord-messenger/auth/internal/server"
	"gitlab.com/bobr-lord-messenger/auth/internal/service"
	"os"
	"os/signal"
	"syscall"
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

	srvr.Run(cfg, handler)

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait
	logrus.Info("Shutting down server...")
	srvr.Shutdown()
	logrus.Info("Server shut down.")
}
