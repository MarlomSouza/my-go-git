package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	GitHubToken string
}

func LoadConfig() (*Config, error) {
	// Load .env file (optional, depending on deployment strategy)
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("no .env file found")
	}

	// Get environment variables
	config := &Config{

		Port:        getEnv("PORT", "5000"),
		GitHubToken: getEnv("GITHUB_TOKEN", ""),
	}

	// Validate required variables
	if config.GitHubToken == "" {
		return nil, errors.New("missing required environment variables")

	}

	return config, nil
}

// Helper function to fetch environment variables with a fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
