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

func (m *GithubServiceMock) FetchOrganization(token string) ([]string, error) {
	args := m.Called(token)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}
