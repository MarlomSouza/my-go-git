package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GitHubClientID     string
	GitHubClientSecret string
	RedirectURL        string
	Port               string
}

func LoadConfig() (*Config, error) {
	// Load .env file (optional, depending on deployment strategy)
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("no .env file found")
	}

	// Get environment variables
	config := &Config{
		GitHubClientID:     getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		RedirectURL:        getEnv("REDIRECT_URL", "http://localhost:5000/login/github/callback"),
		Port:               getEnv("PORT", "5000"),
	}

	// Validate required variables
	if config.GitHubClientID == "" || config.GitHubClientSecret == "" {
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
