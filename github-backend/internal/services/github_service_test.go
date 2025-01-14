package services

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/MarlomSouza/go-git/internal/models"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) R() *resty.Request {
	args := m.Called()
	return args.Get(0).(*resty.Request)
}

// MockRequest is a mock implementation of the resty.Request
type MockRequest struct {
	mock.Mock
	resty.Request
}

func (m *MockRequest) SetHeader(key, value string) *resty.Request {
	args := m.Called(key, value)
	return args.Get(0).(*resty.Request)
}

func (m *MockRequest) SetResult(result interface{}) *resty.Request {
	args := m.Called(result)
	return args.Get(0).(*resty.Request)
}

func (m *MockRequest) Get(url string) (*resty.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*resty.Response), args.Error(1)
}

func TestFetchReposSuccess(t *testing.T) {
	// Arrange
	mockClient := new(MockHTTPClient)
	mockRequest := new(MockRequest)
	mockResponse := &resty.Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`[{"id": 1, "name": "repo1"}, {"id": 2, "name": "repo2"}]`)),
		},
	}

	mockClient.On("R").Return(mockRequest)
	mockRequest.On("SetHeader", "Authorization", "Bearer valid-token").Return(mockRequest)
	mockRequest.On("SetHeader", "Accept", "application/vnd.github.v3+json").Return(mockRequest)
	mockRequest.On("SetResult", mock.Anything).Return(mockRequest)
	mockRequest.On("Get", "https://api.github.com/user/repos").Return(mockResponse, nil)

	svc := &GitHubServiceImp{
		BaseURL:    "https://api.github.com",
		HTTPClient: mockClient,
	}

	url := "https://api.github.com/user/repos"
	token := "valid-token"
	expectedRepos := []models.Repository{
		{Name: "repo1"},
		{Name: "repo2"},
	}

	// Act
	repos, err := svc.fetchRepos(url, token)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedRepos, repos)
}

func TestFetchReposError(t *testing.T) {
	// Arrange
	mockClient := new(MockHTTPClient)
	mockRequest := new(MockRequest)
	mockResponse := &resty.Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Bad credentials"}`)),
		},
	}

	mockClient.On("R").Return(mockRequest)
	mockRequest.On("SetHeader", "Authorization", "Bearer invalid-token").Return(mockRequest)
	mockRequest.On("SetHeader", "Accept", "application/vnd.github.v3+json").Return(mockRequest)
	mockRequest.On("SetResult", mock.Anything).Return(mockRequest)
	mockRequest.On("Get", "https://api.github.com/user/repos").Return(mockResponse, nil)

	svc := &GitHubServiceImp{
		BaseURL:    "https://api.github.com",
		HTTPClient: mockClient,
	}

	url := "https://api.github.com/user/repos"
	token := "invalid-token"

	// Act
	repos, err := svc.fetchRepos(url, token)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, repos)
	assert.Contains(t, err.Error(), "GitHub API error")
	mockClient.AssertExpectations(t)
	mockRequest.AssertExpectations(t)
}
