package main

import (
	"log"
	"net/http"

	"github.com/MarlomSouza/go-git/config"
	"github.com/MarlomSouza/go-git/internal/handlers"
	"github.com/MarlomSouza/go-git/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
