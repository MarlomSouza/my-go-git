package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GitHubClientID      string
	GitHubClientSecret  string
	GitHubRedirectURL   string
	FrontendRedirectURL string
	Port                string
}

func LoadConfig() (*Config, error) {
	// Load .env file (optional, depending on deployment strategy)
	if err := godotenv.Load(); err != nil {
		if os.IsNotExist(err) {
			// Log a warning if .env file is not found, but continue
			fmt.Println("Warning: no .env file found, using environment variables")
		} else {
			return nil, errors.New("no .env file found")
		}
	}

	// Get environment variables
	config := &Config{
		GitHubClientID:      getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret:  getEnv("GITHUB_CLIENT_SECRET", ""),
		GitHubRedirectURL:   getEnv("GITHUB_REDIRECT_URL", "http://localhost:5000/login/github/callback"),
		FrontendRedirectURL: getEnv("FRONTEND_REDIRECT_URL", "http://localhost:3000"),
		Port:                getEnv("PORT", "5000"),
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
