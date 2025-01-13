package services

import (
	"testing"

	"github.com/MarlomSouza/go-git/internal/models"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestFetchReposSuccess(t *testing.T) {
	// Arrange
	mockClient := &resty.Client{}
	svc := &GitHubServiceImp{
		BaseURL:    "https://api.github.com",
		HTTPClient: mockClient,
	}

	expectedRepos := []models.Repository{
		{Name: "repo1", Description: "test repo 1"},
		{Name: "repo2", Description: "test repo 2"},
	}

	token := "valid-token"

	// Act
	repos, err := svc.FetchRepos(token)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedRepos, repos)
}
