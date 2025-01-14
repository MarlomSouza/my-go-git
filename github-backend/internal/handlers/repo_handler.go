package handlers

import (
	"net/http"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/MarlomSouza/go-git/internal/models"
	"github.com/MarlomSouza/go-git/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// RepoHandler handles repository-related endpoints
type RepoHandler struct {
	GitHubService services.GitHubService
}

// NewRepoHandler creates a new instance of RepoHandler
func NewRepoHandler(service services.GitHubService) *RepoHandler {
	return &RepoHandler{GitHubService: service}
}

// RegisterRoutes registers routes for the RepoHandler
func (h *RepoHandler) RegisterRoutes(r chi.Router) {
	r.Route("/repos", func(r chi.Router) {
		r.Use(TokenFromCookieMiddleware)
		r.Get("/", HandlerError(h.GetRepos))
		r.Get("/user", HandlerError(h.GetUser))
	})
}

func (h *RepoHandler) GetRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)

	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	var repos []models.Repository
	var err error

	repos, err = h.GitHubService.FetchRepos(token)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch repositories: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil
}

// GetUser fetches and returns user information
func (h *RepoHandler) GetUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)
	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	user, err := h.GitHubService.FetchUser(token)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch user info: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}
