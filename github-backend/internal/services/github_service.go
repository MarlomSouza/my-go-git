package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/MarlomSouza/go-git/internal/models"
)

type HTTPClient interface {
	R() *resty.Request
}

type GitHubService interface {
	FetchRepos(token string) ([]models.Repository, error)
	FetchUser(token string) (models.User, error)
	FetchOrganization(token string) ([]models.Organization, error)
	FetchOrganizationRepos(token string, org string) ([]models.Repository, error)
	FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error)
}

// GitHubService provides methods for interacting with the GitHub API
type GitHubServiceImp struct {
	HTTPClient HTTPClient
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

// fetchRepos handles the core logic for fetching repositories from the GitHub API
func (s *GitHubServiceImp) fetchRepos(url string, token string) ([]models.Repository, error) {
	var repos []models.Repository
	err := s.makeRequest(url, token, &repos)
	if err != nil {
		return make([]models.Repository, 0), err
	}

	return repos, nil
}

// FetchUser retrieves the authenticated user's profile information
func (s *GitHubServiceImp) FetchUser(token string) (models.User, error) {
	var user models.User
	err := s.makeRequest(fmt.Sprintf("%s/user", s.BaseURL), token, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// fetchOrganizations fetches the list of organizations the authenticated user is a part of
func (s *GitHubServiceImp) FetchOrganization(token string) ([]models.Organization, error) {
	var orgs []models.Organization
	err := s.makeRequest(fmt.Sprintf("%s/user/orgs", s.BaseURL), token, &orgs)
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (s *GitHubServiceImp) FetchOrganizationRepos(token string, org string) ([]models.Repository, error) {
	url := fmt.Sprintf("%s/orgs/%s/repos", s.BaseURL, org)
	return s.fetchRepos(url, token)
}

func (s *GitHubServiceImp) FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error) {
	var members []models.OrganizationMember
	err := s.makeRequest(fmt.Sprintf("%s/orgs/%s/members", s.BaseURL, org), token, &members)

	if err != nil {
		return nil, err
	}

	return members, nil
}

func (s *GitHubServiceImp) makeRequest(url, token string, result interface{}) error {
	resp, err := s.HTTPClient.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(result).
		Get(url)

	if resp.StatusCode() == http.StatusUnauthorized {
		return internalerrors.ErrUnauthorized
	}

	if err != nil {
		return errors.New("request failed: " + err.Error())
	}

	if resp.IsError() {
		return errors.New("GitHub API error: " + resp.String())
	}

	return nil
}
