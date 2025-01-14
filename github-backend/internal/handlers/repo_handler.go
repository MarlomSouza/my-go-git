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
		r.Get("/private", HandlerError(h.GetPrivateRepos))
		r.Get("/user", HandlerError(h.GetUser))
	})

	r.Route("/organization", func(r chi.Router) {
		r.Use(TokenFromCookieMiddleware)
		r.Get("/", HandlerError(h.GetOrganization))
		r.Get("/{org}/repos", HandlerError(h.GetOrganizationRepos))
		r.Get("/{org}/members", HandlerError(h.GetOrganizationMembers))
	})

}

func (h *RepoHandler) GetRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	return h.getAllRepos(w, r, false)
}

// GetPrivateRepos fetches and returns private repositories
func (h *RepoHandler) GetPrivateRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	return h.getAllRepos(w, r, true)
}

func (h *RepoHandler) getAllRepos(w http.ResponseWriter, r *http.Request, private bool) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)

	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	var repos []models.Repository
	var err error

	if private {
		repos, err = h.GitHubService.FetchPrivateRepos(token)
	} else {
		repos, err = h.GitHubService.FetchRepos(token)
	}

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
func (h *RepoHandler) GetOrganization(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)

	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	organizations, err := h.GitHubService.FetchOrganization(token)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch repositories: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return organizations, http.StatusOK, nil
}

// GetOrganizationRepos fetches and returns repositories for a specific organization
func (h *RepoHandler) GetOrganizationRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)
	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	org := chi.URLParam(r, "org")
	repos, err := h.GitHubService.FetchOrganizationRepos(token, org)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch organization repos: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil
}

func (h *RepoHandler) GetOrganizationMembers(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)
	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	org := chi.URLParam(r, "org")
	members, err := h.GitHubService.FetchOrganizationMembers(token, org)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch organization members: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return members, http.StatusOK, nil
}
