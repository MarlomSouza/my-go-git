package services

import (
	"testing"

	"github.com/MarlomSouza/go-git/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type GitHubHTTPClientMock struct {
	mock.Mock
}

func (m *GitHubHTTPClientMock) MakeRequest(url, token string, result interface{}) error {
	args := m.Called(url, token, result)
	return args.Error(0)
}

func Test_FetchUser_Should_fetchUserSuccessfully(t *testing.T) {

	expectedUser := models.User{
		Login:                   "test",
		ID:                      1,
		AvatarURL:               "test",
		Name:                    "test",
		Email:                   "test",
		PublicRepos:             1,
		PublicGists:             1,
		Followers:               1,
		Following:               1,
		CreatedAt:               "test",
		UpdatedAt:               "",
		PrivateGists:            1,
		TotalPrivateRepos:       1,
		OwnedPrivateRepos:       1,
		TwoFactorAuthentication: false,
		Organization:            "",
	}

	gitHubServiceMock := new(GitHubHTTPClientMock)
	gitHubServiceMock.On("MakeRequest", "/user", "token", mock.Anything).
		Run(func(args mock.Arguments) {
			// Populate the third argument (pointer to user) with expectedUser
			result := args.Get(2).(*models.User)
			*result = expectedUser
		}).
		Return(nil)
	service := NewGitHubService(gitHubServiceMock)

	sut, err := service.FetchUser("token")

	assert.Equal(t, expectedUser, sut)
	assert.Nil(t, err)

}
