package internalmock

import (
	"github.com/MarlomSouza/go-git/internal/models"
	"github.com/stretchr/testify/mock"
)

type GithubServiceMock struct {
	mock.Mock
}

func (m *GithubServiceMock) FetchPrivateRepos() ([]models.Repository, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Repository), args.Error(1)
}

func (m *GithubServiceMock) FetchPublicRepos() ([]models.Repository, error) {
	args := m.Called()
	return args.Get(0).([]models.Repository), args.Error(1)
}
