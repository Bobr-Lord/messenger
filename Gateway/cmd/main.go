package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	hand "gitlab.com/bobr-lord-messenger/gateway/internal/handler"
	"gitlab.com/bobr-lord-messenger/gateway/internal/jwtutil"
	"gitlab.com/bobr-lord-messenger/gateway/internal/repository"
	"gitlab.com/bobr-lord-messenger/gateway/internal/server"
	"gitlab.com/bobr-lord-messenger/gateway/internal/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Messenger API
// @version 1.0
// @description Это документация для вашего API

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info(cfg)

	err = jwtutil.LoadKeys("", cfg.PublicKeyPath)
	if err != nil {
		logrus.Fatalf("Failed to load keys: %v", err)
	}
	redisConn := initRedis()

	repo := repository.NewRepository(cfg)
	srvc := service.NewService(repo)
	handler := hand.NewHandler(srvc, redisConn, cfg)
	srvr := server.NewServer()
	srvr.Run(cfg, handler.InitRoutes())

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
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
