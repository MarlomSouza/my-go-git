package services

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/MarlomSouza/go-git/internal/models"
)

type GitHubService interface {
	FetchRepos(token string) ([]models.Repository, error)
	FetchPrivateRepos(token string) ([]models.Repository, error)
}

// GitHubService provides methods for interacting with the GitHub API
type GitHubServiceImp struct {
	HTTPClient *resty.Client
	BaseURL    string
}

// NewGitHubService initializes a new GitHubService
func NewGitHubService() *GitHubServiceImp {
	return &GitHubServiceImp{
		HTTPClient: resty.New(),
		BaseURL:    "https://api.github.com",
	}
}

// FetchPersonalRepos retrieves the authenticated user's personal repositories
func (s *GitHubServiceImp) FetchRepos(token string) ([]models.Repository, error) {
	url := fmt.Sprintf("%s/user/repos?type=all", s.BaseURL)
	return s.fetchRepos(url, token)
}

// FetchPersonalRepos fetches personal repositories of the authenticated user
func (s *GitHubServiceImp) FetchPrivateRepos(token string) ([]models.Repository, error) {
	url := fmt.Sprintf("%s/user/repos?type=private", s.BaseURL)
	return s.fetchRepos(url, token)
}

// fetchRepos handles the core logic for fetching repositories from the GitHub API
func (s *GitHubServiceImp) fetchRepos(url string, token string) ([]models.Repository, error) {
	// Use the injected HTTP client (resty.Client)
	client := s.HTTPClient

	// Store the response in a slice of repositories
	var repos []models.Repository
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
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

// // fetchOrganizations fetches the list of organizations the authenticated user is a part of
// func (s *GitHubServiceImp) fetchOrganizations() ([]string, error) {
// 	client := s.HTTPClient

// 	var orgs []struct {
// 		Login string `json:"login"`
// 	}

// 	resp, err := client.R().
// 		SetHeader("Authorization", "Bearer "+s.Token).
// 		SetHeader("Accept", "application/vnd.github.v3+json").
// 		SetResult(&orgs).
// 		Get(fmt.Sprintf("%s/user/orgs", s.BaseURL))

// 	if err != nil {
// 		return nil, fmt.Errorf("request failed: %w", err)
// 	}

// 	if resp.IsError() {
// 		return nil, fmt.Errorf("GitHub API error: %s", resp.String())
// 	}

// 	// Extract the organization logins from the response
// 	var orgNames []string
// 	for _, org := range orgs {
// 		orgNames = append(orgNames, org.Login)
// 	}

// 	return orgNames, nil
// }
