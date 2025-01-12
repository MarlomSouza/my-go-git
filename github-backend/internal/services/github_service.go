package services

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/MarlomSouza/go-git/internal/models"
)

type GitHubService interface {
	FetchPublicRepos() ([]models.Repository, error)
	FetchPrivateRepos() ([]models.Repository, error)
}

// GitHubService provides methods for interacting with the GitHub API
type GitHubServiceImp struct {
	HTTPClient *resty.Client
	BaseURL    string
	Token      string
}

// NewGitHubService initializes a new GitHubService
func NewGitHubService(token string) *GitHubServiceImp {
	return &GitHubServiceImp{
		HTTPClient: resty.New(),
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
	// Use the injected HTTP client (resty.Client)
	client := s.HTTPClient

	// Store the response in a slice of repositories
	var repos []models.Repository
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+s.Token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(&repos). // Automatically decode JSON into the repos variable
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("GitHub API error: %s", resp.String())
	}

	return repos, nil
}
