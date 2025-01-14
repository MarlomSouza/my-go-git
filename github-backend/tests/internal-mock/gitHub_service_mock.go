package internalmock

import (
	"github.com/MarlomSouza/go-git/internal/models"
	"github.com/stretchr/testify/mock"
)

type GithubServiceMock struct {
	mock.Mock
}

func (m *GithubServiceMock) FetchPrivateRepos(token string) ([]models.Repository, error) {
	args := m.Called(token)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Repository), args.Error(1)
}

func (m *GithubServiceMock) FetchRepos(token string) ([]models.Repository, error) {
	args := m.Called(token)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Repository), args.Error(1)
}

func (m *GithubServiceMock) FetchOrganization(token string) ([]models.Organization, error) {
	args := m.Called(token)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Organization), args.Error(1)
}

func (m *GithubServiceMock) FetchUser(token string) (models.User, error) {
	args := m.Called(token)
	if args.Error(1) != nil {
		return models.User{}, args.Error(1)
	}
	return args.Get(0).(models.User), args.Error(1)
}

func (m *GithubServiceMock) FetchOrganizationRepos(token string, org string) ([]models.Repository, error) {
	args := m.Called(token, org)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Repository), args.Error(1)
}

func (m *GithubServiceMock) FetchOrganizationMembers(token string, org string) ([]models.OrganizationMember, error) {
	args := m.Called(token, org)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.OrganizationMember), args.Error(1)
}
