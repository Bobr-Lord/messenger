package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerHost    string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	ServerPort    string `env:"SERVER_PORT"`
	PublicKeyPath string `env:"PUBLIC_KEY_PATH"  envDefault:"keys/public.pem"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load(".env")
	if err == nil {
		logrus.Info("Loading environment variables")
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("could not read config: %w", err)
	}
	return &cfg, nil
}
