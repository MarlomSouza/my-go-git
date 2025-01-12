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

func Test_GetPrivateRepos_Should_Return_PrivateRepo(t *testing.T) {
	//Arrange
	expectedRepos := []models.Repository{
		{
			Name:         "repo1",
			Description:  "description1",
			LastUpdate:   time.Now(),
			Private:      true,
			Organization: "org1",
		},
	}
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}
	gitHubServiceMock.On("FetchPrivateRepos").Return(expectedRepos, nil)
	req := httptest.NewRequest(http.MethodGet, "/repos/private", nil)
	res := httptest.NewRecorder()

	// Act
	obj, statusCode, _ := h.GetPrivateRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, obj.([]models.Repository), expectedRepos)

}

func Test_GetPrivateRepos_Should_Return_Error(t *testing.T) {
	//Arrange
	expectedError := errors.New("Error when fetching")
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}
	gitHubServiceMock.On("FetchPrivateRepos").Return(nil, expectedError)
	req := httptest.NewRequest(http.MethodGet, "/repos/private", nil)
	res := httptest.NewRecorder()

	// Act
	_, statusCode, err := h.GetPrivateRepos(res, req)

	// Assert
	assert.Equal(t, statusCode, res.Code)
	assert.Equal(t, err.Error(), expectedError.Error())
}
