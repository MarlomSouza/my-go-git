package handlers

import (
	"net/http"

	internalerrors "github.com/MarlomSouza/go-git/internal-errors"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

// OAuthHandler handles GitHub OAuth operations
type OAuthHandler struct {
	OAuthConfig *oauth2.Config
}

func (h *OAuthHandler) RegisterRoutes(r chi.Router) {

	r.Get("/login/github", h.GitHubLogin)
	r.Get("/login/github/callback", HandlerError(h.GitHubCallback))
}

// NewOAuthHandler creates a new instance of OAuthHandler
func NewOAuthHandler(config *oauth2.Config) *OAuthHandler {
	return &OAuthHandler{OAuthConfig: config}
}

// Handle GitHub Login
func (h *OAuthHandler) GitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := h.OAuthConfig.AuthCodeURL("random_state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

// GitHub Callback Handler
func (h *OAuthHandler) GitHubCallback(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	code := r.URL.Query().Get("code")
	if code == "" {
		return nil, http.StatusBadRequest, internalerrors.ErrInternal
	}

	// Exchange the code for an access token
	token, err := h.OAuthConfig.Exchange(r.Context(), code)
	if err != nil {

		return nil, http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}
