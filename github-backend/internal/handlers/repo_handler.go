package handlers

import (
	"net/http"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
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
	r.Get("/repos/private", HandlerError(h.GetPrivateRepos))
	r.Get("/repos/public", HandlerError(h.GetPublicRepos))
}

func (h *RepoHandler) GetPublicRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	//get from header the token

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	repos, err := h.GitHubService.FetchPublicRepos(authHeader)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch public repositories: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil
}

// GetPrivateRepos fetches and returns private repositories
func (h *RepoHandler) GetPrivateRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, http.StatusUnauthorized, nil
	}

	repos, err := h.GitHubService.FetchPrivateRepos(authHeader)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch private repositories: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil
}
