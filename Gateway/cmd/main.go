package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/gateway/internal/config"
	hand "gitlab.com/bobr-lord-messenger/gateway/internal/handler"
	"gitlab.com/bobr-lord-messenger/gateway/internal/jwtutil"
	cnsm "gitlab.com/bobr-lord-messenger/gateway/internal/kafka/consumer"
	prd "gitlab.com/bobr-lord-messenger/gateway/internal/kafka/producer"
	"gitlab.com/bobr-lord-messenger/gateway/internal/repository"
	"gitlab.com/bobr-lord-messenger/gateway/internal/server"
	"gitlab.com/bobr-lord-messenger/gateway/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Messenger API
// @version 1.0
// @description API Gateway for the Messenger service
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
	logrus.Info(cfg)

	err = jwtutil.LoadKeys("", cfg.PublicKeyPath)
	if err != nil {
		logrus.Fatalf("Failed to load keys: %v", err)
	}
	redisConn := initRedis(cfg)

	addrKafka := []string{
		cfg.KafkaHost + ":" + cfg.KafkaPort,
	}
	producer := prd.NewProducerMessage(addrKafka)

	repo := repository.NewRepository(cfg)
	svc := service.NewService(repo)
	handler := hand.NewHandler(svc, redisConn, cfg, producer)

	consumer := cnsm.NewConsumerMessage(addrKafka, handler)
	go func() {
		consumer.Start(context.Background())
	}()

	srv := server.NewServer()
	go func() {
		if err := srv.Run(cfg, handler.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("Error starting server: %v", err)
		}
	}()
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait

	logrus.Info("Shutting down server...")

	if err := consumer.Close(); err != nil {
		logrus.Errorf("Error closing consumer: %v", err)
	}
	if err := producer.Close(); err != nil {
		logrus.Errorf("Error closing producer: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("Server shutdown error: %v", err)
	} else {
		logrus.Info("Server shutdown complete.")
	}
}

func initRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
	})
}
