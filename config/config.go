package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret []byte
}

func Load() (*Config, error) {

	_ = godotenv.Load()

	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if secret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return &Config{
		JWTSecret: []byte(secret),
	}, nil
}
