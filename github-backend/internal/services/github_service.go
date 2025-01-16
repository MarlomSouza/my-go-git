package services

import (
	"fmt"

	"github.com/MarlomSouza/go-git/internal/infra"
	"github.com/MarlomSouza/go-git/internal/models"
)

type GitHubService interface {
	FetchRepos(token string) ([]models.Repository, error)
	FetchUser(token string) (models.User, error)
	FetchOrganization(token string) ([]models.Organization, error)
	FetchOrganizationRepos(token string, org string) ([]models.Repository, error)
	FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error)
}
type GitHubServiceImp struct {
	HTTPClient infra.GitHubHTTPClientInterface
}

func NewGitHubService(httpClient infra.GitHubHTTPClientInterface) *GitHubServiceImp {
	return &GitHubServiceImp{
		HTTPClient: httpClient,
	}
}

// FetchPersonalRepos retrieves the authenticated user's personal repositories
func (s *GitHubServiceImp) FetchRepos(token string) ([]models.Repository, error) {
	return s.fetchRepos("/user/repos?type=all", token)
}

// fetchRepos handles the core logic for fetching repositories from the GitHub API
func (s *GitHubServiceImp) fetchRepos(url string, token string) ([]models.Repository, error) {
	var repos []models.Repository
	err := s.HTTPClient.MakeRequest(url, token, &repos)
	if err != nil {
		return make([]models.Repository, 0), err
	}

	return repos, nil
}

// FetchUser retrieves the authenticated user's profile information
func (s *GitHubServiceImp) FetchUser(token string) (models.User, error) {
	var user models.User
	err := s.HTTPClient.MakeRequest("/user", token, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// fetchOrganizations fetches the list of organizations the authenticated user is a part of
func (s *GitHubServiceImp) FetchOrganization(token string) ([]models.Organization, error) {
	var orgs []models.Organization
	err := s.HTTPClient.MakeRequest("/user/orgs", token, &orgs)
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (s *GitHubServiceImp) FetchOrganizationRepos(token string, org string) ([]models.Repository, error) {
	url := fmt.Sprintf("/orgs/%s/repos", org)
	return s.fetchRepos(url, token)
}

func (s *GitHubServiceImp) FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error) {
	var members []models.OrganizationMember
	err := s.HTTPClient.MakeRequest(fmt.Sprintf("/orgs/%s/members", org), token, &members)

	if err != nil {
		return nil, err
	}

	return members, nil
}
