package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config centralizes runtime configuration loaded from environment variables.
type Config struct {
	Port        string
	DatabaseURL string
}

func Load() Config {
	// Ignore load errors so production/container env vars keep precedence.
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Port:        port,
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
