package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/MarlomSouza/go-git/internal/models"
	internalmock "github.com/MarlomSouza/go-git/tests/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_FetchRepos_Success(t *testing.T) {
	//Arrange
	expectedRepos := []models.Repository{
		{
			Name:        "repo1",
			Description: "description1",
			LastUpdate:  time.Now(),
			Private:     true,
		},
		{
			Name:        "repo2",
			Description: "description2",
			LastUpdate:  time.Now(),
			Private:     false,
		},
	}
	expectedToken := "test-token"
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchRepos", mock.MatchedBy(func(request string) bool { return request == expectedToken })).Return(expectedRepos, nil)
	req, res := NewHttpTest(http.MethodGet, "/repos/", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	obj, statusCode, _ := h.GetRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, obj.([]models.Repository), expectedRepos)

}

func Test_FetchRepos_InternalServerError(t *testing.T) {
	//Arrange
	expectedToken := "test-token"
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchRepos", mock.MatchedBy(func(request string) bool { return request == expectedToken })).Return(nil, internalerrors.ErrInternal)
	req, res := NewHttpTest(http.MethodGet, "/repos/", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	_, statusCode, error := h.GetRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, internalerrors.ErrInternal.Error(), error.Error())
}

func Test_GetRepos_MissingAccessToken(t *testing.T) {
	//Arrange
	h := RepoHandler{
		GitHubService: new(internalmock.GithubServiceMock),
	}

	req := httptest.NewRequest(http.MethodGet, "/repos/private", nil)
	res := httptest.NewRecorder()

	// Act
	_, statusCode, err := h.GetRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, statusCode)
	assert.Equal(t, internalerrors.ErrUnauthorized.Error(), err.Error())
}

func Test_FetchUser_Success(t *testing.T) {
	//Arrange
	expectedUser := models.User{
		Login:                   "test-user",
		ID:                      1,
		AvatarURL:               "http://avatar.com",
		Name:                    "Test User",
		Email:                   "xxx@gmail.com.com",
		PublicRepos:             1,
		PublicGists:             1,
		Followers:               1,
		Following:               1,
		CreatedAt:               "2021-01-01",
		UpdatedAt:               "2021-01-01",
		PrivateGists:            1,
		TotalPrivateRepos:       1,
		OwnedPrivateRepos:       1,
		TwoFactorAuthentication: true,
		Organization:            "test-org",
	}
	expectedToken := "test-token"
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchUser", mock.MatchedBy(func(request string) bool {
		return request == expectedToken
	})).Return(expectedUser, nil)
	req, res := NewHttpTest(http.MethodGet, "/user", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	obj, statusCode, _ := h.GetUser(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, obj.(models.User), expectedUser)

}

func Test_FetchUser_InternalServerError(t *testing.T) {
	//Arrange
	expectedToken := "test-token"
	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := RepoHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchUser", mock.MatchedBy(func(request string) bool {
		return request == expectedToken
	})).Return(models.User{}, internalerrors.ErrInternal)
	req, res := NewHttpTest(http.MethodGet, "/user", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	_, statusCode, error := h.GetUser(res, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, internalerrors.ErrInternal.Error(), error.Error())

}
