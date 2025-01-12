package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// GetPublicRepos fetches and returns public repositories
func (h *RepoHandler) GetPublicRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	repos, err := h.GitHubService.FetchPublicRepos()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch public repositories: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil

}
