package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MarlomSouza/go-git/internal/models"
)

type GitHubService interface {
	FetchPublicRepos() ([]models.Repository, error)
	FetchPrivateRepos() ([]models.Repository, error)
}

// GitHubService provides methods for interacting with the GitHub API
type GitHubServiceImp struct {
	HTTPClient *http.Client
	BaseURL    string
	Token      string
}

// NewGitHubService initializes a new GitHubService
func NewGitHubService(token string) *GitHubServiceImp {
	return &GitHubServiceImp{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    "https://api.github.com",
		Token:      token,
	}
}

// FetchPersonalRepos retrieves the authenticated user's personal repositories
func (s *GitHubServiceImp) FetchPublicRepos() ([]models.Repository, error) {
	url := fmt.Sprintf("%s/user/repos", s.BaseURL)
	return s.fetchRepos(url)
}

// FetchPersonalRepos fetches personal repositories of the authenticated user
func (s *GitHubServiceImp) FetchPrivateRepos() ([]models.Repository, error) {
	url := fmt.Sprintf("%s/user/repos?type=private", s.BaseURL)
	return s.fetchRepos(url)
}

// fetchRepos handles the core logic for fetching repositories from the GitHub API
func (s *GitHubServiceImp) fetchRepos(url string) ([]models.Repository, error) {
	client := s.HTTPClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s", string(body))
	}

	var repos []models.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return repos, nil
}
