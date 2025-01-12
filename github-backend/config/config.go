package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GitHubClientID     string
	GitHubClientSecret string
	AppPrivateKey      string
	AppID              string
	OrgName            string
	Port               string
}

func LoadConfig() *Config {
	// Load .env file (optional, depending on deployment strategy)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get environment variables
	config := &Config{
		GitHubClientID:     getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		AppPrivateKey:      getEnv("APP_PRIVATE_KEY", ""),
		AppID:              getEnv("APP_ID", ""),
		OrgName:            getEnv("GITHUB_ORG_NAME", ""),
		Port:               getEnv("PORT", "8080"),
	}

	// Validate required variables
	if config.GitHubClientID == "" || config.GitHubClientSecret == "" || config.AppPrivateKey == "" || config.AppID == "" {
		log.Fatal("Missing required environment variables")
	}

	return config
}

// Helper function to fetch environment variables with a fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
