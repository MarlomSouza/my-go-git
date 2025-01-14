package services

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/MarlomSouza/go-git/internal/models"
)

type GitHubService interface {
	FetchRepos(token string) ([]models.Repository, error)
	FetchPrivateRepos(token string) ([]models.Repository, error)
	FetchUser(token string) (models.User, error)
	FetchOrganization(token string) ([]models.Organization, error)
	FetchOrganizationRepos(token string, org string) ([]models.Repository, error)
	FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error)
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

// FetchUser retrieves the authenticated user's profile information
func (s *GitHubServiceImp) FetchUser(token string) (models.User, error) {
	client := s.HTTPClient

	var user models.User
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(&user).
		Get(fmt.Sprintf("%s/user", s.BaseURL))

	if err != nil {
		return models.User{}, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return models.User{}, fmt.Errorf("GitHub API error: %s", resp.String())
	}

	return user, nil
}

// fetchOrganizations fetches the list of organizations the authenticated user is a part of
func (s *GitHubServiceImp) FetchOrganization(token string) ([]models.Organization, error) {
	client := s.HTTPClient

	var orgs []models.Organization

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(&orgs).
		Get(fmt.Sprintf("%s/user/orgs", s.BaseURL))

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("GitHub API error: %s", resp.String())
	}

	return orgs, nil
}

func (s *GitHubServiceImp) FetchOrganizationRepos(token string, org string) ([]models.Repository, error) {
	url := fmt.Sprintf("%s/orgs/%s/repos", s.BaseURL, org)
	return s.fetchRepos(url, token)
}

func (s *GitHubServiceImp) FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error) {
	client := s.HTTPClient

	var members []models.OrganizationMember
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(&members).
		Get(fmt.Sprintf("%s/orgs/%s/members", s.BaseURL, org))

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("GitHub API error: %s", resp.String())
	}

	return members, nil
}
