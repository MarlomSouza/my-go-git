package handlers

import (
	"github.com/MarlomSouza/go-git/internal/services"
	"github.com/go-chi/chi/v5"
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
	// r.Get("/repos/org/{orgName}", h.GetOrgRepos)
}
