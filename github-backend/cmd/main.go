package main

import (
	"log"
	"net/http"

	"github.com/MarlomSouza/go-git/config"
	"github.com/MarlomSouza/go-git/internal/handlers"
	"github.com/MarlomSouza/go-git/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	defer log.Printf("Server running on port %s", cfg.Port)

	oauth2Config := &oauth2.Config{
		ClientID:     cfg.GitHubClientID,
		ClientSecret: cfg.GitHubClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       []string{"repo", "user", "read:org"},
		Endpoint:     github.Endpoint,
	}

	// Initialize Chi router
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	gitHubService := services.NewGitHubService()

	repoHandler := handlers.NewRepoHandler(gitHubService)
	repoHandler.RegisterRoutes(r)

	oAuthHandler := handlers.NewOAuthHandler(oauth2Config)
	oAuthHandler.RegisterRoutes(r)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":"+cfg.Port, r)

}
