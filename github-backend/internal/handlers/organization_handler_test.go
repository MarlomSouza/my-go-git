package handlers

import (
	"net/http"
	"testing"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/MarlomSouza/go-git/internal/models"
	internalmock "github.com/MarlomSouza/go-git/tests/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	expectedToken = "test"
)

func Test_FetchOrganization_Success(t *testing.T) {
	//Arrange
	expectedOrganizations := []models.Organization{
		{
			Login:       "org1",
			URL:         "url1",
			AvatarURL:   "avatar1",
			Description: "desc1",
			ID:          1,
		},
		{
			Login:       "org2",
			URL:         "url2",
			AvatarURL:   "avatar2",
			Description: "desc2",
			ID:          2,
		},
	}

	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := OrganizationHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchOrganization", mock.MatchedBy(func(request string) bool {
		return request == expectedToken
	})).Return(expectedOrganizations, nil)
	req, res := NewHttpTest(http.MethodGet, "/organization/", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	obj, statusCode, _ := h.GetOrganization(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, obj.([]models.Organization), expectedOrganizations)
}

func Test_FetchOrganization_InternalServerError(t *testing.T) {
	//Arrange

	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := OrganizationHandler{

		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchOrganization", mock.MatchedBy(func(request string) bool {
		return request == expectedToken
	})).Return(nil, internalerrors.ErrInternal)
	req, res := NewHttpTest(http.MethodGet, "/organization/", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	_, statusCode, error := h.GetOrganization(res, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, internalerrors.ErrInternal.Error(), error.Error())
}

func Test_GetOrganizationRepo_Success(t *testing.T) {
	//Arrange
	expectedRepos := []models.Repository{
		{
			Name:        "repo1",
			Description: "desc1",
		},
	}

	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := OrganizationHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchOrganizationRepos", mock.Anything, mock.Anything).Return(expectedRepos, nil)

	req, res := NewHttpTest(http.MethodGet, "/organization/org1/repos", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	obj, statusCode, _ := h.GetOrganizationRepos(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, obj.([]models.Repository), expectedRepos)
}

func Test_GetOrganizationMember_Success(t *testing.T) {
	//Arrange
	expectedMembers := []models.OrganizationMember{
		{
			Login: "member1",
		},
	}

	gitHubServiceMock := new(internalmock.GithubServiceMock)
	h := OrganizationHandler{
		GitHubService: gitHubServiceMock,
	}

	gitHubServiceMock.On("FetchOrganizationMembers", mock.Anything, mock.Anything).Return(expectedMembers, nil)

	req, res := NewHttpTest(http.MethodGet, "/organization/org1/members", nil)
	req = AddContext(req, "accessToken", expectedToken)

	// Act
	obj, statusCode, _ := h.GetOrganizationMembers(res, req)

	// Assert
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, obj.([]models.OrganizationMember), expectedMembers)

}
