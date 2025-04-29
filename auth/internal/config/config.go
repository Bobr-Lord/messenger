package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost     string `env:"SERVER_HOST"`
	ServerPort     string `env:"SERVER_PORT"`
	PublicKeyPath  string `env:"PUBLIC_KEY_PATH"`
	PrivateKeyPath string `env:"PRIVATE_KEY_PATH"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	godotenv.Load(".env")
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("Error reading config: %s", err)
	}

	return &cfg, nil
}
