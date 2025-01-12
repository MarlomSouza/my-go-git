package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MarlomSouza/go-git/internal/models"
	internalmock "github.com/MarlomSouza/go-git/tests/internal-mock"
	"github.com/stretchr/testify/assert"
)

func Test_GetPublicRepos_Should_Return_PublicRepo(t *testing.T) {
	//Arrange
	expectedRepos := []models.Repository{
		{
			Name:         "repo1",
			Description:  "description1",
			LastUpdate:   time.Now(),
			Private:      false,
			Organization: "org1",
		},
	}
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}
	gitHubServiceMock.On("FetchPublicRepos").Return(expectedRepos, nil)
	req := httptest.NewRequest(http.MethodGet, "/repos/public", nil)
	res := httptest.NewRecorder()

	// Act
	obj, statusCode, _ := h.GetPublicRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.False(t, obj.([]models.Repository)[0].Private)
}

func Test_GetPublicRepos_Should_Return_Error(t *testing.T) {
	//Arrange
	expectedError := errors.New("error when fetching")
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}
	gitHubServiceMock.On("FetchPublicRepos").Return(nil, expectedError)
	req := httptest.NewRequest(http.MethodGet, "/repos/public", nil)
	res := httptest.NewRecorder()

	// Act
	_, statusCode, err := h.GetPublicRepos(res, req)

	// Assert
	assert.Equal(t, statusCode, res.Code)
	assert.Equal(t, err.Error(), expectedError.Error())
}
