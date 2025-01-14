package handlers

import (
	"bytes"
	"context"
	"encoding/json"
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
	req, res := newHttpTest(http.MethodGet, "/repos/", nil)
	req = addContext(req, "accessToken", expectedToken)

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
	req, res := newHttpTest(http.MethodGet, "/repos/", nil)
	req = addContext(req, "accessToken", expectedToken)

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
	_, statusCode, err := h.GetPrivateRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, statusCode)
	assert.Equal(t, internalerrors.ErrUnauthorized.Error(), err.Error())
}

func addContext(req *http.Request, keyParameter string, valueParameter string) *http.Request {
	ctx := context.WithValue(req.Context(), keyParameter, valueParameter)
	return req.WithContext(ctx)
}

func newHttpTest(method string, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {

	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req, _ := http.NewRequest(method, url, &buf)
	rr := httptest.NewRecorder()
	return req, rr
}
