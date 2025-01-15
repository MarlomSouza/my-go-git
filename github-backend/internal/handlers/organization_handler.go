package handlers

import (
	"net/http"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/MarlomSouza/go-git/internal/services"
	"github.com/go-chi/chi/v5"
)

// RepoHandler handles repository-related endpoints
type OrganizationHandler struct {
	GitHubService services.GitHubService
}

// NewRepoHandler creates a new instance of RepoHandler
func NewOrganizationHandler(service services.GitHubService) *OrganizationHandler {
	return &OrganizationHandler{GitHubService: service}
}

// RegisterRoutes registers routes for the RepoHandler
func (h *OrganizationHandler) RegisterRoutes(r chi.Router) {

	r.Route("/organization", func(r chi.Router) {
		r.Use(TokenFromCookieMiddleware)
		r.Get("/", HandlerError(h.GetOrganization))
		r.Get("/{org}/repos", HandlerError(h.GetOrganizationRepos))
		r.Get("/{org}/members", HandlerError(h.GetOrganizationMembers))
	})
}

func (h *OrganizationHandler) GetOrganization(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)

	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	organizations, err := h.GitHubService.FetchOrganization(token)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return organizations, http.StatusOK, nil
}

// GetOrganizationRepos fetches and returns repositories for a specific organization
func (h *OrganizationHandler) GetOrganizationRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)
	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	org := chi.URLParam(r, "org")
	repos, err := h.GitHubService.FetchOrganizationRepos(token, org)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil
}

func (h *OrganizationHandler) GetOrganizationMembers(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	token, ok := r.Context().Value("accessToken").(string)
	if !ok || token == "" {
		return nil, http.StatusUnauthorized, internalerrors.ErrUnauthorized
	}

	org := chi.URLParam(r, "org")
	members, err := h.GitHubService.FetchOrganizationMembers(token, org)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return members, http.StatusOK, nil
}
