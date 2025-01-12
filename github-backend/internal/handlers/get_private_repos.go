package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// GetPrivateRepos fetches and returns personal repositories
func (h *RepoHandler) GetPrivateRepos(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	repos, err := h.GitHubService.FetchPrivateRepos()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to fetch personal repositories: " + err.Error()})
		return nil, http.StatusInternalServerError, err
	}

	return repos, http.StatusOK, nil
}
