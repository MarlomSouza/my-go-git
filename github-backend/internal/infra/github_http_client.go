package infra

import (
	"errors"
	"net/http"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"

	"github.com/go-resty/resty/v2"
)

type GitHubHTTPClientInterface interface {
	MakeRequest(url, token string, result interface{}) error
}

type GitHubHTTPClient struct {
	Client *resty.Client
}

func NewGitHubHTTPClient() *GitHubHTTPClient {
	return &GitHubHTTPClient{
		Client: resty.New(),
	}
}

func (c *GitHubHTTPClient) MakeRequest(url, token string, result interface{}) error {
	resp, err := c.Client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetResult(result).
		Get(url)

	if resp.StatusCode() == http.StatusUnauthorized {
		return internalerrors.ErrUnauthorized
	}

	if err != nil {
		return errors.New("request failed: " + err.Error())
	}

	if resp.IsError() {
		return errors.New("GitHub API error: " + resp.String())
	}

	return nil
}
